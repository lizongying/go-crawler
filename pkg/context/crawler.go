package context

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"time"
)

type Crawler struct {
	Id        string            `json:"id,omitempty"`
	Status    pkg.CrawlerStatus `json:"status,omitempty"`
	StartTime utils.Timestamp   `json:"start_time,omitempty"`
	StopTime  utils.Timestamp   `json:"stop_time,omitempty"`
}

func (c *Crawler) GetId() string {
	return c.Id
}
func (c *Crawler) WithId(id string) pkg.ContextCrawler {
	c.Id = id
	return c
}
func (c *Crawler) GetStatus() pkg.CrawlerStatus {
	return c.Status
}
func (c *Crawler) WithStatus(status pkg.CrawlerStatus) pkg.ContextCrawler {
	c.Status = status
	return c
}
func (c *Crawler) GetStartTime() time.Time {
	return c.StartTime.Time
}
func (c *Crawler) WithStartTime(startTime time.Time) pkg.ContextCrawler {
	c.StartTime = utils.Timestamp{Time: startTime}
	return c
}
func (c *Crawler) GetStopTime() time.Time {
	return c.StopTime.Time
}
func (c *Crawler) WithStopTime(stopTime time.Time) pkg.ContextCrawler {
	c.StopTime = utils.Timestamp{Time: stopTime}
	return c
}
