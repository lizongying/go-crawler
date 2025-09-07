package pkg

import (
	"context"
	"time"
)

// Context represents the execution context for the crawler.
// It holds references to Crawler, Spider, Job, Task, Request, and Item,
// It contains references to different stages of the crawling process.
type Context interface {

	// GetContext returns the current Context itself.
	GetContext() Context

	// CloneContext creates a new empty Context instance.
	CloneContext() Context

	// ----------------- Crawler -----------------

	// GetCrawler returns the crawler context.
	GetCrawler() ContextCrawler

	// WithCrawler sets the crawler context and returns the updated Context.
	WithCrawler(ContextCrawler) Context

	// CloneCrawler creates a new Context with the current crawler context.
	CloneCrawler() Context

	// ----------------- Spider -----------------

	// GetSpider returns the spider context.
	GetSpider() ContextSpider

	// WithSpider sets the spider context and returns the updated Context.
	WithSpider(ContextSpider) Context

	// CloneSpider creates a new Context with the current crawler and spider contexts.
	CloneSpider() Context

	// ----------------- Job -----------------

	// GetJob returns the job context.
	GetJob() ContextJob

	// WithJob sets the job context and returns the updated Context.
	WithJob(ContextJob) Context

	// CloneJob creates a new Context with the current crawler, spider, and job contexts.
	CloneJob() Context

	// ----------------- Task -----------------

	// GetTask returns the task context.
	GetTask() ContextTask

	// WithTask sets the task context and returns the updated Context.
	WithTask(ContextTask) Context

	// CloneTask creates a new Context with the current crawler, spider, job, and task contexts.
	CloneTask() Context

	// ----------------- Request -----------------

	// GetRequest returns the request context.
	GetRequest() ContextRequest

	// WithRequest sets the request context and returns the updated Context.
	WithRequest(ContextRequest) Context

	// CloneRequest creates a new Context with the current crawler, spider, job, task, and request contexts.
	CloneRequest() Context

	// ----------------- Item -----------------

	// GetItem returns the item context.
	GetItem() ContextItem

	// WithItem sets the item context and returns the updated Context.
	WithItem(ContextItem) Context

	// CloneItem creates a new Context with the current crawler, spider, job, task, request, and item contexts.
	CloneItem() Context
}

type ContextCrawler interface {
	GetCrawler() Crawler
	WithCrawler(Crawler) ContextCrawler
	GetContext() context.Context
	WithContext(context.Context) ContextCrawler
	GetId() string
	WithId(string) ContextCrawler
	GetStatus() CrawlerStatus
	WithStatus(CrawlerStatus) ContextCrawler
	GetStartTime() time.Time
	GetStopTime() time.Time
	GetUpdateTime() time.Time
}

type ContextSpider interface {
	GetSpider() Spider
	WithSpider(Spider) ContextSpider
	GetContext() context.Context
	WithContext(context.Context) ContextSpider
	GetId() uint64
	WithId(uint64) ContextSpider
	GetName() string
	WithName(string) ContextSpider
	GetStatus() SpiderStatus
	WithStatus(SpiderStatus) ContextSpider
	GetStartTime() time.Time
	GetStopTime() time.Time
	GetUpdateTime() time.Time
}

type ContextJob interface {
	GetJob() Job
	WithJob(Job) ContextJob
	GetContext() context.Context
	WithContext(context.Context) ContextJob
	GetId() string
	WithId(string) ContextJob
	GetSubId() uint64
	WithSubId(uint64) ContextJob
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
	Stats
	GetTask() Task
	WithTask(Task) ContextTask
	GetContext() context.Context
	WithContext(context.Context) ContextTask
	GetId() string
	WithId(string) ContextTask
	GetJobSubId() uint64
	WithJobSubId(id uint64) ContextTask
	GetStatus() TaskStatus
	WithStatus(TaskStatus) ContextTask
	GetStartTime() time.Time
	WithStartTime(time.Time) ContextTask
	GetStopTime() time.Time
	WithStopTime(time.Time) ContextTask
	GetUpdateTime() time.Time
	WithUpdateTime(time.Time) ContextTask
	GetDeadline() time.Time
	WithDeadline(time.Time) ContextTask
	GetStats() Stats
	WithStats(Stats) ContextTask
	GetStopReason() string
	WithStopReason(stopReason string) ContextTask
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
	GetUpdateTime() time.Time
	GetCookies() map[string]string
	WithCookies(map[string]string) ContextRequest
	GetReferrer() string
	WithReferrer(string) ContextRequest
	GetStopReason() string
	WithStopReason(stopReason string) ContextRequest
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
	GetStopReason() string
	WithStopReason(stopReason string) ContextItem
}
