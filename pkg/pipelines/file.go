package pipelines

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
)

type FilePipeline struct {
	pkg.UnimplementedPipeline
	stats     pkg.Stats
	scheduler pkg.Scheduler
	logger    pkg.Logger
}

func (m *FilePipeline) Start(_ context.Context, crawler pkg.Crawler) error {
	m.scheduler = crawler.GetScheduler()
	return nil
}

func (m *FilePipeline) ProcessItem(ctx context.Context, item pkg.Item) (err error) {
	m.stats.IncItemTotal()

	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		return
	}

	files := item.GetFilesRequest()
	if len(files) == 0 {
		return
	}

	for _, i := range files {
		r, e := m.scheduler.Request(ctx, i)
		if e != nil {
			m.logger.Error(e)
			continue
		}
		item.SetFiles(r.GetFiles())
	}

	return
}

func (m *FilePipeline) FromCrawler(crawler pkg.Crawler) pkg.Pipeline {
	if m == nil {
		return new(FilePipeline).FromCrawler(crawler)
	}

	m.stats = crawler.GetStats()
	m.logger = crawler.GetLogger()
	return m
}
