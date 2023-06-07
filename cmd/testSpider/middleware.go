package main

import (
	"github.com/lizongying/go-crawler/pkg"
)

type Middleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
}

func (m *Middleware) ProcessRequest(c *pkg.Context) (err error) {
	err = c.NextRequest()
	if err != nil {
		m.logger.Error(err)
		return
	}

	r := c.Request
	m.logger.DebugF("request: %+v", r)

	r.SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	return
}

func (m *Middleware) FromCrawler(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(Middleware).FromCrawler(spider)
	}
	m.logger = spider.GetLogger()
	return m
}
