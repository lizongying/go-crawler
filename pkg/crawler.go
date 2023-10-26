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
	GetMode() string
	SetMode(string)
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

	SpiderStart(Context) error
	SpiderStop(Context) error

	GetSignal() Signal
	SetSignal(Signal)

	GetStatistics() Statistics
	SetStatistics(statistics Statistics)

	GetId() string
	GetStatus() CrawlerStatus
	WithStatus(CrawlerStatus) Crawler
	GetStartTime() time.Time
	WithStartTime(t time.Time) Crawler
	GetStopTime() time.Time
	WithStopTime(t time.Time) Crawler
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
func WithMode(mode string) CrawlOption {
	return func(crawler Crawler) {
		crawler.SetMode(mode)
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
