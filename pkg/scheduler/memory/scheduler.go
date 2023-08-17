package memory

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/downloader"
	"github.com/lizongying/go-crawler/pkg/exporter"
	"golang.org/x/time/rate"
	"sync"
	"time"
)

const defaultRequestMax = 1000 * 1000
const defaultChanItemMax = 1000 * 1000

type Scheduler struct {
	pkg.UnimplementedScheduler
	itemConcurrencyChan chan struct{}
	itemTimer           *time.Timer
	itemChan            chan pkg.Item
	requestChan         chan pkg.Request
	extraChanMap        sync.Map

	concurrency uint8
	interval    time.Duration

	crawler pkg.Crawler
	logger  pkg.Logger
	config  pkg.Config
}

func (s *Scheduler) Interval() time.Duration {
	return s.interval
}
func (s *Scheduler) SetInterval(interval time.Duration) {
	s.interval = interval
}
func (s *Scheduler) StartScheduler(ctx context.Context) (err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	s.logger.Info("middlewares", s.MiddlewareNames())
	s.logger.Info("pipelines", s.PipelineNames())

	for _, v := range s.Pipelines() {
		e := v.Start(ctx, s.Spider())
		if errors.Is(e, pkg.BreakErr) {
			s.logger.Debug("pipeline break", v.Name())
			break
		}
	}
	for _, v := range s.Middlewares() {
		e := v.Start(ctx, s.Spider())
		if errors.Is(e, pkg.BreakErr) {
			s.logger.Debug("middlewares break", v.Name())
			break
		}
	}

	s.itemTimer = time.NewTimer(s.GetItemDelay())
	if s.ItemConcurrency() < 1 {
		s.SetItemConcurrencyRaw(1)
	}
	s.SetItemConcurrencyNew(s.ItemConcurrency())
	s.itemConcurrencyChan = make(chan struct{}, s.ItemConcurrency())
	for i := 0; i < s.ItemConcurrency(); i++ {
		s.itemConcurrencyChan <- struct{}{}
	}

	slot := "*"
	if _, ok := s.RequestSlotLoad(slot); !ok {
		requestSlot := rate.NewLimiter(rate.Every(s.interval/time.Duration(s.concurrency)), int(s.concurrency))
		s.RequestSlotStore(slot, requestSlot)
	}

	go s.handleItem(ctx)

	go s.handleRequest(ctx)

	return
}

func (s *Scheduler) StopScheduler(ctx context.Context) (err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	s.logger.Info("Scheduler Stopped")
	return
}
func (s *Scheduler) FromSpider(spider pkg.Spider) pkg.Scheduler {
	if s == nil {
		return new(Scheduler).FromSpider(spider)
	}

	s.UnimplementedScheduler.SetSpider(spider)
	crawler := spider.GetCrawler()
	config := crawler.GetConfig()
	s.config = config
	s.concurrency = config.GetRequestConcurrency()
	s.interval = time.Millisecond * time.Duration(int(config.GetRequestInterval()))
	s.requestChan = make(chan pkg.Request, defaultRequestMax)
	s.itemChan = make(chan pkg.Item, defaultChanItemMax)

	s.SetDownloader(new(downloader.Downloader).FromSpider(spider))
	s.SetExporter(new(exporter.Exporter).FromSpider(spider))
	s.crawler = crawler
	s.logger = crawler.GetLogger()

	return s
}
