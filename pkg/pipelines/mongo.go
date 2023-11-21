package pipelines

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/items"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"time"
)

type MongoPipeline struct {
	pkg.UnimplementedPipeline
	env     string
	logger  pkg.Logger
	mongoDb *mongo.Database
	timeout time.Duration
}

func (m *MongoPipeline) ProcessItem(item pkg.Item) (err error) {
	task := item.GetContext().GetTask()

	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		task.IncItemError()
		return
	}

	if item.Name() != pkg.ItemMongo {
		m.logger.Warn("item not support", pkg.ItemMongo)
		return
	}

	itemMongo, ok := item.GetItem().(*items.ItemMongo)
	if !ok {
		m.logger.Warn("item parsing failed with", pkg.ItemMongo)
		return
	}

	if itemMongo.GetCollection() == "" {
		err = errors.New("collection is empty")
		m.logger.Error(err)
		task.IncItemError()
		return
	}

	data := item.Data()
	if data == nil {
		err = errors.New("nil data")
		m.logger.Error(err)
		task.IncItemError()
		return
	}

	item.GetContext().GetItem().WithSaved(true)

	bs, err := bson.Marshal(data)
	if err != nil {
		m.logger.Error(err)
		task.IncItemError()
		return
	}

	if m.env == "dev" {
		m.logger.Debug("current mode don't need save")
		task.IncItemIgnore()
		return
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, m.timeout)
	defer cancel()

	res, err := m.mongoDb.Collection(itemMongo.GetCollection()).InsertOne(ctx, bs)
	if err != nil {
		if itemMongo.GetUpdate() && !reflect.ValueOf(itemMongo.Id()).IsZero() && mongo.IsDuplicateKeyError(err) {
			_, err = m.mongoDb.Collection(itemMongo.GetCollection()).UpdateOne(ctx, bson.M{"_id": itemMongo.Id()}, bson.M{"$set": itemMongo.Data()})
			if err == nil {
				m.logger.Info(itemMongo.GetCollection(), "update success", itemMongo.Id())
			}
		}
	} else {
		m.logger.Info(itemMongo.GetCollection(), "insert success", res.InsertedID)
	}
	if err != nil {
		m.logger.Error(err)
		task.IncItemError()
		return
	}

	task.IncItemSuccess()
	return
}

func (m *MongoPipeline) FromSpider(spider pkg.Spider) (err error) {
	if m == nil {
		return new(MongoPipeline).FromSpider(spider)
	}

	if err = m.UnimplementedPipeline.FromSpider(spider); err != nil {
		return
	}
	crawler := spider.GetCrawler()
	m.env = spider.GetConfig().GetEnv()
	m.logger = spider.GetLogger()
	m.mongoDb = crawler.GetMongoDb()
	if m.mongoDb == nil {
		err = errors.New("mongoDb nil")
		return
	}
	m.timeout = time.Minute
	return
}
