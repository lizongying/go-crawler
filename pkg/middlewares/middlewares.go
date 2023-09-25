package middlewares

import (
	"github.com/lizongying/go-crawler/pkg"
	"reflect"
	"sort"
	"sync"
)

type Middlewares struct {
	middlewares []pkg.Middleware
	spider      pkg.Spider
	logger      pkg.Logger
	locker      sync.Mutex
}

func (m *Middlewares) MiddlewareNames() (middlewares map[uint8]string) {
	m.locker.Lock()
	defer m.locker.Unlock()

	middlewares = make(map[uint8]string)
	for _, v := range m.middlewares {
		middlewares[v.Order()] = v.Name()
	}

	return
}
func (m *Middlewares) Middlewares() []pkg.Middleware {
	return m.middlewares
}
func (m *Middlewares) SetMiddleware(middleware pkg.Middleware, order uint8) {
	m.locker.Lock()
	defer m.locker.Unlock()

	middleware = middleware.FromSpider(m.spider)

	name := reflect.TypeOf(middleware).Elem().String()
	middleware.SetName(name)
	middleware.SetOrder(order)
	for k, v := range m.middlewares {
		if v.Name() == name && v.Order() != order {
			m.DelMiddleware(k)
			break
		}
	}

	m.middlewares = append(m.middlewares, middleware)

	sort.Slice(m.middlewares, func(i, j int) bool {
		return m.middlewares[i].Order() < m.middlewares[j].Order()
	})
}
func (m *Middlewares) DelMiddleware(index int) {
	m.locker.Lock()
	defer m.locker.Unlock()

	if index < 0 {
		return
	}
	if index >= len(m.middlewares) {
		return
	}

	m.middlewares = append(m.middlewares[:index], m.middlewares[index+1:]...)
	return
}
func (m *Middlewares) CleanMiddlewares() {
	m.locker.Lock()
	defer m.locker.Unlock()

	m.middlewares = make([]pkg.Middleware, 0)
}
func (m *Middlewares) WithCustomMiddleware(middleware pkg.Middleware) {
	m.SetMiddleware(middleware, 10)
}
func (m *Middlewares) WithRetryMiddleware() {
	m.SetMiddleware(new(RetryMiddleware), 20)
}
func (m *Middlewares) WithDumpMiddleware() {
	m.SetMiddleware(new(DumpMiddleware), 30)
}
func (m *Middlewares) WithProxyMiddleware() {
	m.SetMiddleware(new(ProxyMiddleware), 40)
}
func (m *Middlewares) WithRobotsTxtMiddleware() {
	m.SetMiddleware(new(RobotsTxtMiddleware), 50)
}
func (m *Middlewares) WithFilterMiddleware() {
	m.SetMiddleware(new(FilterMiddleware), 60)
}
func (m *Middlewares) WithFileMiddleware() {
	m.SetMiddleware(new(FileMiddleware), 70)
}
func (m *Middlewares) WithImageMiddleware() {
	m.SetMiddleware(new(ImageMiddleware), 80)
}
func (m *Middlewares) WithUrlMiddleware() {
	m.SetMiddleware(new(UrlMiddleware), 90)
}
func (m *Middlewares) WithReferrerMiddleware() {
	m.SetMiddleware(new(ReferrerMiddleware), 100)
}
func (m *Middlewares) WithCookieMiddleware() {
	m.SetMiddleware(new(CookieMiddleware), 110)
}
func (m *Middlewares) WithRedirectMiddleware() {
	m.SetMiddleware(new(RedirectMiddleware), 120)
}
func (m *Middlewares) WithChromeMiddleware() {
	m.SetMiddleware(new(ChromeMiddleware), 130)
}
func (m *Middlewares) WithHttpAuthMiddleware() {
	m.SetMiddleware(new(HttpAuthMiddleware), 140)
}
func (m *Middlewares) WithCompressMiddleware() {
	m.SetMiddleware(new(CompressMiddleware), 150)
}
func (m *Middlewares) WithDecodeMiddleware() {
	m.SetMiddleware(new(DecodeMiddleware), 160)
}
func (m *Middlewares) WithDeviceMiddleware() {
	m.SetMiddleware(new(DeviceMiddleware), 170)
}
func (m *Middlewares) WithHttpMiddleware() {
	m.SetMiddleware(new(HttpMiddleware), 200)
}
func (m *Middlewares) WithStatsMiddleware() {
	m.SetMiddleware(new(StatsMiddleware), 210)
}
func (m *Middlewares) WithRecordErrorMiddleware() {
	m.SetMiddleware(new(RecordErrorMiddleware), 220)
}
func (m *Middlewares) FromSpider(spider pkg.Spider) pkg.Middlewares {
	if m == nil {
		return new(Middlewares).FromSpider(spider)
	}

	m.spider = spider
	m.logger = spider.GetLogger()

	config := spider.GetCrawler().GetConfig()

	// set middlewares
	if config.GetEnableDumpMiddleware() {
		m.WithDumpMiddleware()
	}
	if config.GetEnableProxyMiddleware() {
		m.WithProxyMiddleware()
	}
	if config.GetEnableRobotsTxtMiddleware() {
		m.WithRobotsTxtMiddleware()
	}
	if config.GetEnableFilterMiddleware() {
		m.WithFilterMiddleware()
	}
	if config.GetEnableFileMiddleware() {
		m.WithFileMiddleware()
	}
	if config.GetEnableImageMiddleware() {
		m.WithImageMiddleware()
	}
	if config.GetEnableRetryMiddleware() {
		m.WithRetryMiddleware()
	}
	if config.GetEnableUrlMiddleware() {
		m.WithUrlMiddleware()
	}
	if config.GetEnableReferrerMiddleware() {
		m.WithReferrerMiddleware()
	}
	if config.GetEnableCookieMiddleware() {
		m.WithCookieMiddleware()
	}
	if config.GetEnableRedirectMiddleware() {
		m.WithRedirectMiddleware()
	}
	if config.GetEnableChromeMiddleware() {
		m.WithChromeMiddleware()
	}
	if config.GetEnableHttpAuthMiddleware() {
		m.WithHttpAuthMiddleware()
	}
	if config.GetEnableCompressMiddleware() {
		m.WithCompressMiddleware()
	}
	if config.GetEnableDecodeMiddleware() {
		m.WithDecodeMiddleware()
	}
	if config.GetEnableDeviceMiddleware() {
		m.WithDeviceMiddleware()
	}
	if config.GetEnableHttpMiddleware() {
		m.WithHttpMiddleware()
	}
	if config.GetEnableStatsMiddleware() {
		m.WithStatsMiddleware()
	}
	if config.GetEnableRecordErrorMiddleware() {
		m.WithRecordErrorMiddleware()
	}
	return m
}
