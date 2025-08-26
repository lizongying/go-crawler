package {{.Name}}_spider

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/request"
	"github.com/lizongying/go-crawler/pkg/utils"
	"strconv"
)

const (
	name     = "{{.Name}}"
	host     = "https://httpbin.org"
	okUrl    = "/get"
	jsonName = "{{.Name}}"
)

type Spider struct {
	pkg.Spider
}

func (s *Spider) ParseOk(ctx pkg.Context, response pkg.Response) (err error) {
	var extra ExtraOk
	response.UnsafeExtra(&extra)
	s.Logger().Info("extra", utils.UnsafeJSON(extra))

	count := extra.Count

	s.UnsafeYieldItem(ctx, items.NewItemJsonl(jsonName).
		SetData(&DataOk{
			Id:    strconv.Itoa(count),
			Count: count,
		}))

	if count > 0 {
		s.Logger().Info("response", response.Text())
		return
	}

	s.UnsafeYieldRequest(ctx, request.NewRequest().
		SetUrl(response.Url()).
		SetExtra(&ExtraOk{
			Count: count + 1,
		}).
		SetCallBack(s.ParseOk))

	return
}

// RequestOk go run cmd/{{.Name}}_spider/*.go -c dev.yml -n {{.Name}} -f RequestOk -m once
func (s *Spider) RequestOk(ctx pkg.Context, _ string) (err error) {
	s.UnsafeYieldRequest(ctx, request.NewRequest().
		SetUrl(okUrl).
		SetExtra(&ExtraOk{}).
		SetCallBack(s.ParseOk))

	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	spider = &Spider{
		Spider: baseSpider,
	}

	spider.SetName(name).SetHost(host).WithJsonLinesPipeline()

	return
}

