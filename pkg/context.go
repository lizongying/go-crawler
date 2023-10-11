package pkg

import "context"

type Context interface {
	Global() Context
	GlobalContext() context.Context
	WithGlobalContext(ctx context.Context) Context
	Meta() Meta
	WithMeta(meta Meta) Context
	GetTaskId() string
	WithTaskId(taskId string) Context
	ToContextJson() ContextJson
}

type ContextJson interface {
	ToContext() Context
}
