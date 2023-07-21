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

func (s *Spider) ParseOk(ctx context.Context, response pkg.Response) (err error) {
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

func (s *Spider) ParseHttpAuth(ctx context.Context, response pkg.Response) (err error) {
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

func (s *Spider) ParseCookie(ctx context.Context, response pkg.Response) (err error) {
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

func (s *Spider) ParseGzip(ctx context.Context, response pkg.Response) (err error) {
	s.logger.Info("response", string(response.GetBodyBytes()))

	return
}

func (s *Spider) ParseDeflate(ctx context.Context, response pkg.Response) (err error) {
	s.logger.Info("header", response.GetHeaders())
	s.logger.Info("body", string(response.GetBodyBytes()))

	return
}

func (s *Spider) ParseRedirect(ctx context.Context, response pkg.Response) (err error) {
	s.logger.Info("header", response.GetHeaders())
	s.logger.Info("body", string(response.GetBodyBytes()))

	return
}

func (s *Spider) ParseTimeout(ctx context.Context, response pkg.Response) (err error) {
	s.logger.Info("header", response.GetHeaders())
	s.logger.Info("body", string(response.GetBodyBytes()))

	return
}

func (s *Spider) ParseImages(ctx context.Context, response pkg.Response) (err error) {
	s.logger.Info("Images", utils.JsonStr(response.GetRequest()))
	s.logger.Info("len", len(response.GetBodyBytes()))

	return
}

// TestUrl go run cmd/testSpider/*.go -c dev.yml -f TestUrl -m dev
func (s *Spider) TestUrl(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewHandlerOk(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlOk+strings.Repeat("#", 10000))).
		SetExtra(&ExtraOk{}).
		SetCallBack(s.ParseOk))
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestOk go run cmd/testSpider/*.go -c dev.yml -f TestOk -m dev
func (s *Spider) TestOk(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewHandlerOk(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlOk)).
		SetExtra(&ExtraOk{}).
		SetCallBack(s.ParseOk))
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestHttpAuth go run cmd/testSpider/*.go -c dev.yml -f TestHttpAuth -m dev
func (s *Spider) TestHttpAuth(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewHandlerHttpAuth(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlHttpAuth)).
		SetExtra(&ExtraHttpAuth{}).
		SetCallBack(s.ParseHttpAuth))
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestCookie go run cmd/testSpider/*.go -c dev.yml -f TestCookie -m dev
func (s *Spider) TestCookie(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewHandlerCookie(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlCookie)).
		SetExtra(&ExtraCookie{}).
		SetCallBack(s.ParseCookie))
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestGzip go run cmd/testSpider/*.go -c dev.yml -f TestGzip -m dev
func (s *Spider) TestGzip(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewHandlerGzip(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlGzip)).
		SetCallBack(s.ParseGzip))
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestDeflate go run cmd/testSpider/*.go -c dev.yml -f TestDeflate -m dev
func (s *Spider) TestDeflate(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewHandlerDeflate(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlDeflate)).
		SetCallBack(s.ParseDeflate))
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestRedirect go run cmd/testSpider/*.go -c dev.yml -f TestRedirect -m dev
func (s *Spider) TestRedirect(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewHandlerRedirect(s.logger))
	s.AddDevServerRoutes(devServer.NewHandlerOk(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlRedirect)).
		SetCallBack(s.ParseRedirect))
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestTimeout go run cmd/testSpider/*.go -c dev.yml -f TestTimeout -m dev
func (s *Spider) TestTimeout(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewHandlerTimeout(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlTimeout)).
		SetTimeout(9*time.Second).
		SetCallBack(s.ParseTimeout))
	if err != nil {
		s.logger.Error(err)
	}
	return
}

func (s *Spider) TestImages(ctx context.Context, _ string) (err error) {
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
	//err = pkg.DontStopErr

	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	if baseSpider == nil {
		err = errors.New("nil baseSpider")
		return
	}
	//baseSpider.Username = "username"
	//baseSpider.Password = "password"

	logger := baseSpider.GetLogger()
	if baseSpider.GetMode() == "dev" {
		baseSpider.AddDevServerRoutes(devServer2.NewHandlerCustom(logger))
	}
	//baseSpider.Interval = 0
	//baseSpider.SetRequestRate("*", time.Second*3, 1)
	//baseSpider.AddOkHttpCodes(201)
	//baseSpider.SetMiddleware(new(middlewares.ImageMiddleware), 111)

	baseSpider.SetPlatforms(pkg.Windows, pkg.Mac, pkg.Android, pkg.Iphone, pkg.Ipad)
	baseSpider.SetBrowsers(pkg.Chrome, pkg.Edge, pkg.Safari, pkg.FireFox)

	//baseSpider.Stats = &stats.ImageStats{}

	spider = &Spider{
		Spider: baseSpider,
		logger: baseSpider.GetLogger(),
	}
	spider.SetName("test")
	host, _ := spider.GetConfig().GetDevServer()
	spider.SetHost(host.String())

	return
}

func main() {
	app.NewApp(NewSpider).Run()
}
