package context

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"time"
)

type Spider struct {
	Name      string           `json:"name,omitempty"`
	Status    pkg.SpiderStatus `json:"status,omitempty"`
	StartTime utils.Timestamp  `json:"start_time,omitempty"`
	StopTime  utils.Timestamp  `json:"stop_time,omitempty"`
}

func (c *Spider) GetName() string {
	return c.Name
}
func (c *Spider) WithName(name string) pkg.ContextSpider {
	c.Name = name
	return c
}
func (c *Spider) GetStatus() pkg.SpiderStatus {
	return c.Status
}
func (c *Spider) WithStatus(status pkg.SpiderStatus) pkg.ContextSpider {
	c.Status = status
	return c
}
func (c *Spider) GetStartTime() time.Time {
	return c.StartTime.Time
}
func (c *Spider) WithStartTime(startTime time.Time) pkg.ContextSpider {
	c.StartTime = utils.Timestamp{Time: startTime}
	return c
}
func (c *Spider) GetStopTime() time.Time {
	return c.StopTime.Time
}
func (c *Spider) WithStopTime(stopTime time.Time) pkg.ContextSpider {
	c.StopTime = utils.Timestamp{Time: stopTime}
	return c
}
