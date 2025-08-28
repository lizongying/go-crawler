package spider

import (
	"github.com/lizongying/go-crawler/pkg"
	crawlerContext "github.com/lizongying/go-crawler/pkg/context"
	kafka2 "github.com/lizongying/go-crawler/pkg/scheduler/kafka"
	"github.com/lizongying/go-crawler/pkg/scheduler/memory"
	redis2 "github.com/lizongying/go-crawler/pkg/scheduler/redis"
	"github.com/lizongying/go-crawler/pkg/stats"
	"time"
)

type Task struct {
	context        pkg.Context
	method         *pkg.State
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
func (t *Task) WithContext(ctx pkg.Context) pkg.Task {
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
	id = t.crawler.NextId()
	if t.GetContext() == nil {
		t.WithContext(new(crawlerContext.Context).
			WithCrawler(ctx.GetCrawler()).
			WithSpider(ctx.GetSpider()).
			WithJob(ctx.GetJob()).
			WithTask(new(crawlerContext.Task).
				WithTask(t).
				WithContext(ctx.GetJob().GetContext()).
				WithId(id).
				WithJobSubId(ctx.GetJob().GetSubId()).
				WithStatus(pkg.TaskStatusPending).
				WithStartTime(time.Now()).
				WithStats(new(stats.MediaStats))))
		t.crawler.GetSignal().TaskChanged(t)
	}

	if err = t.StartScheduler(t); err != nil {
		t.logger.Error(err)
		return
	}

	go func() {
		select {
		case <-t.context.GetTask().GetContext().Done():
			if t.context.GetTask().GetStatus() < pkg.TaskStatusSuccess {
				t.stop(t.context.GetTask().GetContext().Err())
			}
			return
		}
	}()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				t.logger.Error(r)
			}
			t.MethodOut()
		}()

		t.MethodIn()
		job := t.context.GetJob()
		var startFunc pkg.StartFunc
		startFunc, err = t.spider.StartFunc(job.GetFunc())
		if err != nil {
			t.logger.Error(err)
			return
		}
		if err = startFunc(t.context, job.GetArgs()); err != nil {
			t.logger.Error(err)
		}
	}()

	return
}
func (t *Task) stop(err error) {
	_ = t.StopScheduler(t)

	if err != nil {
		t.context.GetTask().WithStopReason(err.Error())
		t.context.GetTask().WithStatus(pkg.TaskStatusFailure)
	} else {
		t.context.GetTask().WithStatus(pkg.TaskStatusSuccess)
	}
	t.crawler.GetSignal().TaskChanged(t)
	t.job.TaskStopped(t.context, err)
	return
}
func (t *Task) MethodIn() {
	if !t.method.IsReady() {
		t.method.BeReady()
	}
	t.method.In()
}
func (t *Task) MethodOut() {
	t.method.Out()
}
func (t *Task) RequestIn() {
	if !t.request.IsReady() {
		t.request.BeReady()
	}
	t.request.In()
}
func (t *Task) RequestOut() {
	t.request.Out()
}
func (t *Task) ItemIn() {
	if !t.item.IsReady() {
		t.item.BeReady()
	}
	t.item.In()
}
func (t *Task) ItemOut() {
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
		method:  pkg.NewState("method"),
		request: pkg.NewState("request"),
		item:    pkg.NewState("item"),
	}

	t.requestAndItem = pkg.NewMultiState(t.request, t.item, t.method)

	t.requestAndItem.RegisterIsReadyAndIsZero(func() {
		t.stop(nil)
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
