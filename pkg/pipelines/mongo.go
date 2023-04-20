package pipelines

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

type MongoPipeline struct {
	pkg.UnimplementedPipeline
	logger *logger.Logger

	mongoDb    *mongo.Database
	timeout    time.Duration
	spider     pkg.Spider
	spiderInfo *pkg.SpiderInfo
}

func (p *MongoPipeline) GetName() string {
	return "mongo"
}

func (p *MongoPipeline) SpiderStart(_ context.Context, spider pkg.Spider) (err error) {
	p.spider = spider
	p.spiderInfo = spider.GetInfo()
	p.spiderInfo.Stats.Store("item_error", 0)
	p.spiderInfo.Stats.Store("item_success", 0)
	return
}

func (p *MongoPipeline) ProcessItem(ctx context.Context, item *pkg.Item) (err error) {
	if item.Collection == "" {
		err = errors.New("collection is empty")
		p.logger.Error(err)
		return
	}

	if item == nil {
		err = errors.New("item is empty")
		p.logger.Error(err)
		return
	}
	p.logger.Debug("Data", utils.JsonStr(item.Data))
	bs, err := bson.Marshal(item.Data)
	if err != nil {
		p.logger.Error(err)
		return
	}

	if p.spider.GetInfo().Mode == "test" {
		p.logger.Debug("mode test don't need save")
		return
	}

	if ctx != nil {
		ctx = context.Background()
	}

	c, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()
	res, e := p.mongoDb.Collection(item.Collection).InsertOne(c, bs)
	if e != nil {
		itemError, ok := p.spiderInfo.Stats.Load("item_error")
		if ok {
			itemErrorInt := itemError.(int)
			itemErrorInt++
			p.spiderInfo.Stats.Store("item_error", itemErrorInt)
		}
		p.logger.Error(e)
		return
	}
	itemSuccess, ok := p.spiderInfo.Stats.Load("item_success")
	if ok {
		itemSuccessInt := itemSuccess.(int)
		itemSuccessInt++
		p.spiderInfo.Stats.Store("item_success", itemSuccessInt)
	}
	p.logger.Info(item.Collection, "insert success", res.InsertedID)
	return
}

func (p *MongoPipeline) SpiderStop(_ context.Context) (err error) {
	return
}

func NewMongoPipeline(logger *logger.Logger, mongoDb *mongo.Database) (m pkg.Pipeline) {
	m = &MongoPipeline{
		logger:  logger,
		mongoDb: mongoDb,
		timeout: time.Minute,
	}
	return
}
