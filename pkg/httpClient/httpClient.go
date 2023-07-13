package httpClient

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"github.com/lizongying/go-crawler/static"
	utls "github.com/refraction-networking/utls"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

type HttpClient struct {
	ClientOption
	client           *http.Client
	proxy            *url.URL
	timeout          time.Duration
	httpProto        string
	logger           pkg.Logger
	redirectMaxTimes uint8
	retryMaxTimes    uint8
}

func NewClientJa3(ctx context.Context, conn net.Conn, helloID *utls.ClientHelloID, helloSpec *utls.ClientHelloSpec, serverName string, http2 bool) (net.Conn, error) {
	config := &utls.Config{
		//InsecureSkipVerify: true,
		ServerName: serverName,
	}
	if http2 {
		config.NextProtos = []string{"h2", "http/1.1"}
	} else {
		config.NextProtos = []string{"http/1.1"}
	}

	if helloID == nil {
		helloID = &utls.HelloChrome_Auto
	}
	if helloSpec != nil {
		helloID = &utls.HelloCustom
	}
	c := utls.UClient(conn, config, *helloID)

	if *helloID == utls.HelloCustom {
		if err := c.ApplyPreset(helloSpec); err != nil {
			return nil, err
		}
	}

	if err := c.HandshakeContext(ctx); err != nil {
		return nil, err
	}

	return c, nil
}

func NewClient(conn net.Conn, serverName string, http2 bool) net.Conn {
	config := &tls.Config{
		//InsecureSkipVerify: true,
		ServerName: serverName,
	}
	if http2 {
		config.NextProtos = []string{"h2", "http/1.1"}
	} else {
		config.NextProtos = []string{"http/1.1"}
	}

	return tls.Client(conn, config)
}

type ClientOption struct {
	Ja3         bool
	HelloID     *utls.ClientHelloID
	HelloSpec   *utls.ClientHelloSpec
	DialTimeout *time.Duration
}

func (h *HttpClient) DoRequest(ctx context.Context, request *pkg.Request) (response *pkg.Response, err error) {
	h.logger.DebugF("request: %+v", *request)

	if ctx == nil {
		ctx = context.Background()
	}

	if request.Timeout > 0 {
		//c, cancel := context.WithTimeout(ctx, request.Timeout)
		//defer cancel()
		//request.Request = request.Request.WithContext(c)
	}

	timeout := h.timeout
	if request.Timeout > 0 {
		timeout = request.Timeout
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
	if request.GetProxyEnable() {
		proxy := h.proxy
		if request.Proxy != nil {
			proxy = request.Proxy
		}
		if proxy == nil {
			err = errors.New("nil proxy")
			return
		}
		transport.Proxy = http.ProxyURL(proxy)
	}

	httpProto := h.httpProto
	if request.HttpProto != "" {
		httpProto = request.HttpProto
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

	if request.Request == nil {
		err = errors.New("nil request")
		return
	}

	if h.DialTimeout == nil {
		dialTimeout := 10 * time.Second
		h.DialTimeout = &dialTimeout
	}

	network := request.URL.Scheme
	address := request.URL.Host

	// Check if the URL specifies a port, otherwise use the default port
	if request.URL.Port() == "" {
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

	h.Ja3 = true
	if h.Ja3 {
		conn, _ = NewClientJa3(ctx, conn, h.HelloID, h.HelloSpec, request.URL.Hostname(), h.httpProto == "2.0")
	} else {
		conn = NewClient(conn, request.URL.Hostname(), h.httpProto == "2.0")
	}

	begin := time.Now()
	response = &pkg.Response{
		Request: request,
	}

	if h.httpProto == "2.0" {
		request.Proto = "HTTP/2.0"
		request.ProtoMajor = 2
		request.ProtoMinor = 0

		tr := transport
		tr.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return conn, nil
		}
		tr.DialTLSContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return conn, nil
		}

		response.Response, err = tr.RoundTrip(request.Request)
	} else {
		request.Proto = "HTTP/1.1"
		request.ProtoMajor = 1
		request.ProtoMinor = 1

		tr := transport
		tr.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return conn, nil
		}
		tr.DialTLSContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
			return conn, nil
		}
		response.Response, err = tr.RoundTrip(request.Request)
	}

	response.Request.SpendTime = time.Now().Sub(begin)
	if err != nil {
		retryMaxTimes := h.retryMaxTimes
		if request.RetryMaxTimes != nil {
			retryMaxTimes = *request.RetryMaxTimes
		}
		if request.RetryTimes < retryMaxTimes {
			return
		}
		h.logger.Error(err, "RetryTimes:", request.RetryTimes)
		h.logger.ErrorF("request: %+v", request)
		h.logger.Debug(utils.Request2Curl(request))
		return
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	response.BodyBytes, err = io.ReadAll(response.Body)
	if err != nil {
		h.logger.Error(err)
		return
	}

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

	return h
}
