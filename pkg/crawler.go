package pkg

import (
	"context"
	"database/sql"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Crawler interface {
	GetContext() Context
	WithContext(Context) Crawler

	GetSpiders() []Spider
	AddSpider(Spider)
	Start(context.Context) error
	Stop(Context) error
	RunMockServer() error
	AddMockServerRoutes(...Route) Crawler

	AddDefaultMocks() Crawler

	GetLogger() Logger
	SetLogger(Logger)
	GetConfig() Config
	GetKafkaWriter(name string, topic string) (*kafka.Writer, error)
	GetKafkaReader(name string, topic string) (*kafka.Reader, error)
	GetRedis(name string) (*redis.Client, error)
	GetMongoDb(name string) (*mongo.Database, error)
	GetMysql(name string) (*sql.DB, error)
	GetSqlite(name string) (*sql.DB, error)
	GetStore(name string) (Store, error)

	RunJob(context.Context, string, string, string, JobMode, string) (string, error)
	RerunJob(ctx context.Context, spiderName string, jobId string) (err error)
	KillJob(ctx context.Context, spiderName string, jobId string) (err error)

	SpiderStopped(ctx Context, err error)

	GetSignal() Signal
	SetSignal(Signal)

	GetStatistics() Statistics
	SetStatistics(statistics Statistics)

	GetItemDelay() time.Duration
	WithItemDelay(time.Duration) Crawler
	GetItemConcurrency() uint8
	WithItemConcurrency(uint8) Crawler
	ItemTimer() *time.Timer
	ItemConcurrencyChan() chan struct{}

	NextId() string
	GenUid() uint64

	StartFromCLI() bool

	GetStream() Stream

	GetCDP() bool
	WithCDP(bool) Crawler
}

type CrawlOption func(Crawler)

func WithMockServerRoutes(routes ...NewRoute) CrawlOption {
	return WithMocks(routes...)
}

func WithDefaultMocks() CrawlOption {
	return func(crawler Crawler) {
		crawler.AddDefaultMocks()
	}
}

func WithMocks(routes ...NewRoute) CrawlOption {
	return func(crawler Crawler) {
		if !crawler.GetConfig().MockServerEnable() {
			crawler.GetConfig().SetMockServerEnable(true)
			_ = crawler.RunMockServer()
		}

		for _, v := range routes {
			crawler.AddMockServerRoutes(v(crawler.GetLogger()))
		}
	}
}

func WithLogger(logger Logger) CrawlOption {
	return func(crawler Crawler) {
		crawler.SetLogger(logger)
	}
}
func WithItemDelay(delay time.Duration) CrawlOption {
	return func(crawler Crawler) {
		crawler.WithItemDelay(delay)
	}
}
func WithItemConcurrency(concurrency uint8) CrawlOption {
	return func(crawler Crawler) {
		crawler.WithItemConcurrency(concurrency)
	}
}
func WithCDP(cdp bool) CrawlOption {
	return func(crawler Crawler) {
		crawler.WithCDP(cdp)
	}
}

type CrawlerStatus uint8

const (
	CrawlerStatusUnknown = iota
	CrawlerStatusReady
	CrawlerStatusStarting
	CrawlerStatusRunning
	CrawlerStatusIdle
	CrawlerStatusStopping
	CrawlerStatusStopped
)

func (c CrawlerStatus) String() string {
	switch c {
	case CrawlerStatusReady:
		return "ready"
	case CrawlerStatusStarting:
		return "starting"
	case CrawlerStatusRunning:
		return "running"
	case CrawlerStatusIdle:
		return "idle"
	case CrawlerStatusStopping:
		return "stopping"
	case CrawlerStatusStopped:
		return "stopped"
	default:
		return "unknown"
	}
}
