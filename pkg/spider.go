package pkg

import (
	"context"
	"database/sql"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
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
	SetMiddleware(func() Middleware, int) Spider
	DelMiddleware(string)
	CleanMiddlewares()
	SortedMiddlewares() []Middleware
	SetItemDelay(time.Duration) Spider
	SetItemConcurrency(int) Spider
	SetRequestRate(string, time.Duration, int) Spider
	IsAllowedDomain(*url.URL) bool
	YieldRequest(*Request) error
	YieldItem(Item) error
	RunDevServer() error
	GetDevServerHost() string
	AddDevServerRoutes(routes ...Route) Spider
	AddOkHttpCodes(httpCodes ...int) Spider
	GetOkHttpCodes() []int
	GetPlatform() []Platform
	GetBrowser() []Browser
	GetConfig() *config.Config
	GetLogger() Logger
	GetHttpClient() HttpClient
	GetKafka() *kafka.Writer
	GetMongoDb() *mongo.Database
	GetMysql() *sql.DB
}
