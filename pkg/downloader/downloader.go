package downloader

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/httpClient"
	"github.com/lizongying/go-crawler/pkg/middlewares"
	"reflect"
	"sort"
	"sync"
)

type Downloader struct {
	middlewares        []pkg.Middleware
	processRequestFns  []func(context.Context, *pkg.Request) error
	processResponseFns []func(context.Context, *pkg.Response) error
	httpClient         pkg.HttpClient
	crawler            pkg.Crawler
	logger             pkg.Logger
	locker             sync.Mutex
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

func (d *Downloader) Download(ctx context.Context, request *pkg.Request) (response *pkg.Response, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	err = d.processRequest(ctx, request)
	if err != nil {
		d.logger.Error(err)
		return
	}

	if request == nil {
		err = errors.New("nil request")
		d.logger.Error(err)
		return
	}

	if request.Request == nil {
		err = errors.New("nil request.Request")
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
			return d.Download(request.Context(), request)
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

func (d *Downloader) GetMiddlewareNames() (middlewares map[uint8]string) {
	d.locker.Lock()
	defer d.locker.Unlock()

	middlewares = make(map[uint8]string)
	for _, v := range d.middlewares {
		middlewares[v.GetOrder()] = v.GetName()
	}

	return
}

func (d *Downloader) GetMiddlewares() []pkg.Middleware {
	return d.middlewares
}

func (d *Downloader) SetMiddleware(middleware pkg.Middleware, order uint8) {
	d.locker.Lock()
	defer d.locker.Unlock()

	middleware = middleware.FromCrawler(d.crawler)

	name := reflect.TypeOf(middleware).Elem().String()
	middleware.SetName(name)
	middleware.SetOrder(order)
	for k, v := range d.middlewares {
		if v.GetName() == name && v.GetOrder() != order {
			d.DelMiddleware(k)
			break
		}
	}

	d.middlewares = append(d.middlewares, middleware)
	sort.Slice(d.middlewares, func(i, j int) bool {
		return d.middlewares[i].GetOrder() < d.middlewares[j].GetOrder()
	})

	var processRequestFns []func(context.Context, *pkg.Request) error
	var processResponseFns []func(context.Context, *pkg.Response) error
	for _, v := range d.middlewares {
		processRequestFns = append(processRequestFns, v.ProcessRequest)
		processResponseFns = append(processResponseFns, v.ProcessResponse)
	}
	d.processRequestFns = processRequestFns
	d.processResponseFns = processResponseFns
}

func (d *Downloader) DelMiddleware(index int) {
	d.locker.Lock()
	defer d.locker.Unlock()

	if index < 0 {
		return
	}
	if index >= len(d.middlewares) {
		return
	}

	d.middlewares = append(d.middlewares[:index], d.middlewares[index+1:]...)
	return
}

func (d *Downloader) CleanMiddlewares() {
	d.locker.Lock()
	defer d.locker.Unlock()

	d.middlewares = make([]pkg.Middleware, 0)
}

func (d *Downloader) FromCrawler(crawler pkg.Crawler) pkg.Downloader {
	if d == nil {
		return new(Downloader).FromCrawler(crawler)
	}

	d.httpClient = new(httpClient.HttpClient).FromCrawler(crawler)
	d.crawler = crawler
	d.logger = crawler.GetLogger()
	config := crawler.GetConfig()

	// set middlewares
	if config.GetEnableStats() {
		d.SetMiddleware(new(middlewares.StatsMiddleware), 10)
	}
	if config.GetEnableDumpMiddleware() {
		d.SetMiddleware(new(middlewares.DumpMiddleware), 20)
	}
	if config.GetEnableFilterMiddleware() {
		d.SetMiddleware(new(middlewares.FilterMiddleware), 30)
	}
	if config.GetEnableImageMiddleware() {
		d.SetMiddleware(new(middlewares.ImageMiddleware), 40)
	}
	d.SetMiddleware(new(middlewares.HttpMiddleware), 50)
	if config.GetEnableRetry() {
		d.SetMiddleware(new(middlewares.RetryMiddleware), 60)
	}
	if config.GetEnableUrl() {
		d.SetMiddleware(new(middlewares.UrlMiddleware), 70)
	}
	if config.GetEnableReferer() {
		d.SetMiddleware(new(middlewares.RefererMiddleware), 80)
	}
	if config.GetEnableCookie() {
		d.SetMiddleware(new(middlewares.CookieMiddleware), 90)
	}
	if config.GetEnableRedirect() {
		d.SetMiddleware(new(middlewares.RedirectMiddleware), 100)
	}
	if config.GetEnableChrome() {
		d.SetMiddleware(new(middlewares.ChromeMiddleware), 110)
	}
	if config.GetEnableHttpAuth() {
		d.SetMiddleware(new(middlewares.HttpAuthMiddleware), 120)
	}
	if config.GetEnableCompress() {
		d.SetMiddleware(new(middlewares.CompressMiddleware), 130)
	}
	if config.GetEnableDecode() {
		d.SetMiddleware(new(middlewares.DecodeMiddleware), 140)
	}
	if config.GetEnableDevice() {
		d.SetMiddleware(new(middlewares.DeviceMiddleware), 150)
	}

	return d
}
