package db

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/lizongying/go-crawler/pkg/logger"
	"github.com/segmentio/kafka-go"
	"go.uber.org/fx"
	"strings"
)

func NewKafka(config *config.Config, logger *logger.Logger, lc fx.Lifecycle) (kafkaWriter *kafka.Writer, err error) {
	if !config.KafkaEnable {
		logger.Debug("Kafka Disable")
		return
	}

	uri := config.Kafka.Example.Uri
	if uri == "" {
		logger.Warn("uri is empty")
		return
	}

	kafkaWriter = &kafka.Writer{
		Addr:                   kafka.TCP(strings.Split(uri, ",")...),
		AllowAutoTopicCreation: true,
	}

	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) (err error) {
			if kafkaWriter == nil {
				return
			}

			err = kafkaWriter.Close()
			if err != nil {
				logger.Error(err)
				return
			}
			return
		},
	})
	return
}
