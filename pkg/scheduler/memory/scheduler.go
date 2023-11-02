package memory

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/scheduler"
	"sync"
)

const defaultRequestMax = 1000 * 1000

type Scheduler struct {
	scheduler.UnimplementedScheduler

	requestChan  chan pkg.Request
	extraChanMap sync.Map

	crawler pkg.Crawler
	spider  pkg.Spider
	config  pkg.Config
	logger  pkg.Logger
	task    pkg.Task
}

// StartScheduler
// ctx: ContextTask
func (s *Scheduler) StartScheduler(ctx pkg.Context) (err error) {
	s.task = ctx.GetTask()
	s.UnimplementedScheduler.SetTask(s.task)

	go s.HandleItem(ctx)

	go s.handleRequest(ctx)

	return
}

func (s *Scheduler) StopScheduler(_ pkg.Context) (err error) {
	return
}
func (s *Scheduler) FromSpider(spider pkg.Spider) pkg.Scheduler {
	if s == nil {
		return new(Scheduler).FromSpider(spider)
	}

	s.crawler = spider.GetCrawler()
	s.UnimplementedScheduler.SetCrawler(s.crawler)
	s.spider = spider
	s.UnimplementedScheduler.SetSpider(spider)
	s.config = spider.GetConfig()
	s.logger = spider.GetLogger()
	s.UnimplementedScheduler.SetLogger(s.logger)
	s.UnimplementedScheduler.Init()

	s.requestChan = make(chan pkg.Request, defaultRequestMax)

	return s
}
