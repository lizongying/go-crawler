package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/downloader"
	"github.com/lizongying/go-crawler/pkg/exporter"
	"github.com/redis/go-redis/v9"
	"golang.org/x/time/rate"
	"sync"
	"time"
)

const defaultRequestMax = 1000 * 1000
const defaultChanItemMax = 1000 * 1000
const defaultMaxRequestActive = 1000

type Scheduler struct {
	itemConcurrency     int
	itemConcurrencyNew  int
	itemConcurrencyChan chan struct{}
	itemDelay           time.Duration
	itemTimer           *time.Timer
	itemChan            chan pkg.Item
	itemActiveChan      chan struct{}
	requestKey          string
	requestActiveChan   chan struct{}
	requestSlots        sync.Map

	concurrency uint8
	interval    time.Duration

	pkg.Downloader
	pkg.Exporter

	crawler pkg.Crawler
	logger  pkg.Logger
	redis   *redis.Client
}

func (s *Scheduler) GetDownloader() pkg.Downloader {
	return s.Downloader
}
func (s *Scheduler) SetDownloader(downloader pkg.Downloader) {
	s.Downloader = downloader
}
func (s *Scheduler) GetExporter() pkg.Exporter {
	return s.Exporter
}
func (s *Scheduler) SetExporter(exporter pkg.Exporter) {
	s.Exporter = exporter
}
func (s *Scheduler) GetInterval() time.Duration {
	return s.interval
}
func (s *Scheduler) SetInterval(interval time.Duration) {
	s.interval = interval
}
func (s *Scheduler) Start(ctx context.Context) (err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	s.logger.Info("middlewares", s.GetMiddlewareNames())
	s.logger.Info("pipelines", s.GetPipelineNames())

	for _, v := range s.GetPipelines() {
		e := v.Start(ctx, s.crawler)
		if errors.Is(e, pkg.BreakErr) {
			s.logger.Debug("pipeline break", v.GetName())
			break
		}
	}
	for _, v := range s.GetMiddlewares() {
		e := v.Start(ctx, s.crawler)
		if errors.Is(e, pkg.BreakErr) {
			s.logger.Debug("middlewares break", v.GetName())
			break
		}
	}

	defer func() {
		for _, v := range s.GetMiddlewares() {
			e := v.Stop(ctx)
			if errors.Is(e, pkg.BreakErr) {
				s.logger.Debug("middlewares break", v.GetName())
				break
			}
		}
		for _, v := range s.GetPipelines() {
			e := v.Stop(ctx)
			if errors.Is(e, pkg.BreakErr) {
				s.logger.Debug("pipeline break", v.GetName())
				break
			}
		}
	}()

	s.itemTimer = time.NewTimer(s.itemDelay)
	if s.itemConcurrency < 1 {
		s.itemConcurrency = 1
	}
	s.itemConcurrencyNew = s.itemConcurrency
	s.itemConcurrencyChan = make(chan struct{}, s.itemConcurrency)
	for i := 0; i < s.itemConcurrency; i++ {
		s.itemConcurrencyChan <- struct{}{}
	}

	slot := "*"
	if _, ok := s.requestSlots.Load(slot); !ok {
		requestSlot := rate.NewLimiter(rate.Every(s.interval/time.Duration(s.concurrency)), int(s.concurrency))
		s.requestSlots.Store(slot, requestSlot)
	}

	go s.handleItem(ctx)

	go s.handleRequest(ctx)

	return
}

func (s *Scheduler) Stop(ctx context.Context) (err error) {
	s.logger.Debug("Scheduler wait for stop")
	defer func() {
		s.logger.Info("Scheduler Stopped")
	}()

	if ctx == nil {
		ctx = context.Background()
	}

	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		if len(s.requestActiveChan) > 0 {
			s.logger.Debug("request is active")
			continue
		}
		if len(s.itemActiveChan) > 0 {
			s.logger.Debug("item is active")
			continue
		}
		break
	}

	return
}
func (s *Scheduler) FromCrawler(crawler pkg.Crawler) pkg.Scheduler {
	if s == nil {
		return new(Scheduler).FromCrawler(crawler)
	}

	config := crawler.GetConfig()
	s.concurrency = config.GetRequestConcurrency()
	s.interval = time.Millisecond * time.Duration(int(config.GetRequestInterval()))
	s.requestKey = fmt.Sprintf("crawler:%s:request", crawler.GetSpider().GetName())
	s.requestActiveChan = make(chan struct{}, defaultRequestMax)
	s.itemChan = make(chan pkg.Item, defaultChanItemMax)
	s.itemActiveChan = make(chan struct{}, defaultChanItemMax)

	s.SetDownloader(new(downloader.Downloader).FromCrawler(crawler))
	s.SetExporter(new(exporter.Exporter).FromCrawler(crawler))
	s.crawler = crawler
	s.logger = crawler.GetLogger()
	s.redis = crawler.GetRedis()

	return s
}
