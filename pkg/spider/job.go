package spider

import (
	"context"
	"errors"
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
func (j *Job) start(ctx pkg.Context, jobFunc string, args string, mode pkg.ScheduleMode, spec string, onlyOneTask bool) (err error) {
	j.context = new(crawlerContext.Context).
		WithCrawler(ctx.GetCrawler()).
		WithSpider(ctx.GetSpider()).
		WithSchedule(new(crawlerContext.Schedule).
			WithContext(context.Background()).
			WithId(utils.UUIDV1WithoutHyphens()).
			WithStatus(pkg.ScheduleStatusStarted).
			WithStartTime(time.Now()).
			WithEnable(true).
			WithFunc(jobFunc).
			WithArgs(args).
			WithMode(mode).
			WithSpec(spec).
			WithOnlyOneTask(onlyOneTask))
	j.crawler.GetSignal().ScheduleStarted(j.context)
	return
}
func (j *Job) run(ctx context.Context) (err error) {
	if j.GetContext() == nil {
		err = errors.New("job hasn't started")
		j.logger.Error(err)
		return
	}

	go func() {
		select {
		case <-ctx.Done():
			if j.context.GetScheduleStatus() != pkg.ScheduleStatusStopped {
				j.stop(ctx.Err())
			}
			return
		}
	}()

	j.task.RegisterIsReadyAndIsZero(func() {
		j.stop(nil)
	})
	j.context.WithScheduleContext(ctx)
	//j.crawler.GetSignal().ScheduleStarted(j.context)

	switch j.context.GetScheduleMode() {
	case pkg.ScheduleModeOnce:
		go j.startTask()
	case pkg.ScheduleModeLoop:
		go j.startTask()
	case pkg.ScheduleModeCron:
		if j.context.GetScheduleOnlyOneTask() {
			j.cronJob <- struct{}{}
		}
		cr := cron.New(cron.WithLogger(j.logger))
		cr.MustStart()
		job := new(cron.Job).
			MustEverySpec(j.context.GetScheduleSpec()).
			Callback(func() {
				if j.context.GetScheduleOnlyOneTask() {
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

func (j *Job) stop(err error) {
	if err != nil {
		j.context.GetSchedule().
			WithStatus(pkg.ScheduleStatusStopped).
			WithStopTime(time.Now())
		j.crawler.GetSignal().ScheduleStopped(j.context)

		j.spider.StopSchedule(j.context, err)
		return
	}

	if j.context.GetScheduleStatus() == pkg.ScheduleStatusStopped {
		err = errors.New("the job has been finished early")
		j.logger.Error(err)
		return
	}

	switch j.context.GetScheduleMode() {
	case pkg.ScheduleModeOnce:
		j.context.GetSchedule().
			WithStatus(pkg.ScheduleStatusStopped).
			WithStopTime(time.Now())
		j.crawler.GetSignal().ScheduleStopped(j.context)

		j.spider.StopSchedule(j.context, nil)
	case pkg.ScheduleModeLoop:
		j.startTask()
	case pkg.ScheduleModeCron:
		if j.context.GetScheduleOnlyOneTask() {
			j.cronJob <- struct{}{}
		}
	default:
		// do nothing
	}
	return
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
