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
	name   string
	logger pkg.Logger
}

func (s *BaseSpider) GetName() string {
	return s.name
}
func (s *BaseSpider) SetName(name string) {
	s.name = name
}

func (s *BaseSpider) Start(ctx context.Context) error {
	return nil
}
func (s *BaseSpider) Stop(ctx context.Context) error {
	return nil
}

func NewBaseSpider(crawler pkg.Crawler, logger pkg.Logger) (spider pkg.Spider, err error) {
	spider = &BaseSpider{
		Crawler: crawler,
		logger:  logger,
	}

	return
}
