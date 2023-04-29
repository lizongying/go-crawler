package middlewares

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/logger"
	"sync"
)

type FilterMiddleware struct {
	pkg.UnimplementedMiddleware
	logger *logger.Logger
	info   *pkg.SpiderInfo
	ids    sync.Map
}

func (m *FilterMiddleware) GetName() string {
	return "filter"
}

func (m *FilterMiddleware) SpiderStart(_ context.Context, spider pkg.Spider) (err error) {
	m.info = spider.GetInfo()
	return
}

func (m *FilterMiddleware) ProcessRequest(c *pkg.Context) (err error) {
	r := c.Request
	m.logger.DebugF("request: %+v", r)

	filterBefore, ok := m.info.Stats.Load("filter_before")
	if ok {
		filterBeforeInt := filterBefore.(int)
		filterBeforeInt++
		m.info.Stats.Store("filter_before", filterBeforeInt)
	}

	if r.SkipFilter {
		m.logger.Debug("SkipFilter")
		err = c.NextRequest()
		return
	}

	if r.UniqueKey == "" {
		m.logger.Debug("UniqueKey is empty")
		err = c.NextRequest()
		return
	}

	if _, ok = m.ids.Load(r.UniqueKey); ok {
		err = pkg.ErrIgnoreRequest
		m.logger.InfoF("%s in filter", r.UniqueKey)
		return
	}

	err = c.NextRequest()
	return
}

func (m *FilterMiddleware) ProcessItem(c *pkg.Context) (err error) {
	item := c.Item
	if item.UniqueKey == "" {
		err = c.NextItem()
		return
	}

	m.ids.Store(item.UniqueKey, struct{}{})
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

func NewFilterMiddleware(logger *logger.Logger) (m pkg.Middleware) {
	m = &FilterMiddleware{
		logger: logger,
	}
	return
}
