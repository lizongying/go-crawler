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
func (s *Statistics) crawlerChanged(ctx pkg.Context) {
	if _, ok := s.Nodes[ctx.GetCrawlerId()]; !ok {
		hostname, _ := os.Hostname()
		s.Nodes[ctx.GetCrawlerId()] = &node.Node{
			Hostname: hostname,
			Ip:       utils.LanIp(),
			Enable:   true,
		}
		s.Nodes[ctx.GetCrawlerId()].WithId(ctx.GetCrawlerId())
	}
	s.Nodes[ctx.GetCrawlerId()].
		WithStatusAndTime(ctx.GetCrawlerStatus(), ctx.GetCrawlerUpdateTime())
}
func (s *Statistics) spiderChanged(ctx pkg.Context) {
	spiderOne, ok := s.Spiders[ctx.GetSpiderName()]
	if !ok {
		s.Nodes[ctx.GetCrawlerId()].IncSpider()
		var funcs []string
		for k1, _ := range ctx.GetSpider().GetSpider().StartFuncs() {
			funcs = append(funcs, k1)
		}
		spiderOne = new(statisticsSpider.Spider).
			WithId(ctx.GetSpider().GetId()).
			WithSpider(ctx.GetSpiderName()).
			WithFuncs(funcs).
			WithNode(ctx.GetCrawlerId())
		s.Spiders[ctx.GetSpiderName()] = spiderOne
	}
	spiderOne.
		WithStatusAndTime(ctx.GetSpiderStatus(), ctx.GetSpiderUpdateTime())
}
func (s *Statistics) jobChanged(ctx pkg.Context) {
	_, ok := s.Jobs[ctx.GetJobId()]
	if !ok {
		s.Nodes[ctx.GetCrawlerId()].IncJob()
		s.Spiders[ctx.GetSpiderName()].IncJob()

		var spec string
		mode := ctx.GetJobMode()
		switch ctx.GetJobMode() {
		case pkg.JobModeOnce:
			spec = "once"
		case pkg.JobModeLoop:
			spec = "loop"
		case pkg.JobModeCron:
			spec = fmt.Sprintf("cron (every %s)", ctx.GetJobSpec())
		}

		command := fmt.Sprintf("-n %s -f %s -m %s -s %s -a %s",
			ctx.GetSpiderName(),
			ctx.GetJobFunc(),
			(&mode).String(),
			ctx.GetJobSpec(),
			ctx.GetJobArgs(),
		)
		s.Jobs[ctx.GetJobId()] = new(job.Job).
			WithId(ctx.GetJobId()).
			WithEnable(ctx.GetJobEnable()).
			WithNode(ctx.GetCrawlerId()).
			WithSpider(ctx.GetSpiderName()).
			WithSchedule(spec).
			WithCommand(command)
	}

	s.Jobs[ctx.GetJobId()].
		WithStatusAndTime(ctx.GetJobStatus(), ctx.GetJobUpdateTime())
}
func (s *Statistics) taskChanged(ctx pkg.Context) {
	defer s.mutex.Unlock()
	s.mutex.Lock()

	if ctx.GetTaskStatus() == pkg.TaskStatusPending {
		s.Nodes[ctx.GetCrawlerId()].IncTask()
		s.Spiders[ctx.GetSpiderName()].IncTask()
		s.Jobs[ctx.GetJobId()].IncTask()

		// task
		s.Tasks.Enqueue(ctx.GetJobId(),
			task.WithRecords{
				Task: new(task.Task).
					WithId(ctx.GetTaskId()).
					WithNode(ctx.GetCrawlerId()).
					WithSpider(ctx.GetSpiderName()).
					WithJob(ctx.GetJobId()).
					WithStatus(ctx.GetTaskStatus()).
					WithStartTime(ctx.GetTaskStartTime()),
				Records: queue.NewGroupQueue(10),
			},
			ctx.GetTaskStartTime().UnixNano())

		// spider
		s.Spiders[ctx.GetSpiderName()].
			WithLastTaskId(ctx.GetTaskId()).
			WithLastTaskStatus(ctx.GetTaskStatus()).
			WithLastTaskStartTime(ctx.GetTaskStartTime())

		return
	}

	// task
	for _, v := range s.Tasks.Get(ctx.GetJobId()) {
		t := v.Value().(task.WithRecords).Task
		if ctx.GetTaskId() == t.GetId() {
			t.WithStatus(ctx.GetTaskStatus()).
				WithFinishTime(ctx.GetTaskStopTime())
		}
	}

	// spider
	spider := s.Spiders[ctx.GetSpiderName()]
	if ctx.GetTaskId() == spider.GetLastTaskId() {
		spider.
			WithLastTaskStatus(ctx.GetTaskStatus()).
			WithLastTaskFinishTime(ctx.GetTaskStopTime())
	}
}
func (s *Statistics) itemChanged(item pkg.Item) {
	defer s.mutex.Unlock()
	s.mutex.Lock()

	s.Nodes[item.GetContext().GetCrawlerId()].IncRecord()
	s.Spiders[item.GetContext().GetSpiderName()].IncRecord()
	s.Jobs[item.GetContext().GetJobId()].IncRecord()

	ctx := item.GetContext()

	// task
	for _, v := range s.Tasks.Get(item.GetContext().GetJobId()) {
		t := v.Value().(task.WithRecords)
		if ctx.GetTaskId() == t.Task.GetId() {
			// task
			t.Task.IncRecord()

			//record
			id := item.Id()
			if id == nil {
				id = item.UniqueKey()
			}
			t.Records.Enqueue(ctx.GetTaskId(),
				new(record.Record).
					WithId(ctx.GetItemId()).
					WithUniqueKey(fmt.Sprintf("%v", id)).
					WithSaveTime(ctx.GetItemStopTime()).
					WithNode(ctx.GetCrawlerId()).
					WithSpider(ctx.GetSpiderName()).
					WithJob(ctx.GetJobId()).
					WithTask(ctx.GetTaskId()).
					WithMeta(item.MetaJson()).
					WithData(item.DataJson()),
				ctx.GetItemStopTime().UnixNano())
		}
	}
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
