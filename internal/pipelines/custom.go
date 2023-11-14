package pipelines

import (
	"errors"
	items2 "github.com/lizongying/go-crawler/internal/items"
	"github.com/lizongying/go-crawler/pkg"
)

type CustomPipeline struct {
	pkg.UnimplementedPipeline
	logger pkg.Logger
}

func (m *CustomPipeline) ProcessItem(item pkg.Item) (err error) {
	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		return
	}
	if item.Name() != items2.Custom {
		m.logger.Warn("item not support", items2.Custom)
		return
	}
	itemCustom, ok := item.GetItem().(*items2.ItemCustom)
	if !ok {
		m.logger.Warn("item parsing failed with", items2.Custom)
		return
	}
	m.logger.Debug("itemCustom", itemCustom)
	return nil
}

func (m *CustomPipeline) FromSpider(spider pkg.Spider) (err error) {
	if m == nil {
		return new(CustomPipeline).FromSpider(spider)
	}

	if err = m.UnimplementedPipeline.FromSpider(spider); err != nil {
		return
	}
	m.logger = spider.GetLogger()
	return
}
