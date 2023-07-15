package pkg

import (
	"context"
)

type HttpClient interface {
	DoRequest(context.Context, Request) (response *Response, err error)
	FromCrawler(Crawler) HttpClient
}
