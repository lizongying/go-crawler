package main

import (
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

func (s *Spider) ParseDecode(_ pkg.Context, response pkg.Response) (err error) {
	s.logger.Info("header", response.GetHeaders())
	s.logger.Info("body", string(response.GetBodyBytes()))
	return
}

// TestGbk go run cmd/testDecodeSpider/*.go -c dev.yml -n test-decode -f TestGbk -m dev
func (s *Spider) TestGbk(ctx pkg.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewRouteGbk(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlGbk)).
		SetCallBack(s.ParseDecode))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestGb2312 go run cmd/testDecodeSpider/*.go -c dev.yml -n test-decode -f TestGb2312 -m dev
func (s *Spider) TestGb2312(ctx pkg.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewRouteGb2312(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlGb2312)).
		SetCallBack(s.ParseDecode))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestGb18030 go run cmd/testDecodeSpider/*.go -c dev.yml -n test-decode -f TestGb18030 -m dev
func (s *Spider) TestGb18030(ctx pkg.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewRouteGb18030(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlGb18030)).
		SetCallBack(s.ParseDecode))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestBig5 go run cmd/testDecodeSpider/*.go -c dev.yml -n test-decode -f TestBig5 -m dev
func (s *Spider) TestBig5(ctx pkg.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewRouteBig5(s.logger))

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
	spider = &Spider{
		Spider: baseSpider,
		logger: baseSpider.GetLogger(),
	}
	spider.WithOptions(
		pkg.WithName("test-decode"),
		pkg.WithHost("https://localhost:8081"),
	)
	return
}

func main() {
	app.NewApp(NewSpider).Run()
}
