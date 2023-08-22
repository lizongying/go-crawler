package pkg

import (
	"context"
	"database/sql"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
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
	GetS3() *s3.Client

	SpiderStart(context.Context, ReqSpiderStart) error
	SpiderStop(context.Context, ReqSpiderStop) error
}

type CrawlOption func(Crawler)

func WithMockServerRoute(route func(logger Logger) Route) CrawlOption {
	return func(crawler Crawler) {
		if !crawler.GetConfig().MockServerEnable() {
			crawler.GetConfig().SetMockServerEnable(true)
			_ = crawler.RunMockServer()
		}

		crawler.AddMockServerRoutes(route(crawler.GetLogger()))
	}
}
func WithMode(mode string) CrawlOption {
	return func(crawler Crawler) {
		crawler.SetMode(mode)
	}
}
