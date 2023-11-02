package context

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"time"
)

type Context struct {
	Crawler  pkg.ContextCrawler  `json:"crawler,omitempty"`
	Spider   pkg.ContextSpider   `json:"spider,omitempty"`
	Schedule pkg.ContextSchedule `json:"schedule,omitempty"`
	Task     pkg.ContextTask     `json:"task,omitempty"`
	Request  pkg.ContextRequest  `json:"request,omitempty"`
	Item     pkg.ContextItem     `json:"item,omitempty"`
}

func (c *Context) GetContext() pkg.Context {
	return c
}

func (c *Context) GetCrawler() pkg.ContextCrawler {
	return c.Crawler
}
func (c *Context) WithCrawler(crawler pkg.ContextCrawler) pkg.Context {
	c.Crawler = crawler
	return c
}
func (c *Context) GetCrawlerContext() context.Context {
	return c.Crawler.GetContext()
}
func (c *Context) WithCrawlerContext(ctx context.Context) pkg.Context {
	c.Crawler.WithContext(ctx)
	return c
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
func (c *Context) WithSpider(spider pkg.ContextSpider) pkg.Context {
	c.Spider = spider
	return c
}
func (c *Context) GetSpiderContext() context.Context {
	return c.Spider.GetContext()
}
func (c *Context) WithSpiderContext(ctx context.Context) pkg.Context {
	c.Spider.WithContext(ctx)
	return c
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
func (c *Context) WithSchedule(schedule pkg.ContextSchedule) pkg.Context {
	c.Schedule = schedule
	return c
}
func (c *Context) GetScheduleContext() context.Context {
	return c.Schedule.GetContext()
}
func (c *Context) WithScheduleContext(ctx context.Context) pkg.Context {
	c.Schedule.WithContext(ctx)
	return c
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
func (c *Context) GetScheduleEnable() bool {
	return c.Schedule.GetEnable()
}
func (c *Context) WithScheduleEnable(enable bool) pkg.Context {
	c.Schedule.WithEnable(enable)
	return c
}
func (c *Context) GetScheduleFunc() string {
	return c.Schedule.GetFunc()
}
func (c *Context) WithScheduleFunc(fn string) pkg.Context {
	c.Schedule.WithFunc(fn)
	return c
}
func (c *Context) GetScheduleArgs() string {
	return c.Schedule.GetArgs()
}
func (c *Context) WithScheduleArgs(args string) pkg.Context {
	c.Schedule.WithArgs(args)
	return c
}
func (c *Context) GetScheduleMode() pkg.ScheduleMode {
	return c.Schedule.GetMode()
}
func (c *Context) WithScheduleMode(mode pkg.ScheduleMode) pkg.Context {
	c.Schedule.WithMode(mode)
	return c
}
func (c *Context) GetScheduleSpec() string {
	return c.Schedule.GetSpec()
}
func (c *Context) WithScheduleSpec(spec string) pkg.Context {
	c.Schedule.WithSpec(spec)
	return c
}
func (c *Context) GetScheduleOnlyOneTask() bool {
	return c.Schedule.GetOnlyOneTask()
}
func (c *Context) WithScheduleOnlyOneTask(onlyOneTask bool) pkg.Context {
	c.Schedule.WithOnlyOneTask(onlyOneTask)
	return c
}

func (c *Context) GetTask() pkg.ContextTask {
	return c.Task
}
func (c *Context) WithTask(task pkg.ContextTask) pkg.Context {
	c.Task = task
	return c
}
func (c *Context) GetTaskContext() context.Context {
	return c.Task.GetContext()
}
func (c *Context) WithTaskContext(ctx context.Context) pkg.Context {
	c.Task.WithContext(ctx)
	return c
}
func (c *Context) GetTaskId() string {
	return c.Task.GetId()
}
func (c *Context) WithTaskId(taskId string) pkg.Context {
	c.Task.WithId(taskId)
	return c
}
func (c *Context) GetTaskStatus() pkg.TaskStatus {
	return c.Task.GetStatus()
}
func (c *Context) WithTaskStatus(status pkg.TaskStatus) pkg.Context {
	c.Task.WithStatus(status)
	return c
}
func (c *Context) GetTaskStartTime() time.Time {
	return c.Task.GetStartTime()
}
func (c *Context) WithTaskStartTime(startTime time.Time) pkg.Context {
	c.Task.WithStartTime(startTime)
	return c
}
func (c *Context) GetTaskStopTime() time.Time {
	return c.Task.GetStopTime()
}
func (c *Context) WithTaskStopTime(stopTime time.Time) pkg.Context {
	c.Task.WithStopTime(stopTime)
	return c
}
func (c *Context) GetTaskDeadline() time.Time {
	return c.Task.GetDeadline()
}
func (c *Context) WithTaskDeadline(deadline time.Time) pkg.Context {
	c.Task.WithDeadline(deadline)
	return c
}

func (c *Context) GetRequest() pkg.ContextRequest {
	return c.Request
}
func (c *Context) WithRequest(request pkg.ContextRequest) pkg.Context {
	c.Request = request
	return c
}
func (c *Context) GetRequestId() string {
	return c.Request.GetId()
}
func (c *Context) WithRequestId(taskId string) pkg.Context {
	c.Request.WithId(taskId)
	return c
}
func (c *Context) GetRequestContext() context.Context {
	return c.Request.GetContext()
}
func (c *Context) WithRequestContext(ctx context.Context) pkg.Context {
	c.Request.WithContext(ctx)
	return c
}
func (c *Context) GetRequestStatus() pkg.RequestStatus {
	return c.Request.GetStatus()
}
func (c *Context) WithRequestStatus(status pkg.RequestStatus) pkg.Context {
	c.Request.WithStatus(status)
	return c
}
func (c *Context) GetRequestStartTime() time.Time {
	return c.Request.GetStartTime()
}
func (c *Context) WithRequestStartTime(startTime time.Time) pkg.Context {
	c.Request.WithStartTime(startTime)
	return c
}
func (c *Context) GetRequestStopTime() time.Time {
	return c.Request.GetStopTime()
}
func (c *Context) WithRequestStopTime(stopTime time.Time) pkg.Context {
	c.Request.WithStopTime(stopTime)
	return c
}
func (c *Context) GetRequestDeadline() time.Time {
	return c.Request.GetDeadline()
}
func (c *Context) WithRequestDeadline(deadline time.Time) pkg.Context {
	c.Request.WithDeadline(deadline)
	return c
}
func (c *Context) GetRequestCookies() map[string]string {
	return c.Request.GetCookies()
}
func (c *Context) WithRequestCookies(cookies map[string]string) pkg.Context {
	c.Request.WithCookies(cookies)
	return c
}
func (c *Context) GetRequestReferrer() string {
	return c.Request.GetReferrer()
}
func (c *Context) WithRequestReferrer(referrer string) pkg.Context {
	c.Request.WithReferrer(referrer)
	return c
}

func (c *Context) GetItem() pkg.ContextItem {
	return c.Item
}
func (c *Context) WithItem(request pkg.ContextItem) pkg.Context {
	c.Item = request
	return c
}
func (c *Context) GetItemId() string {
	return c.Item.GetId()
}
func (c *Context) WithItemId(taskId string) pkg.Context {
	c.Item.WithId(taskId)
	return c
}
func (c *Context) GetItemContext() context.Context {
	return c.Item.GetContext()
}
func (c *Context) WithItemContext(ctx context.Context) pkg.Context {
	c.Item.WithContext(ctx)
	return c
}
func (c *Context) GetItemStatus() pkg.ItemStatus {
	return c.Item.GetStatus()
}
func (c *Context) WithItemStatus(status pkg.ItemStatus) pkg.Context {
	c.Item.WithStatus(status)
	return c
}
func (c *Context) GetItemStartTime() time.Time {
	return c.Item.GetStartTime()
}
func (c *Context) WithItemStartTime(startTime time.Time) pkg.Context {
	c.Item.WithStartTime(startTime)
	return c
}
func (c *Context) GetItemStopTime() time.Time {
	return c.Item.GetStopTime()
}
func (c *Context) WithItemStopTime(stopTime time.Time) pkg.Context {
	c.Item.WithStopTime(stopTime)
	return c
}
