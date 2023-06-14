package middlewares

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
)

type UrlMiddleware struct {
	pkg.UnimplementedMiddleware
	urlLengthLimit int
	logger         pkg.Logger
}

func (m *UrlMiddleware) ProcessRequest(_ context.Context, request *pkg.Request) (err error) {
	if m.urlLengthLimit < len(request.Url) {
		err = pkg.ErrUrlLengthLimit
		m.logger.Error(err)
		return
	}

	return
}

func (m *UrlMiddleware) FromCrawler(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(UrlMiddleware).FromCrawler(spider)
	}

	m.urlLengthLimit = spider.GetConfig().GetUrlLengthLimit()
	m.logger = spider.GetLogger()
	return m
}
