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

func (m *CustomPipeline) GetName() string {
	return "custom"
}

func (m *CustomPipeline) Start(_ context.Context, crawler pkg.Crawler) error {
	_ = m.FromCrawler(crawler)
	m.logger.Debug("start")
	return nil
}

func (m *CustomPipeline) ProcessItem(_ context.Context, item pkg.Item) (err error) {
	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		return
	}
	if item.GetName() != items2.Custom {
		m.logger.Warn("item not support", items2.Custom)
		return
	}
	itemCustom, ok := item.(*items2.ItemCustom)
	if !ok {
		m.logger.Warn("item parsing failed with", items2.Custom)
		return
	}
	m.logger.Debug("itemCustom", itemCustom)
	return nil
}

func (m *CustomPipeline) Stop(_ context.Context) error {
	m.logger.Debug("stop")
	return nil
}

func (m *CustomPipeline) FromCrawler(crawler pkg.Crawler) pkg.Pipeline {
	if m == nil {
		return new(CustomPipeline).FromCrawler(crawler)
	}

	m.logger = crawler.GetLogger()
	return m
}
