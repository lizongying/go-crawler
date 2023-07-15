package spider

import (
	"context"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"log"
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
	host      string
	callbacks map[string]pkg.Callback
	errbacks  map[string]pkg.Errback
	logger    pkg.Logger
}

func (s *BaseSpider) GetName() string {
	return s.name
}
func (s *BaseSpider) SetName(name string) pkg.Spider {
	s.name = name
	return s
}
func (s *BaseSpider) GetHost() string {
	return s.host
}
func (s *BaseSpider) SetHost(host string) pkg.Spider {
	s.host = host
	return s
}
func (s *BaseSpider) GetCallbacks() map[string]pkg.Callback {
	return s.callbacks
}
func (s *BaseSpider) GetErrbacks() map[string]pkg.Errback {
	return s.errbacks
}
func (s *BaseSpider) SetCallbacks(callbacks map[string]pkg.Callback) {
	s.callbacks = callbacks
}
func (s *BaseSpider) SetErrbacks(errbacks map[string]pkg.Errback) {
	s.errbacks = errbacks
}
func (s *BaseSpider) Start(ctx context.Context) (err error) {
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
