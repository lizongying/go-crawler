package pkg

import (
	"context"
)

type HttpClient interface {
	DoRequest(context.Context, Request) (Response, error)
	FromCrawler(Crawler) HttpClient
}
