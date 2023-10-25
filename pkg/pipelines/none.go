package pipelines

import (
	"errors"
	"github.com/lizongying/go-crawler/pkg"
)

type NonePipeline struct {
	pkg.UnimplementedPipeline
	logger pkg.Logger
}

func (m *NonePipeline) ProcessItem(itemWithContext pkg.ItemWithContext) (err error) {
	spider := m.GetSpider()

	if itemWithContext == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		spider.IncItemError()
		return
	}

	if itemWithContext.Name() != pkg.ItemNone {
		m.logger.Warn("item not support", pkg.ItemNone)
		return
	}

	data := itemWithContext.Data()
	if data == nil {
		err = errors.New("nil data")
		m.logger.Error(err)
		spider.IncItemError()
		return
	}

	spider.GetCrawler().GetSignal().ItemSaved(itemWithContext)
	spider.IncItemSuccess()
	return
}

func (m *NonePipeline) FromSpider(spider pkg.Spider) pkg.Pipeline {
	if m == nil {
		return new(KafkaPipeline).FromSpider(spider)
	}

	m.UnimplementedPipeline.FromSpider(spider)
	m.logger = spider.GetLogger()
	return m
}
