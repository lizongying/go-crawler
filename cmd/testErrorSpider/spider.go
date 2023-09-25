package main

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/mockServers"
	"github.com/lizongying/go-crawler/pkg/request"
	"strconv"
)

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) ParseOk(ctx pkg.Context, response pkg.Response) (err error) {
	var extra ExtraOk
	response.MustUnmarshalExtra(&extra)

	s.MustYieldItem(ctx, items.NewItemNone().
		SetData(&DataOk{
			Count: extra.Count,
		}))

	if extra.Count > 0 {
		s.logger.Info("manual stop")
		return
	}

	s.MustYieldRequest(ctx, request.NewRequest().
		SetUrl("https://localhost:8081"+mockServers.UrlBadGateway).
		SetExtra(&ExtraOk{
			Count: extra.Count + 1,
		}).
		SetUniqueKey(strconv.Itoa(extra.Count+1)).
		SetCallBack(s.ParseOk))
	return
}

// TestOk go run cmd/testErrorSpider/*.go -c example.yml -n test-error -f TestOk -m once
func (s *Spider) TestOk(ctx pkg.Context, _ string) (err error) {
	s.MustYieldRequest(ctx, request.NewRequest().
		SetUrl("https://localhost:8081"+mockServers.UrlOk).
		SetExtra(&ExtraOk{}).
		SetCallBack(s.ParseOk))
	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	spider = &Spider{
		Spider: baseSpider,
		logger: baseSpider.GetLogger(),
	}
	spider.WithOptions(
		pkg.WithName("test-error"),
		pkg.WithRecordErrorMiddleware(),
	)
	return
}

func main() {
	app.NewApp(NewSpider).Run(
		pkg.WithMockServerRoutes(mockServers.NewRouteOk, mockServers.NewRouteBadGateway),
	)
}
