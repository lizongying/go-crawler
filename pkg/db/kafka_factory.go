package db

import (
	"context"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/segmentio/kafka-go"
	"go.uber.org/fx"
	"strings"
	"sync"
)

type KafkaFactory struct {
	config sync.Map
	logger pkg.Logger

	writers sync.Map
	readers sync.Map

	groupID string
}

func (m *KafkaFactory) GetWriter(name string, topic string) (kafkaWriter *kafka.Writer, err error) {
	if v, ok := m.writers.Load(name + topic); ok {
		return v.(*kafka.Writer), nil
	}

	c, ok := m.config.Load(name)
	if !ok {
		return nil, fmt.Errorf("kafka config %s not found", name)
	}

	conf := c.(pkg.Kafka)

	uri := conf.Uri
	if uri == "" {
		m.logger.Warn("uri is empty")
		return
	}

	kafkaWriter = &kafka.Writer{
		Addr:                   kafka.TCP(strings.Split(uri, ",")...),
		AllowAutoTopicCreation: true,
		Topic:                  topic,
		Logger: kafka.LoggerFunc(func(msg string, a ...interface{}) {
			m.logger.Info(topic, "insert success", a)
		}),
	}

	actual, loaded := m.writers.LoadOrStore(name+topic, kafkaWriter)
	if loaded {
		_ = kafkaWriter.Close()
	}

	return actual.(*kafka.Writer), nil
}

func (m *KafkaFactory) GetReader(name string, topic string) (kafkaReader *kafka.Reader, err error) {
	if v, ok := m.readers.Load(name + topic); ok {
		return v.(*kafka.Reader), nil
	}

	c, ok := m.config.Load(name)
	if !ok {
		return nil, fmt.Errorf("kafka config %s not found", name)
	}

	conf := c.(pkg.Kafka)

	uri := conf.Uri
	if uri == "" {
		m.logger.Warn("uri is empty")
		return
	}

	kafkaReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  strings.Split(uri, ","),
		MaxBytes: 10e6, // 10MB
		Topic:    topic,
		GroupID:  m.groupID,
	})

	actual, loaded := m.readers.LoadOrStore(name+topic, kafkaReader)
	if loaded {
		_ = kafkaReader.Close()
	}

	return actual.(*kafka.Reader), nil
}

func (m *KafkaFactory) Close(_ context.Context) error {
	m.writers.Range(func(key, value interface{}) bool {
		_ = value.(*kafka.Writer).Close()
		return true
	})
	m.readers.Range(func(key, value interface{}) bool {
		_ = value.(*kafka.Reader).Close()
		return true
	})
	return nil
}

func NewKafkaFactory(config *config.Config, logger pkg.Logger, lc fx.Lifecycle) (kafkaFactory *KafkaFactory, err error) {
	kafkaFactory = &KafkaFactory{
		logger:  logger,
		groupID: config.BotName,
	}
	for _, i := range config.KafkaList {
		kafkaFactory.config.Store(i.Name, i)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) (err error) {
			return kafkaFactory.Close(ctx)
		},
	})
	return
}
