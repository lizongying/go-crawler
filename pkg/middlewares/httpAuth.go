package middlewares

import (
	"github.com/lizongying/go-crawler/pkg"
)

type HttpAuthMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
}

func (m *HttpAuthMiddleware) ProcessRequest(ctx pkg.Context, request pkg.Request) (err error) {
	spider := m.GetSpider()
	username := spider.GetUsername()
	if request.GetUsername() != "" {
		username = request.GetUsername()
	}
	password := spider.GetPassword()
	if request.GetPassword() != "" {
		password = request.GetPassword()
	}

	if username != "" && password != "" {
		m.logger.InfoF("BasicAuth %s:%s", password, username)
		request.SetBasicAuth(username, password)
	}

	return
}

func (m *HttpAuthMiddleware) FromSpider(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(HttpAuthMiddleware).FromSpider(spider)
	}

	m.UnimplementedMiddleware.FromSpider(spider)
	m.logger = spider.GetLogger()
	return m
}
