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
	crawler     pkg.Crawler
	stats       pkg.Stats
	logger      pkg.Logger
	kafkaWriter *kafka.Writer
	timeout     time.Duration
}

func (m *KafkaPipeline) ProcessItem(ctx context.Context, item pkg.Item) (err error) {
	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		m.stats.IncItemError()
		return
	}
	if item.GetName() != pkg.ItemKafka {
		m.logger.Warn("item not support", pkg.ItemKafka)
		return
	}
	itemKafka, ok := item.(*items.ItemKafka)
	if !ok {
		m.logger.Warn("item not parsing failed with", pkg.ItemKafka)
		return
	}

	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		m.stats.IncItemError()
		return
	}

	if itemKafka.GetTopic() == "" {
		err = errors.New("topic is empty")
		m.logger.Error(err)
		m.stats.IncItemError()
		return
	}

	data := item.GetData()
	if data == nil {
		err = errors.New("nil data")
		m.logger.Error(err)
		m.stats.IncItemError()
		return
	}

	m.logger.Debug("Data", utils.JsonStr(data))
	bs, err := json.Marshal(data)
	if err != nil {
		m.logger.Error(err)
		m.stats.IncItemError()
		return
	}

	if m.crawler.GetMode() == "test" {
		m.logger.Debug("current mode don't need save")
		m.stats.IncItemIgnore()
		return
	}

	if ctx == nil {
		ctx = context.Background()
	}
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
		m.stats.IncItemError()
		return
	}

	m.stats.IncItemSuccess()
	return
}

func (m *KafkaPipeline) FromCrawler(crawler pkg.Crawler) pkg.Pipeline {
	if m == nil {
		return new(KafkaPipeline).FromCrawler(crawler)
	}

	m.crawler = crawler
	m.stats = crawler.GetStats()
	m.logger = crawler.GetLogger()
	m.kafkaWriter = crawler.GetKafka()
	m.timeout = time.Minute
	return m
}
