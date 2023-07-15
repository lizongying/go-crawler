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
	err = response.Request.UnmarshalExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.JsonStr(extra))
	s.logger.Info("response", string(response.BodyBytes))

	if extra.Count > 0 {
		return
	}
	err = s.YieldRequest(ctx, new(pkg.Request).
		SetUrl(response.Request.GetUrl()).
		SetExtra(&ExtraOk{
			Count: extra.Count + 1,
		}).
		SetCallback(s.ParseMysql))
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
	err = response.Request.UnmarshalExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.JsonStr(extra))
	s.logger.Info("response", string(response.BodyBytes))

	if extra.Count > 0 {
		return
	}
	err = s.YieldRequest(ctx, new(pkg.Request).
		SetUrl(response.Request.GetUrl()).
		SetExtra(&ExtraOk{
			Count: extra.Count + 1,
		}).
		SetCallback(s.ParseKafka))
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
	err = response.Request.UnmarshalExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.JsonStr(extra))
	s.logger.Info("response", string(response.BodyBytes))

	if extra.Count > 0 {
		return
	}
	err = s.YieldRequest(ctx, new(pkg.Request).
		SetUrl(response.Request.GetUrl()).
		SetExtra(&ExtraOk{
			Count: extra.Count + 1,
		}).
		SetCallback(s.ParseMongo))
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
	err = response.Request.UnmarshalExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.JsonStr(extra))
	s.logger.Info("response", string(response.BodyBytes))

	if extra.Count > 2 {
		return
	}
	err = s.YieldRequest(ctx, new(pkg.Request).
		SetUrl(response.Request.GetUrl()).
		SetExtra(&ExtraOk{
			Count: extra.Count + 1,
		}).
		SetCallback(s.ParseCsv))
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
	err = response.Request.UnmarshalExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.JsonStr(extra))
	s.logger.Info("response", string(response.BodyBytes))

	if extra.Count > 2 {
		return
	}
	err = s.YieldRequest(ctx, new(pkg.Request).
		SetUrl(response.Request.GetUrl()).
		SetExtra(&ExtraOk{
			Count: extra.Count + 1,
		}).
		SetCallback(s.ParseJsonl))
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
		return
	}

	return
}

// TestMongo go run cmd/testItemSpider/*.go -c dev.yml -f TestMongo -m dev
func (s *Spider) TestMongo(ctx context.Context, _ string) (err error) {
	err = s.YieldRequest(ctx, new(pkg.Request).
		SetUrl(fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlOk)).
		SetExtra(&ExtraOk{}).
		SetCallback(s.ParseMongo))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestMysql go run cmd/testItemSpider/*.go -c dev.yml -f TestMysql -m dev
func (s *Spider) TestMysql(ctx context.Context, _ string) (err error) {
	err = s.YieldRequest(ctx, new(pkg.Request).
		SetUrl(fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlOk)).
		SetExtra(&ExtraOk{}).
		SetCallback(s.ParseMysql))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestKafka go run cmd/testItemSpider/*.go -c dev.yml -f TestKafka -m dev
func (s *Spider) TestKafka(ctx context.Context, _ string) (err error) {
	err = s.YieldRequest(ctx, new(pkg.Request).
		SetUrl(fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlOk)).
		SetExtra(&ExtraOk{}).
		SetCallback(s.ParseKafka))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestCsv go run cmd/testItemSpider/*.go -c dev.yml -f TestCsv -m dev
func (s *Spider) TestCsv(ctx context.Context, _ string) (err error) {
	err = s.YieldRequest(ctx, new(pkg.Request).
		SetUrl(fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlOk)).
		SetExtra(&ExtraOk{}).
		SetCallback(s.ParseCsv))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestJsonl go run cmd/testItemSpider/*.go -c dev.yml -f TestJsonl -m dev
func (s *Spider) TestJsonl(ctx context.Context, _ string) (err error) {
	err = s.YieldRequest(ctx, new(pkg.Request).
		SetUrl(fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlOk)).
		SetExtra(&ExtraOk{}).
		SetCallback(s.ParseJsonl))
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
