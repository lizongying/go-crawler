package spider

import (
	"context"
	"errors"
	"github.com/lizongying/cron"
	"github.com/lizongying/go-crawler/pkg"
	crawlerContext "github.com/lizongying/go-crawler/pkg/context"
	"github.com/lizongying/go-crawler/pkg/utils"
)

type Job struct {
	context pkg.Context
	task    *pkg.State
	crawler pkg.Crawler
	spider  pkg.Spider
	logger  pkg.Logger
	cron    *cron.Cron
	cronJob chan struct{}
	cancel  context.CancelFunc
}

func (j *Job) GetContext() pkg.Context {
	return j.context
}
func (j *Job) WithContext(ctx pkg.Context) *Job {
	j.context = ctx
	return j
}
func (j *Job) start(ctx pkg.Context, jobFunc string, args string, mode pkg.JobMode, spec string, onlyOneTask bool) (err error) {
	if j.context != nil {
		err = errors.New("the job has started")
		j.logger.Error(err)
		return
	}

	j.context = new(crawlerContext.Context).
		WithCrawler(ctx.GetCrawler()).
		WithSpider(ctx.GetSpider()).
		WithJob(new(crawlerContext.Job).
			WithId(j.crawler.NextId()).
			WithEnable(true).
			WithFunc(jobFunc).
			WithArgs(args).
			WithMode(mode).
			WithSpec(spec).
			WithOnlyOneTask(onlyOneTask))

	j.context.GetJob().WithStatus(pkg.JobStatusReady)
	j.crawler.GetSignal().JobChanged(j.context)

	j.task.RegisterIsReadyAndIsZero(func() {
		j.stop(nil)
	})
	return
}
func (j *Job) kill(_ context.Context) (err error) {
	if j.context == nil {
		err = errors.New("the job hasn't started")
		j.logger.Error(err)
		return
	}

	if !utils.InSlice(j.context.GetJob().GetStatus(), []pkg.JobStatus{
		pkg.JobStatusRunning,
	}) {
		err = errors.New("the job can be killed in the running state")
		j.logger.Error(err)
		return
	}

	j.context.GetJob().WithStatus(pkg.JobStatusStopping)
	j.crawler.GetSignal().JobChanged(j.context)

	j.cancel()
	return
}
func (j *Job) run(ctx context.Context) (err error) {
	if j.context == nil {
		err = errors.New("the job hasn't started")
		j.logger.Error(err)
		return
	}

	if !utils.InSlice(j.context.GetJob().GetStatus(), []pkg.JobStatus{
		pkg.JobStatusReady,
		pkg.JobStatusStopped,
	}) {
		err = errors.New("the job can be started in the ready or stopped state")
		j.logger.Error(err)
		return
	}

	j.context.GetJob().WithStatus(pkg.JobStatusStarting)
	j.crawler.GetSignal().JobChanged(j.context)

	j.task.Clear()

	ctx, j.cancel = context.WithCancel(ctx)
	j.context.GetJob().WithContext(ctx)
	j.context.GetJob().WithSubId(j.crawler.GenUid())

	go func() {
		select {
		case <-j.context.GetJob().GetContext().Done():
			if j.context.GetJob().GetStatus() != pkg.JobStatusStopped {
				j.stop(ctx.Err())
			}
			return
		case <-ctx.Done():
			if j.context.GetJob().GetStatus() != pkg.JobStatusStopped {
				j.stop(ctx.Err())
			}
			return
		}
	}()

	switch j.context.GetJob().GetMode() {
	case pkg.JobModeOnce:
		go j.startTask()
	case pkg.JobModeLoop:
		go j.startTask()
	case pkg.JobModeCron:
		j.cronJob = make(chan struct{}, 1)
		if j.context.GetJob().GetOnlyOneTask() {
			j.cronJob <- struct{}{}
		}
		j.cron = cron.New(cron.WithLogger(j.logger))
		j.cron.MustStart()
		job := new(cron.Job).
			MustEverySpec(j.context.GetJob().GetSpec()).
			Callback(func() {
				if j.context.GetJob().GetOnlyOneTask() {
					<-j.cronJob
					if _, ok := <-j.cronJob; ok {
						return
					}
				}
				j.startTask()
			})
		j.cron.MustAddJob(job)
	default:
		// do nothing
	}
	return
}

func (j *Job) stop(err error) {
	if j.context.GetJob().GetStatus() == pkg.JobStatusStopped {
		err = errors.New("job has been finished")
		j.logger.Error(err)
		return
	}

	j.context.GetJob().WithStatus(pkg.JobStatusIdle)
	j.crawler.GetSignal().JobChanged(j.context)

	if err != nil {
		if j.context.GetJob().GetMode() == pkg.JobModeCron {
			close(j.cronJob)
			j.cron.MustStop()
		}

		j.context.GetJob().WithStopReason(err.Error())
		j.context.GetJob().WithStatus(pkg.JobStatusStopped)
		j.crawler.GetSignal().JobChanged(j.context)

		j.spider.JobStopped(j.context, err)
		return
	}

	switch j.context.GetJob().GetMode() {
	case pkg.JobModeOnce:
		j.context.GetJob().WithStatus(pkg.JobStatusStopped)
		j.crawler.GetSignal().JobChanged(j.context)

		j.spider.JobStopped(j.context, nil)
	case pkg.JobModeLoop:
		j.startTask()
	case pkg.JobModeCron:
		if j.context.GetJob().GetOnlyOneTask() {
			j.cronJob <- struct{}{}
		}
	default:
		// do nothing
	}
	return
}
func (j *Job) startTask() {
	// idle when job stopped
	if j.context.GetJob().GetStatus() != pkg.JobStatusRunning {
		j.context.GetJob().WithStatus(pkg.JobStatusRunning)
		j.crawler.GetSignal().JobChanged(j.context)
	}
	_, _ = new(Task).FromSpider(j.spider).WithJob(j).start(j.context)
	j.task.In()
}
func (j *Job) TaskStopped(ctx pkg.Context, _ error) {
	if ctx.GetTask().GetJobSubId() == j.context.GetJob().GetSubId() {
		j.task.Out()
	}
}

func (j *Job) FromSpider(spider pkg.Spider) *Job {
	*j = Job{
		task:    pkg.NewState(),
		crawler: spider.GetCrawler(),
		spider:  spider,
		logger:  spider.GetLogger(),
	}
	return j
}
