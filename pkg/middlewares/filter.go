package middlewares

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
)

type FilterMiddleware struct {
	pkg.UnimplementedMiddleware
	stats  pkg.Stats
	logger pkg.Logger
	filter pkg.Filter
}

func (m *FilterMiddleware) ProcessRequest(_ context.Context, request *pkg.Request) (err error) {
	if request.SkipFilter {
		m.logger.Debug("SkipFilter")
		return
	}

	if request.UniqueKey == "" {
		m.logger.Debug("UniqueKey is empty")
		return
	}

	if m.filter.ExistsOrStore(request.UniqueKey) {
		err = pkg.ErrIgnoreRequest
		m.logger.InfoF("%s in filter", request.UniqueKey)
		m.stats.IncRequestIgnore()
		return
	}

	return
}

func (m *FilterMiddleware) SpiderStop(_ context.Context) (err error) {
	m.filter.Clean()
	return
}

func (m *FilterMiddleware) FromCrawler(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(FilterMiddleware).FromCrawler(spider)
	}

	m.logger = spider.GetLogger()
	m.filter = spider.GetFilter()
	m.stats = spider.GetStats()
	return m
}
