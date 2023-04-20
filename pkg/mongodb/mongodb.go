package mongodb

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/lizongying/go-crawler/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoDb(config *config.Config, logger *logger.Logger) (mongoDb *mongo.Database, err error) {
	uri := config.Mongo.Example.Uri
	if uri == "" {
		err = errors.New("uri is empty")
		logger.Error(err)
		return
	}

	database := config.Mongo.Example.Database
	if database == "" {
		err = errors.New("database is empty")
		logger.Error(err)
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

	mongoDb = client.Database(database)
	return
}
