package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/devServer"
	"github.com/lizongying/go-crawler/pkg/logger"
	"github.com/lizongying/go-crawler/pkg/spider"
	"github.com/lizongying/go-crawler/pkg/utils"
)

type Spider struct {
	*spider.BaseSpider

	collectionTest string
	tableTest      string
	topicTest      string
	fileNameTest   string
}

func (s *Spider) ParseMysql(ctx context.Context, response *pkg.Response) (err error) {
	extra := response.Request.Extra.(*ExtraOk)
	s.Logger.Info("extra", utils.JsonStr(extra))
	s.Logger.Info("response", string(response.BodyBytes))

	if extra.Count > 0 {
		return
	}
	requestNext := new(pkg.Request)
	requestNext.Url = response.Request.Url
	requestNext.Extra = &ExtraOk{
		Count: extra.Count + 1,
	}
	requestNext.CallBack = s.ParseMysql
	//requestNext.UniqueKey = "1"
	err = s.YieldRequest(ctx, requestNext)
	if err != nil {
		s.Logger.Error(err)
	}
	err = s.YieldItem(ctx, &pkg.ItemMysql{
		Update: true,
		Table:  s.tableTest,
		ItemUnimplemented: pkg.ItemUnimplemented{
			UniqueKey: "1",
			Id:        "3",
			Data: &DataOk{
				Id: "3",
				A:  0,
				B:  2,
				C:  "",
				D:  "2",
			},
		},
	})
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

func (s *Spider) ParseKafka(ctx context.Context, response *pkg.Response) (err error) {
	extra := response.Request.Extra.(*ExtraOk)
	s.Logger.Info("extra", utils.JsonStr(extra))
	s.Logger.Info("response", string(response.BodyBytes))

	if extra.Count > 0 {
		return
	}
	requestNext := new(pkg.Request)
	requestNext.Url = response.Request.Url
	requestNext.Extra = &ExtraOk{
		Count: extra.Count + 1,
	}
	requestNext.CallBack = s.ParseKafka
	//requestNext.UniqueKey = "1"
	err = s.YieldRequest(ctx, requestNext)
	if err != nil {
		s.Logger.Error(err)
	}
	err = s.YieldItem(ctx, &pkg.ItemKafka{
		Topic: s.tableTest,
		ItemUnimplemented: pkg.ItemUnimplemented{
			UniqueKey: "1",
			Id:        "3",
			Data: &DataOk{
				Id: "3",
				A:  0,
				B:  2,
				C:  "",
				D:  "2",
			},
		},
	})
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

func (s *Spider) ParseMongo(ctx context.Context, response *pkg.Response) (err error) {
	extra := response.Request.Extra.(*ExtraOk)
	s.Logger.Info("extra", utils.JsonStr(extra))
	s.Logger.Info("response", string(response.BodyBytes))

	if extra.Count > 0 {
		return
	}
	requestNext := new(pkg.Request)
	requestNext.Url = response.Request.Url
	requestNext.Extra = &ExtraOk{
		Count: extra.Count + 1,
	}
	requestNext.CallBack = s.ParseMongo
	//requestNext.UniqueKey = "1"
	err = s.YieldRequest(ctx, requestNext)
	if err != nil {
		s.Logger.Error(err)
	}
	err = s.YieldItem(ctx, &pkg.ItemMongo{
		Update:     true,
		Collection: s.collectionTest,
		ItemUnimplemented: pkg.ItemUnimplemented{
			UniqueKey: "1",
			Id:        extra.Count,
			Data: &DataOk{
				Id:    fmt.Sprintf(`%d,"%d"`, extra.Count, extra.Count),
				Count: extra.Count,
			},
		},
	})
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

func (s *Spider) ParseCsv(ctx context.Context, response *pkg.Response) (err error) {
	extra := response.Request.Extra.(*ExtraOk)
	s.Logger.Info("extra", utils.JsonStr(extra))
	s.Logger.Info("response", string(response.BodyBytes))

	if extra.Count > 2 {
		return
	}
	requestNext := new(pkg.Request)
	requestNext.Url = response.Request.Url
	requestNext.Extra = &ExtraOk{
		Count: extra.Count + 1,
	}
	requestNext.CallBack = s.ParseCsv
	//requestNext.UniqueKey = "1"
	err = s.YieldRequest(ctx, requestNext)
	if err != nil {
		s.Logger.Error(err)
	}
	err = s.YieldItem(ctx, &pkg.ItemCsv{
		FileName: s.fileNameTest,
		ItemUnimplemented: pkg.ItemUnimplemented{
			UniqueKey: "1",
			Id:        extra.Count,
			Data: &DataOk{
				Id:    fmt.Sprintf("%d,%d", extra.Count, extra.Count),
				Count: extra.Count,
			},
		},
	})
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

func (s *Spider) ParseJsonl(ctx context.Context, response *pkg.Response) (err error) {
	extra := response.Request.Extra.(*ExtraOk)
	s.Logger.Info("extra", utils.JsonStr(extra))
	s.Logger.Info("response", string(response.BodyBytes))

	if extra.Count > 2 {
		return
	}
	requestNext := new(pkg.Request)
	requestNext.Url = response.Request.Url
	requestNext.Extra = &ExtraOk{
		Count: extra.Count + 1,
	}
	requestNext.CallBack = s.ParseJsonl
	//requestNext.UniqueKey = "1"
	err = s.YieldRequest(ctx, requestNext)
	if err != nil {
		s.Logger.Error(err)
	}
	err = s.YieldItem(ctx, &pkg.ItemJsonl{
		FileName: s.fileNameTest,
		ItemUnimplemented: pkg.ItemUnimplemented{
			UniqueKey: "1",
			Id:        extra.Count,
			Data: &DataOk{
				Id:    fmt.Sprintf("%d,%d", extra.Count, extra.Count),
				Count: extra.Count,
			},
		},
	})
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

// TestMongo go run cmd/testItemSpider/*.go -c dev.yml -f TestMongo -m dev
func (s *Spider) TestMongo(ctx context.Context, _ string) (err error) {
	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlOk)
	request.Extra = &ExtraOk{}
	request.CallBack = s.ParseMongo
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

// TestMysql go run cmd/testItemSpider/*.go -c dev.yml -f TestMysql -m dev
func (s *Spider) TestMysql(ctx context.Context, _ string) (err error) {
	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlOk)
	request.Extra = &ExtraOk{}
	request.CallBack = s.ParseMysql
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

// TestKafka go run cmd/testItemSpider/*.go -c dev.yml -f TestKafka -m dev
func (s *Spider) TestKafka(ctx context.Context, _ string) (err error) {
	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlOk)
	request.Extra = &ExtraOk{}
	request.CallBack = s.ParseKafka
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

// TestCsv go run cmd/testItemSpider/*.go -c dev.yml -f TestCsv -m dev
func (s *Spider) TestCsv(ctx context.Context, _ string) (err error) {
	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlOk)
	request.Extra = &ExtraOk{}
	request.CallBack = s.ParseCsv
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

// TestJsonl go run cmd/testItemSpider/*.go -c dev.yml -f TestJsonl -m dev
func (s *Spider) TestJsonl(ctx context.Context, _ string) (err error) {
	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlOk)
	request.Extra = &ExtraOk{}
	request.CallBack = s.ParseJsonl
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

func NewSpider(baseSpider *spider.BaseSpider, logger *logger.Logger) (spider pkg.Spider, err error) {
	if baseSpider == nil {
		err = errors.New("nil baseSpider")
		logger.Error(err)
		return
	}
	baseSpider.Name = "test"
	baseSpider.AddDevServerRoutes(devServer.NewOkHandler(logger))

	//baseSpider.SetMiddleware(new(middlewares.MongoMiddleware), 141)
	//baseSpider.SetMiddleware(new(middlewares.CsvMiddleware), 142)
	//baseSpider.SetMiddleware(new(middlewares.JsonLinesMiddleware), 143)
	//baseSpider.SetMiddleware(new(middlewares.MysqlMiddleware), 144)
	//baseSpider.SetMiddleware(new(middlewares.KafkaMiddleware), 145)

	spider = &Spider{
		BaseSpider:     baseSpider,
		collectionTest: "test",
		tableTest:      "test",
		topicTest:      "test",
		fileNameTest:   "test",
	}

	return
}

func main() {
	app.NewApp(NewSpider).Run()
}
