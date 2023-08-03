package pipelines

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
)

type FilterPipeline struct {
	pkg.UnimplementedPipeline
	filter pkg.Filter
	logger pkg.Logger
}

func (m *FilterPipeline) Start(ctx context.Context, spider pkg.Spider) (err error) {
	err = m.UnimplementedPipeline.Start(ctx, spider)
	m.filter = spider.GetFilter()
	return nil
}

func (m *FilterPipeline) ProcessItem(_ context.Context, item pkg.Item) (err error) {
	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		return
	}

	uniqueKey := item.GetUniqueKey()
	if uniqueKey == "" {
		m.logger.Debug("uniqueKey is empty")
		return
	}
	m.logger.Info("uniqueKey", uniqueKey)

	ctx := pkg.Context{}
	err = m.filter.Store(ctx, uniqueKey)
	return
}

func (m *FilterPipeline) FromSpider(spider pkg.Spider) pkg.Pipeline {
	if m == nil {
		return new(FilterPipeline).FromSpider(spider)
	}

	m.UnimplementedPipeline.FromSpider(spider)
	m.logger = spider.GetLogger()
	return m
}
