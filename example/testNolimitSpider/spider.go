package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/logger"
	"github.com/lizongying/go-crawler/pkg/middlewares"
	"github.com/lizongying/go-crawler/pkg/spider"
	"github.com/lizongying/go-crawler/pkg/utils"
)

type Spider struct {
	*spider.BaseSpider
	host string
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

func (s *Spider) RequestNoLimit(_ context.Context, response *pkg.Response) (err error) {
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
	requestNext.CallBack = s.RequestNoLimit
	//requestNext.UniqueKey = "1"
	err = s.YieldRequest(requestNext)
	if err != nil {
		s.Logger.Error(err)
	}
	err = s.YieldItem(&pkg.Item{
		UniqueKey: "1",
	})
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

func (s *Spider) ParseImages(ctx context.Context, response *pkg.Response) (err error) {
	request := response.Request
	s.Logger.Info("Images", utils.JsonStr(request))

	if ctx == nil {
		ctx = context.Background()
	}

	s.Logger.Info("len", len(response.BodyBytes))

	return
}

func (s *Spider) RequestImagesAsync(ctx context.Context, request *pkg.Request) (err error) {
	s.Logger.Info("Images", utils.JsonStr(request))
	request.CallBack = s.ParseImages

	if ctx == nil {
		ctx = context.Background()
	}

	err = s.YieldRequest(request)
	if err != nil {
		s.Logger.Error(err)
		return err
	}

	//s.Logger.Info("len", len(body))

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
	request.CallBack = s.RequestNoLimit
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
	baseSpider.OkHttpCodes = append(baseSpider.OkHttpCodes, 204)
	baseSpider.SetMiddleware(middlewares.NewImageMiddleware(logger), 3)
	spider = &Spider{
		BaseSpider: baseSpider,
		host:       "http://127.0.0.1:8081/%s",
	}

	return
}

func main() {
	app.NewApp(NewSpider).Run()
}
