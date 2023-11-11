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
	env         string
	logger      pkg.Logger
	kafkaWriter *kafka.Writer
	timeout     time.Duration
}

func (m *KafkaPipeline) ProcessItem(item pkg.Item) (err error) {
	spider := m.GetSpider()
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

	m.kafkaWriter.Topic = itemKafka.GetTopic()
	m.kafkaWriter.Logger = kafka.LoggerFunc(func(msg string, a ...interface{}) {
		m.logger.Info(itemKafka.GetTopic(), "insert success", a)
	})
	err = m.kafkaWriter.WriteMessages(ctx,
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

	item.GetContext().WithItemStopTime(time.Now())
	spider.GetCrawler().GetSignal().ItemChanged(item)
	task.IncItemSuccess()
	return
}

func (m *KafkaPipeline) FromSpider(spider pkg.Spider) pkg.Pipeline {
	if m == nil {
		return new(KafkaPipeline).FromSpider(spider)
	}

	m.UnimplementedPipeline.FromSpider(spider)
	crawler := spider.GetCrawler()
	m.env = spider.GetConfig().GetEnv()
	m.logger = spider.GetLogger()
	m.kafkaWriter = crawler.GetKafka()
	m.timeout = time.Minute
	return m
}
