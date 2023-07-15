package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/devServer"
	"github.com/lizongying/go-crawler/pkg/utils"
	"strconv"
)

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) ParseOk(ctx context.Context, response *pkg.Response) (err error) {
	var extra ExtraOk
	err = response.Request.UnmarshalExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("ExtraOk", utils.JsonStr(extra))
	s.logger.Info("BodyBytes", string(response.BodyBytes))

	if extra.Count > 0 {
		return
	}

	item := pkg.ItemJsonl{
		ItemUnimplemented: pkg.ItemUnimplemented{
			UniqueKey: response.Request.GetUniqueKey(),
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
		return
	}

	//if extra.Count%1000 == 0 {
	//	s.logger.Info("extra", utils.JsonStr(extra))
	//}
	count := extra.Count + 1
	err = s.YieldRequest(ctx, new(pkg.Request).
		SetUrl(response.Request.GetUrl()).
		SetExtra(&ExtraOk{
			Count: count,
		}).
		SetCallback(s.ParseOk).
		SetUniqueKey(strconv.Itoa(count)))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestOk go run cmd/testFileSpider/*.go -c dev.yml -f TestOk -m dev
func (s *Spider) TestOk(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewOkHandler(s.logger))
	s.AddDevServerRoutes(devServer.NewFileHandler(s.logger))

	err = s.YieldRequest(ctx, new(pkg.Request).
		SetUrl(fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlOk)).
		SetExtra(&ExtraOk{}).
		SetUniqueKey("0").
		SetCallback(s.ParseOk))
	if err != nil {
		s.logger.Error(err)
		return
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
	app.NewApp(NewSpider, pkg.WithJsonLinesPipeline()).Run()
}
