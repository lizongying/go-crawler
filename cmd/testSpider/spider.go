package main

import (
	"context"
	"fmt"
	mockServer2 "github.com/lizongying/go-crawler/internal/mockServer"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/media"
	"github.com/lizongying/go-crawler/pkg/mockServer"
	"github.com/lizongying/go-crawler/pkg/request"
	"github.com/lizongying/go-crawler/pkg/utils"
	"strings"
	"time"
)

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) ParseOk(_ pkg.Context, response pkg.Response) (err error) {
	var extra ExtraOk
	if err = response.UnmarshalExtra(&extra); err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.JsonStr(extra))
	s.logger.Info("response", response.BodyStr())
	return
}

func (s *Spider) ParseHttpAuth(_ pkg.Context, response pkg.Response) (err error) {
	var extra ExtraHttpAuth
	if err = response.UnmarshalExtra(&extra); err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.JsonStr(extra))
	s.logger.Info("response", response.BodyStr())
	return
}

func (s *Spider) ParseCookie(ctx pkg.Context, response pkg.Response) (err error) {
	var extra ExtraCookie
	if err = response.UnmarshalExtra(&extra); err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.JsonStr(extra))
	s.logger.Info("response", response.BodyStr())

	if extra.Count > 1 {
		return
	}

	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(response.GetUrl()).
		SetExtra(&ExtraCookie{
			Count: extra.Count + 1,
		}).
		SetCallBack(s.ParseCookie)); err != nil {
		s.logger.Error(err)
	}

	return
}

func (s *Spider) ParseGzip(_ pkg.Context, response pkg.Response) (err error) {
	s.logger.Info("response", response.BodyStr())

	return
}

func (s *Spider) ParseDeflate(_ pkg.Context, response pkg.Response) (err error) {
	s.logger.Info("header", response.Headers())
	s.logger.Info("body", response.BodyStr())

	return
}

func (s *Spider) ParseRedirect(_ pkg.Context, response pkg.Response) (err error) {
	s.logger.Info("header", response.Headers())
	s.logger.Info("body", response.BodyStr())

	return
}

func (s *Spider) ParseTimeout(_ pkg.Context, response pkg.Response) (err error) {
	s.logger.Info("header", response.Headers())
	s.logger.Info("body", response.BodyStr())

	return
}

func (s *Spider) ParseImages(_ pkg.Context, response pkg.Response) (err error) {
	s.logger.Info("Images", utils.JsonStr(response.GetRequest()))
	s.logger.Info("len", len(response.BodyBytes()))

	return
}

// TestUrl go run cmd/testSpider/*.go -c dev.yml -n test -f TestUrl -m dev
func (s *Spider) TestUrl(ctx pkg.Context, _ string) (err error) {
	s.AddMockServerRoutes(mockServer.NewRouteOk(s.logger))

	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mockServer.UrlOk+strings.Repeat("#", 10000))).
		SetExtra(&ExtraOk{}).
		SetCallBack(s.ParseOk)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestOk go run cmd/testSpider/*.go -c dev.yml -n test -f TestOk -m dev
func (s *Spider) TestOk(ctx pkg.Context, _ string) (err error) {
	s.AddMockServerRoutes(mockServer.NewRouteOk(s.logger))

	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mockServer.UrlOk)).
		SetExtra(&ExtraOk{}).
		SetCallBack(s.ParseOk)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestHttpAuth go run cmd/testSpider/*.go -c dev.yml -n test -f TestHttpAuth -m dev
func (s *Spider) TestHttpAuth(ctx pkg.Context, _ string) (err error) {
	s.AddMockServerRoutes(mockServer.NewRouteHttpAuth(s.logger))

	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mockServer.UrlHttpAuth)).
		SetExtra(&ExtraHttpAuth{}).
		SetCallBack(s.ParseHttpAuth)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestCookie go run cmd/testSpider/*.go -c dev.yml -n test -f TestCookie -m dev
func (s *Spider) TestCookie(ctx pkg.Context, _ string) (err error) {
	s.AddMockServerRoutes(mockServer.NewRouteCookie(s.logger))

	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mockServer.UrlCookie)).
		SetExtra(&ExtraCookie{}).
		SetCallBack(s.ParseCookie)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestGzip go run cmd/testSpider/*.go -c dev.yml -n test -f TestGzip -m dev
func (s *Spider) TestGzip(ctx pkg.Context, _ string) (err error) {
	s.AddMockServerRoutes(mockServer.NewRouteGzip(s.logger))

	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mockServer.UrlGzip)).
		SetCallBack(s.ParseGzip)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestDeflate go run cmd/testSpider/*.go -c dev.yml -n test -f TestDeflate -m dev
func (s *Spider) TestDeflate(ctx pkg.Context, _ string) (err error) {
	s.AddMockServerRoutes(mockServer.NewRouteDeflate(s.logger))

	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mockServer.UrlDeflate)).
		SetCallBack(s.ParseDeflate)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestRedirect go run cmd/testSpider/*.go -c dev.yml -n test -f TestRedirect -m dev
func (s *Spider) TestRedirect(ctx pkg.Context, _ string) (err error) {
	s.AddMockServerRoutes(mockServer.NewRouteRedirect(s.logger))
	s.AddMockServerRoutes(mockServer.NewRouteOk(s.logger))

	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mockServer.UrlRedirect)).
		SetCallBack(s.ParseRedirect)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestTimeout go run cmd/testSpider/*.go -c dev.yml -n test -f TestTimeout -m dev
func (s *Spider) TestTimeout(ctx pkg.Context, _ string) (err error) {
	s.AddMockServerRoutes(mockServer.NewRouteTimeout(s.logger))

	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mockServer.UrlTimeout)).
		SetTimeout(9*time.Second).
		SetCallBack(s.ParseTimeout)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestImages go run cmd/testSpider/*.go -c dev.yml -n test -f TestImages -m dev
func (s *Spider) TestImages(ctx pkg.Context, _ string) (err error) {
	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl("https://chinese.aljazeera.net/wp-content/uploads/2023/03/1-126.jpg").
		SetExtra(&ExtraTest{
			Image: new(media.Image),
		}).
		SetCallBack(s.ParseImages)); err != nil {
		s.logger.Error(err)
		return err
	}

	return
}

func (s *Spider) Stop(ctx context.Context) (err error) {
	err = s.Spider.Stop(ctx)
	if err != nil {
		s.logger.Error(err)
		return
	}

	//err = pkg.DontStopErr
	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	spider = &Spider{
		Spider: baseSpider,
		logger: baseSpider.GetLogger(),
	}
	spider.WithOptions(
		pkg.WithName("test"),
		pkg.WithHost("https://localhost:8081"),
		//pkg.WithUsername("username"),
		//pkg.WithPassword("password"),
		//pkg.WithStats(&stats.ImageStats{}),
	)

	return
}

func main() {
	app.NewApp(NewSpider).Run(pkg.WithMockServerRoute(mockServer2.NewRouteCustom))
}
