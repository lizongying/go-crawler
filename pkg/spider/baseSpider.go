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
	name   string
	fns    map[string]func(context.Context, *pkg.Response) error
	logger pkg.Logger
}

func (s *BaseSpider) GetName() string {
	return s.name
}
func (s *BaseSpider) SetName(name string) {
	s.name = name
}
func (s *BaseSpider) register() {
	s.fns = make(map[string]func(context.Context, *pkg.Response) error)
	rt := reflect.TypeOf(s)
	rv := reflect.ValueOf(s)
	l := rt.NumMethod()
	for i := 0; i < l; i++ {
		name := rt.Method(i).Name
		fn, ok := rv.Method(i).Interface().(func(context.Context, *pkg.Response) error)
		if ok {
			s.fns[name] = fn
		}
	}
}
func (s *BaseSpider) GetFn(name string) func(context.Context, *pkg.Response) error {
	fn, ok := s.fns[name]
	if ok {
		return fn
	}
	return nil
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
