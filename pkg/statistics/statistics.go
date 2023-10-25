package statistics

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/statistics/node"
	"github.com/lizongying/go-crawler/pkg/statistics/record"
	statisticsSpider "github.com/lizongying/go-crawler/pkg/statistics/spider"
	"github.com/lizongying/go-crawler/pkg/utils"
	"os"
	"time"
)

type Statistics struct {
	crawler   pkg.Crawler
	logger    pkg.Logger
	Nodes     []pkg.StatisticsNode
	Spiders   []pkg.StatisticsSpider
	Schedules []pkg.StatisticsSchedule
	Tasks     []pkg.StatisticsTask
	Records   []pkg.StatisticsRecord
}

func (s *Statistics) GetNodes() []pkg.StatisticsNode {
	return s.Nodes
}
func (s *Statistics) GetSpiders() []pkg.StatisticsSpider {
	return s.Spiders
}
func (s *Statistics) GetSchedules() []pkg.StatisticsSchedule {
	return s.Schedules
}
func (s *Statistics) GetTasks() []pkg.StatisticsTask {
	return s.Tasks
}
func (s *Statistics) GetRecords() []pkg.StatisticsRecord {
	return s.Records
}

func (s *Statistics) AddNodes(nodes ...pkg.StatisticsNode) {
	s.Nodes = append(s.Nodes, nodes...)
}
func (s *Statistics) AddSpiders(spiders ...pkg.Spider) {
	signal := s.crawler.GetSignal()
	for _, v := range spiders {
		s.Spiders = append(s.Spiders, &statisticsSpider.Spider{
			Spider: v.Name(),
		})
		signal.RegisterSpiderStarted(v.Name(), s.spiderOpened)
		signal.RegisterSpiderStopped(v.Name(), s.spiderClosed)
		signal.RegisterItemSaved(s.itemSaved)
	}
}
func (s *Statistics) AddSchedules(schedules ...pkg.StatisticsSchedule) {
	s.Schedules = append(s.Schedules, schedules...)
}
func (s *Statistics) AddTasks(tasks ...pkg.StatisticsTask) {
	s.Tasks = append(s.Tasks, tasks...)
}
func (s *Statistics) AddRecords(records ...pkg.StatisticsRecord) {
	s.Records = append(s.Records, records...)
}
func (s *Statistics) spiderOpened(spider pkg.Spider) {
	//var sSpider pkg.StatisticsSpider
	//for _, v := range s.Spiders {
	//	if v.GetSpider() == spider.Name() {
	//		sSpider = v
	//		break
	//	}
	//}
	//s.Tasks = append(s.Tasks, &task.Task{
	//	Spider:  sSpider.(*statisticsSpider.Spider),
	//	Started: time.Now(),
	//})
	for _, v := range s.Nodes {
		v.(*node.Node).IncSpider()
	}
	for _, v := range s.Spiders {
		if v.GetSpider() == spider.Name() {
			v.(*statisticsSpider.Spider).
				WithLastRunAt(spider.GetContext().GetStartTime()).
				WithStatus(spider.GetContext().GetStatus())
			break
		}
	}
}

func (s *Statistics) spiderClosed(spider pkg.Spider) {
	//var sTask pkg.StatisticsTask
	//for _, v := range s.Tasks {
	//	if v.GetId() == spider.Name() {
	//		sTask = v
	//		break
	//	}
	//}
	//sTask.SetFinished(time.Now())

	for _, v := range s.Spiders {
		if v.GetSpider() == spider.Name() {
			v.(*statisticsSpider.Spider).
				WithLastFinishAt(spider.GetContext().GetStopTime())
			break
		}
	}
}
func (s *Statistics) taskOpened(spider pkg.Spider) {
	//var sSpider pkg.StatisticsSpider
	//for _, v := range s.Spiders {
	//	if v.GetSpider() == spider.Name() {
	//		sSpider = v
	//		break
	//	}
	//}
	//s.Tasks = append(s.Tasks, &task.Task{
	//	Spider:  sSpider.(*statisticsSpider.Spider),
	//	Started: time.Now(),
	//})
}

func (s *Statistics) taskClosed(spider pkg.Spider) {
	//var sTask pkg.StatisticsTask
	//for _, v := range s.Tasks {
	//	if v.GetId() == spider.Name() {
	//		sTask = v
	//		break
	//	}
	//}
	//sTask.SetFinished(time.Now())
}
func (s *Statistics) itemSaved(itemWithContext pkg.ItemWithContext) {
	s.AddRecords(new(record.Record).
		WithSaveTime(time.Now()).
		WithSpider(itemWithContext.GetSpiderName()).
		WithTaskId(itemWithContext.GetTaskId()).
		WithMeta(itemWithContext.MetaJson()).
		WithData(itemWithContext.DataJson()),
	)
	for _, v := range s.Nodes {
		v.(*node.Node).IncTask()
	}
	for _, v := range s.Spiders {
		if v.GetSpider() == itemWithContext.GetSpiderName() {
			v.(*statisticsSpider.Spider).IncRecord()
			break
		}
	}
}
func (s *Statistics) FromCrawler(crawler pkg.Crawler) pkg.Statistics {
	if s == nil {
		return new(Statistics).FromCrawler(crawler)
	}

	s.crawler = crawler
	s.logger = crawler.GetLogger()

	hostname, _ := os.Hostname()
	s.AddNodes((&node.Node{
		Status:   node.StatusOnline,
		Hostname: hostname,
		Ip:       utils.LanIp(),
		Enable:   true,
	}).WithStartTime(time.Now()))

	return s
}
