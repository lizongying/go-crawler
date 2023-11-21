package middlewares

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"net/http"
	"sort"
)

type StatsMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
}

func (m *StatsMiddleware) taskStopped(c pkg.Context) (err error) {
	task := c.GetTask()
	if c.GetSpider().GetId() != m.GetSpider().GetContext().GetSpider().GetId() {
		return
	}
	if !utils.InSlice(task.GetStatus(), []pkg.TaskStatus{
		pkg.TaskStatusSuccess,
		pkg.TaskStatusFailure,
	}) {
		return
	}

	var sl []any
	sl = append(sl, c.GetSpider().GetName(), c.GetTask().GetId())
	keys := make([]string, 0)
	kv := task.GetStats().GetMap()
	for k := range kv {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		sl = append(sl, fmt.Sprintf("%s: %d,", k, kv[k]))
	}
	m.logger.Info(sl...)
	return
}

func (m *StatsMiddleware) ProcessRequest(ctx pkg.Context, _ pkg.Request) (err error) {
	task := ctx.GetTask()
	task.IncRequestSuccess()
	return
}

func (m *StatsMiddleware) ProcessResponse(ctx pkg.Context, response pkg.Response) (err error) {
	task := ctx.GetTask()
	if response == nil {
		task.IncStatusErr()
	} else {
		if response.GetResponse() != nil && response.StatusCode() == http.StatusOK {
			task.IncStatusOk()
		} else {
			task.IncStatusErr()
		}
	}
	return
}

func (m *StatsMiddleware) FromSpider(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(StatsMiddleware).FromSpider(spider)
	}

	m.UnimplementedMiddleware.FromSpider(spider)
	m.logger = spider.GetLogger()
	spider.GetCrawler().GetSignal().RegisterTaskChanged(m.taskStopped)
	return m
}
