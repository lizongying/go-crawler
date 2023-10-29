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

	scheduleStarted []pkg.FnScheduleStarted
	scheduleStopped []pkg.FnScheduleStopped

	taskStarted []pkg.FnTaskStarted
	taskStopped []pkg.FnTaskStopped

	itemSaved []pkg.FnItemSaved
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
func (s *Signal) RegisterScheduleStarted(fn pkg.FnScheduleStarted) {
	s.scheduleStarted = append(s.scheduleStarted, fn)
}
func (s *Signal) RegisterScheduleStopped(fn pkg.FnScheduleStopped) {
	s.scheduleStopped = append(s.scheduleStopped, fn)
}
func (s *Signal) CrawlerStarted(ctx pkg.Context) {
	for _, v := range s.crawlerStarted {
		v(ctx)
	}
}
func (s *Signal) CrawlerStopped(ctx pkg.Context) {
	for _, v := range s.crawlerStopped {
		v(ctx)
	}
}
func (s *Signal) SpiderStarting(ctx pkg.Context) {
	for _, v := range s.spiderStarting {
		v(ctx)
	}
}
func (s *Signal) SpiderStarted(ctx pkg.Context) {
	for _, v := range s.spiderStarted {
		v(ctx)
	}
}
func (s *Signal) SpiderStopping(ctx pkg.Context) {
	for _, v := range s.spiderStopping {
		v(ctx)
	}
}
func (s *Signal) SpiderStopped(ctx pkg.Context) {
	for _, v := range s.spiderStopped {
		v(ctx)
	}
}
func (s *Signal) ScheduleStarted(ctx pkg.Context) {
	for _, v := range s.scheduleStarted {
		v(ctx)
	}
}
func (s *Signal) ScheduleStopped(ctx pkg.Context) {
	for _, v := range s.scheduleStopped {
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
func (s *Signal) ItemSaved(itemWithContext pkg.ItemWithContext) {
	for _, v := range s.itemSaved {
		v(itemWithContext)
	}
}
func (s *Signal) FromCrawler(crawler pkg.Crawler) pkg.Signal {
	if s == nil {
		return new(Signal).FromCrawler(crawler)
	}

	s.logger = crawler.GetLogger()

	return s
}
