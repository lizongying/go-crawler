package {{.Name}}_spider

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"strconv"
)

const (
	name   = "{{.Name}}"
	host   = "https://httpbin.org"
	okUrl  = "/get"
	output = "{{.Name}}"
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

	if count > 0 {
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

// RequestOk go run cmd/{{.Name}}_spider/*.go -c example.yml -n {{.Name}} -f RequestOk -m once
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
		baseSpider,
	}

	spider.SetName(name).SetHost(host).WithJsonLinesPipeline()

	return
}

