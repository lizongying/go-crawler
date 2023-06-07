package middlewares

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
)

type HttpAuthMiddleware struct {
	pkg.UnimplementedMiddleware
	logger   pkg.Logger
	username string
	password string
}

func (m *HttpAuthMiddleware) SpiderStart(_ context.Context, spider pkg.Spider) (err error) {
	info := spider.GetInfo()
	m.username = info.Username
	m.password = info.Password
	return
}

func (m *HttpAuthMiddleware) ProcessRequest(c *pkg.Context) (err error) {
	m.logger.Debug("enter ProcessRequest")
	defer func() {
		m.logger.Debug("exit ProcessRequest")
	}()

	request := c.Request

	username := m.username
	if request.Username != "" {
		username = request.Username
	}
	password := m.password
	if request.Password != "" {
		password = request.Password
	}

	if username != "" && password != "" {
		request.SetBasicAuth(username, password)
	}
	m.logger.Info("header", request.Header)

	err = c.NextRequest()
	return
}

func (m *HttpAuthMiddleware) FromCrawler(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(HttpAuthMiddleware).FromCrawler(spider)
	}
	m.logger = spider.GetLogger()
	return m
}
