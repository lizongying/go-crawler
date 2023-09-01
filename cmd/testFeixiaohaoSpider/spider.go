package main

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/mockServer"
	"github.com/lizongying/go-crawler/pkg/request"
	"github.com/lizongying/go-crawler/pkg/utils"
)

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) ParseRank(_ pkg.Context, response pkg.Response) (err error) {
	var dataRanks DataRanks
	response.MustUnmarshalData(&dataRanks)
	utils.DumpJson(dataRanks.Data)
	return
}

// TestRank go run cmd/testFeixiaohaoSpider/*.go -c example.yml -n test-feixiaohao -f TestRank -m once
func (s *Spider) TestRank(ctx pkg.Context, _ string) (err error) {
	s.MustYieldRequest(ctx, request.NewRequest().
		SetUrl("https://dncapi.bostonteapartyevent.com/api/coin/web-coinrank?page=1&type=-1&pagesize=100&webp=1").
		SetCallBack(s.ParseRank))
	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	spider = &Spider{
		Spider: baseSpider,
		logger: baseSpider.GetLogger(),
	}
	spider.WithOptions(
		pkg.WithName("test-feixiaohao"),
	)
	return
}

func main() {
	app.NewApp(NewSpider).Run(
		pkg.WithMockServerRoute(mockServer.NewRouteOk),
	)
}
