package context

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"time"
)

type Request struct {
	Context    context.Context     `json:"-"`
	Id         string              `json:"id,omitempty"`
	Status     pkg.RequestStatus   `json:"status,omitempty"`
	StartTime  utils.Timestamp     `json:"start_time,omitempty"`
	StopTime   utils.Timestamp     `json:"stop_time,omitempty"`
	UpdateTime utils.Timestamp     `json:"update_time,omitempty"`
	Deadline   utils.TimestampNano `json:"deadline,omitempty"`
	Cookies    map[string]string   `json:"cookies,omitempty"`
	Referrer   string              `json:"referrer,omitempty"`
	StopReason string              `json:"stop_reason,omitempty"`
}

func (c *Request) GetId() string {
	return c.Id
}
func (c *Request) WithId(id string) pkg.ContextRequest {
	c.Id = id
	return c
}
func (c *Request) GetContext() context.Context {
	return c.Context
}
func (c *Request) WithContext(ctx context.Context) pkg.ContextRequest {
	c.Context = ctx
	return c
}
func (c *Request) GetStatus() pkg.RequestStatus {
	return c.Status
}
func (c *Request) WithStatus(status pkg.RequestStatus) pkg.ContextRequest {
	c.Status = status
	t := time.Now()
	c.withUpdateTime(t)
	switch status {
	case pkg.RequestStatusRunning:
		c.WithStartTime(t)
	case pkg.RequestStatusSuccess:
		c.WithStopTime(t)
	case pkg.RequestStatusFailure:
		c.WithStopTime(t)
	}

	return c
}
func (c *Request) GetStartTime() time.Time {
	return c.StartTime.Time
}
func (c *Request) WithStartTime(startTime time.Time) pkg.ContextRequest {
	c.StartTime = utils.Timestamp{Time: startTime}
	return c
}
func (c *Request) GetStopTime() time.Time {
	return c.StopTime.Time
}
func (c *Request) WithStopTime(stopTime time.Time) pkg.ContextRequest {
	c.StopTime = utils.Timestamp{Time: stopTime}
	return c
}
func (c *Request) GetUpdateTime() time.Time {
	return c.UpdateTime.Time
}
func (c *Request) withUpdateTime(t time.Time) pkg.ContextRequest {
	c.UpdateTime = utils.Timestamp{Time: t}
	return c
}
func (c *Request) GetDeadline() time.Time {
	return c.Deadline.Time
}
func (c *Request) WithDeadline(deadline time.Time) pkg.ContextRequest {
	c.Deadline = utils.TimestampNano{Time: deadline}
	return c
}
func (c *Request) GetCookies() map[string]string {
	return c.Cookies
}
func (c *Request) WithCookies(cookies map[string]string) pkg.ContextRequest {
	c.Cookies = cookies
	return c
}
func (c *Request) GetReferrer() string {
	return c.Referrer
}
func (c *Request) WithReferrer(referrer string) pkg.ContextRequest {
	c.Referrer = referrer
	return c
}
func (c *Request) GetStopReason() string {
	return c.StopReason
}
func (c *Request) WithStopReason(stopReason string) pkg.ContextRequest {
	c.StopReason = stopReason
	return c
}
