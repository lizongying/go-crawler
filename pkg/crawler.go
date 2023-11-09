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
	Stop(context.Context) error
	RunMockServer() error
	AddMockServerRoutes(...Route)
	GetLogger() Logger
	SetLogger(Logger)
	GetConfig() Config
	GetKafka() *kafka.Writer
	GetKafkaReader() *kafka.Reader
	GetRedis() *redis.Client
	GetMongoDb() *mongo.Database
	GetMysql() *sql.DB
	GetSqlite() Sqlite
	GetStore() Store

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
}

type CrawlOption func(Crawler)

func WithMockServerRoutes(routes ...func(logger Logger) Route) CrawlOption {
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

type CrawlerStatus uint8

const (
	CrawlerStatusUnknown = iota
	CrawlerStatusOnline
	CrawlerStatusOffline
)

func (c *CrawlerStatus) String() string {
	switch *c {
	case 1:
		return "online"
	case 2:
		return "offline"
	default:
		return "unknown"
	}
}
