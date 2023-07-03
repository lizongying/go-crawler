package pipelines

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
)

type ImagePipeline struct {
	pkg.UnimplementedPipeline
	stats     pkg.Stats
	scheduler pkg.Scheduler
	logger    pkg.Logger
}

func (m *ImagePipeline) Start(_ context.Context, crawler pkg.Crawler) error {
	m.scheduler = crawler.GetScheduler()
	return nil
}

func (m *ImagePipeline) ProcessItem(ctx context.Context, item pkg.Item) (err error) {
	m.stats.IncItemTotal()

	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		return
	}

	images := item.GetImagesRequest()
	if len(images) == 0 {
		return
	}

	for _, i := range images {
		r, e := m.scheduler.Request(ctx, i)
		if e != nil {
			m.logger.Error(e)
			continue
		}
		item.SetImages(r.Images)
	}

	return
}

func (m *ImagePipeline) FromCrawler(crawler pkg.Crawler) pkg.Pipeline {
	if m == nil {
		return new(ImagePipeline).FromCrawler(crawler)
	}

	m.stats = crawler.GetStats()
	m.logger = crawler.GetLogger()
	return m
}
