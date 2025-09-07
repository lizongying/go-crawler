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

func (s *Spider) ParseImage(ctx pkg.Context, _ pkg.Response) (err error) {
	s.UnsafeYieldItem(ctx, items.NewItemJsonl("image"). // build a jsonl item
								SetData(&DataImage{}).
								SetImagesRequest([]pkg.Request{ // with request list
			request.NewRequest().SetUrl(fmt.Sprintf("%simages/th.jpeg", mock_servers.UrlFile)),
		}))
	return
}

func (s *Spider) ParseFile(ctx pkg.Context, _ pkg.Response) (err error) {
	s.UnsafeYieldItem(ctx, items.NewItemJsonl("file"). // build a jsonl item
								SetData(&DataFile{}).
								SetFilesRequest([]pkg.Request{ // with request list
			request.NewRequest().SetUrl(fmt.Sprintf("%simages/th.jpeg", mock_servers.UrlFile)),
		}))
	return
}

// TestImage go run cmd/test_file_spider/*.go -c example.yml -n test-file -f TestImage -m once
func (s *Spider) TestImage(ctx pkg.Context, _ string) (err error) {

	// request the page
	s.MustYieldRequest(ctx, request.NewRequest().
		SetUrl(mock_servers.UrlOk).
		SetCallBack(s.ParseImage))
	return
}

// TestFile go run cmd/test_file_spider/*.go -c example.yml -n test-file -f TestFile -m once
func (s *Spider) TestFile(ctx pkg.Context, _ string) (err error) {

	// request the page
	s.MustYieldRequest(ctx, request.NewRequest().
		SetUrl(mock_servers.UrlOk).
		SetCallBack(s.ParseFile))
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
