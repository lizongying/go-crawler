package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/devServer"
)

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) ParseDecode(_ context.Context, response *pkg.Response) (err error) {
	s.logger.Info("header", response.Header)
	s.logger.Info("body", string(response.BodyBytes))
	return
}

// TestGbk go run cmd/testDecodeSpider/*.go -c dev.yml -f TestGbk -m dev
func (s *Spider) TestGbk(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewGbkHandler(s.logger))

	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlGbk)
	request.CallBack = s.ParseDecode
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.logger.Error(err)
	}
	s.logger.Info("url", request.Url)
	return
}

// TestGb2312 go run cmd/testDecodeSpider/*.go -c dev.yml -f TestGb2312 -m dev
func (s *Spider) TestGb2312(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewGb2312Handler(s.logger))

	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlGb2312)
	request.CallBack = s.ParseDecode
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.logger.Error(err)
	}
	s.logger.Info("url", request.Url)
	return
}

// TestGb18030 go run cmd/testDecodeSpider/*.go -c dev.yml -f TestGb18030 -m dev
func (s *Spider) TestGb18030(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewGb18030Handler(s.logger))

	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlGb18030)
	request.CallBack = s.ParseDecode
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.logger.Error(err)
	}
	s.logger.Info("url", request.Url)
	return
}

// TestBig5 go run cmd/testDecodeSpider/*.go -c dev.yml -f TestBig5 -m dev
func (s *Spider) TestBig5(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewBig5Handler(s.logger))

	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlBig5)
	request.CallBack = s.ParseDecode
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.logger.Error(err)
	}
	s.logger.Info("url", request.Url)
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
	spider.SetName("test-decode")

	return
}

func main() {
	app.NewApp(NewSpider).Run()
}
