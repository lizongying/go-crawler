package middlewares

import (
	"github.com/lizongying/go-crawler/pkg"
)

type ReferrerMiddleware struct {
	pkg.UnimplementedMiddleware
	logger         pkg.Logger
	referrerPolicy pkg.ReferrerPolicy
}

func (m *ReferrerMiddleware) ProcessRequest(_ pkg.Context, request pkg.Request) (err error) {
	if m.referrerPolicy == pkg.NoReferrerPolicy && request.GetHeaders() != nil && request.GetHeader("Referer") != "" {
		//request.Header.Del("Referer")
	}

	if m.referrerPolicy == pkg.DefaultReferrerPolicy && request.GetHeaders() != nil && request.GetReferrer() != "" {
		request.SetHeader("Referer", request.GetReferrer())
	}

	return
}

func (m *ReferrerMiddleware) ProcessResponse(ctx pkg.Context, response pkg.Response) (err error) {
	// add referrer to context
	if response.GetURL() != nil {
		ctx.Meta.Referrer = response.GetURL()
	}

	return
}

func (m *ReferrerMiddleware) FromSpider(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(ReferrerMiddleware).FromSpider(spider)
	}

	m.UnimplementedMiddleware.FromSpider(spider)
	crawler := spider.GetCrawler()
	m.logger = spider.GetLogger()
	m.referrerPolicy = crawler.GetConfig().GetReferrerPolicy()
	return m
}
