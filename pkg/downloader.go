package pkg

import (
	"context"
)

type Downloader interface {
	Download(context.Context, Request) (Response, error)
	GetMiddlewareNames() map[uint8]string
	GetMiddlewares() []Middleware
	SetMiddleware(Middleware, uint8)
	DelMiddleware(int)
	CleanMiddlewares()
	WithStatsMiddleware()
	WithDumpMiddleware()
	WithProxyMiddleware()
	WithRobotsTxtMiddleware()
	WithFilterMiddleware()
	WithFileMiddleware()
	WithImageMiddleware()
	WithHttpMiddleware()
	WithRetryMiddleware()
	WithUrlMiddleware()
	WithRefererMiddleware()
	WithCookieMiddleware()
	WithRedirectMiddleware()
	WithChromeMiddleware()
	WithHttpAuthMiddleware()
	WithCompressMiddleware()
	WithDecodeMiddleware()
	WithDeviceMiddleware()
	WithCustomMiddleware(Middleware)
	FromCrawler(Crawler) Downloader
}
