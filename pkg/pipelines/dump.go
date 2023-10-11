package pipelines

import (
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
)

type DumpPipeline struct {
	pkg.UnimplementedPipeline
	logger pkg.Logger
}

func (m *DumpPipeline) ProcessItem(itemWithContext pkg.ItemWithContext) (err error) {
	if itemWithContext == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		return
	}

	data := itemWithContext.Data()
	if data == nil {
		err = errors.New("nil data")
		m.logger.Error(err)
		return
	}

	m.logger.Debug("referrer", itemWithContext.Referrer())
	m.logger.Info("item.Data:", utils.JsonStr(data))

	//m.stats.IncItemSuccess()
	return
}

func (m *DumpPipeline) FromSpider(spider pkg.Spider) pkg.Pipeline {
	if m == nil {
		return new(DumpPipeline).FromSpider(spider)
	}

	m.UnimplementedPipeline.FromSpider(spider)
	m.logger = spider.GetLogger()
	return m
}
