package middlewares

import (
	"context"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
	"reflect"
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
	m.chanStop = make(chan struct{})
	m.timer = time.NewTimer(m.interval)
	go m.log(spider)
	return
}

func (m *StatsMiddleware) Stop(_ context.Context) (err error) {
	spider := m.GetSpider()
	m.timer.Stop()
	m.chanStop <- struct{}{}

	kv := make(map[string]uint32)
	getKV(reflect.ValueOf(spider.GetStats()).Elem(), kv)

	var sl []any
	sl = append(sl, spider.GetName())
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

func (m *StatsMiddleware) ProcessRequest(_ pkg.Context, _ pkg.Request) (err error) {
	spider := m.GetSpider()
	spider.IncRequestTotal()
	return
}

func (m *StatsMiddleware) ProcessResponse(_ pkg.Context, response pkg.Response) (err error) {
	spider := m.GetSpider()
	if response == nil {
		spider.IncStatusErr()
	} else {
		if response.GetResponse() != nil && response.GetStatusCode() == http.StatusOK {
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
			kv := make(map[string]uint32)
			getKV(reflect.ValueOf(spider.GetStats()).Elem(), kv)
			var sl []any
			sl = append(sl, spider.GetName())
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

func (m *StatsMiddleware) FromSpider(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(StatsMiddleware).FromSpider(spider)
	}

	m.UnimplementedMiddleware.FromSpider(spider)
	m.logger = spider.GetLogger()
	m.interval = time.Minute
	return m
}
