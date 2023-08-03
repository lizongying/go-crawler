package pkg

import "context"

type Context struct {
	context context.Context
	Spider  Spider
	Meta    Meta
}

func (c *Context) Context() context.Context {
	return c.context
}
