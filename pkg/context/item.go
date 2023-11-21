package context

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"time"
)

type Item struct {
	Context    context.Context     `json:"-"`
	Id         string              `json:"id,omitempty"`
	Status     pkg.ItemStatus      `json:"status,omitempty"`
	StartTime  utils.Timestamp     `json:"start_time,omitempty"`
	StopTime   utils.Timestamp     `json:"stop_time,omitempty"`
	UpdateTime utils.Timestamp     `json:"update_time,omitempty"`
	Deadline   utils.TimestampNano `json:"deadline,omitempty"`
	Cookies    map[string]string   `json:"cookies,omitempty"`
	Referrer   string              `json:"referrer,omitempty"`
	Saved      bool                `json:"saved,omitempty"`
	StopReason string              `json:"stop_reason,omitempty"`
}

func (c *Item) GetId() string {
	return c.Id
}
func (c *Item) WithId(id string) pkg.ContextItem {
	c.Id = id
	return c
}
func (c *Item) GetContext() context.Context {
	return c.Context
}
func (c *Item) WithContext(ctx context.Context) pkg.ContextItem {
	c.Context = ctx
	return c
}
func (c *Item) GetStatus() pkg.ItemStatus {
	return c.Status
}
func (c *Item) WithStatus(status pkg.ItemStatus) pkg.ContextItem {
	c.Status = status
	t := time.Now()
	c.withUpdateTime(t)
	switch status {
	case pkg.ItemStatusRunning:
		c.withStartTime(t)
	case pkg.ItemStatusSuccess:
		c.withStopTime(t)
	case pkg.ItemStatusFailure:
		c.withStopTime(t)
	}

	return c
}
func (c *Item) GetStartTime() time.Time {
	return c.StartTime.Time
}
func (c *Item) withStartTime(startTime time.Time) pkg.ContextItem {
	c.StartTime = utils.Timestamp{Time: startTime}
	return c
}
func (c *Item) GetStopTime() time.Time {
	return c.StopTime.Time
}
func (c *Item) withStopTime(stopTime time.Time) pkg.ContextItem {
	c.StopTime = utils.Timestamp{Time: stopTime}
	return c
}
func (c *Item) GetUpdateTime() time.Time {
	return c.UpdateTime.Time
}
func (c *Item) withUpdateTime(t time.Time) pkg.ContextItem {
	c.UpdateTime = utils.Timestamp{Time: t}
	return c
}
func (c *Item) GetSaved() bool {
	return c.Saved
}
func (c *Item) WithSaved(saved bool) pkg.ContextItem {
	c.Saved = saved
	return c
}
func (c *Item) GetStopReason() string {
	return c.StopReason
}
func (c *Item) WithStopReason(stopReason string) pkg.ContextItem {
	c.StopReason = stopReason
	return c
}
