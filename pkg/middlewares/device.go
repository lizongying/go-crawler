package middlewares

import (
	"context"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/device"
	"github.com/lizongying/go-crawler/static"
	"math/rand"
	"time"
)

type DeviceMiddleware struct {
	pkg.UnimplementedMiddleware
	logger    pkg.Logger
	uaAll     map[string][]device.Device
	platforms []pkg.Platform
	browsers  []pkg.Browser
}

func (m *DeviceMiddleware) Start(ctx context.Context, spider pkg.Spider) (err error) {
	err = m.UnimplementedMiddleware.Start(ctx, spider)
	if len(spider.GetPlatforms()) > 0 {
		m.platforms = spider.GetPlatforms()
	}
	if len(spider.GetBrowsers()) > 0 {
		m.browsers = spider.GetBrowsers()
	}
	return
}

func (m *DeviceMiddleware) ProcessRequest(_ pkg.Context, request pkg.Request) (err error) {
	platforms := m.platforms
	if len(request.GetPlatform()) > 0 {
		platforms = request.GetPlatform()
	}
	browsers := m.browsers
	if len(request.GetBrowser()) > 0 {
		browsers = request.GetBrowser()
	}

	var ua []device.Device
	var uaLen int
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
	uaLen = len(ua)

	if request.GetHeader("User-Agent") != "" && uaLen > 0 {
		u := ua[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(uaLen)]
		if u.UserAgent != "" {
			request.SetHeader("User-Agent", u.UserAgent)
		}
		if u.Fingerprint != "" {
			request.SetFingerprint(u.Fingerprint)
		} else {
			request.SetFingerprint(string(u.Browser))
		}
	}

	return
}

func (m *DeviceMiddleware) FromSpider(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(DeviceMiddleware).FromSpider(spider)
	}

	m.UnimplementedMiddleware.FromSpider(spider)
	m.logger = spider.GetLogger()
	devices, _ := device.NewDevicesFromBytes(static.Devices)
	m.uaAll = devices.Devices
	m.platforms = []pkg.Platform{pkg.Windows, pkg.Mac, pkg.Android, pkg.Iphone, pkg.Ipad, pkg.Linux}
	m.browsers = []pkg.Browser{pkg.Chrome, pkg.Edge, pkg.Safari, pkg.FireFox}
	return m
}
