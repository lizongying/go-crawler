package context

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"time"
)

type Context struct {
	context    context.Context
	spider     pkg.Spider
	meta       pkg.Meta
	taskId     string
	deadline   time.Time
	spiderName string
	startFunc  string
	args       string
	mode       string
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
func (c *Context) Spider() pkg.Spider {
	return c.spider
}
func (c *Context) WithSpider(spider pkg.Spider) pkg.Context {
	c.spider = spider
	return c
}
func (c *Context) Meta() pkg.Meta {
	return c.meta
}
func (c *Context) WithMeta(meta pkg.Meta) pkg.Context {
	c.meta = meta
	return c
}
func (c *Context) TaskId() string {
	return c.taskId
}
func (c *Context) WithTaskId(taskId string) pkg.Context {
	c.taskId = taskId
	return c
}
func (c *Context) Deadline() time.Time {
	return c.deadline
}
func (c *Context) WithDeadline(deadline time.Time) pkg.Context {
	c.deadline = deadline
	return c
}
func (c *Context) SpiderName() string {
	return c.spiderName
}
func (c *Context) WithSpiderName(spiderName string) pkg.Context {
	c.spiderName = spiderName
	return c
}
func (c *Context) StartFunc() string {
	return c.startFunc
}
func (c *Context) WithStartFunc(startFunc string) pkg.Context {
	c.startFunc = startFunc
	return c
}
func (c *Context) Args() string {
	return c.args
}
func (c *Context) WithArgs(args string) pkg.Context {
	c.args = args
	return c
}
func (c *Context) Mode() string {
	return c.mode
}
func (c *Context) WithMode(mode string) pkg.Context {
	c.mode = mode
	return c
}
func (c *Context) ToContextJson() (contextJson pkg.ContextJson) {
	contextJson = &Json{
		TaskId:     c.taskId,
		Deadline:   c.deadline.UnixNano(),
		SpiderName: c.spiderName,
		StartFunc:  c.startFunc,
		Args:       c.args,
		Mode:       c.mode,
	}
	return
}
