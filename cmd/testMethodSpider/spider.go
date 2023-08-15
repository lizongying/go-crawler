package main

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/mockServer"
	"github.com/lizongying/go-crawler/pkg/request"
	"net/http/httputil"
)

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) ParsePost(_ pkg.Context, response pkg.Response) (err error) {
	dumpRequest, err := httputil.DumpRequestOut(response.GetRequest().GetRequest(), false)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.InfoF("request:\n%s", dumpRequest)
	s.logger.InfoF("body:\n%s", response.GetRequest().BodyStr())

	dumpResponse, err := httputil.DumpResponse(response.GetResponse(), false)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.InfoF("response:\n%s", dumpResponse)
	return
}

func (s *Spider) ParseGet(_ pkg.Context, response pkg.Response) (err error) {
	dumpRequest, err := httputil.DumpRequestOut(response.GetRequest().GetRequest(), false)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.InfoF("request:\n%s", dumpRequest)

	dumpResponse, err := httputil.DumpResponse(response.GetResponse(), false)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.InfoF("response:\n%s", dumpResponse)
	return
}

// TestPost go run cmd/testMethodSpider/*.go -c dev.yml -n test-method -f TestPost -m dev
func (s *Spider) TestPost(ctx pkg.Context, _ string) (err error) {
	s.AddMockServerRoutes(mockServer.NewRoutePost(s.logger))

	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), mockServer.UrlPost)).
		SetMethod(pkg.POST).
		SetBodyStr("a=0").
		SetPostForm("a", "1").
		SetPostForm("b", "2").
		SetCallBack(s.ParsePost)); err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestGet go run cmd/testMethodSpider/*.go -c dev.yml -n test-method -f TestGet -m dev
func (s *Spider) TestGet(ctx pkg.Context, _ string) (err error) {
	s.AddMockServerRoutes(mockServer.NewRouteGet(s.logger))

	if err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s?a=0&c=3", s.GetHost(), mockServer.UrlGet)).
		SetForm("a", "1").
		SetForm("b", "2").
		SetCallBack(s.ParseGet)); err != nil {
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
		pkg.WithName("test-method"),
		pkg.WithHost("https://localhost:8081"),
	)

	return
}

func main() {
	app.NewApp(NewSpider).Run()
}
