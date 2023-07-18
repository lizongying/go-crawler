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

func (m *RedirectMiddleware) ProcessRequest(_ context.Context, request pkg.Request) (err error) {
	redirectMaxTimes := m.redirectMaxTimes
	if request.GetRedirectMaxTimes() != nil {
		redirectMaxTimes = *request.GetRedirectMaxTimes()
	}
	if redirectMaxTimes > 0 {
		ctx := context.WithValue(request.Context(), "redirect_max_times", redirectMaxTimes)
		request.WithContext(ctx)
	}
	return
}

func (m *RedirectMiddleware) FromCrawler(crawler pkg.Crawler) pkg.Middleware {
	if m == nil {
		return new(RedirectMiddleware).FromCrawler(crawler)
	}

	m.logger = crawler.GetLogger()
	m.redirectMaxTimes = crawler.GetConfig().GetRedirectMaxTimes()
	return m
}
