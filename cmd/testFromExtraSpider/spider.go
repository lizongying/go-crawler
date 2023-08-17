package main

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/mockServer"
	"github.com/lizongying/go-crawler/pkg/request"
)

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) ParseOk(_ pkg.Context, response pkg.Response) (err error) {
	s.logger.Info(response.BodyStr())
	return
}

// TestMustOk go run cmd/testFromExtraSpider/*.go -c example.yml -n test-from-extra -f TestMustOk -m dev
func (s *Spider) TestMustOk(ctx pkg.Context, _ string) (err error) {
	for _, extra := range []*ExtraOk{{
		Count: 1,
	}, {
		Count: 1,
	}, {
		Count: 1,
	}} {
		s.MustYieldExtra(extra)
	}

	for {
		var extra ExtraOk
		s.MustGetExtra(&extra)

		s.MustYieldRequest(ctx, request.NewRequest().
			SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mockServer.UrlOk)).
			SetExtra(&extra).
			SetCallBack(s.ParseOk))
	}
}

// TestOk go run cmd/testFromExtraSpider/*.go -c example.yml -n test-from-extra -f TestOk -m dev
func (s *Spider) TestOk(ctx pkg.Context, _ string) (err error) {
	for _, extra := range []*ExtraOk{{
		Count: 1,
	}, {
		Count: 1,
	}, {
		Count: 1,
	}, {
		Count: 1,
	}, {
		Count: 1,
	}} {
		if err = s.YieldExtra(extra); err != nil {
			s.logger.Error(err)
			return
		}
	}

	for {
		var extra ExtraOk
		if err = s.GetExtra(&extra); err != nil {
			s.logger.Error(err)
			return
		}

		if err = s.YieldRequest(ctx, request.NewRequest().
			SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mockServer.UrlOk)).
			SetExtra(&extra).
			SetCallBack(s.ParseOk)); err != nil {
			s.logger.Error(err)
			return
		}
	}
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	spider = &Spider{
		Spider: baseSpider,
		logger: baseSpider.GetLogger(),
	}
	spider.WithOptions(
		pkg.WithName("test-from-extra"),
		pkg.WithHost("https://localhost:8081"),
	)

	return
}

func main() {
	app.NewApp(NewSpider).Run(
		pkg.WithMockServerRoute(mockServer.NewRouteOk),
	)
}
