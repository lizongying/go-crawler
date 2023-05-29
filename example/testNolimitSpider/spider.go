package main

import (
	"context"
	"errors"
	"fmt"
	devServer2 "github.com/lizongying/go-crawler/internal/devServer"
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

func (s *Spider) ParseMysql(_ context.Context, response *pkg.Response) (err error) {
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
	err = s.YieldRequest(requestNext)
	if err != nil {
		s.Logger.Error(err)
	}
	err = s.YieldItem(&pkg.ItemMysql{
		Update: true,
		Table:  s.tableTest,
		ItemUnimplemented: pkg.ItemUnimplemented{
			UniqueKey: "1",
			Id:        3,
			Data: &DataMysql{
				Id: 3,
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

func (s *Spider) ParseKafka(_ context.Context, response *pkg.Response) (err error) {
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
	err = s.YieldRequest(requestNext)
	if err != nil {
		s.Logger.Error(err)
	}
	err = s.YieldItem(&pkg.ItemKafka{
		Topic: s.tableTest,
		ItemUnimplemented: pkg.ItemUnimplemented{
			UniqueKey: "1",
			Id:        3,
			Data: &DataMysql{
				Id: 3,
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

func (s *Spider) ParseOk(_ context.Context, response *pkg.Response) (err error) {
	extra := response.Request.Extra.(*ExtraOk)
	s.Logger.Info("extra", utils.JsonStr(extra))
	s.Logger.Debug("response", string(response.BodyBytes))

	if extra.Count > 10 {
		return
	}
	//if extra.Count%1000 == 0 {
	//	s.Logger.Info("extra", utils.JsonStr(extra))
	//}
	requestNext := new(pkg.Request)
	requestNext.Url = response.Request.Url
	requestNext.Extra = &ExtraOk{
		Count: extra.Count + 1,
	}
	requestNext.CallBack = s.ParseOk
	//requestNext.UniqueKey = "1"
	err = s.YieldRequest(requestNext)
	if err != nil {
		s.Logger.Error(err)
	}
	err = s.YieldItem(&pkg.ItemMongo{
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

func (s *Spider) ParseCsv(_ context.Context, response *pkg.Response) (err error) {
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
	err = s.YieldRequest(requestNext)
	if err != nil {
		s.Logger.Error(err)
	}
	err = s.YieldItem(&pkg.ItemCsv{
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

func (s *Spider) ParseJsonl(_ context.Context, response *pkg.Response) (err error) {
	extra := response.Request.Extra.(*ExtraOk)
	s.Logger.Info("request", response.Request.Header)
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
	err = s.YieldRequest(requestNext)
	if err != nil {
		s.Logger.Error(err)
	}
	err = s.YieldItem(&pkg.ItemJsonl{
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

func (s *Spider) TestMysql(_ context.Context, _ string) (err error) {
	if s.Mode == "dev" {
		s.AddDevServerRoutes(devServer.NewOkHandler(s.Logger))
	}
	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlOk)
	request.Extra = &ExtraOk{}
	request.CallBack = s.ParseMysql
	err = s.YieldRequest(request)
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

func (s *Spider) TestKafka(_ context.Context, _ string) (err error) {
	if s.Mode == "dev" {
		s.AddDevServerRoutes(devServer.NewOkHandler(s.Logger))
	}
	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlOk)
	request.Extra = &ExtraOk{}
	request.CallBack = s.ParseKafka
	err = s.YieldRequest(request)
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

// TestOk go run example/testNoLimitSpider/*.go -c dev.yml -f TestOk -m dev
func (s *Spider) TestOk(_ context.Context, _ string) (err error) {
	if s.Mode == "dev" {
		s.AddDevServerRoutes(devServer.NewOkHandler(s.Logger))
	}
	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlOk)
	request.Extra = &ExtraOk{}
	request.CallBack = s.ParseOk
	err = s.YieldRequest(request)
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

func (s *Spider) TestCsv(_ context.Context, _ string) (err error) {
	if s.Mode == "dev" {
		s.AddDevServerRoutes(devServer.NewOkHandler(s.Logger))
	}
	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlOk)
	request.Extra = &ExtraOk{}
	request.CallBack = s.ParseCsv
	err = s.YieldRequest(request)
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

func (s *Spider) TestJsonl(_ context.Context, _ string) (err error) {
	if s.Mode == "dev" {
		s.AddDevServerRoutes(devServer.NewOkHandler(s.Logger))
	}
	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlOk)
	request.Extra = &ExtraOk{}
	request.CallBack = s.ParseJsonl
	err = s.YieldRequest(request)
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

func (s *Spider) Stop(ctx context.Context) (err error) {
	err = s.BaseSpider.Stop(ctx)
	//err = pkg.DontStopErr

	return
}

func NewSpider(baseSpider *spider.BaseSpider, logger *logger.Logger) (spider pkg.Spider, err error) {
	if baseSpider == nil {
		err = errors.New("nil baseSpider")
		logger.Error(err)
		return
	}
	baseSpider.Name = "test"
	if baseSpider.Mode == "dev" {
		baseSpider.AddDevServerRoutes(devServer2.NewCustomHandler(logger))
	}
	//baseSpider.Interval = time.Millisecond
	baseSpider.
		AddOkHttpCodes(201)
	//SetMiddleware(middlewares.NewDeviceMiddleware, 100)
	//SetMiddleware(middlewares.NewMongoMiddleware, 141).
	//SetMiddleware(middlewares.NewCsvMiddleware, 142)
	//SetMiddleware(middlewares.NewJsonLinesMiddleware, 143).
	//SetMiddleware(middlewares.NewMysqlMiddleware, 144).
	//SetMiddleware(middlewares.NewKafkaMiddleware, 145)

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
