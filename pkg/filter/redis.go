package filter

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
	spider pkg.Spider
	logger pkg.Logger
}

func (f *RedisFilter) SpiderOpened() {
	f.key = fmt.Sprintf("crawler:%s:filter", f.spider.Name())
	f.logger.Debug("filter key", f.key)
	ctx := context.Background()
	if CleanFilterOnSpiderOpened {
		err := f.rdb.Del(ctx, f.key).Err()
		if err != nil {
			f.logger.Error(err)
		}
	}
}

func (f *RedisFilter) IsExist(c pkg.Context, uniqueKey any) (ok bool, err error) {
	ctx := c.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	ok, err = f.rdb.SIsMember(ctx, f.key, uniqueKey).Result()
	return
}

func (f *RedisFilter) Store(c pkg.Context, uniqueKey any) (err error) {
	ctx := c.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = f.rdb.SAdd(ctx, f.key, uniqueKey).Err()
	return
}

func (f *RedisFilter) Clean(ctx context.Context) (err error) {
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

	spider.GetSignal().RegisterSpiderOpened(f.SpiderOpened)

	f.spider = spider
	f.rdb = spider.GetCrawler().GetRedis()
	f.logger = spider.GetLogger()
	return f
}
