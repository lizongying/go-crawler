package pipelines

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/segmentio/kafka-go"
	"time"
)

type KafkaPipeline struct {
	pkg.UnimplementedPipeline
	env     string
	logger  pkg.Logger
	timeout time.Duration
	crawler pkg.Crawler
	config  pkg.Config
}

func (m *KafkaPipeline) ProcessItem(item pkg.Item) (err error) {
	task := item.GetContext().GetTask()

	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		task.IncItemError()
		return
	}

	if item.Name() != pkg.ItemKafka {
		m.logger.Warn("item not support", pkg.ItemKafka)
		return
	}

	itemKafka, ok := item.GetItem().(*items.ItemKafka)
	if !ok {
		m.logger.Warn("item not parsing failed with", pkg.ItemKafka)
		return
	}

	if itemKafka.GetTopic() == "" {
		err = errors.New("topic is empty")
		m.logger.Error(err)
		task.IncItemError()
		return
	}

	data := item.Data()
	if data == nil {
		err = errors.New("nil data")
		m.logger.Error(err)
		task.IncItemError()
		return
	}

	item.GetContext().GetItem().WithSaved(true)

	bs, err := json.Marshal(data)
	if err != nil {
		m.logger.Error(err)
		task.IncItemError()
		return
	}

	if m.env == "dev" {
		m.logger.Debug("current mode don't need save")
		task.IncItemIgnore()
		return
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, m.timeout)
	defer cancel()

	kafkaWriter, err := m.crawler.GetKafkaWriter(m.config.GetKafka(), itemKafka.GetTopic())
	if err != nil {
		return
	}

	err = kafkaWriter.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte(fmt.Sprint(itemKafka.Id())),
			Value: bs,
		},
	)
	if err != nil {
		m.logger.Error(err)
		task.IncItemError()
		return
	}

	task.IncItemSuccess()
	return
}

func (m *KafkaPipeline) FromSpider(spider pkg.Spider) (err error) {
	if m == nil {
		return new(KafkaPipeline).FromSpider(spider)
	}

	if err = m.UnimplementedPipeline.FromSpider(spider); err != nil {
		return
	}
	m.crawler = spider.GetCrawler()
	m.config = spider.GetConfig()
	m.env = spider.GetConfig().GetEnv()
	m.logger = spider.GetLogger()
	m.timeout = time.Minute
	return
}
