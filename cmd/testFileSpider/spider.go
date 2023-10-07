package main

import (
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/mockServers"
	"github.com/lizongying/go-crawler/pkg/request"
)

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) ParseOk(ctx pkg.Context, _ pkg.Response) (err error) {
	s.MustYieldItem(ctx, items.NewItemJsonl("image"). // build a jsonl item
								SetImagesRequest([]pkg.Request{ // with request list
			request.NewRequest().SetUrl(fmt.Sprintf("%s%simages/th.jpeg", s.GetHost(), mockServers.UrlFile)),
		}))

	return
}

// TestOk go run cmd/testFileSpider/*.go -c example.yml -n test-file -f TestOk -m once
func (s *Spider) TestOk(ctx pkg.Context, _ string) (err error) {

	// mock a page
	s.AddMockServerRoutes(mockServers.NewRouteOk(s.logger))

	// mock a image
	s.AddMockServerRoutes(mockServers.NewRouteFile(s.logger))

	// request the page
	s.MustYieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mockServers.UrlOk)).
		SetCallBack(s.ParseOk))
	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	if baseSpider == nil {
		err = errors.New("nil baseSpider")
		return
	}

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

func main() {
	app.NewApp(NewSpider).Run()
}
