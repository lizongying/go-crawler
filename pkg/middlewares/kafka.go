package middlewares

import (
	"context"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/logger"
	"github.com/lizongying/go-crawler/pkg/utils"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type KafkaMiddleware struct {
	pkg.UnimplementedMiddleware
	logger *logger.Logger

	kafkaWriter *kafka.Writer
	timeout     time.Duration
	spider      pkg.Spider
	info        *pkg.SpiderInfo
	stats       pkg.Stats
}

func (m *KafkaMiddleware) SpiderStart(_ context.Context, spider pkg.Spider) (err error) {
	m.spider = spider
	m.info = spider.GetInfo()
	m.stats = spider.GetStats()
	return
}

func (m *KafkaMiddleware) ProcessItem(c *pkg.Context) (err error) {
	item, ok := c.Item.(*pkg.ItemKafka)
	if !ok {
		m.logger.Warn("item not support kafka")
		err = c.NextItem()
		return
	}

	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		m.stats.IncItemError()
		err = c.NextItem()
		return
	}

	if item.Topic == "" {
		err = errors.New("topic is empty")
		m.logger.Error(err)
		m.stats.IncItemError()
		err = c.NextItem()
		return
	}

	data := item.GetData()
	if data == nil {
		err = errors.New("nil data")
		m.logger.Error(err)
		m.stats.IncItemError()
		err = c.NextItem()
		return
	}

	m.logger.Debug("Data", utils.JsonStr(data))
	bs, err := bson.Marshal(data)
	if err != nil {
		m.logger.Error(err)
		m.stats.IncItemError()
		err = c.NextItem()
		return
	}

	if m.info.Mode == "test" {
		m.logger.Debug("current mode don't need save")
		m.stats.IncItemIgnore()
		err = c.NextItem()
		return
	}

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, m.timeout)
	defer cancel()

	m.kafkaWriter.Topic = item.Topic
	m.kafkaWriter.Logger = kafka.LoggerFunc(func(msg string, a ...interface{}) {
		m.logger.Info(item.Topic, "insert success", a)
	})
	err = m.kafkaWriter.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte(fmt.Sprint(item.Id)),
			Value: bs,
		},
	)
	if err != nil {
		m.logger.Error(err)
		m.stats.IncItemError()
		err = c.NextItem()
		return
	}

	m.stats.IncItemSuccess()
	err = c.NextItem()
	return
}

func NewKafkaMiddleware(logger *logger.Logger, kafkaWriter *kafka.Writer) (m pkg.Middleware) {
	m = &KafkaMiddleware{
		logger:      logger,
		kafkaWriter: kafkaWriter,
		timeout:     time.Minute,
	}
	return
}
