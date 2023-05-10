package middlewares

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/httpClient"
	"github.com/lizongying/go-crawler/pkg/logger"
)

type HttpMiddleware struct {
	pkg.UnimplementedMiddleware
	httpClient *httpClient.HttpClient
	logger     *logger.Logger
	spider     pkg.Spider
	stats      pkg.Stats
}

func (m *HttpMiddleware) SpiderStart(_ context.Context, spider pkg.Spider) (err error) {
	m.spider = spider
	m.stats = spider.GetStats()
	return
}

func (m *HttpMiddleware) ProcessRequest(c *pkg.Context) (err error) {
	request := c.Request
	m.logger.DebugF("request: %+v", request)

	ctx := context.Background()

	err = m.httpClient.BuildRequest(ctx, request)
	if err != nil {
		m.logger.Error(err)
		m.stats.IncRequestError()
		return
	}

	ok := m.spider.IsAllowedDomain(request.URL)
	if !ok {
		err = errors.New("it's not a allowed domain")
		m.logger.Error(err)
		m.stats.IncRequestError()
		return
	}

	c.Response, err = m.httpClient.BuildResponse(ctx, request)
	if err != nil {
		if request.RetryMaxTimes > 0 && request.RetryTimes < request.RetryMaxTimes {
			err = c.FirstResponse()
			return
		}
		m.logger.Error(err, "RetryTimes:", request.RetryTimes, request.RetryMaxTimes)
		m.stats.IncRequestError()
		return
	}

	m.stats.IncRequestSuccess()
	err = c.FirstResponse()
	return
}

func NewHttpMiddleware(logger *logger.Logger, httpClient *httpClient.HttpClient) (m pkg.Middleware) {
	m = &HttpMiddleware{
		httpClient: httpClient,
		logger:     logger,
	}
	return
}
