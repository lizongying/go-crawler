package pipelines

import (
	"errors"
	"github.com/lizongying/go-crawler/pkg"
)

type NonePipeline struct {
	pkg.UnimplementedPipeline
	logger pkg.Logger
}

func (m *NonePipeline) ProcessItem(item pkg.Item) (err error) {
	spider := m.GetSpider()
	task := item.GetContext().GetTask()

	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		task.IncItemError()
		return
	}

	if item.Name() != pkg.ItemNone {
		m.logger.Warn("item not support", pkg.ItemNone)
		return
	}

	data := item.Data()
	if data == nil {
		err = errors.New("nil data")
		m.logger.Error(err)
		task.IncItemError()
		return
	}

	spider.GetCrawler().GetSignal().ItemStopped(item)
	task.IncItemSuccess()
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
