package middlewares

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/logger"
	"math/rand"
	"time"
)

type platform uint8

const (
	Windows platform = iota
	Mac
	Android
	Iphone
)

type browser uint8

const (
	Chrome browser = iota
	Edge
	FireFox
)

type DeviceMiddleware struct {
	pkg.UnimplementedMiddleware
	logger   *logger.Logger
	spider   pkg.Spider
	uaAll    []string
	uaLen    int
	ua       []string
	platform platform
	browser  browser
}

func (m *DeviceMiddleware) SpiderStart(_ context.Context, spider pkg.Spider) (err error) {
	m.spider = spider
	m.uaAll = []string{}
	m.uaLen = len(m.uaAll)
	m.ua = []string{}
	return
}

func (m *DeviceMiddleware) ProcessRequest(c *pkg.Context) (err error) {
	request := c.Request

	ua := m.ua[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(m.uaLen)]

	if request.UserAgent() == "" {
		request.SetHeader("User-Agent", ua)
	}

	err = c.NextRequest()
	return
}

func NewDeviceMiddleware(logger *logger.Logger) (m pkg.Middleware) {
	m = &DeviceMiddleware{
		logger: logger,
	}
	return
}
