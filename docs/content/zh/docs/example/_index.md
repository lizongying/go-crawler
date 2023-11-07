---
weight: 2
bookFlatSection: true
title: "示例代码"
---

## 示例

example_spider.go

```go
package main

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/mock_servers"
	"github.com/lizongying/go-crawler/pkg/request"
)

type ExtraOk struct {
	Count int
}

type DataOk struct {
	Count int
}

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) ParseOk(ctx pkg.Context, response pkg.Response) (err error) {
	var extra ExtraOk
	response.MustUnmarshalExtra(&extra)

	s.MustYieldItem(ctx, items.NewItemNone().
		SetData(&DataOk{
			Count: extra.Count,
		}))

	if extra.Count > 0 {
		s.logger.Info("manual stop")
		return
	}

	s.MustYieldRequest(ctx, request.NewRequest().
		SetUrl(response.Url()).
		SetExtra(&ExtraOk{
			Count: extra.Count + 1,
		}).
		SetCallBack(s.ParseOk))
	return
}

func (s *Spider) TestOk(ctx pkg.Context, _ string) (err error) {
	s.MustYieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mock_servers.UrlOk)).
		SetExtra(&ExtraOk{}).
		SetCallBack(s.ParseOk))
	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	spider = &Spider{
		Spider: baseSpider,
		logger: baseSpider.GetLogger(),
	}
	spider.WithOptions(
		pkg.WithName("example"),
		pkg.WithHost("https://localhost:8081"),
	)
	return
}

func main() {
	app.NewApp(NewSpider).Run(pkg.WithMockServerRoutes(mock_servers.NewRouteOk))
}

```

### 运行

```shell
go run example_spider.go -c example.yml -n example -f TestOk -m once

```