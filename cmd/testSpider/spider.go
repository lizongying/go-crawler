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
	"github.com/lizongying/go-crawler/pkg/utils"
	"strings"
)

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) ParseOk(ctx context.Context, response *pkg.Response) (err error) {
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
	//if extra.Count%1000 == 0 {
	//	s.logger.Info("extra", utils.JsonStr(extra))
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
	//	s.logger.Error(err)
	//}
	return
}

func (s *Spider) ParseHttpAuth(ctx context.Context, response *pkg.Response) (err error) {
	var extra ExtraHttpAuth
	err = response.Request.GetExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.JsonStr(extra))
	s.logger.Info("response", string(response.BodyBytes))
	return
}

func (s *Spider) ParseCookie(ctx context.Context, response *pkg.Response) (err error) {
	var extra ExtraCookie
	err = response.Request.GetExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.JsonStr(extra))
	s.logger.Info("response", string(response.BodyBytes))

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
		s.logger.Error(err)
	}

	return
}

func (s *Spider) ParseGzip(ctx context.Context, response *pkg.Response) (err error) {
	s.logger.Info("response", string(response.BodyBytes))

	return
}

func (s *Spider) ParseDeflate(ctx context.Context, response *pkg.Response) (err error) {
	s.logger.Info("header", response.Header)
	s.logger.Info("body", string(response.BodyBytes))

	return
}

func (s *Spider) ParseRedirect(ctx context.Context, response *pkg.Response) (err error) {
	s.logger.Info("header", response.Header)
	s.logger.Info("body", string(response.BodyBytes))

	return
}

func (s *Spider) ParseImages(ctx context.Context, response *pkg.Response) (err error) {
	request := response.Request
	s.logger.Info("Images", utils.JsonStr(request))

	if ctx == nil {
		ctx = context.Background()
	}

	s.logger.Info("len", len(response.BodyBytes))

	return
}

// TestUrl go run cmd/testSpider/*.go -c dev.yml -f TestUrl -m dev
func (s *Spider) TestUrl(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewOkHandler(s.logger))
	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlOk+strings.Repeat("#", 10000))
	request.Extra = &ExtraOk{}
	request.CallBack = s.ParseOk
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestOk go run cmd/testSpider/*.go -c dev.yml -f TestOk -m dev
func (s *Spider) TestOk(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewOkHandler(s.logger))
	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlOk)
	request.Extra = &ExtraOk{}
	request.CallBack = s.ParseOk
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestHttpAuth go run cmd/testSpider/*.go -c dev.yml -f TestHttpAuth -m dev
func (s *Spider) TestHttpAuth(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewHttpAuthHandler(s.logger))
	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlHttpAuth)
	request.Extra = &ExtraHttpAuth{}
	request.CallBack = s.ParseHttpAuth
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestCookie go run cmd/testSpider/*.go -c dev.yml -f TestCookie -m dev
func (s *Spider) TestCookie(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewCookieHandler(s.logger))

	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlCookie)
	request.Extra = &ExtraCookie{}
	request.CallBack = s.ParseCookie
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestGzip go run cmd/testSpider/*.go -c dev.yml -f TestGzip -m dev
func (s *Spider) TestGzip(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewGzipHandler(s.logger))

	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlGzip)
	request.CallBack = s.ParseGzip
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestDeflate go run cmd/testSpider/*.go -c dev.yml -f TestDeflate -m dev
func (s *Spider) TestDeflate(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewDeflateHandler(s.logger))

	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlDeflate)
	request.CallBack = s.ParseDeflate
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestRedirect go run cmd/testSpider/*.go -c dev.yml -f TestRedirect -m dev
func (s *Spider) TestRedirect(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewRedirectHandler(s.logger))
	s.AddDevServerRoutes(devServer.NewOkHandler(s.logger))

	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlRedirect)
	request.CallBack = s.ParseRedirect
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.logger.Error(err)
	}
	return
}

func (s *Spider) TestImages(ctx context.Context, _ string) (err error) {
	request := new(pkg.Request)
	request.Url = "https://chinese.aljazeera.net/wp-content/uploads/2023/03/1-126.jpg"
	request.Extra = &ExtraTest{
		Image: new(media.Image),
	}
	request.CallBack = s.ParseImages
	err = s.YieldRequest(ctx, request)
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
		baseSpider.AddDevServerRoutes(devServer2.NewCustomHandler(logger))
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

	return
}

func main() {
	app.NewApp(NewSpider).Run()
}
