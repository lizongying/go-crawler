package test_scheduler_spider

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/mock_servers"
	"github.com/lizongying/go-crawler/pkg/request"
	"github.com/lizongying/go-crawler/pkg/utils"
)

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) ParseOk(ctx pkg.Context, response pkg.Response) (err error) {
	var extra ExtraOk
	if err = response.UnmarshalExtra(&extra); err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.Info("extra", utils.JsonStr(extra))
	//s.logger.Info("response", response.BodyStr())

	if extra.Count > 10 {
		return
	}

	//err = s.YieldItem(ctx, items.NewItemJsonl("image").
	//	SetUniqueKey(response.UniqueKey()).
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
	//	SetUrl(response.Url()).
	//	SetExtra(&ExtraOk{
	//		Count: count,
	//	}).
	//	SetCallBack(s.ParseOk).
	//	SetUniqueKey(strconv.Itoa(count)))
	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(response.Url()).
		SetExtra(&ExtraOk{
			Count: count,
		}).
		SetPriority(uint8(count)).
		SetCallBack(s.ParseOk)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestOk go run cmd/testSchedulerSpider/*.go -c dev.yml -n test-scheduler -f TestOk -m once
func (s *Spider) TestOk(ctx pkg.Context, _ string) (err error) {
	s.AddMockServerRoutes(mock_servers.NewRouteOk(s.logger))

	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mock_servers.UrlOk)).
		SetExtra(&ExtraOk{}).
		//SetUniqueKey("0").
		SetPriority(0).
		SetCallBack(s.ParseOk)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

func (s *Spider) Stop(_ pkg.Context) (err error) {
	//err = pkg.DontStopErr
	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
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
