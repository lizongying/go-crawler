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
	WithCrawlerStartTime(time.Time) Context
	GetCrawlerStopTime() time.Time
	WithCrawlerStopTime(time.Time) Context

	GetSpider() ContextSpider
	WithSpider(ContextSpider) Context
	GetSpiderContext() context.Context
	WithSpiderContext(context.Context) Context
	GetSpiderName() string
	WithSpiderName(string) Context
	GetSpiderStatus() SpiderStatus
	WithSpiderStatus(SpiderStatus) Context
	GetSpiderStartTime() time.Time
	WithSpiderStartTime(time.Time) Context
	GetSpiderStopTime() time.Time
	WithSpiderStopTime(time.Time) Context

	GetSchedule() ContextSchedule
	WithSchedule(ContextSchedule) Context
	GetScheduleContext() context.Context
	WithScheduleContext(context.Context) Context
	GetScheduleId() string
	WithScheduleId(string) Context
	GetScheduleStatus() ScheduleStatus
	WithScheduleStatus(ScheduleStatus) Context
	GetScheduleEnable() bool
	WithScheduleEnable(bool) Context
	GetScheduleStartTime() time.Time
	WithScheduleStartTime(time.Time) Context
	GetScheduleStopTime() time.Time
	WithScheduleStopTime(time.Time) Context
	GetScheduleFunc() string
	WithScheduleFunc(string) Context
	GetScheduleArgs() string
	WithScheduleArgs(string) Context
	GetScheduleMode() ScheduleMode
	WithScheduleMode(ScheduleMode) Context
	GetScheduleSpec() string
	WithScheduleSpec(string) Context
	GetScheduleOnlyOneTask() bool
	WithScheduleOnlyOneTask(bool) Context

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
	WithStartTime(time.Time) ContextCrawler
	GetStopTime() time.Time
	WithStopTime(time.Time) ContextCrawler
}

type ContextSpider interface {
	GetName() string
	WithName(string) ContextSpider
	GetContext() context.Context
	WithContext(context.Context) ContextSpider
	GetStatus() SpiderStatus
	WithStatus(SpiderStatus) ContextSpider
	GetStartTime() time.Time
	WithStartTime(time.Time) ContextSpider
	GetStopTime() time.Time
	WithStopTime(time.Time) ContextSpider
}

type ContextSchedule interface {
	GetId() string
	WithId(string) ContextSchedule
	GetContext() context.Context
	WithContext(context.Context) ContextSchedule
	GetStatus() ScheduleStatus
	WithStatus(ScheduleStatus) ContextSchedule
	GetStartTime() time.Time
	WithStartTime(time.Time) ContextSchedule
	GetStopTime() time.Time
	WithStopTime(time.Time) ContextSchedule
	GetEnable() bool
	WithEnable(bool) ContextSchedule
	GetFunc() string
	WithFunc(string) ContextSchedule
	GetArgs() string
	WithArgs(string) ContextSchedule
	GetMode() ScheduleMode
	WithMode(ScheduleMode) ContextSchedule
	GetSpec() string
	WithSpec(string) ContextSchedule
	GetOnlyOneTask() bool
	WithOnlyOneTask(bool) ContextSchedule
}

type ContextTask interface {
	Task
	Stats
	GetId() string
	WithId(string) ContextTask
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
