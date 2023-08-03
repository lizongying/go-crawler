package main

import (
	"context"
	"errors"
	"fmt"
	devServer2 "github.com/lizongying/go-crawler/internal/devServer"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/devServer"
	"github.com/lizongying/go-crawler/pkg/media"
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
	err = response.UnmarshalExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.JsonStr(extra))
	s.logger.Info("response", string(response.GetBodyBytes()))
	return
}

func (s *Spider) ParseHttpAuth(_ pkg.Context, response pkg.Response) (err error) {
	var extra ExtraHttpAuth
	err = response.UnmarshalExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.JsonStr(extra))
	s.logger.Info("response", string(response.GetBodyBytes()))
	return
}

func (s *Spider) ParseCookie(ctx pkg.Context, response pkg.Response) (err error) {
	var extra ExtraCookie
	err = response.UnmarshalExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.JsonStr(extra))
	s.logger.Info("response", string(response.GetBodyBytes()))

	if extra.Count > 1 {
		return
	}

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(response.GetUrl()).
		SetExtra(&ExtraCookie{
			Count: extra.Count + 1,
		}).
		SetCallBack(s.ParseCookie))
	if err != nil {
		s.logger.Error(err)
	}

	return
}

func (s *Spider) ParseGzip(_ pkg.Context, response pkg.Response) (err error) {
	s.logger.Info("response", string(response.GetBodyBytes()))

	return
}

func (s *Spider) ParseDeflate(_ pkg.Context, response pkg.Response) (err error) {
	s.logger.Info("header", response.GetHeaders())
	s.logger.Info("body", string(response.GetBodyBytes()))

	return
}

func (s *Spider) ParseRedirect(_ pkg.Context, response pkg.Response) (err error) {
	s.logger.Info("header", response.GetHeaders())
	s.logger.Info("body", string(response.GetBodyBytes()))

	return
}

func (s *Spider) ParseTimeout(_ pkg.Context, response pkg.Response) (err error) {
	s.logger.Info("header", response.GetHeaders())
	s.logger.Info("body", string(response.GetBodyBytes()))

	return
}

func (s *Spider) ParseImages(_ pkg.Context, response pkg.Response) (err error) {
	s.logger.Info("Images", utils.JsonStr(response.GetRequest()))
	s.logger.Info("len", len(response.GetBodyBytes()))

	return
}

// TestUrl go run cmd/testSpider/*.go -c dev.yml -n test -f TestUrl -m dev
func (s *Spider) TestUrl(ctx pkg.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewRouteOk(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlOk+strings.Repeat("#", 10000))).
		SetExtra(&ExtraOk{}).
		SetCallBack(s.ParseOk))
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestOk go run cmd/testSpider/*.go -c dev.yml -n test -f TestOk -m dev
func (s *Spider) TestOk(ctx pkg.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewRouteOk(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlOk)).
		SetExtra(&ExtraOk{}).
		SetCallBack(s.ParseOk))
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestHttpAuth go run cmd/testSpider/*.go -c dev.yml -n test -f TestHttpAuth -m dev
func (s *Spider) TestHttpAuth(ctx pkg.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewRouteHttpAuth(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlHttpAuth)).
		SetExtra(&ExtraHttpAuth{}).
		SetCallBack(s.ParseHttpAuth))
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestCookie go run cmd/testSpider/*.go -c dev.yml -n test -f TestCookie -m dev
func (s *Spider) TestCookie(ctx pkg.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewRouteCookie(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlCookie)).
		SetExtra(&ExtraCookie{}).
		SetCallBack(s.ParseCookie))
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestGzip go run cmd/testSpider/*.go -c dev.yml -n test -f TestGzip -m dev
func (s *Spider) TestGzip(ctx pkg.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewRouteGzip(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlGzip)).
		SetCallBack(s.ParseGzip))
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestDeflate go run cmd/testSpider/*.go -c dev.yml -n test -f TestDeflate -m dev
func (s *Spider) TestDeflate(ctx pkg.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewRouteDeflate(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlDeflate)).
		SetCallBack(s.ParseDeflate))
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestRedirect go run cmd/testSpider/*.go -c dev.yml -n test -f TestRedirect -m dev
func (s *Spider) TestRedirect(ctx pkg.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewRouteRedirect(s.logger))
	s.AddDevServerRoutes(devServer.NewRouteOk(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlRedirect)).
		SetCallBack(s.ParseRedirect))
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestTimeout go run cmd/testSpider/*.go -c dev.yml -n test -f TestTimeout -m dev
func (s *Spider) TestTimeout(ctx pkg.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewRouteTimeout(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlTimeout)).
		SetTimeout(9*time.Second).
		SetCallBack(s.ParseTimeout))
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestImages go run cmd/testSpider/*.go -c dev.yml -n test -f TestImages -m dev
func (s *Spider) TestImages(ctx pkg.Context, _ string) (err error) {
	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl("https://chinese.aljazeera.net/wp-content/uploads/2023/03/1-126.jpg").
		SetExtra(&ExtraTest{
			Image: new(media.Image),
		}).
		SetCallBack(s.ParseImages))
	if err != nil {
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
	if baseSpider == nil {
		err = errors.New("nil baseSpider")
		return
	}

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
	app.NewApp(NewSpider).Run(pkg.WithDevServerRoute(devServer2.NewRouteCustom))
}
