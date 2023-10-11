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
	itemWithContextChan chan pkg.ItemWithContext
	requestKey          string
	requestKeySha       string

	concurrency uint8
	interval    time.Duration

	crawler pkg.Crawler
	logger  pkg.Logger

	redis  *redis.Client
	config pkg.Config

	env                 string
	enablePriorityQueue bool
	batch               uint8
}

func (s *Scheduler) Interval() time.Duration {
	return s.interval
}
func (s *Scheduler) SetInterval(interval time.Duration) {
	s.interval = interval
}
func (s *Scheduler) StartScheduler(ctx context.Context) (err error) {
	if s.redis == nil {
		err = errors.New(`redis nil. please check if "redis_enable: false"`)
		s.logger.Error(err)
		return
	}

	if ctx == nil {
		ctx = context.Background()
	}

	s.initScheduler(ctx)

	s.logger.Info("pipelines", s.PipelineNames())

	for _, v := range s.Pipelines() {
		e := v.Start(ctx, s.Spider())
		if errors.Is(e, pkg.BreakErr) {
			s.logger.Debug("pipeline break", v.Name())
			break
		}
	}
	for _, v := range s.Spider().GetMiddlewares().Middlewares() {
		if err = v.Start(ctx, s.Spider()); err != nil {
			s.logger.Error(err)
			return
		}
		s.logger.Info(v.Name(), "started")
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
func (s *Scheduler) initScheduler(_ context.Context) {
	s.requestKey = fmt.Sprintf("%s:%s:request", s.config.GetBotName(), s.Spider().Name())
	if s.enablePriorityQueue {
		s.requestKey = fmt.Sprintf("%s:%s:request:priority", s.config.GetBotName(), s.Spider().Name())
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
		s.logger.Debug("request key sha", s.requestKeySha)
	}

	s.logger.Debug("request key", s.requestKey)
	if s.env == "dev" {
		ctx := context.Background()
		err := s.redis.Del(ctx, s.requestKey).Err()
		if err != nil {
			s.logger.Error(err)
			return
		}
	}

}
func (s *Scheduler) FromSpider(spider pkg.Spider) pkg.Scheduler {
	if s == nil {
		return new(Scheduler).FromSpider(spider)
	}

	s.UnimplementedScheduler.SetSpider(spider)
	crawler := spider.GetCrawler()
	config := crawler.GetConfig()
	s.config = config
	s.env = spider.GetConfig().GetEnv()
	s.enablePriorityQueue = config.GetEnablePriorityQueue()
	s.concurrency = config.GetRequestConcurrency()
	s.interval = time.Millisecond * time.Duration(int(config.GetRequestInterval()))
	s.itemWithContextChan = make(chan pkg.ItemWithContext, defaultChanItemMax)

	s.SetDownloader(new(downloader.Downloader).FromSpider(spider))
	s.SetExporter(new(exporter.Exporter).FromSpider(spider))
	s.crawler = crawler
	s.logger = crawler.GetLogger()
	s.redis = crawler.GetRedis()
	s.batch = 1

	return s
}
