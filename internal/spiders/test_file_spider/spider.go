package test_file_spider

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/mock_servers"
	"github.com/lizongying/go-crawler/pkg/request"
)

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) ParseOk(ctx pkg.Context, _ pkg.Response) (err error) {
	s.UnsafeYieldItem(ctx, items.NewItemJsonl("image"). // build a jsonl item
								SetData(&DataImage{}).
								SetImagesRequest([]pkg.Request{ // with request list
			request.NewRequest().SetUrl(fmt.Sprintf("%s%simages/th.jpeg", s.GetHost(), mock_servers.UrlFile)),
		}))
	return
}

// TestOk go run cmd/test_file_spider/*.go -c example.yml -n test-file -f TestOk -m once
func (s *Spider) TestOk(ctx pkg.Context, _ string) (err error) {

	// mock a page
	s.AddMockServerRoutes(mock_servers.NewRouteOk(s.logger))

	// mock a image
	s.AddMockServerRoutes(mock_servers.NewRouteFile(s.logger))

	// request the page
	s.MustYieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mock_servers.UrlOk)).
		SetCallBack(s.ParseOk))
	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	spider = &Spider{
		Spider: baseSpider,
		logger: baseSpider.GetLogger(),
	}
	spider.WithOptions(
		pkg.WithName("test-file"),
		pkg.WithHost("https://localhost:8081"),
		pkg.WithJsonLinesPipeline(),
	)
	return
}
