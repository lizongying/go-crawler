package crawler

import (
	"context"
	"database/sql"
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/api"
	"github.com/lizongying/go-crawler/pkg/cli"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type Crawler struct {
	spiders     []pkg.Spider
	spiderName  string
	startFunc   string
	args        string
	mode        string
	config      pkg.Config
	logger      pkg.Logger
	MongoDb     *mongo.Database
	Mysql       *sql.DB
	Redis       *redis.Client
	Kafka       *kafka.Writer
	KafkaReader *kafka.Reader
	S3          *s3.Client
	mockServer  pkg.MockServer
	api         *api.Api
}

func (c *Crawler) GetMode() string {
	return c.mode
}
func (c *Crawler) SetMode(mode string) {
	c.mode = mode
}
func (c *Crawler) GetLogger() pkg.Logger {
	return c.logger
}
func (c *Crawler) SetLogger(logger pkg.Logger) {
	c.logger = logger
}
func (c *Crawler) GetSpiders() []pkg.Spider {
	return c.spiders
}
func (c *Crawler) SetSpiders(spiders []pkg.Spider) {
	c.spiders = spiders
}
func (c *Crawler) AddSpider(spider pkg.Spider) {
	c.spiders = append(c.spiders, spider)
}
func (c *Crawler) RunMockServer() (err error) {
	err = c.mockServer.Run()
	if err != nil {
		c.logger.Error(err)
		return
	}

	return
}
func (c *Crawler) AddMockServerRoutes(routes ...pkg.Route) {
	c.mockServer.AddRoutes(routes...)
}
func (c *Crawler) GetConfig() pkg.Config {
	return c.config
}
func (c *Crawler) GetKafka() *kafka.Writer {
	return c.Kafka
}
func (c *Crawler) GetKafkaReader() *kafka.Reader {
	return c.KafkaReader
}
func (c *Crawler) GetMongoDb() *mongo.Database {
	return c.MongoDb
}
func (c *Crawler) GetMysql() *sql.DB {
	return c.Mysql
}
func (c *Crawler) GetRedis() *redis.Client {
	return c.Redis
}
func (c *Crawler) GetS3() *s3.Client {
	return c.S3
}
func (c *Crawler) StartSpider(ctx context.Context, req pkg.ReqStartSpider) (err error) {
	taskId := req.TaskId
	c.logger.Info(taskId)
	spiderName := req.Name
	startFunc := req.Func
	args := req.Args
	var spider pkg.Spider
	for _, v := range c.spiders {
		if v.Name() == spiderName {
			spider = v
			break
		}
	}

	if spider == nil {
		err = errors.New("nil spider")
		c.logger.Error(err)
		return
	}

	c.logger.Info("name", spiderName)
	c.logger.Info("func", startFunc)
	c.logger.Info("args", args)
	c.logger.Info("mode", c.mode)
	c.logger.Info("allowedDomains", spider.GetAllowedDomains())
	c.logger.Info("okHttpCodes", spider.OkHttpCodes())
	c.logger.Info("platforms", spider.GetPlatforms())
	c.logger.Info("browsers", spider.GetBrowsers())
	c.logger.Info("referrerPolicy", c.config.GetReferrerPolicy())
	c.logger.Info("urlLengthLimit", c.config.GetUrlLengthLimit())
	c.logger.Info("redirectMaxTimes", c.config.GetRedirectMaxTimes())
	c.logger.Info("retryMaxTimes", c.config.GetRetryMaxTimes())
	c.logger.Info("filter", c.config.GetFilter())

	if err = spider.Start(ctx, taskId, startFunc, args); err != nil {
		c.logger.Error(err)
		return
	}

	return
}
func (c *Crawler) StopSpider(ctx context.Context, req pkg.ReqStopSpider) (err error) {
	taskId := req.TaskId
	c.logger.Info(taskId)
	return
}
func (c *Crawler) Start(ctx context.Context) (err error) {
	if err = c.api.Run(); err != nil {
		c.logger.Error(err)
		return
	}

	req := pkg.ReqStartSpider{
		Name: c.spiderName,
		Func: c.startFunc,
		Args: c.args,
	}
	switch c.mode {
	case "once":
		req.TaskId = uuid.New().String()
		return c.StartSpider(ctx, req)
	case "loop":
		for {
			req.TaskId = uuid.New().String()
			if err = c.StartSpider(ctx, req); err != nil {
				c.logger.Error(err)
			}
		}
	case "cron":
		return
	default:
		select {}
	}
}

func (c *Crawler) Stop(ctx context.Context) (err error) {
	c.logger.Debug("Crawler wait for stop")
	defer func() {
		if err == nil {
			c.logger.Info("Crawler Stopped")
		}
	}()

	if ctx == nil {
		ctx = context.Background()
	}

	for _, v := range c.spiders {
		if err = v.Stop(ctx); err != nil {
			c.logger.Error(err)
		}
	}

	return
}

func NewCrawler(spiders []pkg.Spider, cli *cli.Cli, config *config.Config, logger pkg.Logger, mongoDb *mongo.Database, mysql *sql.DB, redis *redis.Client, kafka *kafka.Writer, kafkaReader *kafka.Reader, s3 *s3.Client, mockServer pkg.MockServer, httpApi *api.Api) (crawler pkg.Crawler, err error) {
	crawler = &Crawler{
		spiderName:  cli.SpiderName,
		startFunc:   cli.StartFunc,
		args:        cli.Args,
		mode:        cli.Mode,
		config:      config,
		logger:      logger,
		MongoDb:     mongoDb,
		Mysql:       mysql,
		Redis:       redis,
		Kafka:       kafka,
		KafkaReader: kafkaReader,
		S3:          s3,
		mockServer:  mockServer,
		api:         httpApi,
	}

	httpApi.AddRoutes(new(api.RouteHello).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteSpider).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteSpiderRun).FromCrawler(crawler))
	logger.Info("routes", httpApi.GetRoutes())

	for _, v := range spiders {
		v.SetSpider(v)
		v.FromCrawler(crawler)
		for _, option := range v.Options() {
			option(v)
		}

		crawler.AddSpider(v)
	}

	if config.MockServerEnable() {
		err = crawler.RunMockServer()
		if err != nil {
			logger.Error(err)
			return
		}
	}

	return crawler, nil
}
