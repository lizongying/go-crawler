package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/devServer"
	"github.com/lizongying/go-crawler/pkg/request"
)

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) ParseDecode(_ context.Context, response pkg.Response) (err error) {
	s.logger.Info("header", response.GetHeaders())
	s.logger.Info("body", string(response.GetBodyBytes()))
	return
}

// TestGbk go run cmd/testDecodeSpider/*.go -c dev.yml -f TestGbk -m dev
func (s *Spider) TestGbk(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewHandlerGbk(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlGbk)).
		SetCallBack(s.ParseDecode))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestGb2312 go run cmd/testDecodeSpider/*.go -c dev.yml -f TestGb2312 -m dev
func (s *Spider) TestGb2312(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewHandlerGb2312(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlGb2312)).
		SetCallBack(s.ParseDecode))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestGb18030 go run cmd/testDecodeSpider/*.go -c dev.yml -f TestGb18030 -m dev
func (s *Spider) TestGb18030(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewHandlerGb18030(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlGb18030)).
		SetCallBack(s.ParseDecode))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestBig5 go run cmd/testDecodeSpider/*.go -c dev.yml -f TestBig5 -m dev
func (s *Spider) TestBig5(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewHandlerBig5(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlBig5)).
		SetCallBack(s.ParseDecode))
	if err != nil {
		s.logger.Error(err)
		return
	}

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
	host, _ := spider.GetConfig().GetDevServer()
	spider.SetHost(host.String())

	return
}

func main() {
	app.NewApp(NewSpider).Run()
}
