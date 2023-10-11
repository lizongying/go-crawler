package pipelines

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/utils"
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

func (m *MongoPipeline) ProcessItem(itemWithContext pkg.ItemWithContext) (err error) {
	spider := m.GetSpider()
	if itemWithContext == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		spider.IncItemError()
		return
	}
	if itemWithContext.Name() != pkg.ItemMongo {
		m.logger.Warn("item not support", pkg.ItemMongo)
		return
	}
	itemMongo, ok := itemWithContext.GetItem().(*items.ItemMongo)
	if !ok {
		m.logger.Warn("item parsing failed with", pkg.ItemMongo)
		return
	}

	if itemWithContext == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		spider.IncItemError()
		return
	}

	if itemMongo.GetCollection() == "" {
		err = errors.New("collection is empty")
		m.logger.Error(err)
		spider.IncItemError()
		return
	}

	data := itemWithContext.Data()
	if data == nil {
		err = errors.New("nil data")
		m.logger.Error(err)
		spider.IncItemError()
		return
	}

	m.logger.Debug("Data", utils.JsonStr(data))
	bs, err := bson.Marshal(data)
	if err != nil {
		m.logger.Error(err)
		spider.IncItemError()
		return
	}

	if m.env == "dev" {
		m.logger.Debug("current mode don't need save")
		spider.IncItemIgnore()
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
		spider.IncItemError()
		return
	}

	spider.IncItemSuccess()
	return
}

func (m *MongoPipeline) FromSpider(spider pkg.Spider) pkg.Pipeline {
	if m == nil {
		return new(MongoPipeline).FromSpider(spider)
	}

	m.UnimplementedPipeline.FromSpider(spider)
	crawler := spider.GetCrawler()
	m.env = spider.GetConfig().GetEnv()
	m.logger = spider.GetLogger()
	m.mongoDb = crawler.GetMongoDb()
	m.timeout = time.Minute
	return m
}
