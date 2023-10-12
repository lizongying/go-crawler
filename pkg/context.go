package pkg

import "context"

type Context interface {
	Global() Context
	GlobalContext() context.Context
	WithGlobalContext(ctx context.Context) Context
	Spider() Spider
	WithSpider(spider Spider) Context
	Meta() Meta
	WithMeta(meta Meta) Context
	TaskId() string
	WithTaskId(taskId string) Context
	SpiderName() string
	WithSpiderName(spiderName string) Context
	StartFunc() string
	WithStartFunc(startFunc string) Context
	Args() string
	WithArgs(args string) Context
	Mode() string
	WithMode(mode string) Context
	ToContextJson() ContextJson
}

type ContextJson interface {
	ToContext() Context
}
