package pipelines

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
)

type DumpPipeline struct {
	pkg.UnimplementedPipeline
	stats  pkg.Stats
	logger pkg.Logger
}

func (m *DumpPipeline) ProcessItem(_ context.Context, item pkg.Item) (err error) {
	m.stats.IncItemTotal()

	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		return
	}

	data := item.GetData()
	if data == nil {
		err = errors.New("nil data")
		m.logger.Error(err)
		return
	}

	m.logger.Debug("referer", item.GetReferer())
	m.logger.Info("item.Data:", utils.JsonStr(data))

	//m.stats.IncItemSuccess()
	return
}

func (m *DumpPipeline) FromCrawler(crawler pkg.Crawler) pkg.Pipeline {
	if m == nil {
		return new(DumpPipeline).FromCrawler(crawler)
	}

	m.stats = crawler.GetStats()
	m.logger = crawler.GetLogger()
	return m
}
