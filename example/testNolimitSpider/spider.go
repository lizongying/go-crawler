package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/httpServer"
	"github.com/lizongying/go-crawler/pkg/logger"
	"github.com/lizongying/go-crawler/pkg/middlewares"
	"github.com/lizongying/go-crawler/pkg/spider"
	"github.com/lizongying/go-crawler/pkg/utils"
)

type Spider struct {
	*spider.BaseSpider
	host string

	collectionTest string
	fileNameTest   string
}

func (s *Spider) RequestNoLimitSync(ctx context.Context, request *pkg.Request) (err error) {
	extra := request.Extra.(*ExtraNoLimit)
	s.Logger.Info("extra", utils.JsonStr(extra))

	if ctx == nil {
		ctx = context.Background()
	}

	response, err := s.Request(ctx, request)
	if err != nil {
		s.Logger.Error(err)
		return err
	}

	s.Logger.Info("response", string(response.BodyBytes))

	if extra.Count > 1000 {
		return
	}
	requestNext := new(pkg.Request)
	requestNext.Url = fmt.Sprintf(s.host, "no-limit")
	requestNext.Extra = &ExtraNoLimit{
		Count: extra.Count + 1,
	}
	err = s.RequestNoLimitSync(nil, requestNext)
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

func (s *Spider) ResponseNoLimit(_ context.Context, response *pkg.Response) (err error) {
	extra := response.Request.Extra.(*ExtraNoLimit)
	s.Logger.Info("extra", utils.JsonStr(extra))

	s.Logger.Info("response", string(response.BodyBytes))

	if extra.Count > 30 {
		return
	}
	requestNext := new(pkg.Request)
	requestNext.Url = fmt.Sprintf(s.host, "no-limit")
	requestNext.Extra = &ExtraNoLimit{
		Count: extra.Count + 1,
	}
	requestNext.CallBack = s.ResponseNoLimit
	//requestNext.UniqueKey = "1"
	err = s.YieldRequest(requestNext)
	if err != nil {
		s.Logger.Error(err)
	}
	err = s.YieldItem(&pkg.ItemMongo{
		UniqueKey: "1",
	})
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

func (s *Spider) ResponseOk(_ context.Context, response *pkg.Response) (err error) {
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
	requestNext.CallBack = s.ResponseOk
	//requestNext.UniqueKey = "1"
	err = s.YieldRequest(requestNext)
	if err != nil {
		s.Logger.Error(err)
	}
	err = s.YieldItem(&pkg.ItemMongo{
		Collection: s.collectionTest,
		UniqueKey:  "1",
		Id:         extra.Count,
		Update:     true,
		Data: DataOk{
			Id:    fmt.Sprintf(`%d,"%d"`, extra.Count, extra.Count),
			Count: extra.Count,
		},
	})
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

func (s *Spider) ResponseCsv(_ context.Context, response *pkg.Response) (err error) {
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
	requestNext.CallBack = s.ResponseCsv
	//requestNext.UniqueKey = "1"
	err = s.YieldRequest(requestNext)
	if err != nil {
		s.Logger.Error(err)
	}
	err = s.YieldItem(&pkg.ItemCsv{
		FileName:  s.fileNameTest,
		UniqueKey: "1",
		Id:        extra.Count,
		Data: DataOk{
			Id:    fmt.Sprintf("%d,%d", extra.Count, extra.Count),
			Count: extra.Count,
		},
	})
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

func (s *Spider) ResponseJsonl(_ context.Context, response *pkg.Response) (err error) {
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
	requestNext.CallBack = s.ResponseJsonl
	//requestNext.UniqueKey = "1"
	err = s.YieldRequest(requestNext)
	if err != nil {
		s.Logger.Error(err)
	}
	err = s.YieldItem(&pkg.ItemJsonl{
		FileName:  s.fileNameTest,
		UniqueKey: "1",
		Id:        extra.Count,
		Data: DataOk{
			Id:    fmt.Sprintf("%d,%d", extra.Count, extra.Count),
			Count: extra.Count,
		},
	})
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

func (s *Spider) TestNoLimitSync(_ context.Context) (err error) {
	request := new(pkg.Request)
	request.Url = fmt.Sprintf(s.host, "no-limit")
	request.Extra = &ExtraNoLimit{}
	err = s.RequestNoLimitSync(nil, request)
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

func (s *Spider) TestNoLimit(_ context.Context) (err error) {
	request := new(pkg.Request)
	request.Url = fmt.Sprintf(s.host, "no-limit")
	request.Extra = &ExtraNoLimit{}
	request.CallBack = s.ResponseNoLimit
	err = s.YieldRequest(request)
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

func (s *Spider) TestOk(_ context.Context) (err error) {
	if s.Mode == "dev" {
		s.GetDevServer().AddRoutes(httpServer.NewOkHandler(s.Logger))
	}
	request := new(pkg.Request)
	request.Url = fmt.Sprintf(s.host, httpServer.UrlOk)
	request.Extra = &ExtraOk{}
	request.CallBack = s.ResponseOk
	err = s.YieldRequest(request)
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

func (s *Spider) TestCsv(_ context.Context) (err error) {
	if s.Mode == "dev" {
		s.GetDevServer().AddRoutes(httpServer.NewOkHandler(s.Logger))
	}
	request := new(pkg.Request)
	request.Url = fmt.Sprintf(s.host, httpServer.UrlOk)
	request.Extra = &ExtraOk{}
	request.CallBack = s.ResponseCsv
	err = s.YieldRequest(request)
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

func (s *Spider) TestJsonl(_ context.Context) (err error) {
	if s.Mode == "dev" {
		s.GetDevServer().AddRoutes(httpServer.NewOkHandler(s.Logger))
	}
	request := new(pkg.Request)
	request.Url = fmt.Sprintf(s.host, httpServer.UrlOk)
	request.Extra = &ExtraOk{}
	request.CallBack = s.ResponseJsonl
	err = s.YieldRequest(request)
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
	baseSpider.AddOkHttpCodes(201)
	//baseSpider.SetMiddleware(middlewares.NewMongoMiddleware(logger, baseSpider.MongoDb), 141)
	baseSpider.SetMiddleware(middlewares.NewCsvMiddleware(logger), 142)
	baseSpider.SetMiddleware(middlewares.NewJsonLinesMiddleware(logger), 143)

	spider = &Spider{
		BaseSpider:     baseSpider,
		host:           "http://127.0.0.1:8081/%s",
		collectionTest: "test",
		fileNameTest:   "test",
	}

	return
}

func main() {
	app.NewApp(NewSpider).Run()
}
