package context

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"time"
)

type Context struct {
	Crawler pkg.ContextCrawler `json:"crawler,omitempty"`
	Spider  pkg.ContextSpider  `json:"spider,omitempty"`
	Job     pkg.ContextJob     `json:"schedule,omitempty"`
	Task    pkg.ContextTask    `json:"task,omitempty"`
	Request pkg.ContextRequest `json:"request,omitempty"`
	Item    pkg.ContextItem    `json:"item,omitempty"`
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
func (c *Context) GetCrawlerStopTime() time.Time {
	return c.Crawler.GetStopTime()
}
func (c *Context) GetCrawlerUpdateTime() time.Time {
	return c.Crawler.GetUpdateTime()
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
func (c *Context) GetSpiderStopTime() time.Time {
	return c.Spider.GetStopTime()
}
func (c *Context) GetSpiderUpdateTime() time.Time {
	return c.Spider.GetUpdateTime()
}

func (c *Context) GetJob() pkg.ContextJob {
	return c.Job
}
func (c *Context) WithJob(schedule pkg.ContextJob) pkg.Context {
	c.Job = schedule
	return c
}
func (c *Context) GetJobContext() context.Context {
	return c.Job.GetContext()
}
func (c *Context) WithJobContext(ctx context.Context) pkg.Context {
	c.Job.WithContext(ctx)
	return c
}
func (c *Context) GetJobId() string {
	return c.Job.GetId()
}
func (c *Context) WithJobId(id string) pkg.Context {
	c.Job.WithId(id)
	return c
}
func (c *Context) GetJobSubId() uint64 {
	return c.Job.GetSubId()
}
func (c *Context) WithJobSubId(id uint64) pkg.Context {
	c.Job.WithSubId(id)
	return c
}
func (c *Context) GetJobStatus() pkg.JobStatus {
	return c.Job.GetStatus()
}
func (c *Context) WithJobStatus(status pkg.JobStatus) pkg.Context {
	c.Job.WithStatus(status)
	return c
}
func (c *Context) GetJobStartTime() time.Time {
	return c.Job.GetStartTime()
}
func (c *Context) GetJobStopTime() time.Time {
	return c.Job.GetStopTime()
}
func (c *Context) GetJobUpdateTime() time.Time {
	return c.Job.GetUpdateTime()
}
func (c *Context) GetJobEnable() bool {
	return c.Job.GetEnable()
}
func (c *Context) WithJobEnable(enable bool) pkg.Context {
	c.Job.WithEnable(enable)
	return c
}
func (c *Context) GetJobFunc() string {
	return c.Job.GetFunc()
}
func (c *Context) WithJobFunc(fn string) pkg.Context {
	c.Job.WithFunc(fn)
	return c
}
func (c *Context) GetJobArgs() string {
	return c.Job.GetArgs()
}
func (c *Context) WithJobArgs(args string) pkg.Context {
	c.Job.WithArgs(args)
	return c
}
func (c *Context) GetJobMode() pkg.JobMode {
	return c.Job.GetMode()
}
func (c *Context) WithJobMode(mode pkg.JobMode) pkg.Context {
	c.Job.WithMode(mode)
	return c
}
func (c *Context) GetJobSpec() string {
	return c.Job.GetSpec()
}
func (c *Context) WithJobSpec(spec string) pkg.Context {
	c.Job.WithSpec(spec)
	return c
}
func (c *Context) GetJobOnlyOneTask() bool {
	return c.Job.GetOnlyOneTask()
}
func (c *Context) WithJobOnlyOneTask(onlyOneTask bool) pkg.Context {
	c.Job.WithOnlyOneTask(onlyOneTask)
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
