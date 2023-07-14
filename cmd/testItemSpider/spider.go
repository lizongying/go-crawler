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
	requestNext.SetUrl(response.Request.GetUrl())
	requestNext.SetExtra(&ExtraOk{
		Count: extra.Count + 1,
	})
	requestNext.SetCallback(s.ParseMysql)
	//requestNext.SetUniqueKey("1")
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
	requestNext.SetUrl(response.Request.GetUrl())
	requestNext.SetExtra(&ExtraOk{
		Count: extra.Count + 1,
	})
	requestNext.SetCallback(s.ParseKafka)
	//requestNext.SetUniqueKey("1")
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
	requestNext.SetUrl(response.Request.GetUrl())
	requestNext.SetExtra(&ExtraOk{
		Count: extra.Count + 1,
	})
	requestNext.SetCallback(s.ParseMongo)
	//requestNext.SetUniqueKey("1")
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
	requestNext.SetUrl(response.Request.GetUrl())
	requestNext.SetExtra(&ExtraOk{
		Count: extra.Count + 1,
	})
	requestNext.SetCallback(s.ParseCsv)
	//requestNext.SetUniqueKey("1")
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
	requestNext.SetUrl(response.Request.GetUrl())
	requestNext.SetExtra(&ExtraOk{
		Count: extra.Count + 1,
	})
	requestNext.SetCallback(s.ParseJsonl)
	//requestNext.SetUniqueKey("1")
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
	request.SetUrl(fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlOk))
	request.SetExtra(&ExtraOk{})
	request.SetCallback(s.ParseMongo)
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestMysql go run cmd/testItemSpider/*.go -c dev.yml -f TestMysql -m dev
func (s *Spider) TestMysql(ctx context.Context, _ string) (err error) {
	request := new(pkg.Request)
	request.SetUrl(fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlOk))
	request.SetExtra(&ExtraOk{})
	request.SetCallback(s.ParseMysql)
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestKafka go run cmd/testItemSpider/*.go -c dev.yml -f TestKafka -m dev
func (s *Spider) TestKafka(ctx context.Context, _ string) (err error) {
	request := new(pkg.Request)
	request.SetUrl(fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlOk))
	request.SetExtra(&ExtraOk{})
	request.SetCallback(s.ParseKafka)
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestCsv go run cmd/testItemSpider/*.go -c dev.yml -f TestCsv -m dev
func (s *Spider) TestCsv(ctx context.Context, _ string) (err error) {
	request := new(pkg.Request)
	request.SetUrl(fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlOk))
	request.SetExtra(&ExtraOk{})
	request.SetCallback(s.ParseCsv)
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestJsonl go run cmd/testItemSpider/*.go -c dev.yml -f TestJsonl -m dev
func (s *Spider) TestJsonl(ctx context.Context, _ string) (err error) {
	request := new(pkg.Request)
	request.SetUrl(fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlOk))
	request.SetExtra(&ExtraOk{})
	request.SetCallback(s.ParseJsonl)
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
	app.NewApp(NewSpider,
		//pkg.WithMongoPipeline(),
		//pkg.WithCsvPipeline(),
		pkg.WithJsonLinesPipeline(),
		//pkg.WithMysqlPipeline(),
		//pkg.WithKafkaPipeline(),
	).Run()
}
