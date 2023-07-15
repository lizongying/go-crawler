package middlewares

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
)

type CustomMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
}

func (m *CustomMiddleware) GetName() string {
	return "custom"
}

func (m *CustomMiddleware) Start(_ context.Context, crawler pkg.Crawler) error {
	_ = m.FromCrawler(crawler)
	m.logger.Debug("start")
	return nil
}

func (m *CustomMiddleware) ProcessRequest(_ context.Context, request pkg.Request) error {
	m.logger.Debug("request", request)
	return nil
}

func (m *CustomMiddleware) ProcessResponse(_ context.Context, response *pkg.Response) error {
	m.logger.Debug("response", response)
	return nil
}

func (m *CustomMiddleware) Stop(_ context.Context) error {
	m.logger.Debug("stop")
	return nil
}

func (m *CustomMiddleware) FromCrawler(crawler pkg.Crawler) pkg.Middleware {
	m.logger = crawler.GetLogger()
	return m
}
