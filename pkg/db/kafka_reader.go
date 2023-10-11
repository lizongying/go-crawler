package db

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/segmentio/kafka-go"
	"go.uber.org/fx"
	"strings"
)

func NewKafkaReader(config *config.Config, logger pkg.Logger, lc fx.Lifecycle) (kafkaReader *kafka.Reader, err error) {
	if !config.KafkaEnable {
		logger.Debug("Kafka Disable")
		return
	}

	uri := config.Kafka.Example.Uri
	if uri == "" {
		logger.Warn("uri is empty")
		return
	}

	kafkaReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  strings.Split(uri, ","),
		MaxBytes: 10e6, // 10MB
		Topic:    "test",
	})

	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) (err error) {
			if kafkaReader == nil {
				return
			}

			err = kafkaReader.Close()
			if err != nil {
				logger.Error(err)
				return
			}
			return
		},
	})
	return
}
