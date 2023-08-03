package pipelines

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/utils"
	"github.com/segmentio/kafka-go"
	"time"
)

type KafkaPipeline struct {
	pkg.UnimplementedPipeline
	mode        string
	logger      pkg.Logger
	kafkaWriter *kafka.Writer
	timeout     time.Duration
}

func (m *KafkaPipeline) ProcessItem(_ context.Context, item pkg.Item) (err error) {
	spider := m.GetSpider()
	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		spider.IncItemError()
		return
	}
	if item.GetName() != pkg.ItemKafka {
		m.logger.Warn("item not support", pkg.ItemKafka)
		return
	}
	itemKafka, ok := item.GetItem().(*items.ItemKafka)
	if !ok {
		m.logger.Warn("item not parsing failed with", pkg.ItemKafka)
		return
	}

	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		spider.IncItemError()
		return
	}

	if itemKafka.GetTopic() == "" {
		err = errors.New("topic is empty")
		m.logger.Error(err)
		spider.IncItemError()
		return
	}

	data := item.GetData()
	if data == nil {
		err = errors.New("nil data")
		m.logger.Error(err)
		spider.IncItemError()
		return
	}

	m.logger.Debug("Data", utils.JsonStr(data))
	bs, err := json.Marshal(data)
	if err != nil {
		m.logger.Error(err)
		spider.IncItemError()
		return
	}

	if m.mode == "test" {
		m.logger.Debug("current mode don't need save")
		spider.IncItemIgnore()
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
			Key:   []byte(fmt.Sprint(itemKafka.GetId())),
			Value: bs,
		},
	)
	if err != nil {
		m.logger.Error(err)
		spider.IncItemError()
		return
	}

	spider.IncItemSuccess()
	return
}

func (m *KafkaPipeline) FromSpider(spider pkg.Spider) pkg.Pipeline {
	if m == nil {
		return new(KafkaPipeline).FromSpider(spider)
	}

	m.UnimplementedPipeline.FromSpider(spider)
	crawler := spider.GetCrawler()
	m.mode = crawler.GetMode()
	m.logger = spider.GetLogger()
	m.kafkaWriter = crawler.GetKafka()
	m.timeout = time.Minute
	return m
}
