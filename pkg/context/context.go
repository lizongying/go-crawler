package context

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"time"
)

type Context struct {
	Context context.Context `json:"-"`

	// request
	Meta pkg.Meta `json:"meta,omitempty"`

	// task
	TaskId    string              `json:"task_id,omitempty"`
	Deadline  utils.TimestampNano `json:"deadline,omitempty"`
	StartFunc string              `json:"start_func,omitempty"`
	Args      string              `json:"args,omitempty"`
	Mode      string              `json:"mode,omitempty"`

	// spider
	SpiderName string           `json:"spider_name,omitempty"`
	Status     pkg.SpiderStatus `json:"status,omitempty"`
	StartTime  utils.Timestamp  `json:"start_time,omitempty"`
	StopTime   utils.Timestamp  `json:"stop_time,omitempty"`
}

func (c *Context) Global() pkg.Context {
	return c
}
func (c *Context) GlobalContext() context.Context {
	return c.Context
}
func (c *Context) WithGlobalContext(ctx context.Context) pkg.Context {
	c.Context = ctx
	return c
}
func (c *Context) GetMeta() pkg.Meta {
	return c.Meta
}
func (c *Context) WithMeta(meta pkg.Meta) pkg.Context {
	c.Meta = meta
	return c
}
func (c *Context) GetTaskId() string {
	return c.TaskId
}
func (c *Context) WithTaskId(taskId string) pkg.Context {
	c.TaskId = taskId
	return c
}
func (c *Context) GetDeadline() time.Time {
	return c.Deadline.Time
}
func (c *Context) WithDeadline(t time.Time) pkg.Context {
	c.Deadline = utils.TimestampNano{Time: t}
	return c
}
func (c *Context) GetStartFunc() string {
	return c.StartFunc
}
func (c *Context) WithStartFunc(startFunc string) pkg.Context {
	c.StartFunc = startFunc
	return c
}
func (c *Context) GetArgs() string {
	return c.Args
}
func (c *Context) WithArgs(args string) pkg.Context {
	c.Args = args
	return c
}
func (c *Context) GetMode() string {
	return c.Mode
}
func (c *Context) WithMode(mode string) pkg.Context {
	c.Mode = mode
	return c
}
func (c *Context) GetSpiderName() string {
	return c.SpiderName
}
func (c *Context) WithSpiderName(spiderName string) pkg.Context {
	c.SpiderName = spiderName
	return c
}
func (c *Context) GetStatus() pkg.SpiderStatus {
	return c.Status
}
func (c *Context) WithStatus(status pkg.SpiderStatus) pkg.Context {
	c.Status = status
	return c
}
func (c *Context) GetStartTime() time.Time {
	return c.StartTime.Time
}
func (c *Context) WithStartTime(t time.Time) pkg.Context {
	c.StartTime = utils.Timestamp{Time: t}
	return c
}
func (c *Context) GetStopTime() time.Time {
	return c.StopTime.Time
}
func (c *Context) WithStopTime(t time.Time) pkg.Context {
	c.StopTime = utils.Timestamp{Time: t}
	return c
}
