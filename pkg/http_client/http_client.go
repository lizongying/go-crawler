package http_client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/dns_cache"
	response2 "github.com/lizongying/go-crawler/pkg/response"
	"github.com/lizongying/go-crawler/pkg/utils"
	"github.com/lizongying/go-crawler/static"
	utls "github.com/refraction-networking/utls"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
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
	dnsCache         *dns_cache.DnsCache
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
	h.logger.Debugf("request: %s", string(bs))

	if ctx == nil {
		ctx = context.Background()
	}

	timeout := h.timeout
	if request.GetTimeout() > 0 {
		timeout = request.GetTimeout()
	}

	if timeout > 0 {
		c := context.Background()
		c, cancel := context.WithTimeout(c, timeout)
		defer cancel()
		request.WithRequestContext(c)
	}

	// Get a copy of the default root CAs
	defaultCAs, err := x509.SystemCertPool()
	if err != nil {
		defaultCAs = x509.NewCertPool()
	}
	defaultCAs.AppendCertsFromPEM(static.CaCert)

	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 60 * time.Second,
			Resolver: func() *net.Resolver {
				resolver := net.DefaultResolver
				resolver.Dial = func(ctx context.Context, network, address string) (net.Conn, error) {
					var host string
					var port string
					if host, port, err = net.SplitHostPort(address); err != nil {
						return nil, err
					}
					var d net.Dialer
					var c net.Conn
					var e error
					if ip, ok := h.dnsCache.Get(host); ok {
						if strings.Contains(ip.String(), ".") {
							c, e = d.DialContext(ctx, network, fmt.Sprintf("%s:%s", ip.String(), port))
						} else {
							c, e = d.DialContext(ctx, network, fmt.Sprintf("[%s]:%s", ip.String(), port))
						}
					} else {
						c, e = d.DialContext(ctx, network, address)
					}

					h.logger.Info("RemoteAddr", c.RemoteAddr(), network, address)
					return c, e
				}
				return resolver
			}(),
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
	if request.IsProxyEnable() != nil {
		proxyEnable = *request.IsProxyEnable()
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
		request.GetHttpRequest().Proto = "HTTP/1.1"
		request.GetHttpRequest().ProtoMajor = 1
		request.GetHttpRequest().ProtoMinor = 1
		transport.ForceAttemptHTTP2 = false
	} else {
		request.GetHttpRequest().Proto = "HTTP/2.0"
		request.GetHttpRequest().ProtoMajor = 2
		request.GetHttpRequest().ProtoMinor = 0
		transport.ForceAttemptHTTP2 = true
	}

	if requiresHTTP1(request.GetHttpRequest()) {
		transport.TLSClientConfig.NextProtos = nil
	}

	var resp *http.Response
	if h.Ja3 {
		transport.DialTLSContext = func(request pkg.Request) func(ctx context.Context, network, addr string) (net.Conn, error) {
			return func(ctx context.Context, network, addr string) (net.Conn, error) {
				var firstTLSHost string
				var port string
				if firstTLSHost, port, err = net.SplitHostPort(addr); err != nil {
					h.logger.Error(err)
					return nil, err
				}

				// Initiate TLS and check remote host name against certificate.
				cfg := cloneTLSConfig(transport.TLSClientConfig)
				if cfg.ServerName == "" {
					cfg.ServerName = firstTLSHost
				}
				if request.GetHttpProto() == "2.0" {
					cfg.NextProtos = []string{"h2", "http/1.1"}
				} else {
					cfg.NextProtos = []string{"http/1.1"}
				}

				var plainConn net.Conn
				if ip, ok := h.dnsCache.Get(firstTLSHost); ok {
					if strings.Contains(ip.String(), ".") {
						plainConn, err = zeroDialer.DialContext(ctx, network, fmt.Sprintf("%s:%s", ip.String(), port))
					} else {
						plainConn, err = zeroDialer.DialContext(ctx, network, fmt.Sprintf("[%s]:%s", ip.String(), port))
					}
				} else {
					plainConn, err = zeroDialer.DialContext(ctx, network, addr)
					if err != nil {
						h.logger.Error(err)
						return nil, err
					}
					h.dnsCache.ResolveWithRetry(firstTLSHost)
				}
				if err != nil {
					h.logger.Error(err)
					return nil, err
				}

				var helloID *utls.ClientHelloID
				var helloSpec *utls.ClientHelloSpec

				switch pkg.Browser(request.GetFingerprint()) {
				case pkg.BrowserChrome:
					helloID = &utls.HelloChrome_Auto
				case pkg.BrowserEdge:
					helloID = &utls.HelloEdge_Auto
				case pkg.BrowserSafari:
					helloID = &utls.HelloSafari_Auto
				case pkg.BrowserFireFox:
					helloID = &utls.HelloFirefox_Auto
				default:
					if request.GetFingerprint() != "" {
						helloSpec, err = stringToSpec(request.GetFingerprint())
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
	if request.GetRedirectMaxTimes() != nil {
		redirectMaxTimes = *request.GetRedirectMaxTimes()
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
	resp, err = client.Do(request.GetHttpRequest())
	if err != nil {
		h.logger.Error(err)
		return
	}

	response.SetResponse(resp)
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
		h.logger.Errorf("request: %+v", request)
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
	h.dnsCache = dns_cache.NewDnsCache(time.Hour*24, 3)

	return h
}
