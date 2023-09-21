package pkg

type Middlewares interface {
	MiddlewareNames() map[uint8]string
	Middlewares() []Middleware
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
	WithReferrerMiddleware()
	WithCookieMiddleware()
	WithRedirectMiddleware()
	WithChromeMiddleware()
	WithHttpAuthMiddleware()
	WithCompressMiddleware()
	WithDecodeMiddleware()
	WithDeviceMiddleware()
	WithCustomMiddleware(Middleware)
}
