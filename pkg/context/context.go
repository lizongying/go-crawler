package context

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"time"
)

type Context struct {
	Context context.Context `json:"-"`

	// request
	Meta pkg.Meta `json:"meta,omitempty"`

	// task
	TaskId        string              `json:"task_id,omitempty"`
	Deadline      utils.TimestampNano `json:"deadline,omitempty"`
	StartFunc     string              `json:"start_func,omitempty"`
	Args          string              `json:"args,omitempty"`
	Mode          string              `json:"mode,omitempty"`
	TaskStatus    pkg.TaskStatus      `json:"task_status,omitempty"`
	TaskStartTime utils.Timestamp     `json:"task_start_time,omitempty"`
	TaskStopTime  utils.Timestamp     `json:"task_stop_time,omitempty"`

	Crawler  *Crawler  `json:"crawler,omitempty"`
	Spider   *Spider   `json:"spider,omitempty"`
	Schedule *Schedule `json:"schedule,omitempty"`
}

func (c *Context) Global() pkg.Context {
	return c
}
func (c *Context) GlobalContext() context.Context {
	return c.Context
}
func (c *Context) WithGlobalContext(ctx context.Context) pkg.Context {
	c.Context = ctx
	return c
}
func (c *Context) GetMeta() pkg.Meta {
	return c.Meta
}
func (c *Context) WithMeta(meta pkg.Meta) pkg.Context {
	c.Meta = meta
	return c
}
func (c *Context) GetTaskId() string {
	return c.TaskId
}
func (c *Context) WithTaskId(taskId string) pkg.Context {
	c.TaskId = taskId
	return c
}
func (c *Context) GetDeadline() time.Time {
	return c.Deadline.Time
}
func (c *Context) WithDeadline(t time.Time) pkg.Context {
	c.Deadline = utils.TimestampNano{Time: t}
	return c
}
func (c *Context) GetStartFunc() string {
	return c.StartFunc
}
func (c *Context) WithStartFunc(startFunc string) pkg.Context {
	c.StartFunc = startFunc
	return c
}
func (c *Context) GetArgs() string {
	return c.Args
}
func (c *Context) WithArgs(args string) pkg.Context {
	c.Args = args
	return c
}
func (c *Context) GetMode() string {
	return c.Mode
}
func (c *Context) WithMode(mode string) pkg.Context {
	c.Mode = mode
	return c
}
func (c *Context) GetTaskStatus() pkg.TaskStatus {
	return c.TaskStatus
}
func (c *Context) WithTaskStatus(status pkg.TaskStatus) pkg.Context {
	c.TaskStatus = status
	return c
}
func (c *Context) GetTaskStartTime() time.Time {
	return c.TaskStartTime.Time
}
func (c *Context) WithTaskStartTime(t time.Time) pkg.Context {
	c.TaskStartTime = utils.Timestamp{Time: t}
	return c
}
func (c *Context) GetTaskStopTime() time.Time {
	return c.TaskStopTime.Time
}
func (c *Context) WithTaskStopTime(t time.Time) pkg.Context {
	c.TaskStopTime = utils.Timestamp{Time: t}
	return c
}

func (c *Context) GetCrawler() pkg.ContextCrawler {
	return c.Crawler
}
func (c *Context) GetCrawlerId() string {
	return c.Crawler.GetId()
}
func (c *Context) WithCrawlerId(id string) pkg.Context {
	c.Crawler.WithId(id)
	return c
}
func (c *Context) GetCrawlerStatus() pkg.CrawlerStatus {
	return c.Crawler.GetStatus()
}
func (c *Context) WithCrawlerStatus(status pkg.CrawlerStatus) pkg.Context {
	c.Crawler.WithStatus(status)
	return c
}
func (c *Context) GetCrawlerStartTime() time.Time {
	return c.Crawler.GetStartTime()
}
func (c *Context) WithCrawlerStartTime(startTime time.Time) pkg.Context {
	c.Crawler.WithStartTime(startTime)
	return c
}
func (c *Context) GetCrawlerStopTime() time.Time {
	return c.Crawler.GetStopTime()
}
func (c *Context) WithCrawlerStopTime(stopTime time.Time) pkg.Context {
	c.Crawler.WithStopTime(stopTime)
	return c
}

func (c *Context) GetSpider() pkg.ContextSpider {
	return c.Spider
}
func (c *Context) GetSpiderName() string {
	return c.Spider.GetName()
}
func (c *Context) WithSpiderName(name string) pkg.Context {
	c.Spider.WithName(name)
	return c
}
func (c *Context) GetSpiderStatus() pkg.SpiderStatus {
	return c.Spider.GetStatus()
}
func (c *Context) WithSpiderStatus(status pkg.SpiderStatus) pkg.Context {
	c.Spider.WithStatus(status)
	return c
}
func (c *Context) GetSpiderStartTime() time.Time {
	return c.Spider.GetStartTime()
}
func (c *Context) WithSpiderStartTime(startTime time.Time) pkg.Context {
	c.Spider.WithStartTime(startTime)
	return c
}
func (c *Context) GetSpiderStopTime() time.Time {
	return c.Spider.GetStopTime()
}
func (c *Context) WithSpiderStopTime(stopTime time.Time) pkg.Context {
	c.Spider.WithStopTime(stopTime)
	return c
}

func (c *Context) GetSchedule() pkg.ContextSchedule {
	return c.Schedule
}
func (c *Context) GetScheduleId() string {
	return c.Schedule.GetId()
}
func (c *Context) WithScheduleId(id string) pkg.Context {
	c.Schedule.WithId(id)
	return c
}
func (c *Context) GetScheduleStatus() pkg.ScheduleStatus {
	return c.Schedule.GetStatus()
}
func (c *Context) WithScheduleStatus(status pkg.ScheduleStatus) pkg.Context {
	c.Schedule.WithStatus(status)
	return c
}
func (c *Context) GetScheduleEnable() bool {
	return c.Schedule.GetEnable()
}
func (c *Context) WithScheduleEnable(enable bool) pkg.Context {
	c.Schedule.WithEnable(enable)
	return c
}
func (c *Context) GetScheduleStartTime() time.Time {
	return c.Schedule.GetStartTime()
}
func (c *Context) WithScheduleStartTime(startTime time.Time) pkg.Context {
	c.Schedule.WithStartTime(startTime)
	return c
}
func (c *Context) GetScheduleStopTime() time.Time {
	return c.Schedule.GetStopTime()
}
func (c *Context) WithScheduleStopTime(stopTime time.Time) pkg.Context {
	c.Schedule.WithStopTime(stopTime)
	return c
}

func NewContext() pkg.Context {
	c := new(Context)
	c.Crawler = new(Crawler)
	c.Spider = new(Spider)
	c.Schedule = new(Schedule)
	return c
}
