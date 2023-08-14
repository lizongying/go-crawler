package pkg

import (
	"context"
)

type HttpClient interface {
	DoRequest(context.Context, Request) (Response, error)
	Close(context.Context) error
}
