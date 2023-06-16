package pkg

import (
	"context"
)

type Downloader interface {
	Download(context.Context, *Request) (*Response, error)
	GetMiddlewareNames() map[uint8]string
	GetMiddlewares() []Middleware
	SetMiddleware(Middleware, uint8)
	DelMiddleware(int)
	CleanMiddlewares()
	FromCrawler(Crawler) Downloader
}
