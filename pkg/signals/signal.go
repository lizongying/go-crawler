package signals

import (
	"github.com/lizongying/go-crawler/pkg"
)

type Signal struct {
	logger       pkg.Logger
	spiderOpened []pkg.SignalFn
	spiderClosed []pkg.SignalFn
}

func (s *Signal) RegisterSpiderOpened(fn pkg.SignalFn) {
	s.spiderOpened = append(s.spiderOpened, fn)
}
func (s *Signal) RegisterSpiderClosed(fn pkg.SignalFn) {
	s.spiderClosed = append(s.spiderClosed, fn)
}
func (s *Signal) SpiderOpened() {
	for _, v := range s.spiderOpened {
		v()
	}
}
func (s *Signal) SpiderClosed() {
	for _, v := range s.spiderClosed {
		v()
	}
}

func (s *Signal) FromSpider(spider pkg.Spider) pkg.Signal {
	if s == nil {
		return new(Signal).FromSpider(spider)
	}

	s.logger = spider.GetLogger()
	return s
}
