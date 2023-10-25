package signals

import (
	"github.com/lizongying/go-crawler/pkg"
	"sync"
)

type Signal struct {
	logger               pkg.Logger
	crawlerOpened        []pkg.FnCrawlerOpened
	crawlerClosed        []pkg.FnCrawlerClosed
	spiderStarting       map[string][]pkg.FnSpiderStarting
	spiderStarted        map[string][]pkg.FnSpiderStarted
	spiderStopping       map[string][]pkg.FnSpiderStopping
	spiderStopped        map[string][]pkg.FnSpiderStopped
	taskStarted          []pkg.FnTaskStarted
	taskStopped          []pkg.FnTaskStopped
	itemSaved            []pkg.FnItemSaved
	spiderStartingLocker sync.RWMutex
	spiderStartedLocker  sync.RWMutex
	spiderStoppingLocker sync.RWMutex
	spiderStoppedLocker  sync.RWMutex
}

func (s *Signal) RegisterCrawlerOpened(fn pkg.FnCrawlerOpened) {
	s.crawlerOpened = append(s.crawlerOpened, fn)
}
func (s *Signal) RegisterCrawlerClosed(fn pkg.FnCrawlerClosed) {
	s.crawlerClosed = append(s.crawlerClosed, fn)
}
func (s *Signal) RegisterSpiderStarting(spiderName string, fn pkg.FnSpiderStarting) {
	defer s.spiderStartingLocker.Unlock()
	s.spiderStartingLocker.Lock()
	s.spiderStarting[spiderName] = append(s.spiderStarting[spiderName], fn)
}
func (s *Signal) RegisterSpiderStarted(spiderName string, fn pkg.FnSpiderStarted) {
	defer s.spiderStartedLocker.Unlock()
	s.spiderStartedLocker.Lock()
	s.spiderStarted[spiderName] = append(s.spiderStarted[spiderName], fn)
}
func (s *Signal) RegisterSpiderStopping(spiderName string, fn pkg.FnSpiderStopping) {
	defer s.spiderStoppingLocker.Unlock()
	s.spiderStoppingLocker.Lock()
	s.spiderStopping[spiderName] = append(s.spiderStopping[spiderName], fn)
}
func (s *Signal) RegisterSpiderStopped(spiderName string, fn pkg.FnSpiderStopped) {
	defer s.spiderStoppedLocker.Unlock()
	s.spiderStoppedLocker.Lock()
	s.spiderStopped[spiderName] = append(s.spiderStopped[spiderName], fn)
}
func (s *Signal) RegisterTaskStarted(fn pkg.FnTaskStarted) {
	s.taskStarted = append(s.taskStarted, fn)
}
func (s *Signal) RegisterTaskStopped(fn pkg.FnTaskStopped) {
	s.taskStopped = append(s.taskStopped, fn)
}
func (s *Signal) RegisterItemSaved(fn pkg.FnItemSaved) {
	s.itemSaved = append(s.itemSaved, fn)
}
func (s *Signal) CrawlerOpened() {
	for _, v := range s.crawlerOpened {
		v()
	}
}
func (s *Signal) CrawlerClosed() {
	for _, v := range s.crawlerClosed {
		v()
	}
}
func (s *Signal) SpiderStarting(spider pkg.Spider) {
	s.spiderStartingLocker.RLock()
	spiderStarting := s.spiderStarting[spider.Name()]
	s.spiderStartingLocker.RUnlock()
	for _, v := range spiderStarting {
		v(spider)
	}
}
func (s *Signal) SpiderStarted(spider pkg.Spider) {
	s.spiderStartedLocker.RLock()
	spiderStarted := s.spiderStarted[spider.Name()]
	s.spiderStartedLocker.RUnlock()
	for _, v := range spiderStarted {
		v(spider)
	}
}
func (s *Signal) SpiderStopping(spider pkg.Spider) {
	s.spiderStoppingLocker.RLock()
	spiderStopping := s.spiderStopping[spider.Name()]
	s.spiderStoppingLocker.RUnlock()
	for _, v := range spiderStopping {
		v(spider)
	}
}
func (s *Signal) SpiderStopped(spider pkg.Spider) {
	s.spiderStoppedLocker.RLock()
	spiderStopped := s.spiderStopped[spider.Name()]
	s.spiderStoppedLocker.RUnlock()
	for _, v := range spiderStopped {
		v(spider)
	}
}
func (s *Signal) ItemSaved(itemWithContext pkg.ItemWithContext) {
	for _, v := range s.itemSaved {
		v(itemWithContext)
	}
}
func (s *Signal) FromCrawler(crawler pkg.Crawler) pkg.Signal {
	if s == nil {
		return new(Signal).FromCrawler(crawler)
	}

	s.logger = crawler.GetLogger()
	s.spiderStarting = make(map[string][]pkg.FnSpiderStarting)
	s.spiderStarted = make(map[string][]pkg.FnSpiderStarted)
	s.spiderStopping = make(map[string][]pkg.FnSpiderStopping)
	s.spiderStopped = make(map[string][]pkg.FnSpiderStopped)

	return s
}
