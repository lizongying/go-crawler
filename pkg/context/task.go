package context

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"time"
)

type Task struct {
	pkg.Task  `json:"-"`
	pkg.Stats `json:"stats,omitempty"`
	Context   context.Context     `json:"-"`
	Id        string              `json:"id,omitempty"`
	Status    pkg.TaskStatus      `json:"status,omitempty"`
	StartTime utils.Timestamp     `json:"start_time,omitempty"`
	StopTime  utils.Timestamp     `json:"stop_time,omitempty"`
	Deadline  utils.TimestampNano `json:"deadline,omitempty"`
}

func (c *Task) GetTask() pkg.Task {
	return c.Task
}
func (c *Task) WithTask(task pkg.Task) pkg.ContextTask {
	c.Task = task
	return c
}
func (c *Task) GetStats() pkg.Stats {
	return c.Stats
}
func (c *Task) WithStats(stats pkg.Stats) pkg.ContextTask {
	c.Stats = stats
	return c
}
func (c *Task) GetId() string {
	return c.Id
}
func (c *Task) WithId(id string) pkg.ContextTask {
	c.Id = id
	return c
}
func (c *Task) GetContext() context.Context {
	return c.Context
}
func (c *Task) WithContext(ctx context.Context) pkg.ContextTask {
	c.Context = ctx
	return c
}
func (c *Task) GetStatus() pkg.TaskStatus {
	return c.Status
}
func (c *Task) WithStatus(status pkg.TaskStatus) pkg.ContextTask {
	c.Status = status
	return c
}
func (c *Task) GetStartTime() time.Time {
	return c.StartTime.Time
}
func (c *Task) WithStartTime(startTime time.Time) pkg.ContextTask {
	c.StartTime = utils.Timestamp{Time: startTime}
	return c
}
func (c *Task) GetStopTime() time.Time {
	return c.StopTime.Time
}
func (c *Task) WithStopTime(stopTime time.Time) pkg.ContextTask {
	c.StopTime = utils.Timestamp{Time: stopTime}
	return c
}
func (c *Task) GetDeadline() time.Time {
	return c.Deadline.Time
}
func (c *Task) WithDeadline(deadline time.Time) pkg.ContextTask {
	c.Deadline = utils.TimestampNano{Time: deadline}
	return c
}
