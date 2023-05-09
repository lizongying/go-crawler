package db

import (
	"context"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/lizongying/go-crawler/pkg/logger"
	"github.com/segmentio/kafka-go"
	"go.uber.org/fx"
	"strings"
)

func NewKafka(config *config.Config, logger *logger.Logger, lc fx.Lifecycle) (kafkaWriter *kafka.Writer, err error) {
	uri := config.Kafka.Example.Uri
	if uri == "" {
		err = errors.New("uri is empty")
		logger.Error(err)
		return
	}

	kafkaWriter = &kafka.Writer{
		Addr:                   kafka.TCP(strings.Split(uri, ",")...),
		AllowAutoTopicCreation: true,
	}

	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) (err error) {
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
