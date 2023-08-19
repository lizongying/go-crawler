package pkg

import "context"

type Context struct {
	context context.Context
	Spider  Spider
	Meta    Meta
	TaskId  string
}

func (c *Context) Context() context.Context {
	return c.context
}
func (c *Context) WithContext(ctx context.Context) *Context {
	c.context = ctx
	return c
}
