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
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

type HttpClient struct {
	client           *http.Client
	proxy            *url.URL
	timeout          time.Duration
	httpProto        string
	logger           pkg.Logger
	redirectMaxTimes uint8
	retryMaxTimes    uint8
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
	if request.ProxyEnable {
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

	begin := time.Now()
	response = &pkg.Response{
		Request: request,
	}
	response.Response, err = client.Do(request.Request)
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
	h.timeout = config.GetTimeout()
	h.httpProto = config.GetHttpProto()
	h.logger = crawler.GetLogger()
	h.redirectMaxTimes = config.GetRedirectMaxTimes()
	h.retryMaxTimes = config.GetRetryMaxTimes()

	return h
}
