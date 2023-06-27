package pkg

import (
	"context"
	"time"
)

type SchedulerType string

const (
	SchedulerUnknown SchedulerType = ""
	SchedulerMemory  SchedulerType = "memory"
	SchedulerRedis   SchedulerType = "redis"
)

type Scheduler interface {
	GetDownloader() Downloader
	SetDownloader(Downloader)
	GetExporter() Exporter
	SetExporter(Exporter)
	SetItemDelay(time.Duration)
	SetItemConcurrency(int)
	SetRequestRate(string, time.Duration, int)
	YieldItem(context.Context, Item) error
	Request(context.Context, *Request) (*Response, error)
	YieldRequest(context.Context, *Request) error
	Start(context.Context) error
	Stop(context.Context) error
	GetInterval() time.Duration
	SetInterval(time.Duration)
}
