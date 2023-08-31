package pipelines

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
)

type ImagePipeline struct {
	pkg.UnimplementedPipeline
	scheduler pkg.Scheduler
	logger    pkg.Logger
}

func (m *ImagePipeline) Start(ctx context.Context, spider pkg.Spider) (err error) {
	err = m.UnimplementedPipeline.Start(ctx, spider)
	m.scheduler = spider.GetScheduler()
	return nil
}

func (m *ImagePipeline) ProcessItem(_ context.Context, item pkg.Item) (err error) {
	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		return
	}

	images := item.ImagesRequest()
	if len(images) == 0 {
		return
	}

	for _, i := range images {
		ctx := pkg.Context{}
		r, e := m.scheduler.Request(ctx, i)
		if e != nil {
			m.logger.Error(e)
			continue
		}
		item.SetImages(r.Images())
	}

	return
}

func (m *ImagePipeline) FromSpider(spider pkg.Spider) pkg.Pipeline {
	if m == nil {
		return new(ImagePipeline).FromSpider(spider)
	}

	m.UnimplementedPipeline.FromSpider(spider)
	m.logger = spider.GetLogger()
	return m
}
