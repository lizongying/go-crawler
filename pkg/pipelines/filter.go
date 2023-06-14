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

func (m *FilterPipeline) ProcessItem(_ context.Context, item pkg.Item) (err error) {
	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		return
	}

	uniqueKey := item.GetUniqueKey()
	if uniqueKey == "" {
		return
	}

	m.filter.ExistsOrStore(uniqueKey)
	return
}

func (m *FilterPipeline) FromCrawler(spider pkg.Spider) pkg.Pipeline {
	if m == nil {
		return new(FilterPipeline).FromCrawler(spider)
	}

	m.filter = spider.GetFilter()
	m.logger = spider.GetLogger()
	return m
}
