package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/devServer"
	"github.com/lizongying/go-crawler/pkg/pipelines"
	"github.com/lizongying/go-crawler/pkg/utils"
	"strconv"
)

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) ParseOk(ctx context.Context, response *pkg.Response) (err error) {
	var extra ExtraOk
	_ = response.Request.GetExtra(&extra)
	s.logger.Info("extra", utils.JsonStr(extra))
	s.logger.Info("response", string(response.BodyBytes))

	if extra.Count > 0 {
		return
	}

	item := pkg.ItemJsonl{
		ItemUnimplemented: pkg.ItemUnimplemented{
			UniqueKey: response.Request.UniqueKey,
			Data: &DataImage{
				DataOk: DataOk{
					Count: extra.Count,
				},
			},
		},
		FileName: "image",
	}
	item.SetImagesRequest([]*pkg.Request{new(pkg.Request).SetUrl(fmt.Sprintf("%s%simages/th.jpeg", s.GetDevServerHost(), devServer.UrlFile))})
	err = s.YieldItem(ctx, &item)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	//if extra.Count%1000 == 0 {
	//	s.logger.Info("extra", utils.JsonStr(extra))
	//}
	requestNext := new(pkg.Request)
	requestNext.Url = response.Request.Url
	count := extra.Count + 1
	requestNext.Extra = &ExtraOk{
		Count: count,
	}
	requestNext.CallBack = s.ParseOk
	requestNext.UniqueKey = strconv.Itoa(count)
	//requestNext.UniqueKey = "2"
	err = s.YieldRequest(ctx, requestNext)
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestOk go run cmd/testFileSpider/*.go -c dev.yml -f TestOk -m dev
func (s *Spider) TestOk(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewOkHandler(s.logger))
	s.AddDevServerRoutes(devServer.NewFileHandler(s.logger))
	request := new(pkg.Request)
	request.Url = fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlOk)
	request.Extra = &ExtraOk{}
	request.UniqueKey = "0"
	request.CallBack = s.ParseOk
	err = s.YieldRequest(ctx, request)
	if err != nil {
		s.logger.Error(err)
	}
	return
}

func (s *Spider) Stop(ctx context.Context) (err error) {
	//err = pkg.DontStopErr

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
	spider.SetName("test-scheduler")

	return
}

func main() {
	app.NewApp(NewSpider, pkg.WithPipeline(new(pipelines.JsonLinesPipeline), 102)).Run()
}
