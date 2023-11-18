package pkg

import (
	"context"
	"time"
)

type Context interface {
	GetContext() Context

	GetCrawler() ContextCrawler
	WithCrawler(ContextCrawler) Context

	GetSpider() ContextSpider
	WithSpider(ContextSpider) Context

	GetJob() ContextJob
	WithJob(ContextJob) Context

	GetTask() ContextTask
	WithTask(ContextTask) Context

	GetRequest() ContextRequest
	WithRequest(ContextRequest) Context

	GetItem() ContextItem
	WithItem(ContextItem) Context
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
	GetStopReason() string
	WithStopReason(stopReason string) ContextJob
}

type ContextTask interface {
	Task
	Stats
	GetId() string
	WithId(string) ContextTask
	GetJobSubId() uint64
	WithJobSubId(id uint64) ContextTask
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
	GetStopTime() time.Time
	GetUpdateTime() time.Time
	GetSaved() bool
	WithSaved(bool) ContextItem
}
