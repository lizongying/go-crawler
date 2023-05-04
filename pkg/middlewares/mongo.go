package middlewares

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/logger"
	"github.com/lizongying/go-crawler/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MongoMiddleware struct {
	pkg.UnimplementedMiddleware
	logger *logger.Logger

	mongoDb    *mongo.Database
	timeout    time.Duration
	spider     pkg.Spider
	spiderInfo *pkg.SpiderInfo
}

func (m *MongoMiddleware) GetName() string {
	return "mongo"
}

func (m *MongoMiddleware) SpiderStart(_ context.Context, spider pkg.Spider) (err error) {
	m.spider = spider
	m.spiderInfo = spider.GetInfo()
	m.spiderInfo.Stats.Store("item_error", 0)
	m.spiderInfo.Stats.Store("item_success", 0)
	return
}

func (m *MongoMiddleware) ProcessItem(c *pkg.Context) (err error) {
	item, ok := c.Item.(*pkg.ItemMongo)
	if !ok {
		m.logger.Warning("item not support mongo")
		err = c.NextItem()
		return
	}

	if item.Collection == "" {
		err = errors.New("collection is empty")
		m.logger.Error(err)
		err = c.NextItem()
		return
	}

	if item == nil {
		err = errors.New("item is empty")
		m.logger.Error(err)
		err = c.NextItem()
		return
	}
	m.logger.Debug("Data", utils.JsonStr(item.Data))
	bs, err := bson.Marshal(item.Data)
	if err != nil {
		m.logger.Error(err)
		err = c.NextItem()
		return
	}

	if m.spider.GetInfo().Mode == "test" {
		m.logger.Debug("current mode don't need save")
		err = c.NextItem()
		return
	}

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, m.timeout)
	defer cancel()

	res, err := m.mongoDb.Collection(item.Collection).InsertOne(ctx, bs)
	if err != nil {
		if item.Update && mongo.IsDuplicateKeyError(err) {
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
		itemError, ok := m.spiderInfo.Stats.Load("item_error")
		if ok {
			itemErrorInt := itemError.(int)
			itemErrorInt++
			m.spiderInfo.Stats.Store("item_error", itemErrorInt)
		}

		err = c.NextItem()
		return
	}

	itemSuccess, ok := m.spiderInfo.Stats.Load("item_success")
	if ok {
		itemSuccessInt := itemSuccess.(int)
		itemSuccessInt++
		m.spiderInfo.Stats.Store("item_success", itemSuccessInt)
	}

	err = c.NextItem()
	return
}

func NewMongoMiddleware(logger *logger.Logger, mongoDb *mongo.Database) (m pkg.Middleware) {
	m = &MongoMiddleware{
		logger:  logger,
		mongoDb: mongoDb,
		timeout: time.Minute,
	}
	return
}
