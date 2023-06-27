package signals

import (
	"github.com/lizongying/go-crawler/pkg"
)

type Signal struct {
	spider       pkg.Spider
	spiderOpened []func(spider pkg.Spider)
	spiderClosed []func(spider pkg.Spider)
	logger       pkg.Logger
}

func (s *Signal) SetSpider(spider pkg.Spider) {
	s.spider = spider
}
func (s *Signal) RegisterSpiderOpened(fn func(spider pkg.Spider)) {
	s.spiderOpened = append(s.spiderOpened, fn)
}
func (s *Signal) RegisterSpiderClosed(fn func(spider pkg.Spider)) {
	s.spiderClosed = append(s.spiderClosed, fn)
}
func (s *Signal) SpiderOpened() {
	for _, v := range s.spiderOpened {
		v(s.spider)
	}
}
func (s *Signal) SpiderClosed() {
	for _, v := range s.spiderClosed {
		v(s.spider)
	}
}

func (s *Signal) FromCrawler(crawler pkg.Crawler) pkg.Signal {
	if s == nil {
		return new(Signal).FromCrawler(crawler)
	}

	s.logger = crawler.GetLogger()
	return s
}
