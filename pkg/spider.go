package pkg

import (
	"context"
	"net/url"
	"sync"
	"time"
)

type SpiderInfo struct {
	Mode  string
	Name  string
	Stats sync.Map

	Concurrency   int
	Interval      time.Duration
	OkHttpCodes   []int
	RetryMaxTimes int
	Timeout       time.Duration
}

type Spider interface {
	GetInfo() *SpiderInfo
	SetLogger(logger Logger)
	SetSpider(spider Spider)
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	GetAllowedDomains() []string
	ReplaceAllowedDomains([]string) error
	SetAllowedDomain(string)
	DelAllowedDomain(string) error
	CleanAllowedDomains()
	GetMiddlewares() map[int]string
	ReplaceMiddlewares(map[int]Middleware) error
	SetMiddleware(Middleware, int)
	DelMiddleware(string)
	CleanMiddlewares()
	SortedMiddlewares() []Middleware
	SetItemDelay(time.Duration)
	SetItemConcurrency(int)
	SetRequestRate(string, time.Duration, int)
	IsAllowedDomain(*url.URL) bool
	YieldRequest(*Request) error
	YieldItem(*Item) error
	GetDevServer() DevServer
}
