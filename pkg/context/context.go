package context

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
)

type Context struct {
	context context.Context
	Spider  pkg.Spider
	meta    pkg.Meta
	TaskId  string `json:"task_id"`
}

func (c *Context) WithContext(ctx context.Context) pkg.Context {
	c.context = ctx
	return c
}
func (c *Context) Global() pkg.Context {
	return c
}
func (c *Context) GlobalContext() context.Context {
	return c.context
}
func (c *Context) WithGlobalContext(ctx context.Context) pkg.Context {
	c.context = ctx
	return c
}
func (c *Context) Meta() pkg.Meta {
	return c.meta
}
func (c *Context) WithMeta(meta pkg.Meta) pkg.Context {
	c.meta = meta
	return c
}
func (c *Context) GetTaskId() string {
	return c.TaskId
}
func (c *Context) WithTaskId(taskId string) pkg.Context {
	c.TaskId = taskId
	return c
}
func (c *Context) ToContextJson() (contextJson pkg.ContextJson) {
	contextJson = &Json{
		TaskId: c.TaskId,
	}
	return
}
