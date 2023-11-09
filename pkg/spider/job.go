package spider

import (
	"context"
	"errors"
	"github.com/lizongying/cron"
	"github.com/lizongying/go-crawler/pkg"
	crawlerContext "github.com/lizongying/go-crawler/pkg/context"
	"time"
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

	j.task.RegisterIsReadyAndIsZero(func() {
		j.stop(nil)
	})
	return
}
func (j *Job) kill(_ context.Context) (err error) {
	j.context.WithJobStatus(pkg.JobStatusUnknown)
	j.context.WithJobUpdateTime(time.Now())
	j.crawler.GetSignal().JobChanged(j.context, pkg.JobStatusUnknown)

	j.cancel()
	return
}
func (j *Job) run(ctx context.Context) (err error) {
	if j.context == nil {
		err = errors.New("job hasn't started")
		j.logger.Error(err)
		return
	}

	if j.context.GetJobStatus() == pkg.JobStatusStarted {
		err = errors.New("the job has been started")
		j.logger.Error(err)
		return
	}

	j.context.WithJobStatus(pkg.JobStatusUnknown)
	j.context.WithJobUpdateTime(time.Now())
	j.crawler.GetSignal().JobChanged(j.context, pkg.JobStatusUnknown)

	j.task.Clear()

	ctx, j.cancel = context.WithCancel(context.Background())
	j.context.WithJobContext(ctx)
	j.context.WithJobSubId(j.crawler.GenUid())
	j.context.WithJobStatus(pkg.JobStatusStarted)
	j.context.WithJobStartTime(time.Now())
	j.crawler.GetSignal().JobStarted(j.context)

	j.context.WithJobUpdateTime(time.Now())
	j.crawler.GetSignal().JobChanged(j.context, pkg.JobStatusStarted)

	go func() {
		select {
		case <-j.context.GetJobContext().Done():
			if j.context.GetJobStatus() != pkg.JobStatusStopped {
				j.stop(ctx.Err())
			}
			return
		case <-ctx.Done():
			if j.context.GetJobStatus() != pkg.JobStatusStopped {
				j.stop(ctx.Err())
			}
			return
		}
	}()

	switch j.context.GetJobMode() {
	case pkg.JobModeOnce:
		go j.startTask()
	case pkg.JobModeLoop:
		go j.startTask()
	case pkg.JobModeCron:
		j.cronJob = make(chan struct{}, 1)
		if j.context.GetJobOnlyOneTask() {
			j.cronJob <- struct{}{}
		}
		j.cron = cron.New(cron.WithLogger(j.logger))
		j.cron.MustStart()
		job := new(cron.Job).
			MustEverySpec(j.context.GetJobSpec()).
			Callback(func() {
				if j.context.GetJobOnlyOneTask() {
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
	if j.context.GetJobStatus() == pkg.JobStatusStopped {
		err = errors.New("the job has been finished early")
		j.logger.Error(err)
		return
	}

	if err != nil {
		if j.context.GetJobMode() == pkg.JobModeCron {
			close(j.cronJob)
			j.cron.MustStop()
		}

		j.context.GetJob().
			WithStatus(pkg.JobStatusStopped).
			WithStopTime(time.Now())
		j.crawler.GetSignal().JobStopped(j.context)

		j.context.WithJobUpdateTime(time.Now())
		j.crawler.GetSignal().JobChanged(j.context, pkg.JobStatusStopped)

		j.spider.JobStopped(j.context, err)
		return
	}

	switch j.context.GetJobMode() {
	case pkg.JobModeOnce:
		j.context.GetJob().
			WithStatus(pkg.JobStatusStopped).
			WithStopTime(time.Now())
		j.crawler.GetSignal().JobStopped(j.context)

		j.context.WithJobUpdateTime(time.Now())
		j.crawler.GetSignal().JobChanged(j.context, pkg.JobStatusStopped)

		j.spider.JobStopped(j.context, nil)
	case pkg.JobModeLoop:
		j.startTask()
	case pkg.JobModeCron:
		if j.context.GetJobOnlyOneTask() {
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
func (j *Job) TaskStopped(ctx pkg.Context, _ error) {
	if ctx.GetTask().GetJobSubId() == j.context.GetJobSubId() {
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
