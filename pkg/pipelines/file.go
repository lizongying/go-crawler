package pipelines

import (
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	crawlerContext "github.com/lizongying/go-crawler/pkg/context"
)

type FilePipeline struct {
	pkg.UnimplementedPipeline
	logger pkg.Logger
}

func (m *FilePipeline) ProcessItem(item pkg.Item) (err error) {
	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		return
	}

	files := item.FilesRequest()
	if len(files) == 0 {
		return
	}

	for _, i := range files {
		ctx := &crawlerContext.Context{}
		r, e := m.Spider().Request(ctx, i)
		if e != nil {
			m.logger.Error(e)
			continue
		}
		item.SetFiles(r.Files())
	}

	return
}

func (m *FilePipeline) FromSpider(spider pkg.Spider) (err error) {
	if m == nil {
		return new(FilePipeline).FromSpider(spider)
	}

	if err = m.UnimplementedPipeline.FromSpider(spider); err != nil {
		return
	}
	m.logger = spider.GetLogger()
	return
}
