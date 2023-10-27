package signals

import (
	"github.com/lizongying/go-crawler/pkg"
)

type Signal struct {
	logger pkg.Logger

	crawlerStarted []pkg.FnCrawlerStarted
	crawlerStopped []pkg.FnCrawlerStopped

	spiderStarting []pkg.FnSpiderStarting
	spiderStarted  []pkg.FnSpiderStarted
	spiderStopping []pkg.FnSpiderStopping
	spiderStopped  []pkg.FnSpiderStopped

	taskStarted []pkg.FnTaskStarted
	taskStopped []pkg.FnTaskStopped

	itemSaved []pkg.FnItemSaved

	scheduled []pkg.FnScheduled
}

func (s *Signal) RegisterCrawlerStarted(fn pkg.FnCrawlerStarted) {
	s.crawlerStarted = append(s.crawlerStarted, fn)
}
func (s *Signal) RegisterCrawlerStopped(fn pkg.FnCrawlerStopped) {
	s.crawlerStopped = append(s.crawlerStopped, fn)
}
func (s *Signal) RegisterSpiderStarting(fn pkg.FnSpiderStarting) {
	s.spiderStarting = append(s.spiderStarting, fn)
}
func (s *Signal) RegisterSpiderStarted(fn pkg.FnSpiderStarted) {
	s.spiderStarted = append(s.spiderStarted, fn)
}
func (s *Signal) RegisterSpiderStopping(fn pkg.FnSpiderStopping) {
	s.spiderStopping = append(s.spiderStopping, fn)
}
func (s *Signal) RegisterSpiderStopped(fn pkg.FnSpiderStopped) {
	s.spiderStopped = append(s.spiderStopped, fn)
}
func (s *Signal) RegisterTaskStarted(fn pkg.FnTaskStarted) {
	s.taskStarted = append(s.taskStarted, fn)
}
func (s *Signal) RegisterTaskStopped(fn pkg.FnTaskStopped) {
	s.taskStopped = append(s.taskStopped, fn)
}
func (s *Signal) RegisterItemSaved(fn pkg.FnItemSaved) {
	s.itemSaved = append(s.itemSaved, fn)
}
func (s *Signal) RegisterScheduled(fn pkg.FnScheduled) {
	s.scheduled = append(s.scheduled, fn)
}
func (s *Signal) CrawlerStarted(crawler pkg.Crawler) {
	for _, v := range s.crawlerStarted {
		v(crawler)
	}
}
func (s *Signal) CrawlerStopped(crawler pkg.Crawler) {
	for _, v := range s.crawlerStopped {
		v(crawler)
	}
}
func (s *Signal) SpiderStarting(spider pkg.Spider) {
	for _, v := range s.spiderStarting {
		v(spider)
	}
}
func (s *Signal) SpiderStarted(spider pkg.Spider) {
	for _, v := range s.spiderStarted {
		v(spider)
	}
}
func (s *Signal) SpiderStopping(spider pkg.Spider) {
	for _, v := range s.spiderStopping {
		v(spider)
	}
}
func (s *Signal) SpiderStopped(spider pkg.Spider) {
	for _, v := range s.spiderStopped {
		v(spider)
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
func (s *Signal) ItemSaved(itemWithContext pkg.ItemWithContext) {
	for _, v := range s.itemSaved {
		v(itemWithContext)
	}
}
func (s *Signal) Scheduled(ctx pkg.Context) {
	for _, v := range s.scheduled {
		v(ctx)
	}
}
func (s *Signal) FromCrawler(crawler pkg.Crawler) pkg.Signal {
	if s == nil {
		return new(Signal).FromCrawler(crawler)
	}

	s.logger = crawler.GetLogger()

	return s
}
