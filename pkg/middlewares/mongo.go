package middlewares

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"time"
)

type MongoMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger

	mongoDb *mongo.Database
	timeout time.Duration
	spider  pkg.Spider
	info    *pkg.SpiderInfo
	stats   pkg.Stats
}

func (m *MongoMiddleware) SpiderStart(_ context.Context, spider pkg.Spider) (err error) {
	m.spider = spider
	m.info = spider.GetInfo()
	m.stats = spider.GetStats()
	return
}

func (m *MongoMiddleware) ProcessItem(c *pkg.Context) (err error) {
	item, ok := c.Item.(*pkg.ItemMongo)
	if !ok {
		m.logger.Warn("item not support mongo")
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

	if item.Collection == "" {
		err = errors.New("collection is empty")
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

	res, err := m.mongoDb.Collection(item.Collection).InsertOne(ctx, bs)
	if err != nil {
		if item.Update && !reflect.ValueOf(item.Id).IsZero() && mongo.IsDuplicateKeyError(err) {
			_, err = m.mongoDb.Collection(item.Collection).UpdateOne(ctx, bson.M{"_id": item.Id}, bson.M{"$set": item.Data})
			if err == nil {
				m.logger.Info(item.Collection, "update success", item.Id)
			}
		}
	} else {
		m.logger.Info(item.Collection, "insert success", res.InsertedID)
	}
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

func (m *MongoMiddleware) FromCrawler(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(MongoMiddleware).FromCrawler(spider)
	}
	m.logger = spider.GetLogger()
	m.mongoDb = spider.GetMongoDb()
	m.timeout = time.Minute
	return m
}
