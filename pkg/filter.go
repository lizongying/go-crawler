package pkg

import "context"

type FilterType string

const (
	FilterUnknown FilterType = ""
	FilterMemory  FilterType = "memory"
	FilterRedis   FilterType = "redis"
)

type Filter interface {
	IsExist(Context, any) (bool, error)
	Store(Context, any) error
	Clean(context.Context) error
}
