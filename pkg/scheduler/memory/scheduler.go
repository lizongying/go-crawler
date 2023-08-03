package memory

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/downloader"
	"github.com/lizongying/go-crawler/pkg/exporter"
	"golang.org/x/time/rate"
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

	concurrency uint8
	interval    time.Duration

	crawler      pkg.Crawler
	logger       pkg.Logger
	stateRequest *pkg.State
	stateItem    *pkg.State
	stateMethod  *pkg.State
	couldStop    chan struct{}
}

func (s *Scheduler) GetInterval() time.Duration {
	return s.interval
}
func (s *Scheduler) SetInterval(interval time.Duration) {
	s.interval = interval
}
func (s *Scheduler) StartScheduler(ctx context.Context) (err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	s.logger.Info("middlewares", s.GetMiddlewareNames())
	s.logger.Info("pipelines", s.GetPipelineNames())

	for _, v := range s.GetPipelines() {
		e := v.Start(ctx, s.GetSpider())
		if errors.Is(e, pkg.BreakErr) {
			s.logger.Debug("pipeline break", v.GetName())
			break
		}
	}
	for _, v := range s.GetMiddlewares() {
		e := v.Start(ctx, s.GetSpider())
		if errors.Is(e, pkg.BreakErr) {
			s.logger.Debug("middlewares break", v.GetName())
			break
		}
	}

	s.itemTimer = time.NewTimer(s.GetItemDelay())
	if s.GetItemConcurrency() < 1 {
		s.SetItemConcurrencyRaw(1)
	}
	s.SetItemConcurrencyNew(s.GetItemConcurrency())
	s.itemConcurrencyChan = make(chan struct{}, s.GetItemConcurrency())
	for i := 0; i < s.GetItemConcurrency(); i++ {
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

	s.logger.Debug("Scheduler wait for stop")
	states := pkg.NewMultiState(s.stateRequest, s.stateItem, s.stateMethod)
	states.RegisterSetAndZeroFn(func() {
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
		s.couldStop <- struct{}{}
	})
	<-s.couldStop
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
	s.concurrency = config.GetRequestConcurrency()
	s.interval = time.Millisecond * time.Duration(int(config.GetRequestInterval()))
	s.requestChan = make(chan pkg.Request, defaultRequestMax)
	s.itemChan = make(chan pkg.Item, defaultChanItemMax)
	s.couldStop = make(chan struct{})

	s.SetDownloader(new(downloader.Downloader).FromSpider(spider))
	s.SetExporter(new(exporter.Exporter).FromSpider(spider))
	s.crawler = crawler
	s.logger = crawler.GetLogger()
	s.stateRequest = pkg.NewState()
	s.stateItem = pkg.NewState()
	s.stateMethod = pkg.NewState()
	s.stateItem.Set()
	s.stateMethod.Set()

	return s
}
