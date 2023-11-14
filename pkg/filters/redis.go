package filters

import (
	"context"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/redis/go-redis/v9"
)

const CleanFilterOnSpiderOpened = true

type RedisFilter struct {
	key    string
	rdb    *redis.Client
	config pkg.Config
	spider pkg.Spider
	logger pkg.Logger
}

func (f *RedisFilter) SpiderOpened(c pkg.Context) (err error) {
	if c.GetSpiderName() != f.spider.Name() {
		return
	}
	if c.GetSpiderStatus() != pkg.SpiderStatusRunning {
		return
	}

	f.key = fmt.Sprintf("%s:%s:filter", f.config.GetBotName(), f.spider.Name())
	f.logger.Debug("filter key", f.key)
	ctx := context.Background()
	if CleanFilterOnSpiderOpened {
		err = f.rdb.Del(ctx, f.key).Err()
		if err != nil {
			f.logger.Error(err)
		}
	}
	return
}

func (f *RedisFilter) IsExist(c pkg.Context, uniqueKey any) (ok bool, err error) {
	ctx := c.GetRequestContext()
	if ctx == nil {
		ctx = context.Background()
	}

	ok, err = f.rdb.SIsMember(ctx, f.key, uniqueKey).Result()
	return
}

func (f *RedisFilter) Store(c pkg.Context, uniqueKey any) (err error) {
	ctx := c.GetRequestContext()
	if ctx == nil {
		ctx = context.Background()
	}

	err = f.rdb.SAdd(ctx, f.key, uniqueKey).Err()
	return
}

func (f *RedisFilter) Clean(c pkg.Context) (err error) {
	ctx := c.GetRequestContext()
	if ctx == nil {
		ctx = context.Background()
	}

	//err = f.rdb.Del(ctx, f.key).Err()
	return
}

func (f *RedisFilter) FromSpider(spider pkg.Spider) pkg.Filter {
	if f == nil {
		return new(RedisFilter).FromSpider(spider)
	}

	spider.GetCrawler().GetSignal().RegisterSpiderChanged(f.SpiderOpened)

	f.config = spider.GetConfig()
	f.spider = spider
	f.rdb = spider.GetCrawler().GetRedis()
	f.logger = spider.GetLogger()
	return f
}
