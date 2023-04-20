package middlewares

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/internal"
	"github.com/lizongying/go-crawler/internal/httpClient"
	"github.com/lizongying/go-crawler/internal/logger"
	"github.com/lizongying/go-crawler/internal/utils"
)

type HttpMiddleware struct {
	internal.UnimplementedMiddleware
	httpClient *httpClient.HttpClient
	logger     *logger.Logger
	spider     internal.Spider
}

func (m *HttpMiddleware) GetName() string {
	return "http"
}

func (m *HttpMiddleware) SpiderStart(_ context.Context, spider internal.Spider) (err error) {
	m.spider = spider
	return
}

func (m *HttpMiddleware) ProcessRequest(ctx context.Context, r *internal.Request) (request *internal.Request, response *internal.Response, err error) {
	m.logger.Debug("request", utils.JsonStr(r))

	if ctx == nil {
		ctx = context.Background()
	}

	err = m.httpClient.BuildRequest(ctx, r)
	if err != nil {
		m.logger.Error(err)
		return
	}

	ok := m.spider.IsAllowedDomain(r.URL)
	if !ok {
		err = errors.New("it's not a allowed domain")
		m.logger.Error(err)
		return
	}

	response, err = m.httpClient.BuildResponse(ctx, r)
	if err != nil {
		m.logger.Error(err)
		return
	}

	return
}

func (m *HttpMiddleware) ProcessItem(_ context.Context, _ *internal.Item) (err error) {
	return
}

func (m *HttpMiddleware) SpiderStop(_ context.Context) (err error) {
	return
}

func NewHttpMiddleware(logger *logger.Logger, httpClient *httpClient.HttpClient) (m internal.Middleware) {
	m = &HttpMiddleware{
		httpClient: httpClient,
		logger:     logger,
	}
	return
}
