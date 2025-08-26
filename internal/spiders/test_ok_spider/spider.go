package test_ok_spider

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/mock_servers"
	"github.com/lizongying/go-crawler/pkg/request"
)

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) ParseOk(ctx pkg.Context, response pkg.Response) (err error) {
	var extra ExtraOk
	if err = response.Extra(&extra); err != nil {
		s.logger.Error(err)
		return
	}

	if err = s.YieldItem(ctx, items.NewItemNone().
		SetData(&DataOk{
			Count: extra.Count,
		})); err != nil {
		s.logger.Error(err)
		return err
	}

	if extra.Count > 0 {
		s.logger.Info("manual stop")
		return
	}

	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(response.Url()).
		SetExtra(&ExtraOk{
			Count: extra.Count + 1,
		}).
		SetCallBack(s.ParseOk)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestOk go run cmd/test_ok_spider/*.go -c example.yml -n test-ok -f TestOk -m once
func (s *Spider) TestOk(ctx pkg.Context, _ string) (err error) {
	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl("https://localhost:8081"+mock_servers.UrlOk).
		SetExtra(&ExtraOk{}).
		SetCallBack(s.ParseOk)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
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
