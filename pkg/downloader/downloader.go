package downloader

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/httpClient"
	"github.com/lizongying/go-crawler/pkg/httpClient/browser"
	"github.com/lizongying/go-crawler/pkg/middlewares"
	"reflect"
	"sort"
	"sync"
)

type Downloader struct {
	middlewares        []pkg.Middleware
	processRequestFns  []func(pkg.Context, pkg.Request) error
	processResponseFns []func(pkg.Context, pkg.Response) error
	httpClient         pkg.HttpClient
	browserManager     *browser.Manager
	spider             pkg.Spider
	logger             pkg.Logger
	locker             sync.Mutex
}

func (d *Downloader) processRequest(ctx pkg.Context, request pkg.Request) (err error) {
	if request.SkipMiddleware() {
		return
	}
	for k, v := range d.processRequestFns {
		name := d.middlewares[k].Name()
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
	if request.Client() == pkg.ClientBrowser {
		var b *browser.Browser
		b, err = d.browserManager.Pop(context.Background())
		if err != nil {
			d.logger.Error(err)
			return
		}
		client = b
		defer d.browserManager.Put(b)
	}
	response, err = client.DoRequest(request.Context(), request)
	if err != nil {
		d.logger.Error(err)
		return
	}
	if response == nil {
		err = errors.New("response nil")
		d.logger.Error(err)
		return
	}
	d.logger.Debug("StatusCode", response.StatusCode())

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

func (d *Downloader) processResponse(ctx pkg.Context, response pkg.Response) (err error) {
	if response.SkipMiddleware() {
		return
	}
	for k, v := range d.processResponseFns {
		name := d.middlewares[k].Name()
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

func (d *Downloader) MiddlewareNames() (middlewares map[uint8]string) {
	d.locker.Lock()
	defer d.locker.Unlock()

	middlewares = make(map[uint8]string)
	for _, v := range d.middlewares {
		middlewares[v.Order()] = v.Name()
	}

	return
}

func (d *Downloader) Middlewares() []pkg.Middleware {
	return d.middlewares
}

func (d *Downloader) SetMiddleware(middleware pkg.Middleware, order uint8) {
	d.locker.Lock()
	defer d.locker.Unlock()

	middleware = middleware.FromSpider(d.spider)

	name := reflect.TypeOf(middleware).Elem().String()
	middleware.SetName(name)
	middleware.SetOrder(order)
	for k, v := range d.middlewares {
		if v.Name() == name && v.Order() != order {
			d.DelMiddleware(k)
			break
		}
	}

	d.middlewares = append(d.middlewares, middleware)

	sort.Slice(d.middlewares, func(i, j int) bool {
		return d.middlewares[i].Order() < d.middlewares[j].Order()
	})

	var processRequestFns []func(pkg.Context, pkg.Request) error
	var processResponseFns []func(pkg.Context, pkg.Response) error
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
func (d *Downloader) WithCustomMiddleware(middleware pkg.Middleware) {
	d.SetMiddleware(middleware, 10)
}
func (d *Downloader) WithDumpMiddleware() {
	d.SetMiddleware(new(middlewares.DumpMiddleware), 20)
}
func (d *Downloader) WithProxyMiddleware() {
	d.SetMiddleware(new(middlewares.ProxyMiddleware), 30)
}
func (d *Downloader) WithRobotsTxtMiddleware() {
	d.SetMiddleware(new(middlewares.RobotsTxtMiddleware), 40)
}
func (d *Downloader) WithFilterMiddleware() {
	d.SetMiddleware(new(middlewares.FilterMiddleware), 50)
}
func (d *Downloader) WithFileMiddleware() {
	d.SetMiddleware(new(middlewares.FileMiddleware), 60)
}
func (d *Downloader) WithImageMiddleware() {
	d.SetMiddleware(new(middlewares.ImageMiddleware), 70)
}
func (d *Downloader) WithRetryMiddleware() {
	d.SetMiddleware(new(middlewares.RetryMiddleware), 80)
}
func (d *Downloader) WithUrlMiddleware() {
	d.SetMiddleware(new(middlewares.UrlMiddleware), 90)
}
func (d *Downloader) WithReferrerMiddleware() {
	d.SetMiddleware(new(middlewares.ReferrerMiddleware), 100)
}
func (d *Downloader) WithCookieMiddleware() {
	d.SetMiddleware(new(middlewares.CookieMiddleware), 110)
}
func (d *Downloader) WithRedirectMiddleware() {
	d.SetMiddleware(new(middlewares.RedirectMiddleware), 120)
}
func (d *Downloader) WithChromeMiddleware() {
	d.SetMiddleware(new(middlewares.ChromeMiddleware), 130)
}
func (d *Downloader) WithHttpAuthMiddleware() {
	d.SetMiddleware(new(middlewares.HttpAuthMiddleware), 140)
}
func (d *Downloader) WithCompressMiddleware() {
	d.SetMiddleware(new(middlewares.CompressMiddleware), 150)
}
func (d *Downloader) WithDecodeMiddleware() {
	d.SetMiddleware(new(middlewares.DecodeMiddleware), 160)
}
func (d *Downloader) WithDeviceMiddleware() {
	d.SetMiddleware(new(middlewares.DeviceMiddleware), 170)
}
func (d *Downloader) WithHttpMiddleware() {
	d.SetMiddleware(new(middlewares.HttpMiddleware), 200)
}
func (d *Downloader) WithStatsMiddleware() {
	d.SetMiddleware(new(middlewares.StatsMiddleware), 210)
}
func (d *Downloader) Close() {
	d.browserManager.Close()
}
func (d *Downloader) FromSpider(spider pkg.Spider) pkg.Downloader {
	if d == nil {
		return new(Downloader).FromSpider(spider)
	}

	d.spider = spider
	d.httpClient = new(httpClient.HttpClient).FromSpider(spider)
	d.browserManager = new(browser.Manager).FromSpider(spider)
	d.logger = spider.GetLogger()

	spider.GetSignal().RegisterSpiderClosed(d.Close)

	config := spider.GetCrawler().GetConfig()

	// set middlewares
	if config.GetEnableDumpMiddleware() {
		d.WithDumpMiddleware()
	}
	if config.GetEnableProxyMiddleware() {
		d.WithProxyMiddleware()
	}
	if config.GetEnableRobotsTxtMiddleware() {
		d.WithRobotsTxtMiddleware()
	}
	if config.GetEnableFilterMiddleware() {
		d.WithFilterMiddleware()
	}
	if config.GetEnableFileMiddleware() {
		d.WithFileMiddleware()
	}
	if config.GetEnableImageMiddleware() {
		d.WithImageMiddleware()
	}
	if config.GetEnableRetryMiddleware() {
		d.WithRetryMiddleware()
	}
	if config.GetEnableUrlMiddleware() {
		d.WithUrlMiddleware()
	}
	if config.GetEnableReferrerMiddleware() {
		d.WithReferrerMiddleware()
	}
	if config.GetEnableCookieMiddleware() {
		d.WithCookieMiddleware()
	}
	if config.GetEnableRedirectMiddleware() {
		d.WithRedirectMiddleware()
	}
	if config.GetEnableChromeMiddleware() {
		d.WithChromeMiddleware()
	}
	if config.GetEnableHttpAuthMiddleware() {
		d.WithHttpAuthMiddleware()
	}
	if config.GetEnableCompressMiddleware() {
		d.WithCompressMiddleware()
	}
	if config.GetEnableDecodeMiddleware() {
		d.WithDecodeMiddleware()
	}
	if config.GetEnableDeviceMiddleware() {
		d.WithDeviceMiddleware()
	}
	if config.GetEnableHttpMiddleware() {
		d.WithHttpMiddleware()
	}
	if config.GetEnableStatsMiddleware() {
		d.WithStatsMiddleware()
	}
	return d
}
