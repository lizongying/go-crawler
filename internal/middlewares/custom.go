package middlewares

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
)

type CustomMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
}

func (m *CustomMiddleware) Name() string {
	return "custom"
}

func (m *CustomMiddleware) Start(ctx context.Context, spider pkg.Spider) (err error) {
	err = m.UnimplementedMiddleware.Start(ctx, spider)
	m.logger.Debug("start")
	return nil
}

func (m *CustomMiddleware) ProcessRequest(_ pkg.Context, request pkg.Request) error {
	m.logger.Debug("request", request)
	return nil
}

func (m *CustomMiddleware) ProcessResponse(_ pkg.Context, response pkg.Response) error {
	m.logger.Debug("response", response)
	return nil
}

func (m *CustomMiddleware) Stop(_ context.Context) error {
	m.logger.Debug("stop")
	return nil
}

func (m *CustomMiddleware) FromSpider(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(CustomMiddleware).FromSpider(spider)
	}

	m.UnimplementedMiddleware.FromSpider(spider)
	m.logger = spider.GetLogger()
	return m
}
