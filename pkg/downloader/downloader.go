package downloader

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/httpClient"
)

type Downloader struct {
	middlewares        []pkg.Middleware
	processRequestFns  []func(context.Context, *pkg.Request) error
	processResponseFns []func(context.Context, *pkg.Response) error
	httpClient         pkg.HttpClient

	logger pkg.Logger
}

func (d *Downloader) SetMiddlewares(middlewares []pkg.Middleware) {
	d.middlewares = middlewares
	var processRequestFns []func(context.Context, *pkg.Request) error
	var processResponseFns []func(context.Context, *pkg.Response) error
	for _, v := range middlewares {
		processRequestFns = append(processRequestFns, v.ProcessRequest)
		processResponseFns = append(processResponseFns, v.ProcessResponse)
	}
	d.processRequestFns = processRequestFns
	d.processResponseFns = processResponseFns
}

func (d *Downloader) processRequest(ctx context.Context, request *pkg.Request) (err error) {
	for k, v := range d.processRequestFns {
		name := d.middlewares[k].GetName()
		d.logger.Debug("enter", name, "processRequest")
		e := v(ctx, request)
		err = errors.Join(err, e)
		d.logger.Debug("exit", name, "processRequest")
		if errors.Is(e, pkg.ErrIgnoreRequest) {
			break
		}
	}
	return
}

func (d *Downloader) DoRequest(ctx context.Context, request *pkg.Request) (response *pkg.Response, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	err = d.processRequest(ctx, request)
	if err != nil {
		d.logger.Error(err)
		return
	}

	response, err = d.httpClient.DoRequest(request.Context(), request)
	if err != nil {
		d.logger.Error(err)
		return
	}

	err = d.processResponse(ctx, response)
	if err != nil {
		d.logger.Error(err)
		if errors.Is(err, pkg.ErrNeedRetry) {
			return d.DoRequest(request.Context(), request)
		}
		return
	}

	if response == nil {
		err = errors.New("nil response")
		d.logger.Error(err)
		return
	}

	if response != nil && request != nil {
		response.Request = request
	}

	return
}

func (d *Downloader) processResponse(ctx context.Context, response *pkg.Response) (err error) {
	for k, v := range d.processResponseFns {
		name := d.middlewares[k].GetName()
		d.logger.Debug("enter", name, "ProcessResponse")
		e := v(ctx, response)
		err = errors.Join(err, e)
		d.logger.Debug("exit", name, "ProcessResponse")
		if errors.Is(e, pkg.ErrIgnoreResponse) {
			break
		}
	}
	return
}

func (d *Downloader) FromCrawler(spider pkg.Spider) *Downloader {
	if d == nil {
		return new(Downloader).FromCrawler(spider)
	}
	d.logger = spider.GetLogger()
	d.httpClient = new(httpClient.HttpClient).FromCrawler(spider)
	return d
}
