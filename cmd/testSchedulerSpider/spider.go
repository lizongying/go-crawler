package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/devServer"
	"github.com/lizongying/go-crawler/pkg/request"
	"github.com/lizongying/go-crawler/pkg/utils"
	"strconv"
)

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) ParseOk(ctx context.Context, response pkg.Response) (err error) {
	var extra ExtraOk
	err = response.UnmarshalExtra(&extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.JsonStr(extra))
	s.logger.Info("response", string(response.GetBodyBytes()))

	if extra.Count > 0 {
		return
	}

	item := pkg.ItemJsonl{
		ItemUnimplemented: pkg.ItemUnimplemented{
			UniqueKey: response.GetUniqueKey(),
			Data: &DataImage{
				DataOk: DataOk{
					Count: extra.Count,
				},
			},
		},
		FileName: "image",
	}
	item.SetImagesRequest([]pkg.Request{request.NewRequest().SetUrl("https://www.bing.com/th?id=OHR.ClamBears_ZH-CN5686721500_UHD.jpg&w=3840&h=2160&c=8&rs=1&o=3&r=0")})
	err = s.YieldItem(ctx, &item)
	if err != nil {
		s.logger.Error(err)
		return err
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
	}
	return
}

// TestOk go run cmd/testSchedulerSpider/*.go -c dev.yml -f TestOk -m dev
func (s *Spider) TestOk(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewOkHandler(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlOk)).
		SetExtra(&ExtraOk{}).
		SetUniqueKey("0").
		SetCallBack(s.ParseOk))
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
	host, _ := spider.GetConfig().GetDevServer()
	spider.SetHost(host.String())

	return
}

func main() {
	app.NewApp(NewSpider, pkg.WithJsonLinesPipeline()).Run()
}
