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

func (m *DumpPipeline) ProcessItem(item pkg.Item) (err error) {
	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		return
	}

	data := item.Data()
	if data == nil {
		err = errors.New("nil data")
		m.logger.Error(err)
		return
	}

	m.logger.Debug("Data", utils.JsonStr(data))

	m.logger.Debug("referrer", item.Referrer())
	m.logger.Info(m.GetSpider().Name(), item.GetContext().GetTaskId(), "item.Data:", utils.JsonStr(data))
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
