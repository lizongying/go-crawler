package pipelines

import (
	"context"
	"errors"
	items2 "github.com/lizongying/go-crawler/internal/items"
	"github.com/lizongying/go-crawler/pkg"
)

type CustomPipeline struct {
	pkg.UnimplementedPipeline
	logger pkg.Logger
}

func (m *CustomPipeline) Name() string {
	return "custom"
}

func (m *CustomPipeline) Start(ctx context.Context, spider pkg.Spider) (err error) {
	err = m.UnimplementedPipeline.Start(ctx, spider)
	m.logger.Debug("start")
	return nil
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

func (m *CustomPipeline) Stop(_ pkg.Context) error {
	m.logger.Debug("stop")
	return nil
}

func (m *CustomPipeline) FromSpider(spider pkg.Spider) pkg.Pipeline {
	if m == nil {
		return new(CustomPipeline).FromSpider(spider)
	}

	m.UnimplementedPipeline.FromSpider(spider)
	m.logger = spider.GetLogger()
	return m
}
