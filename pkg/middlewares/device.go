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

func (m *DeviceMiddleware) ProcessRequest(_ context.Context, request pkg.Request) (err error) {
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

func (m *DeviceMiddleware) FromCrawler(crawler pkg.Crawler) pkg.Middleware {
	if m == nil {
		return new(DeviceMiddleware).FromCrawler(crawler)
	}

	m.logger = crawler.GetLogger()

	devices, _ := device.NewDevicesFromBytes(static.Devices)
	m.uaAll = devices.Devices
	m.platforms = []pkg.Platform{pkg.Windows, pkg.Mac, pkg.Android, pkg.Iphone, pkg.Ipad, pkg.Linux}
	if len(crawler.GetPlatforms()) > 0 {
		m.platforms = crawler.GetPlatforms()
	}
	m.browsers = []pkg.Browser{pkg.Chrome, pkg.Edge, pkg.Safari, pkg.FireFox}
	if len(crawler.GetBrowsers()) > 0 {
		m.browsers = crawler.GetBrowsers()
	}
	return m
}
