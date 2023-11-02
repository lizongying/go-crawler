package context

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"time"
)

type Schedule struct {
	Context     context.Context    `json:"-"`
	Id          string             `json:"id,omitempty"`
	Status      pkg.ScheduleStatus `json:"status,omitempty"`
	StartTime   utils.Timestamp    `json:"start_time,omitempty"`
	StopTime    utils.Timestamp    `json:"stop_time,omitempty"`
	Enable      bool               `json:"enable,omitempty"`
	Func        string             `json:"func,omitempty"`
	Args        string             `json:"args,omitempty"`
	Mode        pkg.ScheduleMode   `json:"mode,omitempty"`
	Spec        string             `json:"spec,omitempty"`
	OnlyOneTask bool               `json:"only_one_task,omitempty"`
}

func (c *Schedule) GetId() string {
	return c.Id
}
func (c *Schedule) WithId(id string) pkg.ContextSchedule {
	c.Id = id
	return c
}
func (c *Schedule) GetContext() context.Context {
	return c.Context
}
func (c *Schedule) WithContext(ctx context.Context) pkg.ContextSchedule {
	c.Context = ctx
	return c
}
func (c *Schedule) GetStatus() pkg.ScheduleStatus {
	return c.Status
}
func (c *Schedule) WithStatus(status pkg.ScheduleStatus) pkg.ContextSchedule {
	c.Status = status
	return c
}
func (c *Schedule) GetStartTime() time.Time {
	return c.StartTime.Time
}
func (c *Schedule) WithStartTime(startTime time.Time) pkg.ContextSchedule {
	c.StartTime = utils.Timestamp{Time: startTime}
	return c
}
func (c *Schedule) GetStopTime() time.Time {
	return c.StopTime.Time
}
func (c *Schedule) WithStopTime(stopTime time.Time) pkg.ContextSchedule {
	c.StopTime = utils.Timestamp{Time: stopTime}
	return c
}
func (c *Schedule) GetEnable() bool {
	return c.Enable
}
func (c *Schedule) WithEnable(enable bool) pkg.ContextSchedule {
	c.Enable = enable
	return c
}
func (c *Schedule) GetFunc() string {
	return c.Func
}
func (c *Schedule) WithFunc(fn string) pkg.ContextSchedule {
	c.Func = fn
	return c
}
func (c *Schedule) GetArgs() string {
	return c.Args
}
func (c *Schedule) WithArgs(args string) pkg.ContextSchedule {
	c.Args = args
	return c
}
func (c *Schedule) GetMode() pkg.ScheduleMode {
	return c.Mode
}
func (c *Schedule) WithMode(mode pkg.ScheduleMode) pkg.ContextSchedule {
	c.Mode = mode
	return c
}
func (c *Schedule) GetSpec() string {
	return c.Spec
}
func (c *Schedule) WithSpec(spec string) pkg.ContextSchedule {
	c.Spec = spec
	return c
}
func (c *Schedule) GetOnlyOneTask() bool {
	return c.OnlyOneTask
}
func (c *Schedule) WithOnlyOneTask(onlyOneTask bool) pkg.ContextSchedule {
	c.OnlyOneTask = onlyOneTask
	return c
}
