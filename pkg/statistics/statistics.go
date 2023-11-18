package statistics

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/queue"
	"github.com/lizongying/go-crawler/pkg/statistics/job"
	"github.com/lizongying/go-crawler/pkg/statistics/node"
	"github.com/lizongying/go-crawler/pkg/statistics/record"
	statisticsSpider "github.com/lizongying/go-crawler/pkg/statistics/spider"
	"github.com/lizongying/go-crawler/pkg/statistics/task"
	"github.com/lizongying/go-crawler/pkg/utils"
	"os"
	"sync"
)

type Statistics struct {
	crawler pkg.Crawler
	logger  pkg.Logger
	Nodes   map[string]pkg.StatisticsNode
	Spiders map[string]pkg.StatisticsSpider
	Jobs    map[string]pkg.StatisticsJob
	Tasks   *queue.GroupQueue
	mutex   sync.RWMutex
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
func (s *Statistics) GetJobs() (schedules []pkg.StatisticsJob) {
	for _, v := range s.Jobs {
		schedules = append(schedules, v)
	}
	return
}
func (s *Statistics) GetTasks() (tasks []pkg.StatisticsTask) {
	for _, v := range s.Tasks.Get("") {
		tasks = append(tasks, v.Value().(task.WithRecords).Task)
	}
	return
}
func (s *Statistics) GetRecords() (records []pkg.StatisticsRecord) {
	for _, v := range s.Tasks.Get("") {
		for _, v1 := range v.Value().(task.WithRecords).Records.Get("") {
			records = append(records, v1.Value().(pkg.StatisticsRecord))
		}
	}
	return
}
func (s *Statistics) crawlerChanged(ctx pkg.Context) (err error) {
	if _, ok := s.Nodes[ctx.GetCrawler().GetId()]; !ok {
		hostname, _ := os.Hostname()
		s.Nodes[ctx.GetCrawler().GetId()] = &node.Node{
			Hostname: hostname,
			Ip:       utils.LanIp(),
			Enable:   true,
		}
		s.Nodes[ctx.GetCrawler().GetId()].WithId(ctx.GetCrawler().GetId())
	}
	s.Nodes[ctx.GetCrawler().GetId()].
		WithStatusAndTime(ctx.GetCrawler().GetStatus(), ctx.GetCrawler().GetUpdateTime())
	return
}
func (s *Statistics) spiderChanged(ctx pkg.Context) (err error) {
	spiderOne, ok := s.Spiders[ctx.GetSpider().GetName()]
	if !ok {
		s.Nodes[ctx.GetCrawler().GetId()].IncSpider()
		var funcs []string
		for k1, _ := range ctx.GetSpider().GetSpider().StartFuncs() {
			funcs = append(funcs, k1)
		}
		spiderOne = new(statisticsSpider.Spider).
			WithId(ctx.GetSpider().GetId()).
			WithSpider(ctx.GetSpider().GetName()).
			WithFuncs(funcs).
			WithNode(ctx.GetCrawler().GetId())
		s.Spiders[ctx.GetSpider().GetName()] = spiderOne
	}
	spiderOne.
		WithStatusAndTime(ctx.GetSpider().GetStatus(), ctx.GetSpider().GetUpdateTime())
	return
}
func (s *Statistics) jobChanged(ctx pkg.Context) (err error) {
	j := ctx.GetJob()
	id := j.GetId()

	_, ok := s.Jobs[id]
	if !ok {
		s.Nodes[ctx.GetCrawler().GetId()].IncJob()
		s.Spiders[ctx.GetSpider().GetName()].IncJob()

		var spec string
		mode := j.GetMode()
		switch j.GetMode() {
		case pkg.JobModeOnce:
			spec = "once"
		case pkg.JobModeLoop:
			spec = "loop"
		case pkg.JobModeCron:
			spec = fmt.Sprintf("cron (every %s)", j.GetSpec())
		}

		command := fmt.Sprintf("go-crawler -n %s -f %s -m %s -s %s -a '%s'",
			ctx.GetSpider().GetName(),
			j.GetFunc(),
			(&mode).String(),
			j.GetSpec(),
			j.GetArgs(),
		)
		s.Jobs[id] = new(job.Job).
			WithId(id).
			WithEnable(j.GetEnable()).
			WithNode(ctx.GetCrawler().GetId()).
			WithSpider(ctx.GetSpider().GetName()).
			WithSchedule(spec).
			WithCommand(command)
	}

	s.Jobs[id].
		WithStatusAndTime(j.GetStatus(), j.GetUpdateTime())

	switch j.GetStatus() {
	case pkg.JobStatusRunning:
	case pkg.JobStatusStopped:
		s.Jobs[id].WithStopReason(j.GetStopReason())
	}
	return
}
func (s *Statistics) taskChanged(ctx pkg.Context) (err error) {
	defer s.mutex.Unlock()
	s.mutex.Lock()

	if ctx.GetTask().GetStatus() == pkg.TaskStatusPending {
		s.Nodes[ctx.GetCrawler().GetId()].IncTask()
		s.Spiders[ctx.GetSpider().GetName()].IncTask()
		s.Jobs[ctx.GetJob().GetId()].IncTask()

		// task
		s.Tasks.Enqueue(ctx.GetJob().GetId(),
			task.WithRecords{
				Task: new(task.Task).
					WithId(ctx.GetTask().GetId()).
					WithNode(ctx.GetCrawler().GetId()).
					WithSpider(ctx.GetSpider().GetName()).
					WithJob(ctx.GetJob().GetId()).
					WithStatus(ctx.GetTask().GetStatus()).
					WithStartTime(ctx.GetTask().GetStartTime()),
				Records: queue.NewGroupQueue(10),
			},
			ctx.GetTask().GetStartTime().UnixNano())

		// spider
		s.Spiders[ctx.GetSpider().GetName()].
			WithLastTaskId(ctx.GetTask().GetId()).
			WithLastTaskStatus(ctx.GetTask().GetStatus()).
			WithLastTaskStartTime(ctx.GetTask().GetStartTime())

		return
	}

	// task
	for _, v := range s.Tasks.Get(ctx.GetJob().GetId()) {
		t := v.Value().(task.WithRecords).Task
		if ctx.GetTask().GetId() == t.GetId() {
			t.WithStatus(ctx.GetTask().GetStatus())
			t.WithUpdateTime(ctx.GetTask().GetUpdateTime())
			switch ctx.GetTask().GetStatus() {
			case pkg.TaskStatusRunning:
				t.WithStartTime(ctx.GetTask().GetStartTime())
			case pkg.TaskStatusSuccess:
				t.WithFinishTime(ctx.GetTask().GetStopTime())
			case pkg.TaskStatusError:
				t.WithFinishTime(ctx.GetTask().GetStopTime())
				t.WithStopReason(ctx.GetTask().GetStopReason())
			}
		}
	}

	// spider
	spider := s.Spiders[ctx.GetSpider().GetName()]
	if ctx.GetTask().GetId() == spider.GetLastTaskId() {
		spider.
			WithLastTaskStatus(ctx.GetTask().GetStatus()).
			WithLastTaskFinishTime(ctx.GetTask().GetStopTime())
	}
	return
}
func (s *Statistics) itemChanged(item pkg.Item) (err error) {
	defer s.mutex.Unlock()
	s.mutex.Lock()

	s.Nodes[item.GetContext().GetCrawler().GetId()].IncRecord()
	s.Spiders[item.GetContext().GetSpider().GetName()].IncRecord()
	s.Jobs[item.GetContext().GetJob().GetId()].IncRecord()

	ctx := item.GetContext()

	// task
	for _, v := range s.Tasks.Get(item.GetContext().GetJob().GetId()) {
		t := v.Value().(task.WithRecords)
		if ctx.GetTask().GetId() == t.Task.GetId() {
			// task
			t.Task.IncRecord()

			//record
			id := item.Id()
			if id == nil {
				id = item.UniqueKey()
			}
			t.Records.Enqueue(ctx.GetTask().GetId(),
				new(record.Record).
					WithId(ctx.GetItem().GetId()).
					WithUniqueKey(fmt.Sprintf("%v", id)).
					WithSaveTime(ctx.GetItem().GetStopTime()).
					WithNode(ctx.GetCrawler().GetId()).
					WithSpider(ctx.GetSpider().GetName()).
					WithJob(ctx.GetJob().GetId()).
					WithTask(ctx.GetTask().GetId()).
					WithMeta(item.MetaJson()).
					WithData(item.DataJson()),
				ctx.GetItem().GetStopTime().UnixNano())
		}
	}
	return
}
func (s *Statistics) FromCrawler(crawler pkg.Crawler) pkg.Statistics {
	if s == nil {
		return new(Statistics).FromCrawler(crawler)
	}

	s.crawler = crawler
	s.logger = crawler.GetLogger()

	s.Nodes = make(map[string]pkg.StatisticsNode)
	s.Spiders = make(map[string]pkg.StatisticsSpider)
	s.Jobs = make(map[string]pkg.StatisticsJob)
	s.Tasks = queue.NewGroupQueue(10)

	signal := s.crawler.GetSignal()
	signal.RegisterCrawlerChanged(s.crawlerChanged)
	signal.RegisterSpiderChanged(s.spiderChanged)
	signal.RegisterJobChanged(s.jobChanged)
	signal.RegisterTaskChanged(s.taskChanged)
	signal.RegisterItemChanged(s.itemChanged)

	return s
}
