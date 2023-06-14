package pipelines

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
)

type CustomPipeline struct {
	pkg.UnimplementedPipeline
	logger pkg.Logger
}

func (m *CustomPipeline) GetName() string {
	return "custom"
}

func (m *CustomPipeline) SpiderStart(_ context.Context, spider pkg.Spider) error {
	_ = m.FromCrawler(spider)
	m.logger.Debug("start")
	return nil
}

func (m *CustomPipeline) ProcessItem(_ context.Context, _ pkg.Item) error {
	m.logger.Debug("item")
	return nil
}

func (m *CustomPipeline) SpiderStop(_ context.Context) error {
	m.logger.Debug("stop")
	return nil
}

func (m *CustomPipeline) FromCrawler(spider pkg.Spider) pkg.Pipeline {
	if m == nil {
		return new(CustomPipeline).FromCrawler(spider)
	}

	m.logger = spider.GetLogger()
	return m
}
