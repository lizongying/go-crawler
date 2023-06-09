package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/devServer"
	"github.com/lizongying/go-crawler/pkg/logger"
	"github.com/lizongying/go-crawler/pkg/spider"
)

type Spider struct {
	*spider.BaseSpider
}

func (s *Spider) ParseDecode(_ context.Context, response *pkg.Response) (err error) {
	s.Logger.Info("header", response.Header)
	s.Logger.Info("body", string(response.BodyBytes))
	return
}

// TestGbk go run cmd/testDecodeSpider/*.go -c dev.yml -f TestGbk -m dev
func (s *Spider) TestGbk(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewGbkHandler(s.Logger))

	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlGbk)
	request.CallBack = s.ParseDecode
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.Logger.Error(err)
	}
	s.Logger.Info("url", request.Url)
	return
}

// TestGb2312 go run cmd/testDecodeSpider/*.go -c dev.yml -f TestGb2312 -m dev
func (s *Spider) TestGb2312(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewGb2312Handler(s.Logger))

	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlGb2312)
	request.CallBack = s.ParseDecode
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.Logger.Error(err)
	}
	s.Logger.Info("url", request.Url)
	return
}

// TestGb18030 go run cmd/testDecodeSpider/*.go -c dev.yml -f TestGb18030 -m dev
func (s *Spider) TestGb18030(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewGb18030Handler(s.Logger))

	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlGb18030)
	request.CallBack = s.ParseDecode
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.Logger.Error(err)
	}
	s.Logger.Info("url", request.Url)
	return
}

// TestBig5 go run cmd/testDecodeSpider/*.go -c dev.yml -f TestBig5 -m dev
func (s *Spider) TestBig5(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewBig5Handler(s.Logger))

	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlBig5)
	request.CallBack = s.ParseDecode
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.Logger.Error(err)
	}
	s.Logger.Info("url", request.Url)
	return
}

func NewSpider(baseSpider *spider.BaseSpider, logger *logger.Logger) (spider pkg.Spider, err error) {
	if baseSpider == nil {
		err = errors.New("nil baseSpider")
		logger.Error(err)
		return
	}

	baseSpider.Name = "test-decode"

	spider = &Spider{
		BaseSpider: baseSpider,
	}

	return
}

func main() {
	app.NewApp(NewSpider).Run()
}
