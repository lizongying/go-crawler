package middlewares

import (
	"context"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"math/rand"
	"time"
)

type DeviceMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
	spider pkg.Spider
	uaAll  map[string][]map[string]string
	ua     []map[string]string
	uaLen  int
}

func (m *DeviceMiddleware) SpiderStart(_ context.Context, spider pkg.Spider) (err error) {
	m.spider = spider
	platform := spider.GetPlatform()
	browser := spider.GetBrowser()
	var ua []map[string]string
	if len(platform) > 0 {
		if len(browser) > 0 {
			u, ok := m.uaAll[fmt.Sprintf("%d%d", platform, browser)]
			if ok {
				ua = append(ua, u...)
			}
		}
	}
	m.ua = ua
	m.uaLen = len(ua)
	return
}

func (m *DeviceMiddleware) ProcessRequest(c *pkg.Context) (err error) {
	request := c.Request

	platform := request.Platform
	browser := request.Browser
	var ua []map[string]string
	uaLen := 0
	if len(platform) > 0 && len(browser) > 0 {
		u, ok := m.uaAll[fmt.Sprintf("%d%d", platform, browser)]
		if ok {
			ua = append(ua, u...)
		}
		uaLen = len(ua)
	} else {
		ua = m.ua
		uaLen = m.uaLen
	}

	m.logger.Error(1111, request.UserAgent())

	if request.UserAgent() == "" && uaLen > 0 {
		u := ua[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(uaLen)]
		for k, v := range u {
			request.SetHeader(k, v)
		}
	}

	err = c.NextRequest()
	return
}

func (m *DeviceMiddleware) FromCrawler(spider pkg.Spider) pkg.Middleware {
	m.logger = spider.GetLogger()
	return m
}

func NewDeviceMiddleware() pkg.Middleware {
	return &DeviceMiddleware{}
}
