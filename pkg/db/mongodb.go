package db

import (
	"context"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/lizongying/go-crawler/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/fx"
)

func NewMongoDb(config *config.Config, logger *logger.Logger, lc fx.Lifecycle) (db *mongo.Database, err error) {
	if !config.MongoEnable {
		logger.Debug("Mongo Disable")
		return
	}

	uri := config.Mongo.Example.Uri
	if uri == "" {
		logger.Warn("uri is empty")
		return
	}

	database := config.Mongo.Example.Database
	if database == "" {
		logger.Warn("database is empty")
		return
	}

	var ctx = context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		logger.Error(err)
		return
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		logger.Error(err)
		return
	}

	db = client.Database(database)
	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) (err error) {
			if client != nil {
				return
			}

			err = client.Disconnect(ctx)
			if err != nil {
				logger.Error(err)
				return
			}
			return
		},
	})
	return
}
