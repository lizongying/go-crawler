package pipelines

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	crawlerContext "github.com/lizongying/go-crawler/pkg/context"
)

type FilePipeline struct {
	pkg.UnimplementedPipeline
	scheduler pkg.Scheduler
	logger    pkg.Logger
}

func (m *FilePipeline) Start(ctx context.Context, spider pkg.Spider) (err error) {
	err = m.UnimplementedPipeline.Start(ctx, spider)
	m.scheduler = spider.GetScheduler()
	return nil
}

func (m *FilePipeline) ProcessItem(itemWithContext pkg.ItemWithContext) (err error) {
	if itemWithContext == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		return
	}

	files := itemWithContext.FilesRequest()
	if len(files) == 0 {
		return
	}

	for _, i := range files {
		ctx := &crawlerContext.Context{}
		r, e := m.scheduler.Request(ctx, i)
		if e != nil {
			m.logger.Error(e)
			continue
		}
		itemWithContext.SetFiles(r.Files())
	}

	return
}

func (m *FilePipeline) FromSpider(spider pkg.Spider) pkg.Pipeline {
	if m == nil {
		return new(FilePipeline).FromSpider(spider)
	}

	m.UnimplementedPipeline.FromSpider(spider)
	m.logger = spider.GetLogger()
	return m
}
