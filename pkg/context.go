package pkg

import (
	"context"
	"time"
)

type Context interface {
	GetContext() Context
	GetCrawler() ContextCrawler
	WithCrawler(ContextCrawler) Context
	GetCrawlerContext() context.Context
	WithCrawlerContext(context.Context) Context
	GetCrawlerId() string
	WithCrawlerId(string) Context
	GetCrawlerStatus() CrawlerStatus
	WithCrawlerStatus(CrawlerStatus) Context
	GetCrawlerStartTime() time.Time
	GetCrawlerStopTime() time.Time
	GetCrawlerUpdateTime() time.Time

	GetSpider() ContextSpider
	WithSpider(ContextSpider) Context
	GetSpiderContext() context.Context
	WithSpiderContext(context.Context) Context
	GetSpiderName() string
	WithSpiderName(string) Context
	GetSpiderStatus() SpiderStatus
	WithSpiderStatus(SpiderStatus) Context
	GetSpiderStartTime() time.Time
	GetSpiderStopTime() time.Time
	GetSpiderUpdateTime() time.Time

	GetJob() ContextJob
	WithJob(ContextJob) Context
	GetJobContext() context.Context
	WithJobContext(context.Context) Context
	GetJobId() string
	WithJobId(string) Context
	GetJobSubId() uint64
	WithJobSubId(uint64) Context
	GetJobStatus() JobStatus
	WithJobStatus(JobStatus) Context
	GetJobEnable() bool
	WithJobEnable(bool) Context
	GetJobStartTime() time.Time
	GetJobStopTime() time.Time
	GetJobUpdateTime() time.Time
	GetJobFunc() string
	WithJobFunc(string) Context
	GetJobArgs() string
	WithJobArgs(string) Context
	GetJobMode() JobMode
	WithJobMode(JobMode) Context
	GetJobSpec() string
	WithJobSpec(string) Context
	GetJobOnlyOneTask() bool
	WithJobOnlyOneTask(bool) Context

	GetTask() ContextTask
	WithTask(ContextTask) Context
	GetTaskId() string
	WithTaskId(string) Context
	GetTaskContext() context.Context
	WithTaskContext(context.Context) Context
	GetTaskStatus() TaskStatus
	WithTaskStatus(TaskStatus) Context
	GetTaskStartTime() time.Time
	WithTaskStartTime(time.Time) Context
	GetTaskStopTime() time.Time
	WithTaskStopTime(time.Time) Context
	GetTaskDeadline() time.Time
	WithTaskDeadline(time.Time) Context

	GetRequest() ContextRequest
	WithRequest(ContextRequest) Context
	GetRequestId() string
	WithRequestId(string) Context
	GetRequestContext() context.Context
	WithRequestContext(context.Context) Context
	GetRequestStatus() RequestStatus
	WithRequestStatus(RequestStatus) Context
	GetRequestStartTime() time.Time
	WithRequestStartTime(time.Time) Context
	GetRequestStopTime() time.Time
	WithRequestStopTime(time.Time) Context
	GetRequestDeadline() time.Time
	WithRequestDeadline(time.Time) Context
	GetRequestCookies() map[string]string
	WithRequestCookies(map[string]string) Context
	GetRequestReferrer() string
	WithRequestReferrer(string) Context

	GetItem() ContextItem
	WithItem(ContextItem) Context
	GetItemId() string
	WithItemId(string) Context
	GetItemContext() context.Context
	WithItemContext(context.Context) Context
	GetItemStatus() ItemStatus
	WithItemStatus(ItemStatus) Context
	GetItemStartTime() time.Time
	WithItemStartTime(time.Time) Context
	GetItemStopTime() time.Time
	WithItemStopTime(time.Time) Context
}

type ContextCrawler interface {
	GetId() string
	WithId(string) ContextCrawler
	GetContext() context.Context
	WithContext(context.Context) ContextCrawler
	GetStatus() CrawlerStatus
	WithStatus(CrawlerStatus) ContextCrawler
	GetStartTime() time.Time
	GetStopTime() time.Time
	GetUpdateTime() time.Time
}

type ContextSpider interface {
	GetSpider() Spider
	WithSpider(Spider) ContextSpider
	GetId() uint64
	WithId(uint64) ContextSpider
	GetName() string
	WithName(string) ContextSpider
	GetContext() context.Context
	WithContext(context.Context) ContextSpider
	GetStatus() SpiderStatus
	WithStatus(SpiderStatus) ContextSpider
	GetStartTime() time.Time
	GetStopTime() time.Time
	GetUpdateTime() time.Time
}

type ContextJob interface {
	GetId() string
	WithId(string) ContextJob
	GetSubId() uint64
	WithSubId(uint64) ContextJob
	GetContext() context.Context
	WithContext(context.Context) ContextJob
	GetStatus() JobStatus
	WithStatus(JobStatus) ContextJob
	GetStartTime() time.Time
	GetStopTime() time.Time
	GetUpdateTime() time.Time
	GetEnable() bool
	WithEnable(bool) ContextJob
	GetFunc() string
	WithFunc(string) ContextJob
	GetArgs() string
	WithArgs(string) ContextJob
	GetMode() JobMode
	WithMode(JobMode) ContextJob
	GetSpec() string
	WithSpec(string) ContextJob
	GetOnlyOneTask() bool
	WithOnlyOneTask(bool) ContextJob
}

type ContextTask interface {
	Task
	Stats
	GetId() string
	WithId(string) ContextTask
	GetJobSubId() uint64
	WithJobSubId(uint64) ContextTask
	GetContext() context.Context
	WithContext(context.Context) ContextTask
	GetStatus() TaskStatus
	WithStatus(TaskStatus) ContextTask
	GetStartTime() time.Time
	WithStartTime(time.Time) ContextTask
	GetStopTime() time.Time
	WithStopTime(time.Time) ContextTask
	GetDeadline() time.Time
	WithDeadline(time.Time) ContextTask
	GetStats() Stats
	WithStats(Stats) ContextTask
}

type ContextRequest interface {
	GetId() string
	WithId(string) ContextRequest
	GetContext() context.Context
	WithContext(context.Context) ContextRequest
	GetStatus() RequestStatus
	WithStatus(RequestStatus) ContextRequest
	GetStartTime() time.Time
	WithStartTime(time.Time) ContextRequest
	GetStopTime() time.Time
	WithStopTime(time.Time) ContextRequest
	GetDeadline() time.Time
	WithDeadline(time.Time) ContextRequest
	GetCookies() map[string]string
	WithCookies(map[string]string) ContextRequest
	GetReferrer() string
	WithReferrer(string) ContextRequest
}

type ContextItem interface {
	GetId() string
	WithId(string) ContextItem
	GetContext() context.Context
	WithContext(context.Context) ContextItem
	GetStatus() ItemStatus
	WithStatus(ItemStatus) ContextItem
	GetStartTime() time.Time
	WithStartTime(time.Time) ContextItem
	GetStopTime() time.Time
	WithStopTime(time.Time) ContextItem
}
