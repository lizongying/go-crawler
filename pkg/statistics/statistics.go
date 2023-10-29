package statistics

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/statistics/node"
	"github.com/lizongying/go-crawler/pkg/statistics/record"
	"github.com/lizongying/go-crawler/pkg/statistics/schedule"
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
	Schedules map[string]pkg.StatisticsSchedule
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
func (s *Statistics) GetSchedules() (schedules []pkg.StatisticsSchedule) {
	for _, v := range s.Schedules {
		schedules = append(schedules, v)
	}
	return
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
func (s *Statistics) crawlerStarted(ctx pkg.Context) {
	hostname, _ := os.Hostname()
	s.Nodes[ctx.GetCrawlerId()] = &node.Node{
		Hostname: hostname,
		Ip:       utils.LanIp(),
		Enable:   true,
	}
	s.Nodes[ctx.GetCrawlerId()].WithId(ctx.GetCrawlerId())

	s.Nodes[ctx.GetCrawlerId()].(*node.Node).
		WithStatus(ctx.GetCrawlerStatus()).
		WithStartTime(ctx.GetCrawlerStartTime())

	for _, v := range s.crawler.GetSpiders() {
		s.Nodes[ctx.GetCrawlerId()].IncSpider()
		s.Spiders[v.Name()] = &statisticsSpider.Spider{
			Node:   ctx.GetCrawlerId(),
			Spider: v.Name(),
		}
	}
}
func (s *Statistics) crawlerStopped(ctx pkg.Context) {
	s.Nodes[ctx.GetCrawlerId()].(*node.Node).
		WithStatus(ctx.GetCrawlerStatus()).
		WithFinishTime(ctx.GetCrawlerStopTime())
}
func (s *Statistics) spiderStarted(ctx pkg.Context) {
	s.Spiders[ctx.GetSpiderName()].
		WithStatus(ctx.GetSpiderStatus()).
		WithStartTime(ctx.GetSpiderStartTime())
}

func (s *Statistics) spiderStopped(ctx pkg.Context) {
	s.Spiders[ctx.GetSpiderName()].
		WithStatus(ctx.GetSpiderStatus()).
		WithFinishTime(ctx.GetSpiderStopTime())
}
func (s *Statistics) scheduleStarted(ctx pkg.Context) {
	s.Nodes[ctx.GetCrawlerId()].IncSchedule()
	s.Spiders[ctx.GetSpiderName()].IncSchedule()

	s.Schedules[ctx.GetScheduleId()] = new(schedule.Schedule).
		WithId(ctx.GetScheduleId()).
		WithEnable(ctx.GetScheduleEnable()).
		WithNode(ctx.GetCrawlerId()).
		WithSpider(ctx.GetSpiderName())

	s.Schedules[ctx.GetScheduleId()].
		WithStatus(ctx.GetScheduleStatus()).
		WithStartTime(ctx.GetScheduleStartTime())
}
func (s *Statistics) scheduleStopped(ctx pkg.Context) {
	s.Schedules[ctx.GetScheduleId()].
		WithStatus(ctx.GetScheduleStatus()).
		WithFinishTime(ctx.GetScheduleStopTime())
}
func (s *Statistics) taskStarted(ctx pkg.Context) {
	s.Nodes[ctx.GetCrawlerId()].IncTask()
	s.Spiders[ctx.GetSpiderName()].IncTask()
	s.Schedules[ctx.GetScheduleId()].IncTask()

	s.Tasks[ctx.GetTaskId()] = new(task.Task).
		WithId(ctx.GetTaskId()).
		WithNode(ctx.GetCrawlerId()).
		WithSpider(ctx.GetSpiderName()).
		WithSchedule(ctx.GetScheduleId())

	s.Tasks[ctx.GetTaskId()].
		WithStatus(ctx.GetTaskStatus()).
		WithStartTime(ctx.GetTaskStartTime())

	// spider
	s.Spiders[ctx.GetSpiderName()].
		WithLastTaskId(ctx.GetTaskId()).
		WithLastTaskStatus(ctx.GetTaskStatus()).
		WithLastTaskStartTime(ctx.GetTaskStartTime())
}

func (s *Statistics) taskStopped(ctx pkg.Context) {
	s.Tasks[ctx.GetTaskId()].
		WithStatus(ctx.GetTaskStatus()).
		WithFinishTime(ctx.GetTaskStopTime())

	// spider
	spider := s.Spiders[ctx.GetSpiderName()]
	if ctx.GetTaskId() == spider.GetLastTaskId() {
		spider.
			WithLastTaskStatus(ctx.GetTaskStatus()).
			WithLastTaskFinishTime(ctx.GetTaskStopTime())
	}
}
func (s *Statistics) itemSaved(itemWithContext pkg.ItemWithContext) {
	id := itemWithContext.Id()
	if id == nil {
		id = itemWithContext.UniqueKey()
	}
	s.AddRecords(new(record.Record).
		WithId(fmt.Sprintf("%v", id)).
		WithSaveTime(time.Now()).
		WithNode(itemWithContext.GetCrawlerId()).
		WithSpider(itemWithContext.GetSpiderName()).
		WithSchedule(itemWithContext.GetScheduleId()).
		WithTask(itemWithContext.GetTaskId()).
		WithMeta(itemWithContext.MetaJson()).
		WithData(itemWithContext.DataJson()),
	)
	s.Nodes[itemWithContext.GetCrawlerId()].IncRecord()
	s.Spiders[itemWithContext.GetSpiderName()].IncRecord()
	s.Schedules[itemWithContext.GetScheduleId()].IncRecord()
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
	s.Schedules = make(map[string]pkg.StatisticsSchedule)
	s.Tasks = make(map[string]pkg.StatisticsTask)

	signal := s.crawler.GetSignal()
	signal.RegisterCrawlerStarted(s.crawlerStarted)
	signal.RegisterCrawlerStopped(s.crawlerStopped)
	signal.RegisterSpiderStarted(s.spiderStarted)
	signal.RegisterSpiderStopped(s.spiderStopped)
	signal.RegisterScheduleStarted(s.scheduleStarted)
	signal.RegisterScheduleStopped(s.scheduleStopped)
	signal.RegisterTaskStarted(s.taskStarted)
	signal.RegisterTaskStopped(s.taskStopped)
	signal.RegisterItemSaved(s.itemSaved)

	return s
}
