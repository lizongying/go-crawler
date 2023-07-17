package middlewares

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
)

type ReferrerMiddleware struct {
	pkg.UnimplementedMiddleware
	logger         pkg.Logger
	referrerPolicy pkg.ReferrerPolicy
}

func (m *ReferrerMiddleware) ProcessRequest(_ context.Context, request pkg.Request) (err error) {
	if m.referrerPolicy == pkg.NoReferrerPolicy && request.GetHeaders() != nil && request.GetHeader("Referer") != "" {
		//request.Header.Del("Referer")
	}

	if m.referrerPolicy == pkg.DefaultReferrerPolicy && request.GetHeaders() != nil && request.GetReferrer() != "" {
		request.SetHeader("Referer", request.GetReferrer())
	}

	return
}

func (m *ReferrerMiddleware) ProcessResponse(_ context.Context, response pkg.Response) (err error) {
	// add referrer to context
	if response.GetUrl() != "" {
		ctx := context.WithValue(response.Context(), "referrer", response.GetUrl())
		response.WithContext(ctx)
	}

	return
}

func (m *ReferrerMiddleware) FromCrawler(crawler pkg.Crawler) pkg.Middleware {
	if m == nil {
		return new(ReferrerMiddleware).FromCrawler(crawler)
	}

	m.logger = crawler.GetLogger()
	m.referrerPolicy = crawler.GetConfig().GetReferrerPolicy()
	return m
}
