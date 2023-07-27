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
	"time"
)

const defaultRequestMax = 1000 * 1000
const defaultChanItemMax = 1000 * 1000

type Scheduler struct {
	pkg.UnimplementedScheduler
	itemConcurrencyChan chan struct{}
	itemTimer           *time.Timer
	itemChan            chan pkg.Item
	requestKey          string
	requestKeySha       string

	concurrency uint8
	interval    time.Duration

	crawler      pkg.Crawler
	logger       pkg.Logger
	stateRequest *pkg.State
	stateItem    *pkg.State
	stateMethod  *pkg.State
	couldStop    chan struct{}

	redis  *redis.Client
	config pkg.Config

	mode                string
	enablePriorityQueue bool
	batch               uint8
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

func (s *Scheduler) Stop(ctx context.Context) (err error) {
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
func (s *Scheduler) SpiderOpened(spider pkg.Spider) {
	s.requestKey = fmt.Sprintf("%s:%s:request", s.config.GetBotName(), spider.GetName())
	if s.enablePriorityQueue {
		s.requestKey = fmt.Sprintf("%s:%s:request:priority", s.config.GetBotName(), spider.GetName())
		script := `
local r = redis.call("ZRANGEBYSCORE", KEYS[1], 0, 2147483647, "LIMIT", 0, ARGV[1])
for _, v in ipairs(r) do
	redis.call("ZINCRBY", KEYS[1], -2147483648, v)
end
return r
`
		ctx := context.Background()
		r, err := s.redis.Do(ctx, "SCRIPT", "LOAD", script).Result()
		if err != nil {
			s.logger.Error(err)
			return
		}
		var ok bool
		s.requestKeySha, ok = r.(string)
		if !ok {
			s.logger.Error(errors.New("SCRIPT LOAD error"))
			return
		}
		s.logger.Info("request key sha", s.requestKeySha)
	}

	s.logger.Debug("request key", s.requestKey)
	if s.mode == "dev" {
		ctx := context.Background()
		err := s.redis.Del(ctx, s.requestKey).Err()
		if err != nil {
			s.logger.Error(err)
		}
	}

}
func (s *Scheduler) FromCrawler(crawler pkg.Crawler) pkg.Scheduler {
	if s == nil {
		return new(Scheduler).FromCrawler(crawler)
	}

	crawler.GetSignal().RegisterSpiderOpened(s.SpiderOpened)

	config := crawler.GetConfig()
	s.config = config
	s.mode = crawler.GetMode()
	s.enablePriorityQueue = config.GetEnablePriorityQueue()
	s.concurrency = config.GetRequestConcurrency()
	s.interval = time.Millisecond * time.Duration(int(config.GetRequestInterval()))
	s.itemChan = make(chan pkg.Item, defaultChanItemMax)
	s.couldStop = make(chan struct{})

	s.SetDownloader(new(downloader.Downloader).FromCrawler(crawler))
	s.SetExporter(new(exporter.Exporter).FromCrawler(crawler))
	s.crawler = crawler
	s.logger = crawler.GetLogger()
	s.redis = crawler.GetRedis()
	s.batch = 1
	s.stateRequest = pkg.NewState()
	s.stateItem = pkg.NewState()
	s.stateMethod = pkg.NewState()
	s.stateItem.Set()
	s.stateMethod.Set()

	return s
}
