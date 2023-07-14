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
	_ = response.Request.GetExtra(&extra)
	s.logger.Info("extra", utils.JsonStr(extra))
	s.logger.Info("response", string(response.BodyBytes))

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
	item.SetImagesRequest([]*pkg.Request{new(pkg.Request).SetUrl("https://www.bing.com/th?id=OHR.ClamBears_ZH-CN5686721500_UHD.jpg&w=3840&h=2160&c=8&rs=1&o=3&r=0")})
	err = s.YieldItem(ctx, &item)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	//if extra.Count%1000 == 0 {
	//	s.logger.Info("extra", utils.JsonStr(extra))
	//}
	requestNext := new(pkg.Request)
	requestNext.SetUrl(response.Request.GetUrl())
	count := extra.Count + 1
	requestNext.SetExtra(&ExtraOk{
		Count: count,
	})
	requestNext.SetCallback(s.ParseOk)
	requestNext.SetUniqueKey(strconv.Itoa(count))
	//requestNext.SetUniqueKey("2")
	err = s.YieldRequest(ctx, requestNext)
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestOk go run cmd/testSchedulerSpider/*.go -c dev.yml -f TestOk -m dev
func (s *Spider) TestOk(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewOkHandler(s.logger))
	request := new(pkg.Request)
	request.SetUrl(fmt.Sprintf("%s%s", s.GetDevServerHost(), devServer.UrlOk))
	request.SetExtra(&ExtraOk{})
	request.SetUniqueKey("0")
	request.SetCallback(s.ParseOk)
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
	app.NewApp(NewSpider, pkg.WithJsonLinesPipeline()).Run()
}
