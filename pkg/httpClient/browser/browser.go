package browser

import (
	"context"
	"errors"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/launcher/flags"
	"github.com/lizongying/go-crawler/pkg"
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

var browserOptions = []string{
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
	"--no-sandbox",
	"--no-zygote",
}

type Browser struct {
	proxy        *url.URL
	timeout      time.Duration
	browser      *rod.Browser
	hijackRouter *rod.HijackRouter
	logger       pkg.Logger
}

func (b *Browser) init() (err error) {
	l := launcher.New().Env("TZ=Asia/Shanghai").
		Leakless(true).
		Headless(true)
	for _, arg := range browserOptions {
		l.Set(flags.Flag(strings.TrimLeft(arg, "-")))
	}

	u, err := l.Launch()
	if err != nil {
		b.logger.Error(err)
		return
	}

	b.browser = rod.New().ControlURL(u)
	if err = b.browser.Connect(); err != nil {
		b.logger.Error(err)
		return
	}

	if err = b.browser.IgnoreCertErrors(true); err != nil {
		b.logger.Error(err)
		err = b.Close()
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
	if request.Ajax() {
		Url = request.GetReferrer()
	} else {
		for k := range request.GetHeaders() {
			page.MustSetExtraHeaders(k, request.GetHeader(k))
		}

		for _, v := range request.GetCookies() {
			page.MustSetCookies(&proto.NetworkCookieParam{
				Name:  v.Name,
				Value: v.Value,
			})
		}
	}
	//wait := page.WaitNavigation(proto.PageLifecycleEventNameNetworkIdle)
	if err = page.Navigate(Url); err != nil {
		b.logger.Error(err)
		return
	}

	//wait()
	time.Sleep(2 * time.Second)

	response = new(response2.Response)
	response.SetRequest(request)
	response.SetResponse(new(http.Response))

	if request.Ajax() {
		headers := make(map[string]string)
		for k := range request.GetHeaders() {
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
}`, request.GetUrl(), request.GetMethod(), headers, request.GetBody(), int(timeout/time.Millisecond))
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

func (b *Browser) Close() (err error) {
	if b == nil {
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

	return
}

func (b *Browser) FromSpider(spider pkg.Spider) pkg.HttpClient {
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
