package pkg

import (
	"context"
	"net/url"
	"time"
)

type SpiderInfo struct {
	Mode string
	Name string

	Concurrency   int
	Interval      time.Duration
	RetryMaxTimes int
	Timeout       time.Duration
}

type Spider interface {
	GetInfo() *SpiderInfo
	GetStats() Stats
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
	SetMiddleware(Middleware, int) Spider
	DelMiddleware(string)
	CleanMiddlewares()
	SortedMiddlewares() []Middleware
	SetItemDelay(time.Duration) Spider
	SetItemConcurrency(int) Spider
	SetRequestRate(string, time.Duration, int) Spider
	IsAllowedDomain(*url.URL) bool
	YieldRequest(*Request) error
	YieldItem(Item) error
	GetDevServer() DevServer
	AddOkHttpCodes(httpCodes ...int) Spider
	GetOkHttpCodes() []int
}
