package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/devServer"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/request"
	"github.com/lizongying/go-crawler/pkg/utils"
)

type Spider struct {
	pkg.Spider
	logger         pkg.Logger
	collectionTest string
	tableTest      string
	topicTest      string
	fileNameTest   string
}

func (s *Spider) ParseMysql(ctx context.Context, response pkg.Response) (err error) {
	var extra ExtraOk
	err = response.UnmarshalExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.JsonStr(extra))
	s.logger.Info("response", string(response.GetBodyBytes()))

	if extra.Count > 0 {
		return
	}
	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(response.GetUrl()).
		SetExtra(&ExtraOk{
			Count: extra.Count + 1,
		}).
		SetCallBack(s.ParseMysql))
	if err != nil {
		s.logger.Error(err)
	}
	err = s.YieldItem(ctx, items.NewItemMysql(s.tableTest, true).
		SetUniqueKey("1").
		SetId("3").
		SetData(&DataOk{
			Id: "3",
			A:  0,
			B:  2,
			C:  "",
			D:  "2",
		}))
	if err != nil {
		s.logger.Error(err)
	}
	return
}

func (s *Spider) ParseKafka(ctx context.Context, response pkg.Response) (err error) {
	var extra ExtraOk
	err = response.UnmarshalExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.JsonStr(extra))
	s.logger.Info("response", string(response.GetBodyBytes()))

	if extra.Count > 0 {
		return
	}
	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(response.GetUrl()).
		SetExtra(&ExtraOk{
			Count: extra.Count + 1,
		}).
		SetCallBack(s.ParseKafka))
	if err != nil {
		s.logger.Error(err)
	}
	err = s.YieldItem(ctx, items.NewItemKafka(s.topicTest).
		SetUniqueKey("1").
		SetId("3").
		SetData(&DataOk{
			Id: "3",
			A:  0,
			B:  2,
			C:  "",
			D:  "2",
		}))
	if err != nil {
		s.logger.Error(err)
	}
	return
}

func (s *Spider) ParseMongo(ctx context.Context, response pkg.Response) (err error) {
	var extra ExtraOk
	err = response.UnmarshalExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.JsonStr(extra))
	s.logger.Info("response", string(response.GetBodyBytes()))

	if extra.Count > 0 {
		return
	}
	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(response.GetUrl()).
		SetExtra(&ExtraOk{
			Count: extra.Count + 1,
		}).
		SetCallBack(s.ParseMongo))
	if err != nil {
		s.logger.Error(err)
	}
	err = s.YieldItem(ctx, items.NewItemMongo(s.collectionTest, true).
		SetUniqueKey("1").
		SetId(extra.Count).
		SetData(&DataOk{
			Id:    fmt.Sprintf("%d,%d", extra.Count, extra.Count),
			Count: extra.Count,
		}))
	if err != nil {
		s.logger.Error(err)
	}
	return
}

func (s *Spider) ParseCsv(ctx context.Context, response pkg.Response) (err error) {
	var extra ExtraOk
	err = response.UnmarshalExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.JsonStr(extra))
	s.logger.Info("response", string(response.GetBodyBytes()))

	if extra.Count > 2 {
		return
	}
	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(response.GetUrl()).
		SetExtra(&ExtraOk{
			Count: extra.Count + 1,
		}).
		SetCallBack(s.ParseCsv))
	if err != nil {
		s.logger.Error(err)
	}
	err = s.YieldItem(ctx, items.NewItemCsv(s.fileNameTest).
		SetUniqueKey("1").
		SetId(extra.Count).
		SetData(&DataOk{
			Id:    fmt.Sprintf("%d,%d", extra.Count, extra.Count),
			Count: extra.Count,
		}))
	if err != nil {
		s.logger.Error(err)
	}
	return
}

func (s *Spider) ParseJsonl(ctx context.Context, response pkg.Response) (err error) {
	var extra ExtraOk
	err = response.UnmarshalExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.JsonStr(extra))
	s.logger.Info("response", string(response.GetBodyBytes()))

	if extra.Count > 2 {
		return
	}
	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(response.GetUrl()).
		SetExtra(&ExtraOk{
			Count: extra.Count + 1,
		}).
		SetCallBack(s.ParseJsonl))
	if err != nil {
		s.logger.Error(err)
	}
	err = s.YieldItem(ctx, items.NewItemJsonl(s.fileNameTest).
		SetUniqueKey("1").
		SetId(extra.Count).
		SetData(&DataOk{
			Id:    fmt.Sprintf("%d,%d", extra.Count, extra.Count),
			Count: extra.Count,
		}))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestMongo go run cmd/testItemSpider/*.go -c dev.yml -f TestMongo -m dev
func (s *Spider) TestMongo(ctx context.Context, _ string) (err error) {
	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlOk)).
		SetExtra(&ExtraOk{}).
		SetCallBack(s.ParseMongo))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestMysql go run cmd/testItemSpider/*.go -c dev.yml -f TestMysql -m dev
func (s *Spider) TestMysql(ctx context.Context, _ string) (err error) {
	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlOk)).
		SetExtra(&ExtraOk{}).
		SetCallBack(s.ParseMysql))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestKafka go run cmd/testItemSpider/*.go -c dev.yml -f TestKafka -m dev
func (s *Spider) TestKafka(ctx context.Context, _ string) (err error) {
	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlOk)).
		SetExtra(&ExtraOk{}).
		SetCallBack(s.ParseKafka))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestCsv go run cmd/testItemSpider/*.go -c dev.yml -f TestCsv -m dev
func (s *Spider) TestCsv(ctx context.Context, _ string) (err error) {
	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlOk)).
		SetExtra(&ExtraOk{}).
		SetCallBack(s.ParseCsv))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestJsonl go run cmd/testItemSpider/*.go -c dev.yml -f TestJsonl -m dev
func (s *Spider) TestJsonl(ctx context.Context, _ string) (err error) {
	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlOk)).
		SetExtra(&ExtraOk{}).
		SetCallBack(s.ParseJsonl))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	if baseSpider == nil {
		err = errors.New("nil baseSpider")
		return
	}

	logger := baseSpider.GetLogger()
	baseSpider.AddDevServerRoutes(devServer.NewHandlerOk(logger))

	spider = &Spider{
		Spider:         baseSpider,
		logger:         logger,
		collectionTest: "test",
		tableTest:      "test",
		topicTest:      "test",
		fileNameTest:   "test",
	}
	spider.SetName("test-item")
	host, _ := spider.GetConfig().GetDevServer()
	spider.SetHost(host.String())

	return
}

func main() {
	app.NewApp(NewSpider,
		//pkg.WithMongoPipeline(),
		//pkg.WithCsvPipeline(),
		pkg.WithJsonLinesPipeline(),
		//pkg.WithMysqlPipeline(),
		//pkg.WithKafkaPipeline(),
	).Run()
}
