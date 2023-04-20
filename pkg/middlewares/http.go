package middlewares

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/httpClient"
	"github.com/lizongying/go-crawler/pkg/logger"
	"github.com/lizongying/go-crawler/pkg/utils"
)

type HttpMiddleware struct {
	pkg.UnimplementedMiddleware
	httpClient *httpClient.HttpClient
	logger     *logger.Logger
	spider     pkg.Spider
}

func (m *HttpMiddleware) GetName() string {
	return "http"
}

func (m *HttpMiddleware) SpiderStart(_ context.Context, spider pkg.Spider) (err error) {
	m.spider = spider
	return
}

func (m *HttpMiddleware) ProcessRequest(ctx context.Context, r *pkg.Request) (request *pkg.Request, response *pkg.Response, err error) {
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

func (m *HttpMiddleware) ProcessItem(_ context.Context, _ *pkg.Item) (err error) {
	return
}

func (m *HttpMiddleware) SpiderStop(_ context.Context) (err error) {
	return
}

func NewHttpMiddleware(logger *logger.Logger, httpClient *httpClient.HttpClient) (m pkg.Middleware) {
	m = &HttpMiddleware{
		httpClient: httpClient,
		logger:     logger,
	}
	return
}
