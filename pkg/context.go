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
	GetSpiderName() string
	WithSpiderName(spiderName string) Context
	GetStartFunc() string
	WithStartFunc(startFunc string) Context
	GetArgs() string
	WithArgs(args string) Context
	GetMode() string
	WithMode(mode string) Context

	GetStatus() SpiderStatus
	WithStatus(SpiderStatus) Context
	GetStartTime() time.Time
	WithStartTime(t time.Time) Context
	GetStopTime() time.Time
	WithStopTime(t time.Time) Context

	GetCrawlerId() string
	WithCrawlerId(crawlerId string) Context
}
