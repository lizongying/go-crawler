package middlewares

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
)

type CookieMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
}

func (m *CookieMiddleware) ProcessResponse(_ context.Context, response pkg.Response) (err error) {
	// add cookies to context
	cookies := response.GetCookies()
	if len(cookies) > 0 {
		ctx := context.WithValue(response.Context(), "cookies", cookies)
		response.GetRequest().WithContext(ctx)
	}

	return
}

func (m *CookieMiddleware) FromCrawler(crawler pkg.Crawler) pkg.Middleware {
	if m == nil {
		return new(CookieMiddleware).FromCrawler(crawler)
	}

	m.logger = crawler.GetLogger()
	return m
}
