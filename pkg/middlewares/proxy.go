package middlewares

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
)

type ProxyMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
}

func (m *ProxyMiddleware) ProcessRequest(_ context.Context, request *pkg.Request) (err error) {

	return
}

func (m *ProxyMiddleware) FromCrawler(crawler pkg.Crawler) pkg.Middleware {
	if m == nil {
		return new(ProxyMiddleware).FromCrawler(crawler)
	}

	m.logger = crawler.GetLogger()
	return m
}
