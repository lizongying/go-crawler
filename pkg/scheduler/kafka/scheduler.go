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

	go s.HandleItem(ctx)

	go s.handleRequest(ctx)
	return
}

func (s *Scheduler) StopScheduler(_ pkg.Task) (err error) {
	return
}
func (s *Scheduler) FromSpider(spider pkg.Spider) (scheduler pkg.Scheduler, err error) {
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
	k := s.config.GetKafka()
	topic := fmt.Sprintf("%s-%s-request", s.config.GetBotName(), s.spider.Name())
	s.logger.Info("topic", topic)
	s.kafkaWriter, err = s.crawler.GetKafkaWriter(k, topic)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.kafkaReader, err = s.crawler.GetKafkaReader(k, topic)
	if err != nil {
		s.logger.Error(err)
		return
	}

	return s, nil
}
