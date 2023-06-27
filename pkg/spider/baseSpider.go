package spider

import (
	"context"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"log"
	"reflect"
	"runtime"
)

var (
	buildBranch string
	buildCommit string
	buildTime   string
)

func init() {
	info := fmt.Sprintf("Branch: %s, Commit: %s, Time: %s, GOVersion: %s, OS: %s, ARCH: %s", buildBranch, buildCommit, buildTime, runtime.Version(), runtime.GOOS, runtime.GOARCH)
	log.Println(info)
}

type BaseSpider struct {
	pkg.Crawler
	name      string
	callbacks map[string]pkg.Callback
	errbacks  map[string]pkg.Errback
	logger    pkg.Logger
}

func (s *BaseSpider) GetName() string {
	return s.name
}
func (s *BaseSpider) SetName(name string) {
	s.name = name
}
func (s *BaseSpider) register() {
	s.callbacks = make(map[string]pkg.Callback)
	s.errbacks = make(map[string]pkg.Errback)
	rt := reflect.TypeOf(s)
	rv := reflect.ValueOf(s)
	l := rt.NumMethod()
	for i := 0; i < l; i++ {
		name := rt.Method(i).Name
		callback, ok := rv.Method(i).Interface().(pkg.Callback)
		if ok {
			s.callbacks[name] = callback
		}
		errback, ok := rv.Method(i).Interface().(pkg.Errback)
		if ok {
			s.errbacks[name] = errback
		}
	}
}
func (s *BaseSpider) GetCallbacks() map[string]pkg.Callback {
	return s.callbacks
}
func (s *BaseSpider) GetErrbacks() map[string]pkg.Errback {
	return s.errbacks
}
func (s *BaseSpider) Start(ctx context.Context) (err error) {
	s.register()
	return
}
func (s *BaseSpider) Stop(ctx context.Context) (err error) {
	s.logger.Debug("BaseSpider Wait for stop")
	defer func() {
		s.logger.Info("BaseSpider Stopped")
	}()

	return
}

func NewBaseSpider(crawler pkg.Crawler, logger pkg.Logger) (spider pkg.Spider, err error) {
	spider = &BaseSpider{
		Crawler: crawler,
		logger:  logger,
	}

	return
}
