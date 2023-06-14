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
	"github.com/lizongying/go-crawler/pkg/stats"
	"github.com/lizongying/go-crawler/pkg/utils"
	"strings"
)

type Spider struct {
	*spider.BaseSpider
}

func (s *Spider) ParseOk(ctx context.Context, response *pkg.Response) (err error) {
	extra := response.Request.Extra.(*ExtraOk)
	s.Logger.Info("extra", utils.JsonStr(extra))
	s.Logger.Info("response", string(response.BodyBytes))

	if extra.Count > 0 {
		return
	}
	//if extra.Count%1000 == 0 {
	//	s.Logger.Info("extra", utils.JsonStr(extra))
	//}
	//requestNext := new(pkg.Request)
	//requestNext.Url = response.Request.Url
	//requestNext.Extra = &ExtraOk{
	//	Count: extra.Count + 1,
	//}
	//requestNext.CallBack = s.ParseOk
	////requestNext.UniqueKey = "1"
	//err = s.YieldRequest(ctx, requestNext)
	//if err != nil {
	//	s.Logger.Error(err)
	//}
	return
}

func (s *Spider) ParseHttpAuth(ctx context.Context, response *pkg.Response) (err error) {
	extra := response.Request.Extra.(*ExtraHttpAuth)
	s.Logger.Info("extra", utils.JsonStr(extra))
	s.Logger.Info("response", string(response.BodyBytes))
	return
}

func (s *Spider) ParseCookie(ctx context.Context, response *pkg.Response) (err error) {
	extra := response.Request.Extra.(*ExtraCookie)
	s.Logger.Info("extra", utils.JsonStr(extra))
	s.Logger.Info("response", string(response.BodyBytes))

	if extra.Count > 1 {
		return
	}

	requestNext := new(pkg.Request)
	requestNext.Url = response.Request.Url
	requestNext.Extra = &ExtraCookie{
		Count: extra.Count + 1,
	}
	requestNext.CallBack = s.ParseCookie
	//requestNext.UniqueKey = "1"
	err = s.YieldRequest(ctx, requestNext)
	if err != nil {
		s.Logger.Error(err)
	}

	return
}

func (s *Spider) ParseGzip(ctx context.Context, response *pkg.Response) (err error) {
	s.Logger.Info("response", string(response.BodyBytes))

	return
}

func (s *Spider) ParseDeflate(ctx context.Context, response *pkg.Response) (err error) {
	s.Logger.Info("header", response.Header)
	s.Logger.Info("body", string(response.BodyBytes))

	return
}

func (s *Spider) ParseRedirect(ctx context.Context, response *pkg.Response) (err error) {
	s.Logger.Info("header", response.Header)
	s.Logger.Info("body", string(response.BodyBytes))

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

// TestUrl go run cmd/testSpider/*.go -c dev.yml -f TestUrl -m dev
func (s *Spider) TestUrl(ctx context.Context, _ string) (err error) {
	if s.Mode == "dev" {
		s.AddDevServerRoutes(devServer.NewOkHandler(s.Logger))
	}
	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlOk+strings.Repeat("#", 10000))
	request.Extra = &ExtraOk{}
	request.CallBack = s.ParseOk
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

// TestOk go run cmd/testSpider/*.go -c dev.yml -f TestOk -m dev
func (s *Spider) TestOk(ctx context.Context, _ string) (err error) {
	if s.Mode == "dev" {
		s.AddDevServerRoutes(devServer.NewOkHandler(s.Logger))
	}
	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlOk)
	request.Extra = &ExtraOk{}
	request.CallBack = s.ParseOk
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

// TestHttpAuth go run cmd/testSpider/*.go -c dev.yml -f TestHttpAuth -m dev
func (s *Spider) TestHttpAuth(ctx context.Context, _ string) (err error) {
	if s.Mode == "dev" {
		s.AddDevServerRoutes(devServer.NewHttpAuthHandler(s.Logger))
	}
	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlHttpAuth)
	request.Extra = &ExtraHttpAuth{}
	request.CallBack = s.ParseHttpAuth
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

// TestCookie go run cmd/testSpider/*.go -c dev.yml -f TestCookie -m dev
func (s *Spider) TestCookie(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewCookieHandler(s.Logger))

	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlCookie)
	request.Extra = &ExtraCookie{}
	request.CallBack = s.ParseCookie
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

// TestGzip go run cmd/testSpider/*.go -c dev.yml -f TestGzip -m dev
func (s *Spider) TestGzip(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewGzipHandler(s.Logger))

	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlGzip)
	request.CallBack = s.ParseGzip
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

// TestDeflate go run cmd/testSpider/*.go -c dev.yml -f TestDeflate -m dev
func (s *Spider) TestDeflate(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewDeflateHandler(s.Logger))

	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlDeflate)
	request.CallBack = s.ParseDeflate
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

// TestRedirect go run cmd/testSpider/*.go -c dev.yml -f TestRedirect -m dev
func (s *Spider) TestRedirect(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewRedirectHandler(s.Logger))
	s.AddDevServerRoutes(devServer.NewOkHandler(s.Logger))

	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlRedirect)
	request.CallBack = s.ParseRedirect
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.Logger.Error(err)
	}
	return
}

func (s *Spider) TestImages(ctx context.Context, _ string) (err error) {
	request := new(pkg.Request)
	request.Url = "https://chinese.aljazeera.net/wp-content/uploads/2023/03/1-126.jpg"
	request.Extra = &ExtraTest{
		Image: new(pkg.Image),
	}
	request.CallBack = s.ParseImages
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.Logger.Error(err)
		return err
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
	baseSpider.Username = "username"
	baseSpider.Password = "password"
	if baseSpider.Mode == "dev" {
		baseSpider.AddDevServerRoutes(devServer2.NewCustomHandler(logger))
	}
	//baseSpider.Interval = 0
	//baseSpider.SetRequestRate("*", time.Second*3, 1)
	//baseSpider.AddOkHttpCodes(201)
	//baseSpider.SetMiddleware(new(middlewares.ImageMiddleware), 111)

	baseSpider.SetPlatforms(pkg.Windows, pkg.Mac, pkg.Android, pkg.Iphone, pkg.Ipad)
	baseSpider.SetBrowsers(pkg.Chrome, pkg.Edge, pkg.Safari, pkg.FireFox)

	baseSpider.Stats = &stats.ImageStats{}

	spider = &Spider{
		BaseSpider: baseSpider,
	}

	return
}

func main() {
	app.NewApp(NewSpider).Run()
}
