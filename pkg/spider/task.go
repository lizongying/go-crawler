package spider

import (
	"context"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	crawlerContext "github.com/lizongying/go-crawler/pkg/context"
	kafka2 "github.com/lizongying/go-crawler/pkg/scheduler/kafka"
	"github.com/lizongying/go-crawler/pkg/scheduler/memory"
	redis2 "github.com/lizongying/go-crawler/pkg/scheduler/redis"
	"github.com/lizongying/go-crawler/pkg/stats"
	"github.com/lizongying/go-crawler/pkg/utils"
	"reflect"
	"time"
)

type Task struct {
	context        pkg.Context
	request        *pkg.State
	item           *pkg.State
	requestAndItem *pkg.MultiState
	crawler        pkg.Crawler
	spider         pkg.Spider
	logger         pkg.Logger
	job            *Job
	pkg.Stats
	pkg.Scheduler
}

func (t *Task) GetContext() pkg.Context {
	return t.context
}
func (t *Task) WithContext(ctx pkg.Context) *Task {
	t.context = ctx
	return t
}
func (t *Task) GetScheduler() pkg.Scheduler {
	return t.Scheduler
}
func (t *Task) WithScheduler(scheduler pkg.Scheduler) pkg.Task {
	t.Scheduler = scheduler
	return t
}
func (t *Task) start(ctx pkg.Context) (id string, err error) {
	id = utils.UUIDV1WithoutHyphens()
	if t.GetContext() == nil {
		t.WithContext(new(crawlerContext.Context).
			WithCrawler(ctx.GetCrawler()).
			WithSpider(ctx.GetSpider()).
			WithSchedule(ctx.GetSchedule()).
			WithTask(new(crawlerContext.Task).
				WithTask(t).
				WithContext(context.Background()).
				WithId(id).
				WithStatus(pkg.TaskStatusPending).
				WithStartTime(time.Now()).
				WithStats(&stats.MediaStats{})))
		t.crawler.GetSignal().TaskStarted(t.context)
	}

	if err = t.StartScheduler(t.context); err != nil {
		t.logger.Error(err)
		return
	}

	go func() {

		defer func() {
			//if r := recover(); r != nil {
			//	s.logger.Error(r)
			//}
		}()

		t.logger.Info(t.spider.Name(), id, "task started")

		t.context.WithTaskStartTime(time.Now())
		t.context.WithTaskStatus(pkg.TaskStatusRunning)
		t.crawler.GetSignal().TaskStarted(t.context)

		params := []reflect.Value{
			reflect.ValueOf(t.context),
			reflect.ValueOf(t.context.GetScheduleArgs()),
		}
		caller := reflect.ValueOf(t.spider).MethodByName(t.context.GetScheduleFunc())
		if !caller.IsValid() {
			err = errors.New(fmt.Sprintf("schedule func is invalid: %s", t.context.GetScheduleFunc()))
			t.logger.Error(err)
			return
		}

		res := caller.Call(params)
		if len(res) != 1 {
			err = errors.New(fmt.Sprintf("%s has too many return values", t.context.GetScheduleFunc()))
			t.logger.Error(err)
			return
		}

		if res[0].Type().Name() != "error" {
			err = errors.New(fmt.Sprintf("%s should return an error", t.context.GetScheduleFunc()))
			t.logger.Error(err)
			return
		}

		if !res[0].IsNil() {
			err = res[0].Interface().(error)
			t.logger.Error(err)
			return
		}
	}()

	return
}
func (t *Task) stop() (err error) {
	if err = t.StopScheduler(t.context); err != nil {
		t.logger.Error(err)
		return
	}

	stopTime := time.Now()
	t.context.WithTaskStatus(pkg.TaskStatusSuccess)
	t.context.WithTaskStopTime(stopTime)
	t.crawler.GetSignal().TaskStopped(t.context)
	t.logger.Info(t.spider.Name(), t.context.GetTaskId(), "task finished. spend time:", stopTime.Sub(t.context.GetSpiderStartTime()))
	t.job.StopTask()
	return
}
func (t *Task) ReadyRequest() {
	t.request.BeReady()
}
func (t *Task) StartRequest() {
	t.request.In()
}
func (t *Task) StopRequest() {
	t.request.Out()
}
func (t *Task) ReadyItem() {
	t.item.BeReady()
}
func (t *Task) StartItem() {
	t.item.In()
}
func (t *Task) StopItem() {
	t.item.Out()
}
func (t *Task) WithJob(job *Job) *Task {
	t.job = job
	return t
}
func (t *Task) FromSpider(spider pkg.Spider) *Task {
	*t = Task{
		crawler: spider.GetCrawler(),
		spider:  spider,
		logger:  spider.GetLogger(),
		request: pkg.NewState(),
		item:    pkg.NewState(),
	}

	t.requestAndItem = pkg.NewMultiState(t.request, t.item)

	t.requestAndItem.RegisterIsReadyAndIsZero(func() {
		if err := t.stop(); err != nil {
			t.logger.Error(err)
		}
	})

	config := spider.GetConfig()

	switch config.GetScheduler() {
	case pkg.SchedulerMemory:
		t.WithScheduler(new(memory.Scheduler).FromSpider(spider))
	case pkg.SchedulerRedis:
		t.WithScheduler(new(redis2.Scheduler).FromSpider(spider))
	case pkg.SchedulerKafka:
		t.WithScheduler(new(kafka2.Scheduler).FromSpider(spider))
	default:
	}

	return t
}
