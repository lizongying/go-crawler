package crawler

import (
	"context"
	"database/sql"
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/lizongying/go-crawler/pkg"
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
	devServer   pkg.DevServer
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
func (c *Crawler) RunDevServer() (err error) {
	err = c.devServer.Run()
	if err != nil {
		c.logger.Error(err)
		return
	}

	return
}
func (c *Crawler) AddDevServerRoutes(routes ...pkg.Route) {
	c.devServer.AddRoutes(routes...)
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
func (c *Crawler) Start(ctx context.Context) (err error) {
	var spider pkg.Spider
	for _, v := range c.spiders {
		if v.GetName() == c.spiderName {
			spider = v
			break
		}
	}

	if spider == nil {
		err = errors.New("nil spider")
		c.logger.Error(err)
		return
	}

	c.logger.Info("name", c.spiderName)
	c.logger.Info("start func", c.startFunc)
	c.logger.Info("args", c.args)
	c.logger.Info("mode", c.mode)
	c.logger.Info("allowedDomains", spider.GetAllowedDomains())
	c.logger.Info("okHttpCodes", spider.GetOkHttpCodes())
	c.logger.Info("platforms", spider.GetPlatforms())
	c.logger.Info("browsers", spider.GetBrowsers())
	c.logger.Info("referrerPolicy", c.config.GetReferrerPolicy())
	c.logger.Info("urlLengthLimit", c.config.GetUrlLengthLimit())
	c.logger.Info("redirectMaxTimes", c.config.GetRedirectMaxTimes())
	c.logger.Info("retryMaxTimes", c.config.GetRetryMaxTimes())
	c.logger.Info("filter", c.config.GetFilter())

	err = spider.Start(ctx, c.startFunc, c.args)
	if err != nil {
		c.logger.Error(err)
		return
	}

	return
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
		err = v.Stop(ctx)
		if err != nil {
			c.logger.Error(err)
			continue
		}
	}

	return
}

func NewCrawler(spiders []pkg.Spider, cli *cli.Cli, config *config.Config, logger pkg.Logger, mongoDb *mongo.Database, mysql *sql.DB, redis *redis.Client, kafka *kafka.Writer, kafkaReader *kafka.Reader, s3 *s3.Client, devServer pkg.DevServer) (crawler pkg.Crawler, err error) {
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
		devServer:   devServer,
	}

	for _, v := range spiders {
		v.SetSpider(v)
		v.FromCrawler(crawler)
		for _, option := range v.Options() {
			option(v)
		}

		crawler.AddSpider(v)
	}

	if config.GetEnableDevServer() {
		err = crawler.RunDevServer()
		if err != nil {
			logger.Error(err)
			return
		}
	}

	return crawler, nil
}
