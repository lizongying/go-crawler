package middlewares

import (
	"github.com/lizongying/go-crawler/pkg"
)

type CookieMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
}

func (m *CookieMiddleware) ProcessResponse(ctx pkg.Context, response pkg.Response) (err error) {
	if response.GetResponse() == nil {
		m.logger.Debug("response nil")
		return
	}

	// add cookies to context
	cookies := response.Cookies()
	if len(cookies) > 0 {
		request := ctx.GetRequest()
		if request.GetCookies() == nil {
			request.WithCookies(make(map[string]string))
		}
		for _, v := range cookies {
			request.GetCookies()[v.Name] = v.Value
		}
		ctx.WithRequest(request)
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
