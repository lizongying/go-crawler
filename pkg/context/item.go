package context

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"time"
)

type Item struct {
	Context   context.Context     `json:"-"`
	Id        string              `json:"id,omitempty"`
	Status    pkg.ItemStatus      `json:"status,omitempty"`
	StartTime utils.Timestamp     `json:"start_time,omitempty"`
	StopTime  utils.Timestamp     `json:"stop_time,omitempty"`
	Deadline  utils.TimestampNano `json:"deadline,omitempty"`
	Cookies   map[string]string   `json:"cookies,omitempty"`
	Referrer  string              `json:"referrer,omitempty"`
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
	return c
}
func (c *Item) GetStartTime() time.Time {
	return c.StartTime.Time
}
func (c *Item) WithStartTime(startTime time.Time) pkg.ContextItem {
	c.StartTime = utils.Timestamp{Time: startTime}
	return c
}
func (c *Item) GetStopTime() time.Time {
	return c.StopTime.Time
}
func (c *Item) WithStopTime(stopTime time.Time) pkg.ContextItem {
	c.StopTime = utils.Timestamp{Time: stopTime}
	return c
}
