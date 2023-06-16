package pkg

import (
	"context"
	"database/sql"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"net/url"
	"time"
)

type Crawler interface {
	GetMode() string
	SetMode(string)
	GetAllowedDomains() []string
	ReplaceAllowedDomains([]string) error
	SetAllowedDomain(string)
	DelAllowedDomain(string) error
	CleanAllowedDomains()
	IsAllowedDomain(*url.URL) bool

	SetItemDelay(time.Duration) Crawler
	SetItemConcurrency(int) Crawler
	SetRequestRate(string, time.Duration, int) Crawler

	Request(context.Context, *Request) (*Response, error)
	YieldRequest(context.Context, *Request) error
	YieldItem(context.Context, Item) error
	SetDownloader(Downloader)
	SetExporter(Exporter)

	GetInfo() *SpiderInfo
	GetStats() Stats
	GetLogger() Logger
	SetLogger(logger Logger)
	SetSpider(spider Spider)
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	RunDevServer() error
	GetDevServerHost() string
	AddDevServerRoutes(routes ...Route)
	AddOkHttpCodes(httpCodes ...int)
	GetOkHttpCodes() []int
	GetPlatforms() []Platform
	SetPlatforms(...Platform)
	GetBrowsers() []Browser
	SetBrowsers(...Browser)
	GetConfig() Config
	GetKafka() *kafka.Writer
	GetMongoDb() *mongo.Database
	GetMysql() *sql.DB
	GetFilter() Filter
	SetMiddleware(Middleware, uint8)
	SetPipeline(Pipeline, uint8)
	GetRetryMaxTimes() uint8
	SetRetryMaxTimes(uint8)
	GetTimeout() time.Duration
	SetTimeout(time.Duration)
	GetInterval() time.Duration
	SetInterval(time.Duration)
}

type CrawlOption func(Crawler)

func WithLogger(logger Logger) CrawlOption {
	return func(crawler Crawler) {
		crawler.SetLogger(logger)
	}
}
func WithMode(mode string) CrawlOption {
	return func(crawler Crawler) {
		crawler.SetMode(mode)
	}
}
func WithDownloader(downloader Downloader) CrawlOption {
	return func(crawler Crawler) {
		crawler.SetDownloader(downloader)
	}
}
func WithExporter(exporter Exporter) CrawlOption {
	return func(crawler Crawler) {
		crawler.SetExporter(exporter)
	}
}
func WithMiddleware(middleware Middleware, order uint8) CrawlOption {
	return func(crawler Crawler) {
		crawler.SetMiddleware(middleware, order)
	}
}
func WithPipeline(pipeline Pipeline, order uint8) CrawlOption {
	return func(crawler Crawler) {
		crawler.SetPipeline(pipeline, order)
	}
}
func WithRetryMaxTimes(retryMaxTimes uint8) CrawlOption {
	return func(crawler Crawler) {
		crawler.SetRetryMaxTimes(retryMaxTimes)
	}
}
func WithTimeout(timeout time.Duration) CrawlOption {
	return func(crawler Crawler) {
		crawler.SetTimeout(timeout)
	}
}
func WithInterval(timeout time.Duration) CrawlOption {
	return func(crawler Crawler) {
		crawler.SetInterval(timeout)
	}
}
