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
	m.logger.Info("data", utils.JsonStr(data))

	//m.stats.IncItemSuccess()
	return
}

func (m *DumpPipeline) FromCrawler(spider pkg.Spider) pkg.Pipeline {
	if m == nil {
		return new(DumpPipeline).FromCrawler(spider)
	}

	m.stats = spider.GetStats()
	m.logger = spider.GetLogger()
	return m
}
