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

func (m *RefererMiddleware) ProcessRequest(_ context.Context, request pkg.Request) (err error) {
	if m.refererPolicy == pkg.NoReferrerPolicy && request.GetHeaders() != nil && request.GetHeader("Referer") != "" {
		//request.Header.Del("Referer")
	}

	if m.refererPolicy == pkg.DefaultReferrerPolicy && request.GetHeaders() != nil && request.GetReferer() != "" {
		request.SetHeader("Referer", request.GetReferer())
	}

	return
}

func (m *RefererMiddleware) ProcessResponse(_ context.Context, response pkg.Response) (err error) {
	// add referer to context
	if response.GetUrl() != "" {
		ctx := context.WithValue(response.Context(), "referer", response.GetUrl())
		response.WithContext(ctx)
	}

	return
}

func (m *RefererMiddleware) FromCrawler(crawler pkg.Crawler) pkg.Middleware {
	if m == nil {
		return new(RefererMiddleware).FromCrawler(crawler)
	}

	m.logger = crawler.GetLogger()
	m.refererPolicy = crawler.GetConfig().GetReferrerPolicy()
	return m
}
