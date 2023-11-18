package context

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"time"
)

type Job struct {
	Context     context.Context `json:"-"`
	Id          string          `json:"id,omitempty"`
	SubId       uint64          `json:"sub_id,omitempty"`
	Status      pkg.JobStatus   `json:"status,omitempty"`
	StartTime   utils.Timestamp `json:"start_time,omitempty"`
	StopTime    utils.Timestamp `json:"stop_time,omitempty"`
	UpdateTime  utils.Timestamp `json:"update_time,omitempty"`
	Enable      bool            `json:"enable,omitempty"`
	Func        string          `json:"func,omitempty"`
	Args        string          `json:"args,omitempty"`
	Mode        pkg.JobMode     `json:"mode,omitempty"`
	Spec        string          `json:"spec,omitempty"`
	OnlyOneTask bool            `json:"only_one_task,omitempty"`
	StopReason  string          `json:"stop_reason,omitempty"`
}

func (c *Job) GetId() string {
	return c.Id
}
func (c *Job) WithId(id string) pkg.ContextJob {
	c.Id = id
	return c
}
func (c *Job) GetSubId() uint64 {
	return c.SubId
}
func (c *Job) WithSubId(id uint64) pkg.ContextJob {
	c.SubId = id
	return c
}
func (c *Job) GetContext() context.Context {
	return c.Context
}
func (c *Job) WithContext(ctx context.Context) pkg.ContextJob {
	c.Context = ctx
	return c
}
func (c *Job) GetStatus() pkg.JobStatus {
	return c.Status
}
func (c *Job) WithStatus(status pkg.JobStatus) pkg.ContextJob {
	c.Status = status
	t := time.Now()
	c.withUpdateTime(t)
	switch status {
	case pkg.JobStatusRunning:
		c.withStartTime(t)
	case pkg.JobStatusStopped:
		c.withStopTime(t)
	}

	return c
}
func (c *Job) GetStartTime() time.Time {
	return c.StartTime.Time
}
func (c *Job) withStartTime(startTime time.Time) pkg.ContextJob {
	c.StartTime = utils.Timestamp{Time: startTime}
	return c
}
func (c *Job) GetStopTime() time.Time {
	return c.StopTime.Time
}
func (c *Job) withStopTime(stopTime time.Time) pkg.ContextJob {
	c.StopTime = utils.Timestamp{Time: stopTime}
	return c
}
func (c *Job) GetUpdateTime() time.Time {
	return c.UpdateTime.Time
}
func (c *Job) withUpdateTime(t time.Time) pkg.ContextJob {
	c.UpdateTime = utils.Timestamp{Time: t}
	return c
}
func (c *Job) GetEnable() bool {
	return c.Enable
}
func (c *Job) WithEnable(enable bool) pkg.ContextJob {
	c.Enable = enable
	return c
}
func (c *Job) GetFunc() string {
	return c.Func
}
func (c *Job) WithFunc(fn string) pkg.ContextJob {
	c.Func = fn
	return c
}
func (c *Job) GetArgs() string {
	return c.Args
}
func (c *Job) WithArgs(args string) pkg.ContextJob {
	c.Args = args
	return c
}
func (c *Job) GetMode() pkg.JobMode {
	return c.Mode
}
func (c *Job) WithMode(mode pkg.JobMode) pkg.ContextJob {
	c.Mode = mode
	return c
}
func (c *Job) GetSpec() string {
	return c.Spec
}
func (c *Job) WithSpec(spec string) pkg.ContextJob {
	c.Spec = spec
	return c
}
func (c *Job) GetOnlyOneTask() bool {
	return c.OnlyOneTask
}
func (c *Job) WithOnlyOneTask(onlyOneTask bool) pkg.ContextJob {
	c.OnlyOneTask = onlyOneTask
	return c
}
func (c *Job) GetStopReason() string {
	return c.StopReason
}
func (c *Job) WithStopReason(stopReason string) pkg.ContextJob {
	c.StopReason = stopReason
	return c
}
