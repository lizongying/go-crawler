package context

import (
	"github.com/lizongying/go-crawler/pkg"
)

// Context represents the execution context for the crawler.
// It holds references to Crawler, Spider, Job, Task, Request, and Item,
// It contains references to different stages of the crawling process.
//
// Fields:
//   - Crawler: The global crawler context (pkg.ContextCrawler)
//   - Spider : The spider context (pkg.ContextSpider)
//   - Job    : The job context (pkg.ContextJob)
//   - Task   : The task context (pkg.ContextTask)
//   - Request: The request context (pkg.ContextRequest)
//   - Item   : The item context (pkg.ContextItem)
//
// Each field can be set, retrieved, and cloned individually using
// the corresponding Get/With/Clone methods.
type Context struct {
	Crawler pkg.ContextCrawler `json:"crawler,omitempty"` // The global crawler context
	Spider  pkg.ContextSpider  `json:"spider,omitempty"`  // The current spider context
	Job     pkg.ContextJob     `json:"job,omitempty"`     // The job context
	Task    pkg.ContextTask    `json:"task,omitempty"`    // The task context
	Request pkg.ContextRequest `json:"request,omitempty"` // The request context
	Item    pkg.ContextItem    `json:"item,omitempty"`    // The item context
}

// GetContext returns the current Context itself.
func (c *Context) GetContext() pkg.Context {
	return c
}

// CloneContext creates a new empty Context instance.
func (c *Context) CloneContext() pkg.Context {
	return new(Context)
}

// ----------------- Crawler -----------------

// GetCrawler returns the crawler context.
func (c *Context) GetCrawler() pkg.ContextCrawler {
	return c.Crawler
}

// WithCrawler sets the crawler context and returns the updated Context.
func (c *Context) WithCrawler(crawler pkg.ContextCrawler) pkg.Context {
	c.Crawler = crawler
	return c
}

// CloneCrawler creates a new Context with the current crawler context.
func (c *Context) CloneCrawler() pkg.Context {
	return new(Context).
		WithCrawler(c.GetCrawler())
}

// ----------------- Spider -----------------

// GetSpider returns the spider context.
func (c *Context) GetSpider() pkg.ContextSpider {
	return c.Spider
}

// WithSpider sets the spider context and returns the updated Context.
func (c *Context) WithSpider(spider pkg.ContextSpider) pkg.Context {
	c.Spider = spider
	return c
}

// CloneSpider creates a new Context with the current crawler and spider contexts.
func (c *Context) CloneSpider() pkg.Context {
	return new(Context).
		WithCrawler(c.GetCrawler()).
		WithSpider(c.GetSpider())
}

// ----------------- Job -----------------

// GetJob returns the job context.
func (c *Context) GetJob() pkg.ContextJob {
	return c.Job
}

// WithJob sets the job context and returns the updated Context.
func (c *Context) WithJob(schedule pkg.ContextJob) pkg.Context {
	c.Job = schedule
	return c
}

// CloneJob creates a new Context with the current crawler, spider, and job contexts.
func (c *Context) CloneJob() pkg.Context {
	return new(Context).
		WithCrawler(c.GetCrawler()).
		WithSpider(c.GetSpider()).
		WithJob(c.GetJob())
}

// ----------------- Task -----------------

// GetTask returns the task context.
func (c *Context) GetTask() pkg.ContextTask {
	return c.Task
}

// WithTask sets the task context and returns the updated Context.
func (c *Context) WithTask(task pkg.ContextTask) pkg.Context {
	c.Task = task
	return c
}

// CloneTask creates a new Context with the current crawler, spider, job, and task contexts.
func (c *Context) CloneTask() pkg.Context {
	return new(Context).
		WithCrawler(c.GetCrawler()).
		WithSpider(c.GetSpider()).
		WithJob(c.GetJob()).
		WithTask(c.GetTask())
}

// ----------------- Request -----------------

// GetRequest returns the request context.
func (c *Context) GetRequest() pkg.ContextRequest {
	return c.Request
}

// WithRequest sets the request context and returns the updated Context.
func (c *Context) WithRequest(request pkg.ContextRequest) pkg.Context {
	c.Request = request
	return c
}

// CloneRequest creates a new Context with the current crawler, spider, job, task, and request contexts.
func (c *Context) CloneRequest() pkg.Context {
	return new(Context).
		WithCrawler(c.GetCrawler()).
		WithSpider(c.GetSpider()).
		WithJob(c.GetJob()).
		WithTask(c.GetTask()).
		WithRequest(c.GetRequest())
}

// ----------------- Item -----------------

// GetItem returns the item context.
func (c *Context) GetItem() pkg.ContextItem {
	return c.Item
}

// WithItem sets the item context and returns the updated Context.
func (c *Context) WithItem(request pkg.ContextItem) pkg.Context {
	c.Item = request
	return c
}

// CloneItem creates a new Context with the current crawler, spider, job, task, request, and item contexts.
func (c *Context) CloneItem() pkg.Context {
	return new(Context).
		WithCrawler(c.GetCrawler()).
		WithSpider(c.GetSpider()).
		WithJob(c.GetJob()).
		WithTask(c.GetTask()).
		WithRequest(c.GetRequest()).
		WithItem(c.GetItem())
}
