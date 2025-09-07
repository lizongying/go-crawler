package db

import (
	"context"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/fx"
	"sync"
)

type MongoFactory struct {
	config sync.Map
	logger pkg.Logger

	clients sync.Map
}

func (m *MongoFactory) GetClient(name string) (db *mongo.Database, err error) {
	if v, ok := m.clients.Load(name); ok {
		return v.(*mongo.Database), nil
	}

	c, ok := m.config.Load(name)
	if !ok {
		return nil, fmt.Errorf("mongo config %s not found", name)
	}

	conf := c.(pkg.Mongo)

	uri := conf.Uri
	if uri == "" {
		m.logger.Warn("uri is empty")
		return
	}

	database := conf.Database
	if database == "" {
		m.logger.Warn("database is empty")
		return
	}

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		m.logger.Error(err)
		return
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		m.logger.Error(err)
		return
	}

	db = client.Database(database)

	actual, loaded := m.clients.LoadOrStore(name, db)
	if loaded {
		_ = client.Disconnect(ctx)
	}

	return actual.(*mongo.Database), nil
}

func (m *MongoFactory) Close(_ context.Context) error {
	ctx := context.Background()
	m.clients.Range(func(key, value interface{}) bool {
		_ = value.(*mongo.Database).Client().Disconnect(ctx)
		return true
	})
	return nil
}

func NewMongoFactory(config *config.Config, logger pkg.Logger, lc fx.Lifecycle) (mongoFactory *MongoFactory, err error) {
	mongoFactory = &MongoFactory{
		logger: logger,
	}
	for _, i := range config.MongoList {
		mongoFactory.config.Store(i.Name, i)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) (err error) {
			return mongoFactory.Close(ctx)
		},
	})
	return
}
