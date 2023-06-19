package pipelines

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

type MongoPipeline struct {
	pkg.UnimplementedPipeline
	crawler pkg.Crawler
	stats   pkg.Stats
	logger  pkg.Logger
	mongoDb *mongo.Database
	timeout time.Duration
}

func (m *MongoPipeline) ProcessItem(ctx context.Context, item pkg.Item) (err error) {
	itemMongo, ok := item.(*pkg.ItemMongo)
	if !ok {
		m.logger.Warn("item not support mongo")
		return
	}

	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		m.stats.IncItemError()
		return
	}

	if itemMongo.Collection == "" {
		err = errors.New("collection is empty")
		m.logger.Error(err)
		m.stats.IncItemError()
		return
	}

	data := item.GetData()
	if data == nil {
		err = errors.New("nil data")
		m.logger.Error(err)
		m.stats.IncItemError()
		return
	}

	m.logger.Debug("Data", utils.JsonStr(data))
	bs, err := bson.Marshal(data)
	if err != nil {
		m.logger.Error(err)
		m.stats.IncItemError()
		return
	}

	if m.crawler.GetMode() == "test" {
		m.logger.Debug("current mode don't need save")
		m.stats.IncItemIgnore()
		return
	}

	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancel := context.WithTimeout(ctx, m.timeout)
	defer cancel()

	res, err := m.mongoDb.Collection(itemMongo.Collection).InsertOne(ctx, bs)
	if err != nil {
		if itemMongo.Update && !reflect.ValueOf(itemMongo.Id).IsZero() && mongo.IsDuplicateKeyError(err) {
			_, err = m.mongoDb.Collection(itemMongo.Collection).UpdateOne(ctx, bson.M{"_id": itemMongo.Id}, bson.M{"$set": itemMongo.Data})
			if err == nil {
				m.logger.Info(itemMongo.Collection, "update success", itemMongo.Id)
			}
		}
	} else {
		m.logger.Info(itemMongo.Collection, "insert success", res.InsertedID)
	}
	if err != nil {
		m.logger.Error(err)
		m.stats.IncItemError()
		return
	}

	m.stats.IncItemSuccess()
	return
}

func (m *MongoPipeline) FromCrawler(crawler pkg.Crawler) pkg.Pipeline {
	if m == nil {
		return new(MongoPipeline).FromCrawler(crawler)
	}

	m.crawler = crawler
	m.stats = crawler.GetStats()
	m.logger = crawler.GetLogger()
	m.mongoDb = crawler.GetMongoDb()
	m.timeout = time.Minute
	return m
}
