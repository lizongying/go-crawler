package middlewares

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
)

type FilterMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
	filter pkg.Filter
}

func (m *FilterMiddleware) Start(ctx context.Context, spider pkg.Spider) (err error) {
	err = m.UnimplementedMiddleware.Start(ctx, spider)
	m.filter = spider.GetFilter()
	return
}

func (m *FilterMiddleware) ProcessRequest(ctx pkg.Context, request pkg.Request) (err error) {
	task := ctx.GetTask()
	skipFilter := false
	if request.IsSkipFilter() != nil {
		skipFilter = *request.IsSkipFilter()
	}
	if skipFilter {
		m.logger.Debug("SkipFilter")
		return
	}

	if request.GetUniqueKey() == "" {
		m.logger.Debug("UniqueKey is empty")
		return
	}

	ok, e := m.filter.IsExist(ctx, request.GetUniqueKey())
	if err != nil {
		err = e
		return
	}

	if ok {
		err = pkg.ErrIgnoreRequest
		m.logger.Infof("%s in filter", request.GetUniqueKey())
		task.IncRequestIgnore()
		return
	}

	return
}

func (m *FilterMiddleware) Stop(ctx pkg.Context) (err error) {
	err = m.filter.Clean(ctx)
	return
}

func (m *FilterMiddleware) FromSpider(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(FilterMiddleware).FromSpider(spider)
	}

	m.UnimplementedMiddleware.FromSpider(spider)
	m.logger = spider.GetLogger()
	return m
}
