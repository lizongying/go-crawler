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

func (m *HttpAuthMiddleware) ProcessRequest(_ context.Context, request *pkg.Request) (err error) {
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
	m.logger.InfoF("BasicAuth %s:%s", password, username)

	return
}

func (m *HttpAuthMiddleware) FromCrawler(crawler pkg.Crawler) pkg.Middleware {
	if m == nil {
		return new(HttpAuthMiddleware).FromCrawler(crawler)
	}

	m.logger = crawler.GetLogger()
	info := crawler.GetInfo()
	m.username = info.Username
	m.password = info.Password
	return m
}
