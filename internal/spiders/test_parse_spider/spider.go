package test_parse_spider

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/mock_servers"
	"github.com/lizongying/go-crawler/pkg/request"
	"github.com/lizongying/go-crawler/pkg/utils"
)

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) Parse(_ pkg.Context, response pkg.Response) (err error) {
	var dataParse DataParse
	response.MustUnmarshalData(&dataParse)
	utils.DumpJsonPretty(dataParse.Data)
	return
}

// TestOk go run cmd/test_parse_spider/*.go -c example.yml -n test-parse -f TestOk -m once
func (s *Spider) TestOk(ctx pkg.Context, _ string) (err error) {
	s.MustYieldRequest(ctx, request.NewRequest().
		SetUrl("https://localhost:8081"+mock_servers.UrlHtml+"index.html").
		SetCallBack(s.Parse))
	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	spider = &Spider{
		Spider: baseSpider,
		logger: baseSpider.GetLogger(),
	}
	spider.WithOptions(
		pkg.WithName("test-parse"),
	)
	return
}
