package browser

import (
	"context"
	"errors"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/launcher/flags"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/media"
	response2 "github.com/lizongying/go-crawler/pkg/response"
	"github.com/lizongying/go-crawler/pkg/utils"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"

	"github.com/go-rod/stealth"
)

var browserFlags = []string{
	"--disable-gpu",
	"--disable-demo-mode",
	"--disable-cookie-encryption",
	"--disable-setuid-sandbox",
	"--disable-dev-shm-usage",
	"--disable-background-timer-throttling",
	"--disable-backgrounding-occluded-windows",
	"--disable-breakpad",
	"--disable-component-extensions-with-background-pages",
	"--disable-extensions",
	"--disable-ipc-flooding-protection",
	"--disable-renderer-backgrounding",
	"--hide-scrollbars",
	"--metrics-recording-only",
	"--mute-audio",
	"--no-zygote",
}

type Browser struct {
	proxy        *url.URL
	timeout      time.Duration
	cancel       context.CancelFunc
	browser      *rod.Browser
	hijackRouter *rod.HijackRouter
	logger       pkg.Logger
	launcher     *Launcher
}

func (b *Browser) init() (err error) {
	l := &Launcher{
		Launcher: launcher.New().
			//StartURL("").
			NoSandbox(true).
			Env("TZ=Asia/Shanghai").
			Leakless(true).
			Headless(true),
	}
	ctx := context.Background()
	ctx, b.cancel = context.WithTimeout(ctx, b.timeout)
	l.Context(ctx)

	l.managed = true
	if !l.managed {
		err = errors.New("managed")
		return
	}

	if b.proxy != nil {
		l.Proxy(b.proxy.String())
	}

	if len(browserFlags) != 0 {
		for _, flag := range browserFlags {
			l.Set(flags.Flag(strings.TrimLeft(flag, "-")))
		}
	}

	u, err := l.Launch()
	if err != nil {
		b.logger.Error(err)
		return
	}

	b.launcher = l
	b.browser = rod.New().ControlURL(u)
	if err = b.browser.Connect(); err != nil {
		b.logger.Error(err)
		return
	}

	if err = b.browser.IgnoreCertErrors(true); err != nil {
		b.logger.Error(err)
		err = b.Close(nil)
		if err != nil {
			b.logger.Error(err)
			return
		}

		return
	}

	b.hijackRouter = b.browser.HijackRequests()
	b.hijackRouter.MustAdd("*.png|jpg|jpeg|gif|mp4|webm|avi|wav|mp3", func(ctx *rod.Hijack) {
		if utils.InSlice(ctx.Request.Type(), []proto.NetworkResourceType{
			proto.NetworkResourceTypeImage,
			proto.NetworkResourceTypeMedia,
		}) {
			ctx.Response.Fail(proto.NetworkErrorReasonBlockedByClient)
			return
		}
		ctx.ContinueRequest(&proto.FetchContinueRequest{})
	})
	go b.hijackRouter.Run()

	return
}

func (b *Browser) DoRequest(ctx context.Context, request pkg.Request) (response pkg.Response, err error) {
	if b == nil {
		err = errors.New("browser nil")
		return
	}

	page, err := stealth.Page(b.browser)
	if err != nil {
		b.logger.Error(err)
		return
	}

	defer func() {
		err = proto.NetworkClearBrowserCache{}.Call(page)
		if err != nil {
			b.logger.Error(err)
		}
		err = proto.NetworkClearBrowserCookies{}.Call(page)
		if err != nil {
			b.logger.Error(err)
		}
		if page == nil {
			return
		}
		err = page.Close()
		if err != nil {
			b.logger.Error(err)
		}
	}()

	start := time.Now()
	page = page.Context(ctx)
	Url := request.GetUrl()
	if request.IsAjax() {
		Url = request.GetReferrer()
	} else {
		for k := range request.Headers() {
			page.MustSetExtraHeaders(k, request.GetHeader(k))
		}

		for _, v := range request.Cookies() {
			page.MustSetCookies(&proto.NetworkCookieParam{
				Name:   v.Name,
				Value:  v.Value,
				Domain: request.GetURL().Host,
			})
		}
	}
	//wait := page.WaitNavigation(proto.PageLifecycleEventNameNetworkIdle)
	if err = page.Navigate(Url); err != nil {
		b.logger.Error(err)
		return
	}

	//wait()

	response = new(response2.Response)
	response.SetRequest(request)
	response.SetResponse(new(http.Response))

	storePath := request.GetScreenshot()
	if request.GetScreenshot() != "" {
		page.MustWaitLoad()
		_ = page.MustScreenshot(storePath)
		response.SetImages([]pkg.Image{&media.Image{
			File: media.File{
				StorePath: storePath,
			},
		}})
	}

	if request.IsAjax() {
		headers := make(map[string]string)
		for k := range request.Headers() {
			headers[k] = request.GetHeader(k)
		}
		timeout := b.timeout
		if request.GetTimeout() > 0 {
			timeout = request.GetTimeout()
		}
		res, e := page.Eval(`
(url, method, headers, body, timeout) => {
	return new Promise((resolve, reject) => {
		const xhr = new XMLHttpRequest();
		xhr.timeout = timeout;
		xhr.open(method, url, true);
		for (k in headers) {
			xhr.setRequestHeader(k, headers[k]);
		};
        xhr.onload = function() {
           if (xhr.status >= 200 && xhr.status < 300) {
               resolve({
					status: xhr.status,
					body: xhr.responseText,
				});
           } else {
               reject(xhr.statusText);
           }
        };
		xhr.ontimeout = function () {
		   reject('timeout');
		};
        xhr.onerror = function() {
           reject(xhr.statusText);
        };
		xhr.send(body);
	})
}`, request.GetUrl(), request.GetMethod(), headers, request.GetBodyStr(), int(timeout/time.Millisecond))
		if e != nil {
			err = e
			return
		}

		response.SetStatusCode(res.Value.Get("status").Int())
		response.SetBodyBytes([]byte(res.Value.Get("body").Str()))
		return
	}

	response.SetStatusCode(200)

	if source, _ := page.HTML(); source != "" {
		response.SetBodyBytes([]byte(source))
	}

	// cookie
	cookies, err := page.Cookies([]string{})
	if err != nil {
		b.logger.Error(err)
		return
	}

	for _, c := range cookies {
		response.SetCookies(&http.Cookie{
			Name:       c.Name,
			Value:      c.Value,
			Path:       c.Path,
			Domain:     c.Domain,
			Expires:    c.Expires.Time(),
			RawExpires: c.Expires.String(),
			Secure:     c.Secure,
			HttpOnly:   c.HTTPOnly,
		})
	}
	response.SetSpendTime(time.Now().Sub(start))

	return
}

func (b *Browser) Close(_ context.Context) (err error) {
	if b == nil {
		err = errors.New("browser nil")
		return
	}

	if b.hijackRouter != nil {
		_ = b.hijackRouter.Stop()
		b.hijackRouter = nil
	}

	if err = b.browser.Close(); err != nil {
		b.logger.Error(err)
		return
	}

	if b.launcher.Has(flags.Leakless) {
		b.launcher.Kill()
	}

	if !b.launcher.Has(flags.KeepUserDataDir) {
		b.launcher.Cleanup()
	}

	return
}

func NewBrowser(logger pkg.Logger, proxy *url.URL, timeout time.Duration) (b *Browser, err error) {
	b = &Browser{
		logger:  logger,
		proxy:   proxy,
		timeout: timeout,
	}

	if err = b.init(); err != nil {
		b.logger.Error(err)
		return
	}

	return
}

func (b *Browser) FromSpider(spider pkg.Spider) *Browser {
	if b == nil {
		return new(Browser).FromSpider(spider)
	}

	b.logger = spider.GetLogger()
	config := spider.GetCrawler().GetConfig()
	b.proxy = config.GetProxy()
	b.timeout = config.GetRequestTimeout()

	err := b.init()
	if err != nil {
		b.logger.Error(err)
	}

	return b
}
