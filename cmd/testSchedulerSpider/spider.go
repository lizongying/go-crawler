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
	s.logger.Info("extra", utils.JsonStr(extra))
	//s.logger.Info("response", string(response.GetBodyBytes()))

	if extra.Count > 10 {
		return
	}

	//err = s.YieldItem(ctx, items.NewItemJsonl("image").
	//	SetUniqueKey(response.GetUniqueKey()).
	//	SetData(&DataImage{
	//		DataOk: DataOk{
	//			Count: extra.Count,
	//		},
	//
	//	}).SetImagesRequest([]pkg.Request{
	//	request.NewRequest().SetUrl("https://www.bing.com/th?id=OHR.ClamBears_ZH-CN5686721500_UHD.jpg&w=3840&h=2160&c=8&rs=1&o=3&r=0"),
	//}))

	//if extra.Count%1000 == 0 {
	//	s.logger.Info("extra", utils.JsonStr(extra))
	//}
	count := extra.Count + 1
	//err = s.YieldRequest(ctx, request.NewRequest().
	//	SetUrl(response.GetUrl()).
	//	SetExtra(&ExtraOk{
	//		Count: count,
	//	}).
	//	SetCallBack(s.ParseOk).
	//	SetUniqueKey(strconv.Itoa(count)))
	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(response.GetUrl()).
		SetExtra(&ExtraOk{
			Count: count,
		}).
		SetPriority(uint8(count)).
		SetCallBack(s.ParseOk))
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestOk go run cmd/testSchedulerSpider/*.go -c dev.yml -n test-scheduler -f TestOk -m dev
func (s *Spider) TestOk(ctx pkg.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewRouteOk(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlOk)).
		SetExtra(&ExtraOk{}).
		//SetUniqueKey("0").
		SetPriority(0).
		SetCallBack(s.ParseOk))

	if err != nil {
		s.logger.Error(err)
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
		pkg.WithName("test-scheduler"),
		pkg.WithHost("https://localhost:8081"),
	)

	return
}

func main() {
	app.NewApp(NewSpider).Run()
}
