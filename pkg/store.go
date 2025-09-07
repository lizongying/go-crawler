package pkg

import (
	"golang.org/x/net/context"
)

type Store interface {
	// Save
	// is s3:
	//
	//	prefix=bucket
	//
	// is file:
	//
	//	prefix=dir
	Save(prefix string, key string, body []byte) (storePath string, err error)
	Close(ctx context.Context) error
}
