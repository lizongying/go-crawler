package httpClient

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/lizongying/go-crawler/pkg/logger"
	"github.com/lizongying/go-crawler/pkg/utils"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const defaultTimeout = time.Minute

type HttpClient struct {
	Client          *http.Client
	ClientWithProxy *http.Client
	Proxy           *url.URL
	Timeout         time.Duration
	logger          *logger.Logger
}

func (h *HttpClient) getClient() {
	h.Client = &http.Client{
		Timeout: h.Timeout,
	}

	tr := &http.Transport{
		Proxy: http.ProxyURL(h.Proxy),
	}
	h.ClientWithProxy = &http.Client{
		Transport: tr,
		Timeout:   h.Timeout,
	}

	return
}

func (h *HttpClient) BuildRequest(ctx context.Context, request *pkg.Request) (err error) {
	h.logger.Debug("request", utils.JsonStr(request))

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
	h.logger.Debug("request", utils.JsonStr(request))

	if ctx == nil {
		ctx = context.Background()
	}

	if request.TimeoutRequest > 0 {
		c, cancel := context.WithTimeout(ctx, request.TimeoutRequest)
		defer cancel()
		request.Request = request.Request.WithContext(c)
	}

	client := h.Client
	if request.ProxyEnable {
		client = h.ClientWithProxy
	}

	resp, err := client.Do(request.Request)
	if err != nil {
		h.logger.Error(err)
		h.logger.Debug(utils.Request2Curl(request))
		return
	}

	response = &pkg.Response{
		Response: resp,
		Request:  request,
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	response.BodyBytes, err = io.ReadAll(response.Body)
	if err != nil {
		h.logger.Error(err)
		return
	}

	return
}

func NewHttpClient(config *config.Config, logger *logger.Logger) (httpClient *HttpClient, err error) {
	proxyExample := config.Proxy.Example
	if proxyExample == "" {
		err = errors.New("proxy is empty")
		logger.Error(err)
		return
	}

	proxy, err := url.Parse(proxyExample)
	if err != nil {
		logger.Error(err)
		return
	}

	httpClient = &HttpClient{
		Proxy:   proxy,
		Timeout: defaultTimeout,
		logger:  logger,
	}

	httpClient.getClient()

	return
}
