package main

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/mockServers"
	"github.com/lizongying/go-crawler/pkg/request"
)

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) ParseCompress(_ pkg.Context, response pkg.Response) (err error) {
	s.logger.Info("header", response.Headers())
	s.logger.Info("body", response.BodyStr())
	return
}

// TestGzip go run cmd/testCompressSpider/*.go -c dev.yml -n test-compress -f TestGzip -m once
func (s *Spider) TestGzip(ctx pkg.Context, _ string) (err error) {
	s.AddMockServerRoutes(mockServers.NewRouteGzip(s.logger))

	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mockServers.UrlGzip)).
		SetCallBack(s.ParseCompress)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestDeflate go run cmd/testCompressSpider/*.go -c dev.yml -n test-compress -f TestDeflate -m once
func (s *Spider) TestDeflate(ctx pkg.Context, _ string) (err error) {
	s.AddMockServerRoutes(mockServers.NewRouteDeflate(s.logger))

	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mockServers.UrlDeflate)).
		SetCallBack(s.ParseCompress)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestBrotli go run cmd/testCompressSpider/*.go -c dev.yml -n test-compress -f TestBrotli -m once
func (s *Spider) TestBrotli(ctx pkg.Context, _ string) (err error) {
	s.AddMockServerRoutes(mockServers.NewRouteBrotli(s.logger))

	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mockServers.UrlBrotli)).
		SetCallBack(s.ParseCompress)); err != nil {
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
		pkg.WithName("test-compress"),
		pkg.WithHost("https://localhost:8081"),
	)
	return
}

func main() {
	app.NewApp(NewSpider).Run()
}
