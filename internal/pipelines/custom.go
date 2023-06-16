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

func (m *CustomPipeline) Start(_ context.Context, crawler pkg.Crawler) error {
	_ = m.FromCrawler(crawler)
	m.logger.Debug("start")
	return nil
}

func (m *CustomPipeline) ProcessItem(_ context.Context, _ pkg.Item) error {
	m.logger.Debug("item")
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
