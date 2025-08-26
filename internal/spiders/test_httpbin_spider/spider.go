package test_httpbin_spider

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"strconv"
	"time"
)

const (
	name   = "test_httpbin"
	host   = "https://httpbin.org"
	okUrl  = "/get"
	output = "httpbin"
)

type Spider struct {
	pkg.Spider
}

func (s *Spider) ParseOk(ctx pkg.Context, response pkg.Response) (err error) {
	var extra ExtraOk
	response.UnsafeExtra(&extra)
	s.Logger().Info("extra", utils.UnsafeJSON(extra))

	count := extra.Count

	s.NewItemJsonl(ctx, output).
		SetData(&DataOk{
			Id:    strconv.Itoa(count),
			Count: count,
		}).
		UnsafeYield()

	if count > 5 {
		s.Logger().Info("response", response.Text())
		return
	}

	s.NewRequest(ctx).
		SetUrl(response.Url()).
		SetExtra(&ExtraOk{
			Count: count + 1,
		}).
		SetCallBack(s.ParseOk).
		UnsafeYield()

	return
}

// RequestOk go run cmd/test_httpbin_spider/*.go -c dev.yml -n test_httpbin -f RequestOk -m once
func (s *Spider) RequestOk(ctx pkg.Context, _ string) (err error) {
	s.NewRequest(ctx).
		SetUrl(okUrl).
		SetExtra(&ExtraOk{}).
		SetCallBack(s.ParseOk).
		UnsafeYield()

	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	spider = &Spider{
		Spider: baseSpider,
	}

	spider.SetName(name).SetHost(host).WithInterval(time.Second * 5).WithJsonLinesPipeline()

	return
}
