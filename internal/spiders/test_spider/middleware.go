package test_spider

import (
	"github.com/lizongying/go-crawler/pkg"
)

type Middleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
}

func (m *Middleware) ProcessRequest(_ pkg.Context, request pkg.Request) (err error) {
	request.SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	return
}

func (m *Middleware) FromSpider(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(Middleware).FromSpider(spider)
	}

	m.UnimplementedMiddleware.FromSpider(spider)
	m.logger = spider.GetLogger()
	return m
}
