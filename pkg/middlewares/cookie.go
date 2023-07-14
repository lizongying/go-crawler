package middlewares

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
)

type CookieMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
}

func (m *CookieMiddleware) ProcessResponse(_ context.Context, response *pkg.Response) (err error) {
	// add cookies to context
	cookies := response.Cookies()
	if len(cookies) > 0 {
		ctx := context.WithValue(response.Request.Context(), "cookies", cookies)
		response.Request.Request = *response.Request.WithContext(ctx)
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
