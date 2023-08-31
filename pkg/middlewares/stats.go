package middlewares

import (
	"context"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
	"sort"
	"time"
)

type StatsMiddleware struct {
	pkg.UnimplementedMiddleware
	logger   pkg.Logger
	interval time.Duration
	timer    *time.Timer
	chanStop chan struct{}
}

func (m *StatsMiddleware) Start(ctx context.Context, spider pkg.Spider) (err error) {
	err = m.UnimplementedMiddleware.Start(ctx, spider)
	m.chanStop = make(chan struct{}, 1)
	m.timer = time.NewTimer(m.interval)
	go m.log(spider)
	return
}

func (m *StatsMiddleware) Stop(_ context.Context) (err error) {
	spider := m.GetSpider()
	m.timer.Stop()
	m.chanStop <- struct{}{}

	var sl []any
	sl = append(sl, spider.Name())
	keys := make([]string, 0)
	kv := spider.GetStats().GetMap()
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

func (m *StatsMiddleware) ProcessRequest(_ pkg.Context, _ pkg.Request) (err error) {
	spider := m.GetSpider()
	spider.IncRequestSuccess()
	return
}

func (m *StatsMiddleware) ProcessResponse(_ pkg.Context, response pkg.Response) (err error) {
	spider := m.GetSpider()
	if response == nil {
		spider.IncStatusErr()
	} else {
		if response.GetResponse() != nil && response.StatusCode() == http.StatusOK {
			spider.IncStatusOk()
		} else {
			spider.IncStatusErr()
		}
	}
	return
}

func (m *StatsMiddleware) log(spider pkg.Spider) {
	for {
		m.timer.Reset(m.interval)
		select {
		case <-m.timer.C:
			var sl []any
			sl = append(sl, spider.Name())
			keys := make([]string, 0)
			kv := spider.GetStats().GetMap()
			for k := range kv {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			for _, k := range keys {
				sl = append(sl, fmt.Sprintf("%s:", k), kv[k])
			}
			m.logger.Info(sl...)
		case <-m.chanStop:
			return
		}
	}
}

func (m *StatsMiddleware) FromSpider(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(StatsMiddleware).FromSpider(spider)
	}

	m.UnimplementedMiddleware.FromSpider(spider)
	m.logger = spider.GetLogger()
	m.interval = time.Minute
	return m
}
