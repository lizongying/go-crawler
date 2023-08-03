package middlewares

import (
	"github.com/lizongying/go-crawler/pkg"
)

type UrlMiddleware struct {
	pkg.UnimplementedMiddleware
	urlLengthLimit int
	logger         pkg.Logger
}

func (m *UrlMiddleware) ProcessRequest(_ pkg.Context, request pkg.Request) (err error) {
	if m.urlLengthLimit < len(request.GetUrl()) {
		err = pkg.ErrUrlLengthLimit
		m.logger.Error(err)
		return
	}

	return
}

func (m *UrlMiddleware) FromSpider(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(UrlMiddleware).FromSpider(spider)
	}

	m.UnimplementedMiddleware.FromSpider(spider)
	crawler := spider.GetCrawler()
	m.urlLengthLimit = crawler.GetConfig().GetUrlLengthLimit()
	m.logger = spider.GetLogger()
	return m
}
