package redis

import (
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/scheduler"
	"github.com/redis/go-redis/v9"
)

const defaultRequestMax = 1000 * 1000

type Scheduler struct {
	scheduler.UnimplementedScheduler

	// only redis
	redis         *redis.Client
	requestKey    string
	requestKeySha string

	crawler pkg.Crawler
	spider  pkg.Spider
	config  pkg.Config
	logger  pkg.Logger
	task    pkg.Task

	env                 string
	enablePriorityQueue bool
	batch               uint8
}

func (s *Scheduler) StartScheduler(task pkg.Task) (err error) {
	if s.redis == nil {
		err = errors.New(`redis nil`)
		s.logger.Error(err)
		return
	}

	s.task = task
	s.UnimplementedScheduler.SetTask(s.task)

	ctx := task.GetContext()

	s.initScheduler(ctx)

	go s.HandleItem(ctx)

	go s.handleRequest(ctx)
	return
}

func (s *Scheduler) StopScheduler(_ pkg.Task) (err error) {
	return
}
func (s *Scheduler) initScheduler(ctx pkg.Context) {
	s.requestKey = fmt.Sprintf("%s:%s:request", s.config.GetBotName(), s.spider.Name())
	if s.enablePriorityQueue {
		s.requestKey = fmt.Sprintf("%s:%s:request:priority", s.config.GetBotName(), s.spider.Name())
		script := `
local r = redis.call("ZRANGEBYSCORE", KEYS[1], 0, 2147483647, "LIMIT", 0, ARGV[1])
for _, v in ipairs(r) do
	redis.call("ZINCRBY", KEYS[1], -2147483648, v)
end
return r
`
		r, err := s.redis.Do(ctx.GetTask().GetContext(), "SCRIPT", "LOAD", script).Result()
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
		err := s.redis.Del(ctx.GetTask().GetContext(), s.requestKey).Err()
		if err != nil {
			s.logger.Error(err)
			return
		}
	}

}
func (s *Scheduler) FromSpider(spider pkg.Spider) (scheduler pkg.Scheduler, err error) {
	if s == nil {
		return new(Scheduler).FromSpider(spider)
	}

	s.crawler = spider.GetCrawler()
	s.UnimplementedScheduler.SetCrawler(s.crawler)
	s.spider = spider
	s.UnimplementedScheduler.SetSpider(spider)
	s.config = spider.GetConfig()
	s.logger = spider.GetLogger()
	s.UnimplementedScheduler.SetLogger(s.logger)
	s.UnimplementedScheduler.Init()

	s.redis, err = s.crawler.GetRedis(s.config.GetRedis())
	if err != nil {
		s.logger.Error(err)
		return
	}

	s.env = s.config.GetEnv()
	s.enablePriorityQueue = s.config.GetEnablePriorityQueue()
	s.batch = 1

	return s, nil
}
