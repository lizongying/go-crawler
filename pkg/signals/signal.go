package signals

import (
	"github.com/lizongying/go-crawler/pkg"
)

type Signal struct {
	logger pkg.Logger

	crawlerChanged []pkg.FnCrawlerChanged

	spiderChanged []pkg.FnSpiderChanged

	jobChanged []pkg.FnJobChanged

	taskChanged []pkg.FnTaskChanged

	requestChanged []pkg.FnRequestChanged

	itemChanged []pkg.FnItemChanged
}

func (s *Signal) RegisterCrawlerChanged(fn pkg.FnCrawlerChanged) {
	s.crawlerChanged = append(s.crawlerChanged, fn)
}
func (s *Signal) RegisterSpiderChanged(fn pkg.FnSpiderChanged) {
	s.spiderChanged = append(s.spiderChanged, fn)
}
func (s *Signal) RegisterJobChanged(fn pkg.FnJobChanged) {
	s.jobChanged = append(s.jobChanged, fn)
}
func (s *Signal) RegisterTaskChanged(fn pkg.FnTaskChanged) {
	s.taskChanged = append(s.taskChanged, fn)
}
func (s *Signal) RegisterRequestChanged(fn pkg.FnRequestChanged) {
	s.requestChanged = append(s.requestChanged, fn)
}
func (s *Signal) RegisterItemChanged(fn pkg.FnItemChanged) {
	s.itemChanged = append(s.itemChanged, fn)
}
func (s *Signal) CrawlerChanged(ctx pkg.Context) {
	for _, v := range s.crawlerChanged {
		_ = v(ctx)
	}
}
func (s *Signal) SpiderChanged(ctx pkg.Context) {
	for _, v := range s.spiderChanged {
		_ = v(ctx)
	}
}
func (s *Signal) JobChanged(ctx pkg.Context) {
	for _, v := range s.jobChanged {
		_ = v(ctx)
	}
}
func (s *Signal) TaskChanged(task pkg.Task) {
	for _, v := range s.taskChanged {
		_ = v(task)
	}
}
func (s *Signal) RequestChanged(request pkg.Request) {
	for _, v := range s.requestChanged {
		_ = v(request)
	}
}
func (s *Signal) ItemChanged(item pkg.Item) {
	for _, v := range s.itemChanged {
		_ = v(item)
	}
}
func (s *Signal) FromCrawler(crawler pkg.Crawler) pkg.Signal {
	if s == nil {
		return new(Signal).FromCrawler(crawler)
	}

	s.logger = crawler.GetLogger()

	return s
}
