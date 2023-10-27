package statistics

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/statistics/node"
	"github.com/lizongying/go-crawler/pkg/statistics/record"
	statisticsSpider "github.com/lizongying/go-crawler/pkg/statistics/spider"
	"github.com/lizongying/go-crawler/pkg/statistics/task"
	"github.com/lizongying/go-crawler/pkg/utils"
	"os"
	"time"
)

type Statistics struct {
	crawler   pkg.Crawler
	logger    pkg.Logger
	Nodes     map[string]pkg.StatisticsNode
	Spiders   map[string]pkg.StatisticsSpider
	Schedules []pkg.StatisticsSchedule
	Tasks     map[string]pkg.StatisticsTask
	Records   []pkg.StatisticsRecord
}

func (s *Statistics) GetNodes() (nodes []pkg.StatisticsNode) {
	for _, v := range s.Nodes {
		nodes = append(nodes, v)
	}
	return
}
func (s *Statistics) GetSpiders() (spiders []pkg.StatisticsSpider) {
	for _, v := range s.Spiders {
		spiders = append(spiders, v)
	}
	return
}
func (s *Statistics) GetSchedules() []pkg.StatisticsSchedule {
	return s.Schedules
}
func (s *Statistics) GetTasks() (tasks []pkg.StatisticsTask) {
	for _, v := range s.Tasks {
		tasks = append(tasks, v)
	}
	return
}
func (s *Statistics) GetRecords() []pkg.StatisticsRecord {
	return s.Records
}
func (s *Statistics) AddSchedules(schedules ...pkg.StatisticsSchedule) {
	s.Schedules = append(s.Schedules, schedules...)
}
func (s *Statistics) AddSpiders(spiders ...pkg.Spider) {
	for _, v := range spiders {
		s.Nodes[v.GetCrawler().GetId()].IncSpider()
		s.Spiders[v.Name()] = &statisticsSpider.Spider{
			Spider: v.Name(),
		}
	}
}
func (s *Statistics) AddTasks(tasks ...pkg.StatisticsTask) {
	for _, v := range tasks {
		//s.Tasks[v.GetCrawler().GetId()].IncSpider()
		s.Tasks[v.GetId()] = &task.Task{
			Id: v.GetId(),
		}
	}
}
func (s *Statistics) AddRecords(records ...pkg.StatisticsRecord) {
	s.Records = append(s.Records, records...)
}
func (s *Statistics) crawlerStarted(crawler pkg.Crawler) {
	s.Nodes[crawler.GetId()].(*node.Node).
		WithStatus(crawler.GetStatus()).
		WithStartTime(crawler.GetStartTime())
}
func (s *Statistics) crawlerStopped(crawler pkg.Crawler) {
	s.Nodes[crawler.GetId()].(*node.Node).
		WithStatus(crawler.GetStatus()).
		WithFinishTime(crawler.GetStopTime())
}
func (s *Statistics) spiderStarted(spider pkg.Spider) {
	s.Spiders[spider.Name()].
		WithStatus(spider.GetContext().GetStatus()).
		WithLastRunAt(spider.GetContext().GetStartTime())
}

func (s *Statistics) spiderStopped(spider pkg.Spider) {
	s.Spiders[spider.Name()].
		WithStatus(spider.GetContext().GetStatus()).
		WithLastFinishAt(spider.GetContext().GetStopTime())
}
func (s *Statistics) taskStarted(ctx pkg.Context) {
	s.Nodes[ctx.GetCrawlerId()].IncTask()
	s.Spiders[ctx.GetSpiderName()].IncTask()

	s.Tasks[ctx.GetTaskId()] = new(task.Task).
		WithId(ctx.GetTaskId()).
		WithNode(ctx.GetCrawlerId()).
		WithSpider(ctx.GetSpiderName())

	s.Tasks[ctx.GetTaskId()].
		WithStatus(ctx.GetTaskStatus()).
		WithStartTime(ctx.GetTaskStartTime())
}

func (s *Statistics) taskStopped(ctx pkg.Context) {
	s.Tasks[ctx.GetTaskId()].
		WithStatus(ctx.GetTaskStatus()).
		WithFinishTime(ctx.GetTaskStopTime())
}
func (s *Statistics) itemSaved(itemWithContext pkg.ItemWithContext) {
	s.AddRecords(new(record.Record).
		WithSaveTime(time.Now()).
		WithSpider(itemWithContext.GetSpiderName()).
		WithTaskId(itemWithContext.GetTaskId()).
		WithMeta(itemWithContext.MetaJson()).
		WithData(itemWithContext.DataJson()),
	)
	s.Nodes[itemWithContext.GetCrawlerId()].IncRecord()
	s.Spiders[itemWithContext.GetSpiderName()].IncRecord()
	s.Tasks[itemWithContext.GetTaskId()].IncRecord()
}
func (s *Statistics) FromCrawler(crawler pkg.Crawler) pkg.Statistics {
	if s == nil {
		return new(Statistics).FromCrawler(crawler)
	}

	s.crawler = crawler
	s.logger = crawler.GetLogger()

	s.Nodes = make(map[string]pkg.StatisticsNode)
	s.Spiders = make(map[string]pkg.StatisticsSpider)
	s.Tasks = make(map[string]pkg.StatisticsTask)

	hostname, _ := os.Hostname()
	s.Nodes[crawler.GetId()] = &node.Node{
		Hostname: hostname,
		Ip:       utils.LanIp(),
		Enable:   true,
	}
	s.Nodes[crawler.GetId()].WithId(crawler.GetId())

	signal := s.crawler.GetSignal()
	signal.RegisterCrawlerStarted(s.crawlerStarted)
	signal.RegisterCrawlerStopped(s.crawlerStopped)
	signal.RegisterSpiderStarted(s.spiderStarted)
	signal.RegisterSpiderStopped(s.spiderStopped)
	signal.RegisterTaskStarted(s.taskStarted)
	signal.RegisterTaskStopped(s.taskStopped)
	signal.RegisterItemSaved(s.itemSaved)

	return s
}
