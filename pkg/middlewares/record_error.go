package middlewares

import (
	"context"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"strings"
	"time"
)

type RecordError struct {
	Id      string `bson:"_id" json:"id"`
	TaskId  string `bson:"task_id" json:"task_id"`
	Request string `bson:"request" json:"request"`
	Error   string `bson:"error" json:"error"`
}

type RecordErrorMiddleware struct {
	pkg.UnimplementedMiddleware
	logger  pkg.Logger
	mongoDb *mongo.Database
}

func (m *RecordErrorMiddleware) ProcessError(ctx pkg.Context, response pkg.Response, err error) (next bool) {
	request, e := response.GetRequest().Marshal()
	if e != nil {
		return true
	}

	recordError := &RecordError{
		// TODO
		//Id:      fmt.Sprintf("%s-%s", c.TaskId, response.UniqueKey()),
		Id:      response.UniqueKey(),
		TaskId:  ctx.GetTaskId(),
		Request: string(request),
		Error:   err.Error(),
	}

	item := items.NewItemMongo(fmt.Sprintf("%s_%s", strings.ReplaceAll(m.GetSpider().Name(), "-", "_"), "error"), false)
	item.SetId(response.UniqueKey())
	item.SetData(recordError)
	if m.mongoDb != nil {
		e = m.ToMongo(item)
		if e != nil {
			return true
		}
	}

	return
}

func (m *RecordErrorMiddleware) ToMongo(item pkg.Item) (err error) {
	spider := m.GetSpider()
	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		spider.IncItemError()
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

	if item == nil {
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

	data := item.Data()
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

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	res, err := m.mongoDb.Collection(itemMongo.GetCollection()).InsertOne(ctx, bs)
	if err != nil {
		if itemMongo.GetUpdate() && !reflect.ValueOf(itemMongo.Id()).IsZero() && mongo.IsDuplicateKeyError(err) {
			_, err = m.mongoDb.Collection(itemMongo.GetCollection()).UpdateOne(ctx, bson.M{"_id": itemMongo.Id()}, bson.M{"$set": itemMongo.Data()})
			if err == nil {
				m.logger.Info("error", itemMongo.GetCollection(), "update success", itemMongo.Id())
			}
		}
	} else {
		m.logger.Info("error", itemMongo.GetCollection(), "insert success", res.InsertedID)
	}
	if err != nil {
		m.logger.Error(err)
		//spider.IncItemError()
		return
	}

	//spider.IncItemSuccess()
	return
}

func (m *RecordErrorMiddleware) FromSpider(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(RecordErrorMiddleware).FromSpider(spider)
	}

	m.UnimplementedMiddleware.FromSpider(spider)
	m.logger = spider.GetLogger()
	crawler := spider.GetCrawler()
	m.mongoDb = crawler.GetMongoDb()
	return m
}
