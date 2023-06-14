package middlewares

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
)

type DumpMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
}

func (m *DumpMiddleware) SpiderStart(_ context.Context, spider pkg.Spider) (err error) {
	m.logger = spider.GetLogger()
	return
}

func (m *DumpMiddleware) ProcessRequest(_ context.Context, request *pkg.Request) (err error) {
	m.logger.InfoF("request: %+v", *request)
	return
}

func (m *DumpMiddleware) FromCrawler(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(DumpMiddleware).FromCrawler(spider)
	}

	m.logger = spider.GetLogger()
	return m
}
