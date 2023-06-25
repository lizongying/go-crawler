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

func (m *FilterPipeline) ProcessItem(ctx context.Context, item pkg.Item) (err error) {
	if ctx == nil {
		ctx = context.Background()
	}

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

	err = m.filter.Store(ctx, uniqueKey)
	return
}

func (m *FilterPipeline) FromCrawler(crawler pkg.Crawler) pkg.Pipeline {
	if m == nil {
		return new(FilterPipeline).FromCrawler(crawler)
	}

	m.filter = crawler.GetFilter()
	m.logger = crawler.GetLogger()
	return m
}
