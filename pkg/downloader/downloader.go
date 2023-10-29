package downloader

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/http_client"
	"github.com/lizongying/go-crawler/pkg/http_client/browser"
	response2 "github.com/lizongying/go-crawler/pkg/response"
)

type Downloader struct {
	httpClient     pkg.HttpClient
	browserManager *browser.Manager
	spider         pkg.Spider
	logger         pkg.Logger
	middlewares    []pkg.Middleware
}

func (d *Downloader) Download(ctx pkg.Context, request pkg.Request) (response pkg.Response, err error) {
	err = d.processRequest(ctx, request)
	if err != nil {
		if errors.Is(err, pkg.ErrIgnoreRequest) {
			d.logger.Debug(err)
		} else {
			d.logger.Error(err)
		}
		return
	}

	if request == nil {
		err = errors.New("nil request")
		d.logger.Error(err)
		return
	}

	client := d.httpClient
	if request.GetClient() == pkg.ClientBrowser {
		var b *browser.Browser
		b, err = d.browserManager.Pop(context.Background())
		if err != nil {
			d.logger.Error(err)
			return
		}
		client = b
		defer d.browserManager.Put(b)
	}
	response, err = client.DoRequest(request.RequestContext(), request)
	if err != nil {
		d.logger.Error(err)
	}

	if response == nil {
		response = new(response2.Response)
		response.SetRequest(request)
	}

	err = d.processResponse(ctx, response)
	if err != nil {
		d.logger.Error(err)
		if errors.Is(err, pkg.ErrNeedRetry) {
			return d.Download(ctx, request)
		}
		return
	}

	if response == nil {
		err = errors.New("nil response")
		d.logger.Error(err)
		return
	}

	if response != nil && request != nil {
		response.SetRequest(request)
	}

	return
}
func (d *Downloader) processRequest(ctx pkg.Context, request pkg.Request) (err error) {
	if request.IsSkipMiddleware() {
		return
	}
	for _, v := range d.middlewares {
		name := v.Name()
		d.logger.Debug("enter", name, "processRequest")
		e := v.ProcessRequest(ctx, request)
		err = errors.Join(err, e)
		d.logger.Debug("exit", name, "processRequest")
		if errors.Is(e, pkg.ErrIgnoreRequest) {
			break
		}
	}
	return
}
func (d *Downloader) processResponse(ctx pkg.Context, response pkg.Response) (err error) {
	if response.SkipMiddleware() {
		return
	}
	for _, v := range d.middlewares {
		name := v.Name()
		d.logger.Debug("enter", name, "ProcessResponse")
		e := v.ProcessResponse(ctx, response)
		err = errors.Join(err, e)
		d.logger.Debug("exit", name, "ProcessResponse")
		if errors.Is(e, pkg.ErrIgnoreResponse) {
			break
		}
	}
	return
}
func (d *Downloader) Close(ctx pkg.Context) {
	if ctx.GetSpiderName() != d.spider.Name() {
		return
	}
	d.browserManager.Close()
}
func (d *Downloader) FromSpider(spider pkg.Spider) pkg.Downloader {
	if d == nil {
		return new(Downloader).FromSpider(spider)
	}

	d.spider = spider
	d.httpClient = new(http_client.HttpClient).FromSpider(spider)
	d.browserManager = new(browser.Manager).FromSpider(spider)
	d.logger = spider.GetLogger()
	d.middlewares = spider.GetMiddlewares().Middlewares()

	spider.GetCrawler().GetSignal().RegisterSpiderStopped(d.Close)
	return d
}
