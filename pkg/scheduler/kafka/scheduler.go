package kafka

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/scheduler"
	"github.com/segmentio/kafka-go"
)

type Scheduler struct {
	scheduler.UnimplementedScheduler

	kafkaReader *kafka.Reader
	kafkaWriter *kafka.Writer
	requestKey  string

	crawler pkg.Crawler
	spider  pkg.Spider
	config  pkg.Config
	logger  pkg.Logger
	task    pkg.Task
}

func (s *Scheduler) StartScheduler(task pkg.Task) (err error) {
	s.task = task
	s.UnimplementedScheduler.SetTask(s.task)

	ctx := task.GetContext()

	s.initScheduler(ctx)

	go s.HandleItem(ctx)

	go s.handleRequest(ctx)
	return
}

func (s *Scheduler) StopScheduler(_ pkg.Task) (err error) {
	return
}
func (s *Scheduler) initScheduler(_ pkg.Context) {
	s.requestKey = fmt.Sprintf("%s-%s-request", s.config.GetBotName(), s.spider.Name())
	s.logger.Info("request key", s.requestKey)
	s.kafkaWriter.Topic = s.requestKey
	s.kafkaReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  s.kafkaReader.Config().Brokers,
		MaxBytes: 10e6, // 10MB
		Topic:    s.requestKey,
		GroupID:  s.config.GetBotName(),
	})
}
func (s *Scheduler) FromSpider(spider pkg.Spider) pkg.Scheduler {
	if s == nil {
		return new(Scheduler).FromSpider(spider)
	}

	s.crawler = spider.GetCrawler()
	s.UnimplementedScheduler.SetCrawler(s.crawler)
	s.spider = spider
	s.UnimplementedScheduler.SetSpider(spider)
	s.config = spider.GetConfig()
	s.logger = spider.GetLogger()
	s.UnimplementedScheduler.SetLogger(s.logger)
	s.UnimplementedScheduler.Init()

	s.kafkaWriter = s.crawler.GetKafka()
	s.kafkaReader = s.crawler.GetKafkaReader()

	return s
}
