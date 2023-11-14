package pipelines

import (
	"errors"
	"github.com/lizongying/go-crawler/pkg"
)

type FilterPipeline struct {
	pkg.UnimplementedPipeline
	filter pkg.Filter
	logger pkg.Logger
}

func (m *FilterPipeline) ProcessItem(item pkg.Item) (err error) {
	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		return
	}

	uniqueKey := item.UniqueKey()
	if uniqueKey == "" {
		m.logger.Debug("uniqueKey is empty")
		return
	}
	m.logger.Info("uniqueKey", uniqueKey)

	err = m.filter.Store(item.GetContext(), uniqueKey)
	return
}

func (m *FilterPipeline) FromSpider(spider pkg.Spider) (err error) {
	if m == nil {
		return new(FilterPipeline).FromSpider(spider)
	}

	if err = m.UnimplementedPipeline.FromSpider(spider); err != nil {
		return
	}
	m.logger = spider.GetLogger()
	m.filter = spider.GetFilter()
	return
}
