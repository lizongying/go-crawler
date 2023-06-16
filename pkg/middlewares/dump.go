package middlewares

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
)

type DumpMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
}

func (m *DumpMiddleware) ProcessRequest(_ context.Context, request *pkg.Request) (err error) {
	m.logger.InfoF("request: %+v", *request)
	return
}

func (m *DumpMiddleware) FromCrawler(crawler pkg.Crawler) pkg.Middleware {
	if m == nil {
		return new(DumpMiddleware).FromCrawler(crawler)
	}

	m.logger = crawler.GetLogger()
	return m
}
