package signals

import (
	"github.com/lizongying/go-crawler/pkg"
)

type Signal struct {
	logger pkg.Logger

	crawlerChanged []pkg.FnCrawlerChanged

	spiderChanged []pkg.FnSpiderChanged

	jobChanged []pkg.FnJobChanged

	taskStarted []pkg.FnTaskStarted
	taskStopped []pkg.FnTaskStopped

	requestStarted []pkg.FnRequestStarted
	requestStopped []pkg.FnRequestStopped

	itemStarted []pkg.FnItemStarted
	itemStopped []pkg.FnItemStopped
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
func (s *Signal) RegisterTaskStarted(fn pkg.FnTaskStarted) {
	s.taskStarted = append(s.taskStarted, fn)
}
func (s *Signal) RegisterTaskStopped(fn pkg.FnTaskStopped) {
	s.taskStopped = append(s.taskStopped, fn)
}
func (s *Signal) RegisterRequestStarted(fn pkg.FnRequestStarted) {
	s.requestStarted = append(s.requestStarted, fn)
}
func (s *Signal) RegisterRequestStopped(fn pkg.FnRequestStopped) {
	s.requestStopped = append(s.requestStopped, fn)
}
func (s *Signal) RegisterItemStarted(fn pkg.FnItemStarted) {
	s.itemStarted = append(s.itemStarted, fn)
}
func (s *Signal) RegisterItemStopped(fn pkg.FnItemStopped) {
	s.itemStopped = append(s.itemStopped, fn)
}
func (s *Signal) CrawlerChanged(ctx pkg.Context) {
	for _, v := range s.crawlerChanged {
		v(ctx)
	}
}
func (s *Signal) SpiderChanged(ctx pkg.Context) {
	for _, v := range s.spiderChanged {
		v(ctx)
	}
}
func (s *Signal) JobChanged(ctx pkg.Context) {
	for _, v := range s.jobChanged {
		v(ctx)
	}
}
func (s *Signal) TaskStarted(ctx pkg.Context) {
	for _, v := range s.taskStarted {
		v(ctx)
	}
}
func (s *Signal) TaskStopped(ctx pkg.Context) {
	for _, v := range s.taskStopped {
		v(ctx)
	}
}
func (s *Signal) RequestStarted(ctx pkg.Context) {
	for _, v := range s.requestStarted {
		v(ctx)
	}
}
func (s *Signal) RequestStopped(ctx pkg.Context) {
	for _, v := range s.requestStopped {
		v(ctx)
	}
}
func (s *Signal) ItemStarted(ctx pkg.Context) {
	for _, v := range s.itemStarted {
		v(ctx)
	}
}
func (s *Signal) ItemStopped(item pkg.Item) {
	for _, v := range s.itemStopped {
		v(item)
	}
}
func (s *Signal) FromCrawler(crawler pkg.Crawler) pkg.Signal {
	if s == nil {
		return new(Signal).FromCrawler(crawler)
	}

	s.logger = crawler.GetLogger()

	return s
}
