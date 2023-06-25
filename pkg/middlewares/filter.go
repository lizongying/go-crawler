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

func (m *FilterMiddleware) ProcessRequest(ctx context.Context, request *pkg.Request) (err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if request.SkipFilter {
		m.logger.Debug("SkipFilter")
		return
	}

	if request.UniqueKey == "" {
		m.logger.Debug("UniqueKey is empty")
		return
	}

	ok, e := m.filter.IsExist(ctx, request.UniqueKey)
	if err != nil {
		err = e
		return
	}

	if ok {
		err = pkg.ErrIgnoreRequest
		m.logger.InfoF("%s in filter", request.UniqueKey)
		m.stats.IncRequestIgnore()
		return
	}

	return
}

func (m *FilterMiddleware) Stop(ctx context.Context) (err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	err = m.filter.Clean(ctx)
	return
}

func (m *FilterMiddleware) FromCrawler(crawler pkg.Crawler) pkg.Middleware {
	if m == nil {
		return new(FilterMiddleware).FromCrawler(crawler)
	}

	m.logger = crawler.GetLogger()
	m.filter = crawler.GetFilter()
	m.stats = crawler.GetStats()
	return m
}
