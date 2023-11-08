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
			WithId(utils.UUIDV1WithoutHyphens()).
			WithStatus(pkg.JobStatusStarted).
			WithStartTime(time.Now()).
			WithEnable(true).
			WithFunc(jobFunc).
			WithArgs(args).
			WithMode(mode).
			WithSpec(spec).
			WithOnlyOneTask(onlyOneTask))
	j.crawler.GetSignal().JobStarted(j.context)
	return
}
func (j *Job) kill(_ context.Context) (err error) {
	j.cancel()
	return
}
func (j *Job) run(ctx context.Context) (err error) {
	if j.GetContext() == nil {
		err = errors.New("job hasn't started")
		j.logger.Error(err)
		return
	}

	go func() {
		j.logger.Infof("2222222222%+v\n", j.context.GetJobContext())
		select {
		case <-j.context.GetJobContext().Done():
			if j.context.GetJobStatus() != pkg.JobStatusStopped {
				j.logger.Info(333333333333)
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

	j.task.RegisterIsReadyAndIsZero(func() {
		j.stop(nil)
	})

	ctx, j.cancel = context.WithCancel(context.Background())
	j.context.WithJobContext(ctx)

	//j.crawler.GetSignal().JobStarted(j.context)

	switch j.context.GetJobMode() {
	case pkg.JobModeOnce:
		go j.startTask()
	case pkg.JobModeLoop:
		go j.startTask()
	case pkg.JobModeCron:
		if j.context.GetJobOnlyOneTask() {
			j.cronJob <- struct{}{}
		}
		cr := cron.New(cron.WithLogger(j.logger))
		cr.MustStart()
		job := new(cron.Job).
			MustEverySpec(j.context.GetJobSpec()).
			Callback(func() {
				if j.context.GetJobOnlyOneTask() {
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
		j.context.GetJob().
			WithStatus(pkg.JobStatusStopped).
			WithStopTime(time.Now())
		j.crawler.GetSignal().JobStopped(j.context)

		j.spider.JobStopped(j.context, err)
		return
	}

	if j.context.GetJobStatus() == pkg.JobStatusStopped {
		err = errors.New("the job has been finished early")
		j.logger.Error(err)
		return
	}

	switch j.context.GetJobMode() {
	case pkg.JobModeOnce:
		j.context.GetJob().
			WithStatus(pkg.JobStatusStopped).
			WithStopTime(time.Now())
		j.crawler.GetSignal().JobStopped(j.context)

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
