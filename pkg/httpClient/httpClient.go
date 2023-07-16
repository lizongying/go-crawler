package httpClient

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	response2 "github.com/lizongying/go-crawler/pkg/response"
	"github.com/lizongying/go-crawler/pkg/utils"
	"github.com/lizongying/go-crawler/static"
	utls "github.com/refraction-networking/utls"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

type Option struct {
	HelloID            *utls.ClientHelloID
	HelloSpec          *utls.ClientHelloSpec
	ServerName         string
	InsecureSkipVerify bool
	RootCAs            *x509.CertPool
	Http2              bool
}
type HttpClient struct {
	Ja3 bool
	Option
	DialTimeout      *time.Duration
	client           *http.Client
	proxy            *url.URL
	timeout          time.Duration
	httpProto        string
	logger           pkg.Logger
	redirectMaxTimes uint8
	retryMaxTimes    uint8
}

func NewClientJa3(ctx context.Context, conn net.Conn, option Option) (net.Conn, error) {
	config := &utls.Config{
		RootCAs:            option.RootCAs,
		InsecureSkipVerify: option.InsecureSkipVerify,
		ServerName:         option.ServerName,
	}
	if option.Http2 {
		config.NextProtos = []string{"h2", "http/1.1"}
	} else {
		config.NextProtos = []string{"http/1.1"}
	}

	if option.HelloID == nil {
		option.HelloID = &utls.HelloChrome_Auto
	}
	if option.HelloSpec != nil {
		option.HelloID = &utls.HelloCustom
	}
	c := utls.UClient(conn, config, *option.HelloID)

	if *option.HelloID == utls.HelloCustom {
		if err := c.ApplyPreset(option.HelloSpec); err != nil {
			return nil, err
		}
	}

	if err := c.HandshakeContext(ctx); err != nil {
		return nil, err
	}

	return c, nil
}

func NewClient(conn net.Conn, option Option) net.Conn {
	config := &tls.Config{
		RootCAs:            option.RootCAs,
		InsecureSkipVerify: option.InsecureSkipVerify,
		ServerName:         option.ServerName,
	}
	if option.Http2 {
		config.NextProtos = []string{"h2", "http/1.1"}
	} else {
		config.NextProtos = []string{"http/1.1"}
	}

	return tls.Client(conn, config)
}

func (h *HttpClient) DoRequest(ctx context.Context, request pkg.Request) (response pkg.Response, err error) {
	h.logger.DebugF("request: %+v", request.GetRequest())

	if ctx == nil {
		ctx = context.Background()
	}

	if request.GetTimeout() > 0 {
		//c, cancel := context.WithTimeout(ctx, request.Timeout)
		//defer cancel()
		//request.Request = request.Request.WithContext(c)
	}

	timeout := h.timeout
	if request.GetTimeout() > 0 {
		timeout = request.GetTimeout()
	}

	// Get a copy of the default root CAs
	defaultCAs, err := x509.SystemCertPool()
	if err != nil {
		defaultCAs = x509.NewCertPool()
	}
	defaultCAs.AppendCertsFromPEM(static.Cert)

	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 60 * time.Second,
		}).DialContext,
		DisableKeepAlives:     true,
		IdleConnTimeout:       180 * time.Second,
		TLSHandshakeTimeout:   20 * time.Second,
		ExpectContinueTimeout: 2 * time.Second,
		ResponseHeaderTimeout: 30 * time.Second,

		MaxConnsPerHost:     1000,
		MaxIdleConns:        1000,
		MaxIdleConnsPerHost: 1000,
		TLSClientConfig: &tls.Config{
			RootCAs: defaultCAs,
			//InsecureSkipVerify: true,
		},
	}
	proxyEnable := false
	if request.GetProxyEnable() != nil {
		proxyEnable = *request.GetProxyEnable()
	}
	if proxyEnable {
		proxy := h.proxy
		if request.GetProxy() != nil {
			proxy = request.GetProxy()
		}
		if proxy == nil {
			err = errors.New("nil proxy")
			return
		}
		transport.Proxy = http.ProxyURL(proxy)
	}

	httpProto := h.httpProto
	if request.GetHttpProto() != "" {
		httpProto = request.GetHttpProto()
	}
	if httpProto != "2.0" {
		transport.ForceAttemptHTTP2 = false
	} else {
		transport.ForceAttemptHTTP2 = true
	}

	client := h.client
	client.Transport = transport

	if timeout > 0 {
		client.Timeout = timeout
	}

	if h.DialTimeout == nil {
		dialTimeout := 10 * time.Second
		h.DialTimeout = &dialTimeout
	}

	network := request.GetURL().Scheme
	address := request.GetURL().Host

	// Check if the URL specifies a port, otherwise use the default port
	if request.GetURL().Port() == "" {
		switch network {
		case "http":
			address += ":80"
		case "https":
			address += ":443"
		default:
			return nil, errors.New(fmt.Sprintf("Unsupported network: %s\n", network))
		}
	}

	conn, err := net.DialTimeout("tcp", address, *h.DialTimeout)
	if err != nil {
		return nil, err
	}

	option := Option{
		RootCAs:    defaultCAs,
		ServerName: request.GetURL().Hostname(),
		Http2:      h.httpProto == "2.0",
	}
	if h.Ja3 {
		conn, _ = NewClientJa3(ctx, conn, option)
	} else {
		conn = NewClient(conn, option)
	}

	begin := time.Now()
	response = new(response2.Response).SetRequest(request)

	var resp *http.Response
	if h.httpProto == "2.0" {
		request.GetRequest().Proto = "HTTP/2.0"
		request.GetRequest().ProtoMajor = 2
		request.GetRequest().ProtoMinor = 0

		tr := transport
		tr.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return conn, nil
		}
		tr.DialTLSContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return conn, nil
		}

		resp, err = tr.RoundTrip(request.GetRequest())
		if err != nil {
			h.logger.Error(err)
			return
		}
		response.SetResponse(resp)
	} else {
		request.GetRequest().Proto = "HTTP/1.1"
		request.GetRequest().ProtoMajor = 1
		request.GetRequest().ProtoMinor = 1

		tr := transport
		tr.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return conn, nil
		}
		tr.DialTLSContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return conn, nil
		}

		resp, err = tr.RoundTrip(request.GetRequest())
		if err != nil {
			h.logger.Error(err)
			return
		}
		response.SetResponse(resp)
	}

	response.SetSpendTime(time.Now().Sub(begin))
	if err != nil {
		retryMaxTimes := h.retryMaxTimes
		if request.GetRetryMaxTimes() != nil {
			retryMaxTimes = *request.GetRetryMaxTimes()
		}
		if request.GetRetryTimes() < retryMaxTimes {
			return
		}
		h.logger.Error(err, "RetryTimes:", request.GetRetryTimes())
		h.logger.ErrorF("request: %+v", request)
		h.logger.Debug(utils.Request2Curl(request))
		return
	}

	defer func(body io.ReadCloser) {
		err = body.Close()
		if err != nil {
			h.logger.Error(err)
		}
	}(response.GetBody())

	var bodyBytes []byte
	bodyBytes, err = io.ReadAll(response.GetBody())
	if err != nil {
		h.logger.Error(err)
		return
	}
	response.SetBodyBytes(bodyBytes)

	return
}

func (h *HttpClient) FromCrawler(crawler pkg.Crawler) pkg.HttpClient {
	if h == nil {
		return new(HttpClient).FromCrawler(crawler)
	}

	config := crawler.GetConfig()
	h.redirectMaxTimes = config.GetRedirectMaxTimes()

	h.client = http.DefaultClient
	h.client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		redirectMaxTimes := h.redirectMaxTimes

		ctx := req.Context()
		redirectMaxTimesVal := ctx.Value("redirect_max_times")
		if redirectMaxTimesVal != nil {
			redirectMaxTimes, _ = redirectMaxTimesVal.(uint8)
		}

		if uint8(len(via)) > redirectMaxTimes {
			return errors.New(fmt.Sprintf("stopped after %d redirects", redirectMaxTimes))
		}
		return nil
	}
	h.proxy = config.GetProxy()
	h.timeout = config.GetRequestTimeout()
	h.httpProto = config.GetHttpProto()
	h.logger = crawler.GetLogger()
	h.redirectMaxTimes = config.GetRedirectMaxTimes()
	h.retryMaxTimes = config.GetRetryMaxTimes()
	h.Ja3 = true

	return h
}
