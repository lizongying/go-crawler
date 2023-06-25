package pkg

import (
	"context"
	"database/sql"
	"github.com/redis/go-redis/v9"
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
	GetUsername() string
	SetUsername(string)
	GetPassword() string
	SetPassword(string)
	GetSpider() Spider
	SetSpider(Spider)
	Start(context.Context) error
	Stop(context.Context) error
	RunDevServer() error
	GetDevServerHost() string
	AddDevServerRoutes(...Route)
	GetLogger() Logger
	SetLogger(Logger)
	GetPlatforms() []Platform
	SetPlatforms(...Platform)
	GetBrowsers() []Browser
	SetBrowsers(...Browser)
	GetConfig() Config
	GetKafka() *kafka.Writer
	GetRedis() *redis.Client
	GetMongoDb() *mongo.Database
	GetMysql() *sql.DB
	GetFilter() Filter
	SetFilter(Filter)
	GetRetryMaxTimes() uint8
	SetRetryMaxTimes(uint8)
	GetTimeout() time.Duration
	SetTimeout(time.Duration)
	GetOkHttpCodes() []int
	SetOkHttpCodes(...int)
	GetScheduler() Scheduler
	SetScheduler(Scheduler)
	YieldItem(context.Context, Item) error
	Request(context.Context, *Request) (*Response, error)
	YieldRequest(context.Context, *Request) error
	GetStats() Stats
}

type CrawlOption func(Crawler)

func WithMode(mode string) CrawlOption {
	return func(crawler Crawler) {
		crawler.SetMode(mode)
	}
}
func WithPlatforms(platforms ...Platform) CrawlOption {
	return func(crawler Crawler) {
		crawler.SetPlatforms(platforms...)
	}
}
func WithBrowsers(browsers ...Browser) CrawlOption {
	return func(crawler Crawler) {
		crawler.SetBrowsers(browsers...)
	}
}
func WithLogger(logger Logger) CrawlOption {
	return func(crawler Crawler) {
		crawler.SetLogger(logger)
	}
}
func WithFilter(filter Filter) CrawlOption {
	return func(crawler Crawler) {
		crawler.SetFilter(filter)
	}
}
func WithDownloader(downloader Downloader) CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().SetDownloader(downloader)
	}
}
func WithExporter(exporter Exporter) CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().SetExporter(exporter)
	}
}
func WithMiddleware(middleware Middleware, order uint8) CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetDownloader().SetMiddleware(middleware, order)
	}
}
func WithPipeline(pipeline Pipeline, order uint8) CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetExporter().SetPipeline(pipeline, order)
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
		crawler.GetScheduler().SetInterval(timeout)
	}
}
func WithOkHttpCodes(httpCodes ...int) CrawlOption {
	return func(crawler Crawler) {
		crawler.SetOkHttpCodes(httpCodes...)
	}
}
