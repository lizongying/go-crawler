package middlewares

import (
	"github.com/lizongying/go-crawler/pkg"
	"math/rand"
)

type ProxyMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
}

func (m *ProxyMiddleware) ProcessRequest(_ pkg.Context, request pkg.Request) (err error) {
	if request.IsProxyEnable() == nil || !*request.IsProxyEnable() {
		return
	}

	spider := m.GetSpider()

	config := spider.GetConfig()
	proxies := config.GetProxyList()
	length := len(proxies)
	if length > 0 {
		proxy := proxies[rand.Intn(length)]
		request.SetProxy(proxy.Uri)
		m.logger.Debugf("proxy: %s", proxy)
	}

	return
}

func (m *ProxyMiddleware) FromSpider(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(ProxyMiddleware).FromSpider(spider)
	}

	m.UnimplementedMiddleware.FromSpider(spider)
	m.logger = spider.GetLogger()
	return m
}
