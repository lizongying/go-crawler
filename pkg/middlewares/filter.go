package middlewares

import (
	"context"
	"github.com/lizongying/go-crawler/internal"
	"github.com/lizongying/go-crawler/internal/logger"
	"github.com/lizongying/go-crawler/internal/utils"
	"sync"
)

type FilterMiddleware struct {
	internal.UnimplementedMiddleware
	logger *logger.Logger
	info   *internal.SpiderInfo
	ids    sync.Map
}

func (m *FilterMiddleware) GetName() string {
	return "filter"
}

func (m *FilterMiddleware) SpiderStart(_ context.Context, spider internal.Spider) (err error) {
	m.info = spider.GetInfo()
	return
}

func (m *FilterMiddleware) ProcessRequest(_ context.Context, r *internal.Request) (request *internal.Request, response *internal.Response, err error) {
	m.logger.Debug("request", utils.JsonStr(r))

	filterBefore, ok := m.info.Stats.Load("filter_before")
	if ok {
		filterBeforeInt := filterBefore.(int)
		filterBeforeInt++
		m.info.Stats.Store("filter_before", filterBeforeInt)
	}

	if r.SkipFilter {
		return
	}

	if r.UniqueKey == "" {
		return
	}

	if _, ok = m.ids.Load(r.UniqueKey); ok {
		err = internal.ErrIgnoreRequest
		m.logger.InfoF("%s in filter", r.UniqueKey)
		return
	}

	return
}

func (m *FilterMiddleware) ProcessItem(_ context.Context, item *internal.Item) (err error) {
	if item.UniqueKey == "" {
		return
	}

	m.ids.Store(item.UniqueKey, struct{}{})
	return
}

func (m *FilterMiddleware) SpiderStop(_ context.Context) (err error) {
	m.ids.Range(func(key, _ any) bool {
		m.ids.Delete(key)
		return true
	})
	return
}

func NewFilterMiddleware(logger *logger.Logger) (m internal.Middleware) {
	m = &FilterMiddleware{
		logger: logger,
	}
	return
}
