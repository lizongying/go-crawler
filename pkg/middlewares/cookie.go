package middlewares

import (
	"github.com/lizongying/go-crawler/pkg"
)

type CookieMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
}

func (m *CookieMiddleware) ProcessResponse(ctx pkg.Context, response pkg.Response) (err error) {
	// add cookies to context
	cookies := response.GetCookies()
	if len(cookies) > 0 {
		ctx.Meta.Cookies = cookies
	}

	return
}

func (m *CookieMiddleware) FromSpider(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(CookieMiddleware).FromSpider(spider)
	}

	m.UnimplementedMiddleware.FromSpider(spider)
	m.logger = spider.GetLogger()
	return m
}
