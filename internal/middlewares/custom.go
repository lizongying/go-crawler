package middlewares

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
)

type CustomMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger

	spider pkg.Spider
}

func (m *CustomMiddleware) GetName() string {
	return "custom"
}

func (m *CustomMiddleware) SpiderStart(_ context.Context, spider pkg.Spider) (err error) {
	m.spider = spider
	m.logger.Debug("start")
	return
}

func (m *CustomMiddleware) ProcessRequest(_ context.Context, request *pkg.Request) (err error) {
	m.logger.Debug("request", request)
	return
}

func (m *CustomMiddleware) ProcessResponse(_ context.Context, response *pkg.Response) (err error) {
	m.logger.Debug("response", response)
	return
}

func (m *CustomMiddleware) ProcessItem(c *pkg.Context) (err error) {
	if err = c.NextItem(); err != nil {
		m.logger.Error(err)
	}
	return
}

func (m *CustomMiddleware) SpiderStop(_ context.Context) (err error) {
	m.logger.Debug("stop")
	return
}

func (m *CustomMiddleware) FromCrawler(spider pkg.Spider) pkg.Middleware {
	m.logger = spider.GetLogger()
	return m
}

func NewCustomMiddleware() pkg.Middleware {
	return &CustomMiddleware{}
}
