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
	logger pkg.Logger
}

func (f *RedisFilter) SpiderOpened(spider pkg.Spider) {
	f.key = fmt.Sprintf("crawler:%s:filter", spider.GetName())
	f.logger.Debug("filter key", f.key)
	ctx := context.Background()
	if CleanFilterOnSpiderOpened {
		err := f.rdb.Del(ctx, f.key).Err()
		if err != nil {
			f.logger.Error(err)
		}
	}
}

func (f *RedisFilter) IsExist(ctx context.Context, uniqueKey any) (ok bool, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	ok, err = f.rdb.SIsMember(ctx, f.key, uniqueKey).Result()
	return
}

func (f *RedisFilter) Store(ctx context.Context, uniqueKey any) (err error) {
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

func (f *RedisFilter) FromCrawler(crawler pkg.Crawler) pkg.Filter {
	if f == nil {
		return new(RedisFilter).FromCrawler(crawler)
	}

	crawler.GetSignal().RegisterSpiderOpened(f.SpiderOpened)

	f.rdb = crawler.GetRedis()
	f.logger = crawler.GetLogger()
	return f
}
