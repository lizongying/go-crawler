package middlewares

import (
	"github.com/lizongying/go-crawler/pkg"
)

type ProxyMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
}

func (m *ProxyMiddleware) ProcessRequest(_ pkg.Context, request pkg.Request) (err error) {

	return
}

func (m *ProxyMiddleware) FromSpider(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(ProxyMiddleware).FromSpider(spider)
	}

	m.UnimplementedMiddleware.FromSpider(spider)
	m.logger = spider.GetLogger()
	return m
}
