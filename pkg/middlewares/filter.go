package middlewares

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"sync"
)

type FilterMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
	info   *pkg.SpiderInfo
	stats  pkg.Stats
	ids    sync.Map
}

func (m *FilterMiddleware) SpiderStart(_ context.Context, spider pkg.Spider) (err error) {
	m.info = spider.GetInfo()
	m.stats = spider.GetStats()
	return
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

	if _, ok := m.ids.Load(request.UniqueKey); ok {
		err = pkg.ErrIgnoreRequest
		m.logger.InfoF("%s in filter", request.UniqueKey)
		m.stats.IncRequestIgnore()
		return
	}

	return
}

func (m *FilterMiddleware) ProcessItem(c *pkg.Context) (err error) {
	item := c.Item
	if item.GetUniqueKey() == "" {
		err = c.NextItem()
		return
	}

	m.ids.Store(item.GetUniqueKey(), struct{}{})
	err = c.NextItem()
	return
}

func (m *FilterMiddleware) SpiderStop(_ context.Context) (err error) {
	m.ids.Range(func(key, _ any) bool {
		m.ids.Delete(key)
		return true
	})
	return
}

func (m *FilterMiddleware) FromCrawler(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(FilterMiddleware).FromCrawler(spider)
	}
	m.logger = spider.GetLogger()
	return m
}
