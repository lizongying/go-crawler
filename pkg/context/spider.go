package context

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"time"
)

type Spider struct {
	Spider     pkg.Spider       `json:"-"`
	Context    context.Context  `json:"-"`
	Id         uint64           `json:"id,omitempty"`
	Name       string           `json:"name,omitempty"`
	Status     pkg.SpiderStatus `json:"status,omitempty"`
	StartTime  utils.Timestamp  `json:"start_time,omitempty"`
	StopTime   utils.Timestamp  `json:"stop_time,omitempty"`
	UpdateTime utils.Timestamp  `json:"update_time,omitempty"`
}

func (c *Spider) GetSpider() pkg.Spider {
	return c.Spider
}
func (c *Spider) WithSpider(spider pkg.Spider) pkg.ContextSpider {
	c.Spider = spider
	return c
}
func (c *Spider) GetId() uint64 {
	return c.Id
}
func (c *Spider) WithId(id uint64) pkg.ContextSpider {
	c.Id = id
	return c
}
func (c *Spider) GetName() string {
	return c.Name
}
func (c *Spider) WithName(name string) pkg.ContextSpider {
	c.Name = name
	return c
}
func (c *Spider) GetContext() context.Context {
	return c.Context
}
func (c *Spider) WithContext(ctx context.Context) pkg.ContextSpider {
	c.Context = ctx
	return c
}
func (c *Spider) GetStatus() pkg.SpiderStatus {
	return c.Status
}
func (c *Spider) WithStatus(status pkg.SpiderStatus) pkg.ContextSpider {
	c.Status = status
	t := time.Now()
	c.withUpdateTime(t)
	switch status {
	case pkg.SpiderStatusRunning:
		c.withStartTime(t)
	case pkg.SpiderStatusStopped:
		c.withStopTime(t)
	}
	return c
}
func (c *Spider) GetStartTime() time.Time {
	return c.StartTime.Time
}
func (c *Spider) withStartTime(startTime time.Time) pkg.ContextSpider {
	c.StartTime = utils.Timestamp{Time: startTime}
	return c
}
func (c *Spider) GetStopTime() time.Time {
	return c.StopTime.Time
}
func (c *Spider) withStopTime(stopTime time.Time) pkg.ContextSpider {
	c.StopTime = utils.Timestamp{Time: stopTime}
	return c
}
func (c *Spider) GetUpdateTime() time.Time {
	return c.UpdateTime.Time
}
func (c *Spider) withUpdateTime(t time.Time) pkg.ContextSpider {
	c.UpdateTime = utils.Timestamp{Time: t}
	return c
}
