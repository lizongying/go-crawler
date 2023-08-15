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
	RunDevServer() error
	AddDevServerRoutes(...Route)
	GetLogger() Logger
	SetLogger(Logger)
	GetConfig() Config
	GetKafka() *kafka.Writer
	GetKafkaReader() *kafka.Reader
	GetRedis() *redis.Client
	GetMongoDb() *mongo.Database
	GetMysql() *sql.DB
	GetS3() *s3.Client
}

type CrawlOption func(Crawler)

func WithDevServerRoute(route func(logger Logger) Route) CrawlOption {
	return func(crawler Crawler) {
		if !crawler.GetConfig().DevServerEnable() {
			crawler.GetConfig().SetDevServerEnable(true)
			_ = crawler.RunDevServer()
		}

		crawler.AddDevServerRoutes(route(crawler.GetLogger()))
	}
}
func WithMode(mode string) CrawlOption {
	return func(crawler Crawler) {
		crawler.SetMode(mode)
	}
}
