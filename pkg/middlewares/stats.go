package middlewares

import (
	"context"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"net/http"
	"reflect"
	"sort"
	"time"
)

type StatsMiddleware struct {
	pkg.UnimplementedMiddleware
	logger       pkg.Logger
	interval     time.Duration
	timer        *time.Timer
	chanStop     chan struct{}
	statusOK     int
	statusIgnore int
	statusErr    int
	crawler      pkg.Crawler
	stats        pkg.Stats
}

func (m *StatsMiddleware) Start(_ context.Context, crawler pkg.Crawler) (err error) {
	m.stats = crawler.GetStats()
	m.chanStop = make(chan struct{})
	m.timer = time.NewTimer(m.interval)
	go m.log()
	return
}

func (m *StatsMiddleware) ProcessRequest(_ context.Context, request *pkg.Request) (err error) {
	m.stats.IncRequestTotal()
	return
}

func (m *StatsMiddleware) ProcessResponse(_ context.Context, response *pkg.Response) (err error) {
	if response == nil {
		m.stats.IncStatusErr()
	} else {
		if response.Response != nil && response.StatusCode == http.StatusOK {
			m.stats.IncStatusOk()
		} else {
			m.stats.IncStatusErr()
		}
	}
	return
}

func (m *StatsMiddleware) SpiderStop(_ context.Context) (err error) {
	m.timer.Stop()
	m.chanStop <- struct{}{}
	m.logger.Debug(utils.JsonStr(m.stats))
	kv := make(map[string]uint32)
	getKV(reflect.ValueOf(m.stats).Elem(), kv)
	var sl []any
	sl = append(sl, m.crawler.GetSpider().GetName())
	keys := make([]string, 0)
	for k := range kv {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		sl = append(sl, fmt.Sprintf("%s:", k), kv[k])
	}
	m.logger.Info(sl...)
	return
}

func (m *StatsMiddleware) log() {
	for {
		m.timer.Reset(m.interval)
		select {
		case <-m.timer.C:
			m.logger.Debug(utils.JsonStr(m.stats))
			kv := make(map[string]uint32)
			getKV(reflect.ValueOf(m.stats).Elem(), kv)
			var sl []any
			sl = append(sl, m.crawler.GetSpider().GetName())
			keys := make([]string, 0)
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

func getKV(v reflect.Value, m map[string]uint32) {
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := t.Field(i).Name

		switch field.Kind() {
		case reflect.Ptr:
			if field.IsNil() {
				continue
			}
			getKV(reflect.Indirect(field), m)
		case reflect.Struct:
			getKV(field, m)
		default:
			m[fieldName] = field.Interface().(uint32)
		}
	}
}

func (m *StatsMiddleware) FromCrawler(crawler pkg.Crawler) pkg.Middleware {
	if m == nil {
		return new(StatsMiddleware).FromCrawler(crawler)
	}

	m.crawler = crawler
	m.logger = crawler.GetLogger()
	m.interval = time.Minute
	return m
}
