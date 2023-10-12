package test_from_request_spider

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

func (s *Spider) ParseOk(_ pkg.Context, response pkg.Response) (err error) {
	s.logger.Info(response.BodyStr())
	return
}

// TestMustOk
// go run cmd/testFromRequestSpider/*.go -c example.yml -n test-from-request -f TestMustOk -m once
// go run cmd/testFromRequestSpider/*.go -c example.yml -n test-from-request -f TestMustOk -m manual
// curl -H "Content-Type: application/json" -X POST -d ' {"timeout": 1, "name": "test-from-request", "func":"TestMustOk", "args":"" }' "http://127.0.0.1:8080/spider/run"
func (s *Spider) TestMustOk(ctx pkg.Context, _ string) (err error) {
	for _, r := range []pkg.Request{
		request.NewRequest().SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mock_servers.UrlOk)).SetCallBack(s.ParseOk),
		request.NewRequest().SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mock_servers.UrlOk)).SetCallBack(s.ParseOk),
		request.NewRequest().SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mock_servers.UrlOk)).SetCallBack(s.ParseOk),
	} {
		s.MustYieldRequest(ctx, r)
	}
	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	spider = &Spider{
		Spider: baseSpider,
		logger: baseSpider.GetLogger(),
	}
	spider.WithOptions(
		pkg.WithName("test-from-request"),
		pkg.WithHost("https://localhost:8081"),
	)
	return
}
