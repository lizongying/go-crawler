package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/devServer"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/request"
	"github.com/lizongying/go-crawler/pkg/utils"
	"strconv"
)

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) ParseOk(ctx pkg.Context, response pkg.Response) (err error) {
	var extra ExtraOk
	err = response.UnmarshalExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("ExtraOk", utils.JsonStr(extra))
	s.logger.Info("BodyBytes", string(response.GetBodyBytes()))

	if extra.Count > 0 {
		return
	}

	err = s.YieldItem(ctx, items.NewItemJsonl("image").
		SetUniqueKey(response.GetUniqueKey()).
		SetData(&DataImage{
			DataOk: DataOk{
				Count: extra.Count,
			},
		}).
		SetImagesRequest([]pkg.Request{
			request.NewRequest().SetUrl(fmt.Sprintf("%s%simages/th.jpeg", s.GetHost(), devServer.UrlFile)),
		}))
	if err != nil {
		s.logger.Error(err)
		return
	}

	//if extra.Count%1000 == 0 {
	//	s.logger.Info("extra", utils.JsonStr(extra))
	//}
	count := extra.Count + 1
	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(response.GetUrl()).
		SetExtra(&ExtraOk{
			Count: count,
		}).
		SetCallBack(s.ParseOk).
		SetUniqueKey(strconv.Itoa(count)))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestOk go run cmd/testFileSpider/*.go -c dev.yml -n test-file -f TestOk -m dev
func (s *Spider) TestOk(ctx pkg.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewRouteOk(s.logger))
	s.AddDevServerRoutes(devServer.NewRouteFile(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlOk)).
		SetExtra(&ExtraOk{}).
		SetUniqueKey("0").
		SetCallBack(s.ParseOk))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

func (s *Spider) Stop(ctx context.Context) (err error) {
	err = s.Spider.Stop(ctx)
	if err != nil {
		s.logger.Error(err)
		return
	}

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
