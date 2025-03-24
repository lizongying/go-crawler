package statistics

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/queue"
	statisticsCrawler "github.com/lizongying/go-crawler/pkg/statistics/crawler"
	statisticsItem "github.com/lizongying/go-crawler/pkg/statistics/item"
	statisticsJob "github.com/lizongying/go-crawler/pkg/statistics/job"
	statisticsRequest "github.com/lizongying/go-crawler/pkg/statistics/request"
	statisticsSpider "github.com/lizongying/go-crawler/pkg/statistics/spider"
	statisticsTask "github.com/lizongying/go-crawler/pkg/statistics/task"
	"github.com/lizongying/go-crawler/pkg/utils"
	"os"
	"sync"
)

type Statistics struct {
	crawler  pkg.Crawler
	logger   pkg.Logger
	Crawlers map[string]pkg.StatisticsCrawler
	Spiders  map[string]pkg.StatisticsSpider
	Jobs     map[string]pkg.StatisticsJob
	Tasks    *queue.GroupQueue
	mutex    sync.RWMutex
}

func (s *Statistics) GetCrawlers() (nodes []pkg.StatisticsCrawler) {
	for _, v := range s.Crawlers {
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
		tasks = append(tasks, v.Value().(statisticsTask.WithItems).Task)
	}
	return
}
func (s *Statistics) GetRequests() (records []pkg.StatisticsRequest) {
	for _, v := range s.Tasks.Get("") {
		for _, v1 := range v.Value().(statisticsTask.WithItems).Requests.Get("") {
			records = append(records, v1.Value().(pkg.StatisticsRequest))
		}
	}
	return
}
func (s *Statistics) GetItems() (records []pkg.StatisticsItem) {
	for _, v := range s.Tasks.Get("") {
		for _, v1 := range v.Value().(statisticsTask.WithItems).Items.Get("") {
			records = append(records, v1.Value().(pkg.StatisticsItem))
		}
	}
	return
}
func (s *Statistics) crawlerChanged(ctx pkg.Context) (err error) {
	if _, ok := s.Crawlers[ctx.GetCrawler().GetId()]; !ok {
		hostname, _ := os.Hostname()
		s.Crawlers[ctx.GetCrawler().GetId()] = &statisticsCrawler.Crawler{
			Hostname: hostname,
			Ip:       utils.LanIp(),
			Enable:   true,
		}
		s.Crawlers[ctx.GetCrawler().GetId()].WithId(ctx.GetCrawler().GetId())
	}
	s.Crawlers[ctx.GetCrawler().GetId()].
		WithStatusAndTime(ctx.GetCrawler().GetStatus(), ctx.GetCrawler().GetUpdateTime())
	return
}
func (s *Statistics) spiderChanged(ctx pkg.Context) (err error) {
	spiderOne, ok := s.Spiders[ctx.GetSpider().GetName()]
	if !ok {
		s.Crawlers[ctx.GetCrawler().GetId()].IncSpider()
		var funcs []string
		for k1 := range ctx.GetSpider().GetSpider().StartFuncs() {
			funcs = append(funcs, k1)
		}
		spiderOne = new(statisticsSpider.Spider).
			WithId(ctx.GetSpider().GetId()).
			WithSpider(ctx.GetSpider().GetName()).
			WithFuncs(funcs).
			WithCrawler(ctx.GetCrawler().GetId())
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
		s.Crawlers[ctx.GetCrawler().GetId()].IncJob()
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
		s.Jobs[id] = new(statisticsJob.Job).
			WithId(id).
			WithEnable(j.GetEnable()).
			WithCrawler(ctx.GetCrawler().GetId()).
			WithSpider(ctx.GetSpider().GetName()).
			WithSchedule(spec).
			WithCommand(command)
	}

	s.Jobs[id].
		WithStatusAndTime(j.GetStatus(), j.GetUpdateTime())

	switch j.GetStatus() {
	case pkg.JobStatusRunning:
	case pkg.JobStatusFailure:
		s.Jobs[id].WithStopReason(j.GetStopReason())
	}
	return
}
func (s *Statistics) taskChanged(task pkg.Task) (err error) {
	defer s.mutex.Unlock()
	s.mutex.Lock()

	ctx := task.GetContext()

	if ctx.GetTask().GetStatus() == pkg.TaskStatusPending {
		s.Crawlers[ctx.GetCrawler().GetId()].IncTask()
		s.Spiders[ctx.GetSpider().GetName()].IncTask()
		s.Jobs[ctx.GetJob().GetId()].IncTask()

		// task
		s.Tasks.Enqueue(ctx.GetJob().GetId(),
			statisticsTask.WithItems{
				Task: new(statisticsTask.Task).
					WithId(ctx.GetTask().GetId()).
					WithCrawler(ctx.GetCrawler().GetId()).
					WithSpider(ctx.GetSpider().GetName()).
					WithJob(ctx.GetJob().GetId()).
					WithStatus(ctx.GetTask().GetStatus()).
					WithStartTime(ctx.GetTask().GetStartTime()),
				Items:    queue.NewGroupQueue(10),
				Requests: queue.NewGroupQueue(10),
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
		t := v.Value().(statisticsTask.WithItems).Task
		if ctx.GetTask().GetId() == t.GetId() {
			t.WithStatus(ctx.GetTask().GetStatus())
			t.WithUpdateTime(ctx.GetTask().GetUpdateTime())
			switch ctx.GetTask().GetStatus() {
			case pkg.TaskStatusRunning:
				t.WithStartTime(ctx.GetTask().GetStartTime())
			case pkg.TaskStatusSuccess:
				t.WithFinishTime(ctx.GetTask().GetStopTime())
			case pkg.TaskStatusFailure:
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
func (s *Statistics) requestChanged(request pkg.Request) (err error) {
	defer s.mutex.Unlock()
	s.mutex.Lock()

	ctx := request.GetContext()
	contextRequest := ctx.GetRequest()

	// task
	for _, v := range s.Tasks.Get(ctx.GetJob().GetId()) {
		t := v.Value().(statisticsTask.WithItems)
		if ctx.GetTask().GetId() == t.Task.GetId() {
			var sr *statisticsRequest.Request
			for _, v1 := range t.Requests.Get(ctx.GetTask().GetId()) {
				r := v1.Value().(*statisticsRequest.Request)
				if r.Id == contextRequest.GetId() {
					sr = r
					break
				}
			}

			if sr == nil {
				s.Crawlers[ctx.GetCrawler().GetId()].IncRequest()
				s.Spiders[ctx.GetSpider().GetName()].IncRequest()
				s.Jobs[ctx.GetJob().GetId()].IncRequest()

				// task
				t.Task.IncRequest()

				dataRequest, _ := request.Marshal()

				//request
				sr = new(statisticsRequest.Request).
					WithId(contextRequest.GetId()).
					WithUniqueKey(request.GetUniqueKey()).
					WithCrawler(ctx.GetCrawler().GetId()).
					WithSpider(ctx.GetSpider().GetName()).
					WithJob(ctx.GetJob().GetId()).
					WithTask(ctx.GetTask().GetId()).
					WithMeta(request.GetExtra()).
					WithData(string(dataRequest))

				t.Requests.Enqueue(ctx.GetTask().GetId(),
					sr,
					contextRequest.GetUpdateTime().UnixNano())
			}
			sr.WithUpdateTime(contextRequest.GetUpdateTime())
			sr.WithStatus(contextRequest.GetStatus())
			switch contextRequest.GetStatus() {
			case pkg.RequestStatusPending:
			case pkg.RequestStatusRunning:
				sr.WithStartTime(contextRequest.GetStartTime())
			case pkg.RequestStatusSuccess:
				sr.WithFinishTime(contextRequest.GetStopTime())
			case pkg.RequestStatusFailure:
				sr.WithFinishTime(contextRequest.GetStopTime())
			}
			break
		}
	}
	return
}
func (s *Statistics) itemChanged(item pkg.Item) (err error) {
	defer s.mutex.Unlock()
	s.mutex.Lock()

	ctx := item.GetContext()
	contextItem := ctx.GetItem()

	// task
	for _, v := range s.Tasks.Get(ctx.GetJob().GetId()) {
		t := v.Value().(statisticsTask.WithItems)
		if ctx.GetTask().GetId() == t.Task.GetId() {
			var statisticsRecord *statisticsItem.Item
			for _, v1 := range t.Items.Get(ctx.GetTask().GetId()) {
				r := v1.Value().(*statisticsItem.Item)
				if r.Id == contextItem.GetId() {
					statisticsRecord = r
					break
				}
			}

			if statisticsRecord == nil {
				s.Crawlers[ctx.GetCrawler().GetId()].IncItem()
				s.Spiders[ctx.GetSpider().GetName()].IncItem()
				s.Jobs[ctx.GetJob().GetId()].IncItem()

				// task
				t.Task.IncItem()

				//record
				id := item.Id()
				if id == nil {
					id = item.UniqueKey()
				}

				statisticsRecord = new(statisticsItem.Item).
					WithId(contextItem.GetId()).
					WithUniqueKey(fmt.Sprintf("%v", id)).
					WithCrawler(ctx.GetCrawler().GetId()).
					WithSpider(ctx.GetSpider().GetName()).
					WithJob(ctx.GetJob().GetId()).
					WithTask(ctx.GetTask().GetId()).
					WithMeta(item.MetaJson()).
					WithData(item.DataJson())

				t.Items.Enqueue(ctx.GetTask().GetId(),
					statisticsRecord,
					contextItem.GetUpdateTime().UnixNano())
			}
			statisticsRecord.WithUpdateTime(contextItem.GetUpdateTime())
			statisticsRecord.WithStatus(contextItem.GetStatus())
			switch contextItem.GetStatus() {
			case pkg.ItemStatusPending:
			case pkg.ItemStatusRunning:
				statisticsRecord.WithStartTime(contextItem.GetStartTime())
			case pkg.ItemStatusSuccess:
				statisticsRecord.WithFinishTime(contextItem.GetStopTime())
			case pkg.ItemStatusFailure:
				statisticsRecord.WithFinishTime(contextItem.GetStopTime())
			}
			break
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

	s.Crawlers = make(map[string]pkg.StatisticsCrawler)
	s.Spiders = make(map[string]pkg.StatisticsSpider)
	s.Jobs = make(map[string]pkg.StatisticsJob)
	s.Tasks = queue.NewGroupQueue(10)

	signal := s.crawler.GetSignal()
	signal.RegisterCrawlerChanged(s.crawlerChanged)
	signal.RegisterSpiderChanged(s.spiderChanged)
	signal.RegisterJobChanged(s.jobChanged)
	signal.RegisterTaskChanged(s.taskChanged)
	signal.RegisterRequestChanged(s.requestChanged)
	signal.RegisterItemChanged(s.itemChanged)
	return s
}
