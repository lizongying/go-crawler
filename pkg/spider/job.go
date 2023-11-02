package spider

import (
	"context"
	"github.com/lizongying/cron"
	"github.com/lizongying/go-crawler/pkg"
	crawlerContext "github.com/lizongying/go-crawler/pkg/context"
	"github.com/lizongying/go-crawler/pkg/utils"
	"time"
)

type Job struct {
	context pkg.Context
	task    *pkg.State
	crawler pkg.Crawler
	spider  pkg.Spider
	logger  pkg.Logger
	cronJob chan struct{}
}

func (j *Job) GetContext() pkg.Context {
	return j.context
}
func (j *Job) WithContext(ctx pkg.Context) *Job {
	j.context = ctx
	return j
}
func (j *Job) start(ctx pkg.Context, jobFunc string, args string, mode pkg.ScheduleMode, spec string, onlyOneTask bool, resultChan chan struct{}) (id string) {
	j.task.RegisterIsReadyAndIsZero(func() {
		j.stop(resultChan)
	})
	id = utils.UUIDV1WithoutHyphens()
	j.context = new(crawlerContext.Context).
		WithCrawler(ctx.GetCrawler()).
		WithSpider(ctx.GetSpider()).
		WithSchedule(new(crawlerContext.Schedule).
			WithContext(context.Background()).
			WithId(id).
			WithStatus(pkg.ScheduleStatusStarted).
			WithStartTime(time.Now()).
			WithEnable(true).
			WithFunc(jobFunc).
			WithArgs(args).
			WithMode(mode).
			WithSpec(spec))
	j.crawler.GetSignal().ScheduleStarted(j.context)

	switch mode {
	case pkg.ScheduleModeOnce:
		go j.startTask()
	case pkg.ScheduleModeLoop:
		go j.startTask()
	case pkg.ScheduleModeCron:
		if onlyOneTask {
			j.cronJob <- struct{}{}
		}
		cr := cron.New(cron.WithLogger(j.logger))
		cr.MustStart()
		job := new(cron.Job).
			MustEverySpec(spec).
			Callback(func() {
				if onlyOneTask {
					<-j.cronJob
				}
				j.startTask()
			})
		cr.MustAddJob(job)
	default:
		// do nothing
	}
	return
}

func (j *Job) stop(resultChan chan struct{}) {
	switch j.context.GetScheduleMode() {
	case pkg.ScheduleModeOnce:
		j.context.GetSchedule().
			WithStatus(pkg.ScheduleStatusStopped).
			WithStopTime(time.Now())
		j.crawler.GetSignal().ScheduleStopped(j.context)

		resultChan <- struct{}{}
		j.spider.StopSchedule()
	case pkg.ScheduleModeLoop:
		j.startTask()
	case pkg.ScheduleModeCron:
		if j.context.GetScheduleOnlyOneTask() {
			j.cronJob <- struct{}{}
		}
	default:
		// do nothing
	}
}
func (j *Job) startTask() {
	_, _ = new(Task).FromSpider(j.spider).WithJob(j).start(j.context)
	j.task.In()
}
func (j *Job) StopTask() {
	j.task.Out()
}

func (j *Job) FromSpider(spider pkg.Spider) *Job {
	*j = Job{
		task:    pkg.NewState(),
		crawler: spider.GetCrawler(),
		spider:  spider,
		logger:  spider.GetLogger(),
		cronJob: make(chan struct{}, 1),
	}
	return j
}
