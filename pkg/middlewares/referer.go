package middlewares

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
)

type RefererMiddleware struct {
	pkg.UnimplementedMiddleware
	logger        pkg.Logger
	refererPolicy pkg.ReferrerPolicy
}

func (m *RefererMiddleware) ProcessRequest(_ context.Context, request *pkg.Request) (err error) {
	if m.refererPolicy == pkg.NoReferrerPolicy && request.Header != nil && request.Header.Get("Referer") != "" {
		//request.Header.Del("Referer")
	}

	if m.refererPolicy == pkg.DefaultReferrerPolicy && request.Header != nil && request.Referer != "" {
		request.SetHeader("Referer", request.Referer)
	}

	return
}

func (m *RefererMiddleware) ProcessResponse(_ context.Context, response *pkg.Response) (err error) {
	// add referer to context
	if response.Request.Url != "" {
		ctx := context.WithValue(response.Request.Context(), "referer", response.Request.Url)
		response.Request.Request = response.Request.WithContext(ctx)
	}

	return
}

func (m *RefererMiddleware) FromCrawler(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(RefererMiddleware).FromCrawler(spider)
	}

	m.logger = spider.GetLogger()
	m.refererPolicy = spider.GetConfig().GetReferrerPolicy()
	return m
}
