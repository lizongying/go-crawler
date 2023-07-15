package middlewares

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
)

type HttpAuthMiddleware struct {
	pkg.UnimplementedMiddleware
	username string
	password string
	logger   pkg.Logger
}

func (m *HttpAuthMiddleware) ProcessRequest(_ context.Context, request *pkg.Request) (err error) {
	username := m.username
	if request.GetUsername() != "" {
		username = request.GetUsername()
	}
	password := m.password
	if request.GetPassword() != "" {
		password = request.GetPassword()
	}

	if username != "" && password != "" {
		m.logger.InfoF("BasicAuth %s:%s", password, username)
		request.SetBasicAuth(username, password)
	}

	return
}

func (m *HttpAuthMiddleware) FromCrawler(crawler pkg.Crawler) pkg.Middleware {
	if m == nil {
		return new(HttpAuthMiddleware).FromCrawler(crawler)
	}

	m.username = crawler.GetUsername()
	m.password = crawler.GetPassword()
	m.logger = crawler.GetLogger()
	return m
}
