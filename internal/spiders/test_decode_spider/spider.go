package test_decode_spider

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

func (s *Spider) ParseDecode(_ pkg.Context, response pkg.Response) (err error) {
	s.logger.Info("header", response.Headers())
	s.logger.Info("body", response.BodyStr())
	return
}

// TestGbk go run cmd/testDecodeSpider/*.go -c dev.yml -n test-decode -f TestGbk -m once
func (s *Spider) TestGbk(ctx pkg.Context, _ string) (err error) {
	s.AddMockServerRoutes(mock_servers.NewRouteGbk(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mock_servers.UrlGbk)).
		SetCallBack(s.ParseDecode))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestGb2312 go run cmd/testDecodeSpider/*.go -c dev.yml -n test-decode -f TestGb2312 -m once
func (s *Spider) TestGb2312(ctx pkg.Context, _ string) (err error) {
	s.AddMockServerRoutes(mock_servers.NewRouteGb2312(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mock_servers.UrlGb2312)).
		SetCallBack(s.ParseDecode))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestGb18030 go run cmd/testDecodeSpider/*.go -c dev.yml -n test-decode -f TestGb18030 -m once
func (s *Spider) TestGb18030(ctx pkg.Context, _ string) (err error) {
	s.AddMockServerRoutes(mock_servers.NewRouteGb18030(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mock_servers.UrlGb18030)).
		SetCallBack(s.ParseDecode))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestBig5 go run cmd/testDecodeSpider/*.go -c dev.yml -n test-decode -f TestBig5 -m once
func (s *Spider) TestBig5(ctx pkg.Context, _ string) (err error) {
	s.AddMockServerRoutes(mock_servers.NewRouteBig5(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mock_servers.UrlBig5)).
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
