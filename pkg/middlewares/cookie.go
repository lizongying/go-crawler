package middlewares

import (
	"github.com/lizongying/go-crawler/pkg"
)

type CookieMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger

	enableCookie bool
}

func (m *CookieMiddleware) ProcessRequest(c *pkg.Context) (err error) {
	m.logger.Debug("enter ProcessRequest")
	defer func() {
		m.logger.Debug("exit ProcessRequest")
	}()

	request := c.Request
	if m.enableCookie && len(request.Cookies) > 0 {
		for _, cookie := range request.Cookies {
			request.AddCookie(cookie)
		}
	}

	err = c.NextRequest()
	if err != nil {
		m.logger.Debug(err)
		return
	}

	return
}

func (m *CookieMiddleware) FromCrawler(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(CookieMiddleware).FromCrawler(spider)
	}
	m.logger = spider.GetLogger()
	m.enableCookie = spider.GetConfig().GetEnableCookie()
	return m
}
