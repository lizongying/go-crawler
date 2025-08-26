package test_compress_spider

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/mock_servers"
	"github.com/lizongying/go-crawler/pkg/request"
)

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) ParseCompress(_ pkg.Context, response pkg.Response) (err error) {
	s.logger.Info("header", response.Headers())
	s.logger.Info("body", response.Text())
	return
}

// TestGzip go run cmd/test_compress_spider/*.go -c dev.yml -n test-compress -f TestGzip -m once
func (s *Spider) TestGzip(ctx pkg.Context, _ string) (err error) {
	s.AddMockServerRoutes(mock_servers.NewRouteGzip(s.logger))

	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mock_servers.UrlGzip)).
		SetCallBack(s.ParseCompress)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestDeflate go run cmd/testCompressSpider/*.go -c dev.yml -n test-compress -f TestDeflate -m once
func (s *Spider) TestDeflate(ctx pkg.Context, _ string) (err error) {
	s.AddMockServerRoutes(mock_servers.NewRouteDeflate(s.logger))

	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mock_servers.UrlDeflate)).
		SetCallBack(s.ParseCompress)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestBrotli go run cmd/testCompressSpider/*.go -c dev.yml -n test-compress -f TestBrotli -m once
func (s *Spider) TestBrotli(ctx pkg.Context, _ string) (err error) {
	s.AddMockServerRoutes(mock_servers.NewRouteBrotli(s.logger))

	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mock_servers.UrlBrotli)).
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
