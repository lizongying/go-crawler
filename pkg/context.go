package pkg

import (
	"context"
	"time"
)

type Context interface {
	Global() Context
	GlobalContext() context.Context
	WithGlobalContext(ctx context.Context) Context
	GetMeta() Meta
	WithMeta(meta Meta) Context
	GetTaskId() string
	WithTaskId(taskId string) Context

	GetStartFunc() string
	WithStartFunc(startFunc string) Context
	GetArgs() string
	WithArgs(args string) Context
	GetMode() string
	WithMode(mode string) Context
	GetTaskStatus() TaskStatus
	WithTaskStatus(TaskStatus) Context
	GetTaskStartTime() time.Time
	WithTaskStartTime(t time.Time) Context
	GetTaskStopTime() time.Time
	WithTaskStopTime(t time.Time) Context

	GetCrawler() ContextCrawler
	GetCrawlerId() string
	WithCrawlerId(string) Context
	GetCrawlerStatus() CrawlerStatus
	WithCrawlerStatus(CrawlerStatus) Context
	GetCrawlerStartTime() time.Time
	WithCrawlerStartTime(time.Time) Context
	GetCrawlerStopTime() time.Time
	WithCrawlerStopTime(time.Time) Context

	GetSpider() ContextSpider
	GetSpiderName() string
	WithSpiderName(string) Context
	GetSpiderStatus() SpiderStatus
	WithSpiderStatus(SpiderStatus) Context
	GetSpiderStartTime() time.Time
	WithSpiderStartTime(time.Time) Context
	GetSpiderStopTime() time.Time
	WithSpiderStopTime(time.Time) Context

	GetSchedule() ContextSchedule
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
}

type ContextCrawler interface {
	GetId() string
	WithId(string) ContextCrawler
}

type ContextSpider interface {
	GetName() string
	WithName(string) ContextSpider
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
	GetStatus() ScheduleStatus
	WithStatus(ScheduleStatus) ContextSchedule
	GetStartTime() time.Time
	WithStartTime(time.Time) ContextSchedule
	GetStopTime() time.Time
	WithStopTime(time.Time) ContextSchedule
}
