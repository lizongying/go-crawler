package context

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"time"
)

type Schedule struct {
	Id        string             `json:"id,omitempty"`
	Status    pkg.ScheduleStatus `json:"status,omitempty"`
	Enable    bool               `json:"enable,omitempty"`
	StartTime utils.Timestamp    `json:"start_time,omitempty"`
	StopTime  utils.Timestamp    `json:"stop_time,omitempty"`
}

func (c *Schedule) GetId() string {
	return c.Id
}
func (c *Schedule) WithId(id string) pkg.ContextSchedule {
	c.Id = id
	return c
}
func (c *Schedule) GetStatus() pkg.ScheduleStatus {
	return c.Status
}
func (c *Schedule) WithStatus(status pkg.ScheduleStatus) pkg.ContextSchedule {
	c.Status = status
	return c
}
func (c *Schedule) GetEnable() bool {
	return c.Enable
}
func (c *Schedule) WithEnable(enable bool) pkg.ContextSchedule {
	c.Enable = enable
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
