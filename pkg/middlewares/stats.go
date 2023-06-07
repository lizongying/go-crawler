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
	spiderInfo   *pkg.SpiderInfo
	stats        pkg.Stats
}

func (m *StatsMiddleware) SpiderStart(_ context.Context, spider pkg.Spider) (err error) {
	m.spiderInfo = spider.GetInfo()
	m.stats = spider.GetStats()
	m.chanStop = make(chan struct{})
	m.timer = time.NewTimer(m.interval)
	go m.log()
	return
}

func (m *StatsMiddleware) ProcessRequest(c *pkg.Context) (err error) {
	m.logger.Debug("enter ProcessRequest")
	defer func() {
		m.logger.Debug("exit ProcessRequest")
	}()

	r := c.Request
	m.logger.DebugF("request: %+v", r)

	m.stats.IncRequestTotal()
	err = c.NextRequest()
	return
}

func (m *StatsMiddleware) ProcessResponse(c *pkg.Context) (err error) {
	response := c.Response
	m.logger.Debug("before response")

	if response == nil {
		m.stats.IncStatusErr()
	} else {
		if response.Response != nil && response.StatusCode == http.StatusOK {
			m.stats.IncStatusOk()
		} else {
			m.stats.IncStatusErr()
		}
	}
	err = c.NextResponse()
	return
}

func (m *StatsMiddleware) SpiderStop(_ context.Context) (err error) {
	m.timer.Stop()
	m.chanStop <- struct{}{}
	m.logger.Debug(utils.JsonStr(m.stats))
	kv := make(map[string]uint32)
	getKV(reflect.ValueOf(m.stats).Elem(), kv)
	var sl []any
	sl = append(sl, m.spiderInfo.Name)
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
			sl = append(sl, m.spiderInfo.Name)
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

func (m *StatsMiddleware) FromCrawler(spider pkg.Spider) pkg.Middleware {
	m.logger = spider.GetLogger()
	m.interval = time.Minute
	return m
}

func NewStatsMiddleware() pkg.Middleware {
	return &StatsMiddleware{}
}
