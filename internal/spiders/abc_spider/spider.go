package abc_spider

import (
	"encoding/json"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"strconv"
)

const (
	name   = "abc"
	host   = "https://localhost:8081"
	okUrl  = "/get"
	output = "abc"
)

type Spider struct {
	pkg.SimpleSpider
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

// RequestOk go run cmd/abc_spider/*.go -c example.yml -n abc -f RequestOk -m once
func (s *Spider) RequestOk(ctx pkg.Context, _ string) (err error) {
	s.NewRequest(ctx).
		SetUrl(okUrl).
		SetExtra(&ExtraOk{}).
		SetCallBack(s.ParseOk).
		UnsafeYield()

	return
}

// FromExtra go run cmd/abc_spider/*.go -c example.yml -n abc -f FromExtra -m once
func (s *Spider) FromExtra(ctx pkg.Context, extraStr string) (err error) {
	var extra ExtraOk
	_ = json.Unmarshal([]byte(extraStr), &extra)

	s.NewRequest(ctx).
		SetUrl(okUrl).
		SetExtra(&extra).
		SetCallBack(s.ParseOk).
		UnsafeYield()

	return
}

func NewSpider(simpleSpider pkg.SimpleSpider) (spider pkg.SimpleSpider, err error) {
	spider = &Spider{
		simpleSpider,
	}

	spider.SetName(name).SetHost(host).SetRatePerHour("*", 1800, 1).WithJsonLinesPipeline()

	return
}
