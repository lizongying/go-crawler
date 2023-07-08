package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/devServer"
)

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) ParseOk(ctx context.Context, response *pkg.Response) (err error) {
	var extra ExtraOk
	_ = response.Request.GetExtra(&extra)

	item := pkg.ItemNone{
		ItemUnimplemented: pkg.ItemUnimplemented{
			Data: &DataOk{
				Count: extra.Count,
			},
		},
	}
	err = s.YieldItem(ctx, &item)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	if extra.Count > 0 {
		return
	}

	requestNext := new(pkg.Request)
	//requestNext.Url = response.Request.Url
	requestNext.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlGb2312)
	requestNext.Extra = &ExtraOk{
		Count: extra.Count + 1,
	}
	requestNext.CallBack = s.ParseOk
	err = s.YieldRequest(ctx, requestNext)
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestOk go run cmd/testOkSpider/*.go -c example.yml -f TestOk -m dev
func (s *Spider) TestOk(ctx context.Context, _ string) (err error) {
	// mock server
	s.AddDevServerRoutes(devServer.NewOkHandler(s.logger))

	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlOk)
	request.Extra = &ExtraOk{}
	request.CallBack = s.ParseOk
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.logger.Error(err)
	}
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
	spider.SetName("test-ok")
	spider.SetHost(spider.GetDevServerHost())
	spider.AddDevServerRoutes(devServer.NewRobotsTxtHandler(baseSpider.GetLogger()))

	return
}

func main() {
	app.NewApp(NewSpider, pkg.WithRobotsTxtMiddleware()).Run()
}
