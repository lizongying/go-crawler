package httpClient

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/lizongying/go-crawler/pkg/logger"
	"github.com/lizongying/go-crawler/pkg/utils"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const defaultTimeout = time.Minute
const defaultHttpProto = "2.0"

type HttpClient struct {
	client      *http.Client
	proxy       *url.URL
	timeout     time.Duration
	httpProto   string
	logger      pkg.Logger
	middlewares map[int]pkg.Middleware
}

func (h *HttpClient) BuildRequest(ctx context.Context, request *pkg.Request) (err error) {
	h.logger.DebugF("request: %+v", request)

	if ctx == nil {
		ctx = context.Background()
	}

	if request.Method == "" {
		request.Method = "GET"
	}
	request.CreateTime = utils.NowStr()
	request.Checksum = utils.StrMd5(request.Method, request.Url, request.BodyStr)
	if request.CanonicalHeaderKey {
		headers := make(map[string][]string)
		for k, v := range request.Header {
			headers[http.CanonicalHeaderKey(k)] = v
		}
		request.Header = headers
	}

	if request.Request == nil {
		Url, e := url.Parse(request.Url)
		if e != nil {
			err = e
			h.logger.Error(err)
			return
		}

		var body io.Reader
		if request.BodyStr != "" {
			body = strings.NewReader(request.BodyStr)
		}

		request.Request, err = http.NewRequest(request.Method, Url.String(), body)
		if err != nil {
			h.logger.Error(err)
			return
		}

		request.Request.Header = request.Header
	}

	return
}

func (h *HttpClient) BuildResponse(ctx context.Context, request *pkg.Request) (response *pkg.Response, err error) {
	h.logger.DebugF("request: %+v", request)

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

	response = &pkg.Response{
		Request: request,
	}

	begin := time.Now()
	response.Response, err = client.Do(request.Request)
	response.Request.SpendTime = time.Now().Sub(begin)
	if err != nil {
		if request.RetryTimes < request.RetryMaxTimes {
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

func NewHttpClient(config *config.Config, logger *logger.Logger) (httpClient pkg.HttpClient, err error) {
	proxyExample := config.Proxy.Example
	var proxy *url.URL
	if proxyExample != "" {
		proxy, err = url.Parse(proxyExample)
		if err != nil {
			logger.Error(err)
			return
		}
	}

	timeout := defaultTimeout
	if config.Request.Timeout > 0 {
		timeout = time.Second * time.Duration(config.Request.Timeout)
	}

	httpProto := defaultHttpProto
	if config.Request.HttpProto != "" {
		httpProto = config.Request.HttpProto
	}

	httpClient = &HttpClient{
		client:    http.DefaultClient,
		proxy:     proxy,
		timeout:   timeout,
		httpProto: httpProto,
		logger:    logger,
	}

	return
}
