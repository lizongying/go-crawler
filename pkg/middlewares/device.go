package middlewares

import (
	"context"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/device"
	"github.com/lizongying/go-crawler/static"
	"math/rand"
	"reflect"
	"time"
)

type DeviceMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
	uaAll  map[string][]device.Device
	ua     []device.Device
	uaLen  int
}

func (m *DeviceMiddleware) ProcessRequest(_ context.Context, request *pkg.Request) (err error) {
	platform := request.Platform
	browser := request.Browser
	var ua []device.Device
	uaLen := 0
	if len(platform) > 0 && len(browser) > 0 {
		u, ok := m.uaAll[fmt.Sprintf("%s-%s", platform, browser)]
		if ok {
			ua = append(ua, u...)
		}
		uaLen = len(ua)
	} else {
		ua = m.ua
		uaLen = m.uaLen
	}

	if len(request.Header) == 0 && uaLen > 0 {
		u := ua[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(uaLen)]
		rt := reflect.TypeOf(u)
		rv := reflect.ValueOf(u)
		for i := 0; i < rt.NumField(); i++ {
			request.SetHeader(rt.Field(i).Tag.Get("name"), rv.Field(i).String())
		}
	}

	return
}

func (m *DeviceMiddleware) FromCrawler(crawler pkg.Crawler) pkg.Middleware {
	if m == nil {
		return new(DeviceMiddleware).FromCrawler(crawler)
	}

	m.logger = crawler.GetLogger()
	platforms := crawler.GetPlatforms()
	browsers := crawler.GetBrowsers()

	devices, _ := device.NewDevicesFromBytes(static.Devices)
	m.uaAll = devices.Devices

	var ua []device.Device
	if len(platforms) > 0 && len(browsers) > 0 {
		for _, platform := range platforms {
			for _, browser := range browsers {
				u, ok := m.uaAll[fmt.Sprintf("%s-%s", platform, browser)]
				if ok {
					ua = append(ua, u...)
				}
			}
		}
	}
	m.ua = ua
	m.uaLen = len(ua)
	return m
}
