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

type HttpClient struct {
	Ja3              bool
	client           *http.Client
	proxy            *url.URL
	timeout          time.Duration
	httpProto        string
	logger           pkg.Logger
	redirectMaxTimes uint8
	retryMaxTimes    uint8
}

func NewClientJa3(ctx context.Context, conn net.Conn, cfg *tls.Config, helloID *utls.ClientHelloID, helloSpec *utls.ClientHelloSpec) (net.Conn, error) {
	config := &utls.Config{
		Rand:                        cfg.Rand,
		Time:                        cfg.Time,
		VerifyPeerCertificate:       cfg.VerifyPeerCertificate,
		RootCAs:                     cfg.RootCAs,
		NextProtos:                  cfg.NextProtos,
		ServerName:                  cfg.ServerName,
		ClientCAs:                   cfg.ClientCAs,
		InsecureSkipVerify:          cfg.InsecureSkipVerify,
		CipherSuites:                cfg.CipherSuites,
		PreferServerCipherSuites:    cfg.PreferServerCipherSuites,
		SessionTicketsDisabled:      cfg.SessionTicketsDisabled,
		SessionTicketKey:            cfg.SessionTicketKey,
		MinVersion:                  cfg.MinVersion,
		MaxVersion:                  cfg.MaxVersion,
		DynamicRecordSizingDisabled: cfg.DynamicRecordSizingDisabled,
		KeyLogWriter:                cfg.KeyLogWriter,
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

var zeroDialer net.Dialer

func (h *HttpClient) DoRequest(ctx context.Context, request pkg.Request) (response pkg.Response, err error) {
	bs, _ := request.Marshal()
	h.logger.DebugF("request: %s", string(bs))

	if ctx == nil {
		ctx = context.Background()
	}

	timeout := h.timeout
	if request.Timeout() > 0 {
		timeout = request.Timeout()
	}

	if timeout > 0 {
		c := context.Background()
		if meta, ok := ctx.Value("meta").(pkg.Meta); ok {
			c = context.WithValue(c, "meta", meta)
		}
		c, cancel := context.WithTimeout(c, timeout)
		defer cancel()
		request.WithContext(c)
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
	if request.ProxyEnable() != nil {
		proxyEnable = *request.ProxyEnable()
	}
	if proxyEnable {
		proxy := h.proxy
		if request.Proxy() != nil {
			proxy = request.Proxy()
		}
		if proxy == nil {
			err = errors.New("nil proxy")
			return
		}
		transport.Proxy = http.ProxyURL(proxy)
	}

	httpProto := h.httpProto
	if request.HttpProto() != "" {
		httpProto = request.HttpProto()
	}
	if httpProto != "2.0" {
		request.GetRequest().Proto = "HTTP/1.1"
		request.GetRequest().ProtoMajor = 1
		request.GetRequest().ProtoMinor = 1
		transport.ForceAttemptHTTP2 = false
	} else {
		request.GetRequest().Proto = "HTTP/2.0"
		request.GetRequest().ProtoMajor = 2
		request.GetRequest().ProtoMinor = 0
		transport.ForceAttemptHTTP2 = true
	}

	if requiresHTTP1(request.GetRequest()) {
		transport.TLSClientConfig.NextProtos = nil
	}

	var resp *http.Response
	if h.Ja3 {
		transport.DialTLSContext = func(request pkg.Request) func(ctx context.Context, network, addr string) (net.Conn, error) {
			return func(ctx context.Context, network, addr string) (net.Conn, error) {
				var firstTLSHost string
				if firstTLSHost, _, err = net.SplitHostPort(addr); err != nil {
					h.logger.Error(err)
					return nil, err
				}

				// Initiate TLS and check remote host name against certificate.
				cfg := cloneTLSConfig(transport.TLSClientConfig)
				if cfg.ServerName == "" {
					cfg.ServerName = firstTLSHost
				}
				if request.HttpProto() == "2.0" {
					cfg.NextProtos = []string{"h2", "http/1.1"}
				} else {
					cfg.NextProtos = []string{"http/1.1"}
				}

				plainConn, err := zeroDialer.DialContext(ctx, "tcp", addr)
				if err != nil {
					h.logger.Error(err)
					return nil, err
				}

				var helloID *utls.ClientHelloID
				var helloSpec *utls.ClientHelloSpec

				switch pkg.Browser(request.Fingerprint()) {
				case pkg.Chrome:
					helloID = &utls.HelloChrome_Auto
				case pkg.Edge:
					helloID = &utls.HelloEdge_Auto
				case pkg.Safari:
					helloID = &utls.HelloSafari_Auto
				case pkg.FireFox:
					helloID = &utls.HelloFirefox_Auto
				default:
					if request.Fingerprint() != "" {
						helloSpec, err = stringToSpec(request.Fingerprint())
						if err != nil {
							h.logger.Error(err)
							helloID = &utls.HelloChrome_Auto
						}
					}
				}

				tlsConn, err := NewClientJa3(ctx, plainConn, cfg, helloID, helloSpec)
				if err != nil {
					h.logger.Error(err)
					return nil, err
				}

				return tlsConn, nil
			}
		}(request)
	}

	client := h.client

	redirectMaxTimes := h.redirectMaxTimes
	if request.RedirectMaxTimes() != nil {
		redirectMaxTimes = *request.RedirectMaxTimes()
	}
	client.CheckRedirect = func(redirectMaxTimes uint8) func(req *http.Request, via []*http.Request) error {
		return func(req *http.Request, via []*http.Request) error {
			if uint8(len(via)) > redirectMaxTimes {
				return errors.New(fmt.Sprintf("stopped after %d redirects", redirectMaxTimes))
			}
			return nil
		}
	}(redirectMaxTimes)

	client.Transport = transport

	if timeout > 0 {
		client.Timeout = timeout
	}
	response = new(response2.Response).SetRequest(request)
	begin := time.Now()
	resp, err = client.Do(request.GetRequest())
	if err != nil {
		h.logger.Error(err)
		return
	}

	response.SetResponse(resp)
	response.SetSpendTime(time.Now().Sub(begin))

	if err != nil {
		retryMaxTimes := h.retryMaxTimes
		if request.RetryMaxTimes() != nil {
			retryMaxTimes = *request.RetryMaxTimes()
		}
		if request.RetryTimes() < retryMaxTimes {
			return
		}
		h.logger.Error(err, "RetryTimes:", request.RetryTimes())
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
func (h *HttpClient) Close(_ context.Context) (err error) {
	return
}
func (h *HttpClient) FromSpider(spider pkg.Spider) pkg.HttpClient {
	if h == nil {
		return new(HttpClient).FromSpider(spider)
	}

	config := spider.GetCrawler().GetConfig()
	h.redirectMaxTimes = config.GetRedirectMaxTimes()

	h.client = http.DefaultClient
	h.proxy = config.GetProxy()
	h.timeout = config.GetRequestTimeout()
	h.httpProto = config.GetHttpProto()
	h.logger = spider.GetLogger()
	h.redirectMaxTimes = config.GetRedirectMaxTimes()
	h.retryMaxTimes = config.GetRetryMaxTimes()
	h.Ja3 = config.GetEnableJa3()

	return h
}
