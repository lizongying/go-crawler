package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/app"
	"github.com/lizongying/go-crawler/pkg/devServer"
	"github.com/lizongying/go-crawler/pkg/request"
	"net/http/httputil"
)

type Spider struct {
	pkg.Spider
	logger pkg.Logger
}

func (s *Spider) ParsePost(ctx context.Context, response pkg.Response) (err error) {
	dumpRequest, err := httputil.DumpRequestOut(response.GetRequest().GetRequest(), false)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.InfoF("request:\n%s", dumpRequest)
	s.logger.InfoF("body:\n%s", response.GetRequest().GetBody())

	dumpResponse, err := httputil.DumpResponse(response.GetResponse(), false)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.logger.InfoF("response:\n%s", dumpResponse)
	return
}

func (s *Spider) ParseGet(ctx context.Context, response pkg.Response) (err error) {
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

// TestPost go run cmd/testMethodSpider/*.go -c dev.yml -f TestPost -m dev
func (s *Spider) TestPost(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewPostHandler(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s", s.GetHost(), devServer.UrlPost)).
		SetMethod("POST").
		SetBody("a=0").
		SetPostForm("a", "1").
		SetPostForm("b", "2").
		SetCallBack(s.ParsePost))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

// TestGet go run cmd/testMethodSpider/*.go -c dev.yml -f TestGet -m dev
func (s *Spider) TestGet(ctx context.Context, _ string) (err error) {
	s.AddDevServerRoutes(devServer.NewGetHandler(s.logger))

	err = s.YieldRequest(ctx, request.NewRequest().
		SetUrl(fmt.Sprintf("%s%s?a=0&c=3", s.GetHost(), devServer.UrlGet)).
		SetForm("a", "1").
		SetForm("b", "2").
		SetCallBack(s.ParseGet))
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

func NewSpider(baseSpider pkg.Spider) (spider pkg.Spider, err error) {
	if baseSpider == nil {
		err = errors.New("nil baseSpider")
		return
	}

	logger := baseSpider.GetLogger()
	spider = &Spider{
		Spider: baseSpider,
		logger: logger,
	}
	spider.SetName("test-method")
	host, _ := spider.GetConfig().GetDevServer()
	spider.SetHost(host.String())

	return
}

func main() {
	app.NewApp(NewSpider).Run()
}
