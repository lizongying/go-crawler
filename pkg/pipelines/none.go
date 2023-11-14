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
	spider := m.Spider()
	task := item.GetContext().GetTask()

	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		task.IncItemError()
		return
	}

	if item.GetContext().GetItemProcessed() {
		return
	}

	data := item.Data()
	if data == nil {
		err = errors.New("nil data")
		m.logger.Error(err)
		task.IncItemError()
		return
	}

	item.GetContext().WithItemStatus(pkg.ItemStatusSuccess)
	spider.GetCrawler().GetSignal().ItemChanged(item)
	task.IncItemSuccess()
	return
}

func (m *NonePipeline) FromSpider(spider pkg.Spider) (err error) {
	if m == nil {
		return new(KafkaPipeline).FromSpider(spider)
	}

	if err = m.UnimplementedPipeline.FromSpider(spider); err != nil {
		return
	}
	m.logger = spider.GetLogger()
	return
}
