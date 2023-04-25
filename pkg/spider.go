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
	GetPipelines() map[int]string
	ReplacePipelines(map[int]Pipeline) error
	SetPipeline(Pipeline, int)
	DelPipeline(string)
	CleanPipelines()
	SetItemDelay(time.Duration)
	SetItemConcurrency(int)
	SetRate(string, time.Duration, int)
	IsAllowedDomain(*url.URL) bool
}
