package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/devServer"
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

func (s *Spider) ParseMysql(ctx context.Context, response *pkg.Response) (err error) {
	var extra ExtraOk
	err = response.Request.GetExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.JsonStr(extra))
	s.logger.Info("response", string(response.BodyBytes))

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
		s.logger.Error(err)
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
		s.logger.Error(err)
	}
	return
}

func (s *Spider) ParseKafka(ctx context.Context, response *pkg.Response) (err error) {
	var extra ExtraOk
	err = response.Request.GetExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.JsonStr(extra))
	s.logger.Info("response", string(response.BodyBytes))

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
		s.logger.Error(err)
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
		s.logger.Error(err)
	}
	return
}

func (s *Spider) ParseMongo(ctx context.Context, response *pkg.Response) (err error) {
	var extra ExtraOk
	err = response.Request.GetExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.JsonStr(extra))
	s.logger.Info("response", string(response.BodyBytes))

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
		s.logger.Error(err)
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
		s.logger.Error(err)
	}
	return
}

func (s *Spider) ParseCsv(ctx context.Context, response *pkg.Response) (err error) {
	var extra ExtraOk
	err = response.Request.GetExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.JsonStr(extra))
	s.logger.Info("response", string(response.BodyBytes))

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
		s.logger.Error(err)
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
		s.logger.Error(err)
	}
	return
}

func (s *Spider) ParseJsonl(ctx context.Context, response *pkg.Response) (err error) {
	var extra ExtraOk
	err = response.Request.GetExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.JsonStr(extra))
	s.logger.Info("response", string(response.BodyBytes))

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
		s.logger.Error(err)
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
		s.logger.Error(err)
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
		s.logger.Error(err)
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
		s.logger.Error(err)
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
		s.logger.Error(err)
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
		s.logger.Error(err)
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
		s.logger.Error(err)
	}
	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	if baseSpider == nil {
		err = errors.New("nil baseSpider")
		return
	}

	logger := baseSpider.GetLogger()
	baseSpider.AddDevServerRoutes(devServer.NewOkHandler(logger))

	//baseSpider.SetPipeline(new(pipelines.MongoPipeline), 141)
	//baseSpider.SetPipeline(new(pipelines.CsvPipeline), 142)
	//baseSpider.SetPipeline(new(pipelines.JsonLinesPipeline), 143)
	//baseSpider.SetPipeline(new(pipelines.MysqlPipeline), 144)
	//baseSpider.SetPipeline(new(pipelines.KafkaPipeline), 145)

	spider = &Spider{
		Spider:         baseSpider,
		logger:         logger,
		collectionTest: "test",
		tableTest:      "test",
		topicTest:      "test",
		fileNameTest:   "test",
	}
	spider.SetName("test-item")

	return
}

func main() {
	app.NewApp(NewSpider).Run()
}
