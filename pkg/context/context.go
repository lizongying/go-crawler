package context

import (
	"github.com/lizongying/go-crawler/pkg"
)

type Context struct {
	Crawler pkg.ContextCrawler `json:"crawler,omitempty"`
	Spider  pkg.ContextSpider  `json:"spider,omitempty"`
	Job     pkg.ContextJob     `json:"job,omitempty"`
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

func (c *Context) GetSpider() pkg.ContextSpider {
	return c.Spider
}
func (c *Context) WithSpider(spider pkg.ContextSpider) pkg.Context {
	c.Spider = spider
	return c
}

func (c *Context) GetJob() pkg.ContextJob {
	return c.Job
}
func (c *Context) WithJob(schedule pkg.ContextJob) pkg.Context {
	c.Job = schedule
	return c
}

func (c *Context) GetTask() pkg.ContextTask {
	return c.Task
}
func (c *Context) WithTask(task pkg.ContextTask) pkg.Context {
	c.Task = task
	return c
}

func (c *Context) GetRequest() pkg.ContextRequest {
	return c.Request
}
func (c *Context) WithRequest(request pkg.ContextRequest) pkg.Context {
	c.Request = request
	return c
}

func (c *Context) GetItem() pkg.ContextItem {
	return c.Item
}
func (c *Context) WithItem(request pkg.ContextItem) pkg.Context {
	c.Item = request
	return c
}
