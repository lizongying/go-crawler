package middlewares

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
)

type RedirectMiddleware struct {
	pkg.UnimplementedMiddleware
	logger           pkg.Logger
	redirectMaxTimes uint8
}

func (m *RedirectMiddleware) ProcessRequest(_ pkg.Context, request pkg.Request) (err error) {
	redirectMaxTimes := m.redirectMaxTimes
	if request.GetRedirectMaxTimes() != nil {
		redirectMaxTimes = *request.GetRedirectMaxTimes()
	}
	if redirectMaxTimes > 0 {
		ctx := context.WithValue(request.RequestContext(), "redirect_max_times", redirectMaxTimes)
		request.WithContext(ctx)
	}
	return
}

func (m *RedirectMiddleware) FromSpider(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(RedirectMiddleware).FromSpider(spider)
	}

	m.UnimplementedMiddleware.FromSpider(spider)
	crawler := spider.GetCrawler()
	m.logger = spider.GetLogger()
	m.redirectMaxTimes = crawler.GetConfig().GetRedirectMaxTimes()
	return m
}
