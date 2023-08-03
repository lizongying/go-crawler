package main

import (
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/devServer"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/request"
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

	err = s.YieldItem(ctx, items.NewItemNone().
		SetData(&DataOk{
			Count: extra.Count,
		}))
	if err != nil {
		s.logger.Error(err)
		return err
	}

	if extra.Count > 0 {
		return
	}

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(response.GetUrl()).
		SetExtra(&ExtraOk{
			Count: extra.Count + 1,
		}).
		SetCallBack(s.ParseOk))
	if err != nil {
		s.logger.Error(err)
	}
	return
}

// TestOk go run cmd/testOkSpider/*.go -c example.yml -n test-ok -f TestOk -m dev
func (s *Spider) TestOk(ctx pkg.Context, _ string) (err error) {
	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlOk)).
		SetExtra(&ExtraOk{}).
		SetCallBack(s.ParseOk))
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
	spider.WithOptions(
		pkg.WithName("test-ok"),
		pkg.WithHost("https://localhost:8081"),
	)

	return
}

func main() {
	app.NewApp(NewSpider).Run(
		//pkg.WithDevServerRoute(devServer.NewRouteRobotsTxt),
		pkg.WithDevServerRoute(devServer.NewRouteOk),
	)
}
