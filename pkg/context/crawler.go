package context

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"time"
)

type Crawler struct {
	Context    context.Context   `json:"-"`
	Id         string            `json:"id,omitempty"`
	Status     pkg.CrawlerStatus `json:"status,omitempty"`
	StartTime  utils.Timestamp   `json:"start_time,omitempty"`
	StopTime   utils.Timestamp   `json:"stop_time,omitempty"`
	UpdateTime utils.Timestamp   `json:"update_time,omitempty"`
	StopReason string            `json:"stop_reason,omitempty"`
}

func (c *Crawler) GetId() string {
	return c.Id
}
func (c *Crawler) WithId(id string) pkg.ContextCrawler {
	c.Id = id
	return c
}
func (c *Crawler) GetContext() context.Context {
	return c.Context
}
func (c *Crawler) WithContext(ctx context.Context) pkg.ContextCrawler {
	c.Context = ctx
	return c
}
func (c *Crawler) GetStatus() pkg.CrawlerStatus {
	return c.Status
}
func (c *Crawler) WithStatus(status pkg.CrawlerStatus) pkg.ContextCrawler {
	c.Status = status
	t := time.Now()
	c.withUpdateTime(t)
	switch status {
	case pkg.CrawlerStatusRunning:
		c.withStartTime(t)
	case pkg.CrawlerStatusStopped:
		c.withStopTime(t)
	}

	return c
}
func (c *Crawler) GetStartTime() time.Time {
	return c.StartTime.Time
}
func (c *Crawler) withStartTime(startTime time.Time) pkg.ContextCrawler {
	c.StartTime = utils.Timestamp{Time: startTime}
	return c
}
func (c *Crawler) GetStopTime() time.Time {
	return c.StopTime.Time
}
func (c *Crawler) withStopTime(stopTime time.Time) pkg.ContextCrawler {
	c.StopTime = utils.Timestamp{Time: stopTime}
	return c
}
func (c *Crawler) GetUpdateTime() time.Time {
	return c.UpdateTime.Time
}
func (c *Crawler) withUpdateTime(t time.Time) pkg.ContextCrawler {
	c.UpdateTime = utils.Timestamp{Time: t}
	return c
}
