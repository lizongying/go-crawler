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

	m.logger.Debug("Data", utils.UnsafeJSON(data))

	m.logger.Debug("referrer", item.Referrer())
	m.logger.Debug(m.Spider().Name(), item.GetContext().GetTask().GetId(), "item.Data:", utils.UnsafeJSON(data))
	return
}

func (m *DumpPipeline) FromSpider(spider pkg.Spider) (err error) {
	if m == nil {
		return new(DumpPipeline).FromSpider(spider)
	}

	if err = m.UnimplementedPipeline.FromSpider(spider); err != nil {
		return
	}
	m.logger = spider.GetLogger()
	return
}
