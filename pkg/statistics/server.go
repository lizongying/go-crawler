package statistics

import (
	"context"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	pb "github.com/lizongying/go-crawler/pkg/api/proto"
	"github.com/lizongying/go-crawler/pkg/queue"
	statisticsCrawler "github.com/lizongying/go-crawler/pkg/statistics/crawler"
	statisticsItem "github.com/lizongying/go-crawler/pkg/statistics/item"
	statisticsJob "github.com/lizongying/go-crawler/pkg/statistics/job"
	statisticsRequest "github.com/lizongying/go-crawler/pkg/statistics/request"
	statisticsTask "github.com/lizongying/go-crawler/pkg/statistics/task"
	"github.com/lizongying/go-crawler/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"os"
)

type Server struct {
	pb.UnimplementedStatisticsServer
	statistics *Statistics
}

func (s *Server) CrawlerChanged(ctx context.Context, crawler *pb.Crawler) (response *pb.Response, err error) {
	if _, ok := s.statistics.Crawlers[crawler.GetId()]; !ok {
		hostname, _ := os.Hostname()
		s.statistics.Crawlers[crawler.GetId()] = &statisticsCrawler.Crawler{
			Hostname: hostname,
			Ip:       utils.LanIp(),
			Enable:   true,
		}
		s.statistics.Crawlers[crawler.GetId()].WithId(crawler.GetId())
	}
	s.statistics.Crawlers[crawler.GetId()].
		WithStatusAndTime(pkg.CrawlerStatus(crawler.GetStatus()), crawler.GetUpdateTime().AsTime())

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		log.Printf("metadata: %v", md)
	}

	response = &pb.Response{Message: "Hello "}
	return
}

func (s *Server) jobChanged(_ context.Context, job *pb.Job) (response *pb.Response, err error) {
	id := job.GetId()

	ctx := job.GetContext()

	_, ok := s.statistics.Jobs[id]
	if !ok {
		s.statistics.Crawlers[ctx.GetCrawler().GetId()].IncJob()
		s.statistics.Spiders[ctx.GetSpider().GetName()].IncJob()

		var spec string
		mode := job.GetMode()
		switch job.GetMode() {
		case pkg.JobModeOnce:
			spec = "once"
		case pkg.JobModeLoop:
			spec = "loop"
		case pkg.JobModeCron:
			spec = fmt.Sprintf("cron (every %s)", job.GetSpec())
		}

		command := fmt.Sprintf("go-crawler -n %s -f %s -m %s -s %s -a '%s'",
			ctx.GetSpider().GetName(),
			job.GetFunc(),
			(&mode).String(),
			job.GetSpec(),
			job.GetArgs(),
		)
		s.statistics.Jobs[id] = new(statisticsJob.Job).
			WithCrawler(ctx.GetCrawler().GetId()).
			WithSpider(ctx.GetSpider().GetName()).
			WithId(id).
			WithEnable(job.GetEnable()).
			WithSchedule(spec).
			WithCommand(command)
	}

	s.statistics.Jobs[id].
		WithStatusAndTime(pkg.JobStatus(job.GetStatus()), job.GetUpdateTime().AsTime())

	switch job.GetStatus() {
	case pb.JobStatus_JobStatusRunning:
	case pb.JobStatus_JobStatusFailure:
		s.statistics.Jobs[id].WithStopReason(job.GetStopReason())
	}
	return
}

func (s *Server) taskChanged(_ context.Context, task *pb.Task) (response *pb.Response, err error) {
	defer s.statistics.mutex.Unlock()
	s.statistics.mutex.Lock()

	ctx := task.GetContext()

	if task.GetStatus() == pkg.TaskStatusPending {
		s.statistics.Crawlers[ctx.GetCrawler().GetId()].IncTask()
		s.statistics.Spiders[ctx.GetSpider().GetName()].IncTask()
		s.statistics.Jobs[ctx.GetJob().GetId()].IncTask()

		// task
		s.statistics.Tasks.Enqueue(ctx.GetJob().GetId(),
			statisticsTask.WithItems{
				Task: new(statisticsTask.Task).
					WithCrawler(ctx.GetCrawler().GetId()).
					WithSpider(ctx.GetSpider().GetName()).
					WithJob(ctx.GetJob().GetId()).
					WithId(task.GetId()).
					WithStatus(pkg.TaskStatus(task.GetStatus())).
					WithStartTime(task.GetStartTime().AsTime()),
				Items:    queue.NewGroupQueue(10),
				Requests: queue.NewGroupQueue(10),
			},
			task.GetStartTime().AsTime().UnixNano())

		// spider
		s.statistics.Spiders[ctx.GetSpider().GetName()].
			WithLastTaskId(task.GetId()).
			WithLastTaskStatus(pkg.TaskStatus(task.GetStatus())).
			WithLastTaskStartTime(task.GetStartTime().AsTime())

		return
	}

	// task
	for _, v := range s.statistics.Tasks.Get(ctx.GetJob().GetId()) {
		t := v.Value().(statisticsTask.WithItems).Task
		if task.GetId() == t.GetId() {
			t.WithStatus(pkg.TaskStatus(task.GetStatus()))
			t.WithUpdateTime(task.GetUpdateTime().AsTime())
			switch task.GetStatus() {
			case pb.TaskStatus_TaskStatusRunning:
				t.WithStartTime(task.GetStartTime().AsTime())
			case pb.TaskStatus_TaskStatusSuccess:
				t.WithFinishTime(task.GetStopTime().AsTime())
			case pb.TaskStatus_TaskStatusFailure:
				t.WithFinishTime(task.GetStopTime().AsTime())
				t.WithStopReason(task.GetStopReason())
			}
		}
	}

	// spider
	spider := s.statistics.Spiders[ctx.GetSpider().GetName()]
	if task.GetId() == spider.GetLastTaskId() {
		spider.
			WithLastTaskStatus(pkg.TaskStatus(task.GetStatus())).
			WithLastTaskFinishTime(task.GetStopTime().AsTime())
	}
	return
}

func (s *Server) requestChanged(_ context.Context, request *pb.Request) (response *pb.Response, err error) {
	defer s.statistics.mutex.Unlock()
	s.statistics.mutex.Lock()

	ctx := request.GetContext()

	// task
	for _, v := range s.statistics.Tasks.Get(ctx.GetJob().GetId()) {
		t := v.Value().(statisticsTask.WithItems)
		if ctx.GetTask().GetId() == t.Task.GetId() {
			var sr *statisticsRequest.Request
			for _, v1 := range t.Requests.Get(ctx.GetTask().GetId()) {
				r := v1.Value().(*statisticsRequest.Request)
				if r.Id == request.GetId() {
					sr = r
					break
				}
			}

			if sr == nil {
				s.statistics.Crawlers[ctx.GetCrawler().GetId()].IncRequest()
				s.statistics.Spiders[ctx.GetSpider().GetName()].IncRequest()
				s.statistics.Jobs[ctx.GetJob().GetId()].IncRequest()

				// task
				t.Task.IncRequest()

				//request
				sr = new(statisticsRequest.Request).
					WithCrawler(ctx.GetCrawler().GetId()).
					WithSpider(ctx.GetSpider().GetName()).
					WithJob(ctx.GetJob().GetId()).
					WithTask(ctx.GetTask().GetId()).
					WithId(request.GetId()).
					WithUniqueKey(request.GetUniqueKey()).
					WithMeta(request.GetExtra()).
					WithData(request.GetData())

				t.Requests.Enqueue(ctx.GetTask().GetId(),
					sr,
					request.GetUpdateTime().AsTime().UnixNano())
			}
			sr.WithUpdateTime(request.GetUpdateTime().AsTime())
			sr.WithStatus(pkg.RequestStatus(request.GetStatus()))
			switch request.GetStatus() {
			case pb.RequestStatus_RequestStatusPending:
			case pb.RequestStatus_RequestStatusRunning:
				sr.WithStartTime(request.GetStartTime().AsTime())
			case pb.RequestStatus_RequestStatusSuccess:
				sr.WithFinishTime(request.GetStopTime().AsTime())
			case pb.RequestStatus_RequestStatusFailure:
				sr.WithFinishTime(request.GetStopTime().AsTime())
			}
			break
		}
	}
	return
}
func (s *Server) itemChanged(_ context.Context, item *pb.Item) (response *pb.Response, err error) {
	defer s.statistics.mutex.Unlock()
	s.statistics.mutex.Lock()

	ctx := item.GetContext()

	// task
	for _, v := range s.statistics.Tasks.Get(ctx.GetJob().GetId()) {
		t := v.Value().(statisticsTask.WithItems)
		if ctx.GetTask().GetId() == t.Task.GetId() {
			var statisticsRecord *statisticsItem.Item
			for _, v1 := range t.Items.Get(ctx.GetTask().GetId()) {
				r := v1.Value().(*statisticsItem.Item)
				if r.Id == item.GetId() {
					statisticsRecord = r
					break
				}
			}

			if statisticsRecord == nil {
				s.statistics.Crawlers[ctx.GetCrawler().GetId()].IncItem()
				s.statistics.Spiders[ctx.GetSpider().GetName()].IncItem()
				s.statistics.Jobs[ctx.GetJob().GetId()].IncItem()

				// task
				t.Task.IncItem()

				//record
				statisticsRecord = new(statisticsItem.Item).
					WithCrawler(ctx.GetCrawler().GetId()).
					WithSpider(ctx.GetSpider().GetName()).
					WithJob(ctx.GetJob().GetId()).
					WithTask(ctx.GetTask().GetId()).
					WithId(item.GetId()).
					WithUniqueKey(item.GetUniqueKey()).
					WithMeta(item.GetMeta()).
					WithData(item.GetData())

				t.Items.Enqueue(ctx.GetTask().GetId(),
					statisticsRecord,
					item.GetUpdateTime().AsTime().UnixNano())
			}
			statisticsRecord.WithUpdateTime(item.GetUpdateTime().AsTime())
			statisticsRecord.WithStatus(pkg.ItemStatus(item.GetStatus()))
			switch item.GetStatus() {
			case pb.ItemStatus_ItemStatusPending:
			case pb.ItemStatus_ItemStatusRunning:
				statisticsRecord.WithStartTime(item.GetStartTime().AsTime())
			case pb.ItemStatus_ItemStatusSuccess:
				statisticsRecord.WithFinishTime(item.GetStopTime().AsTime())
			case pb.ItemStatus_ItemStatusFailure:
				statisticsRecord.WithFinishTime(item.GetStopTime().AsTime())
			}
			break
		}
	}
	return
}

func NewServer(server *grpc.Server) (statisticsServer *Server) {
	statisticsServer = &Server{
		statistics: &Statistics{},
	}
	pb.RegisterStatisticsServer(server, statisticsServer)
	return
}
