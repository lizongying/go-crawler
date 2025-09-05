package db

import (
	"context"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"sync"
)

type RedisFactory struct {
	config *config.Config
	logger pkg.Logger

	once sync.Once
	rdb  *redis.Client
	err  error
}

func NewRedisFactory(config *config.Config, logger pkg.Logger, lc fx.Lifecycle) (redisFactory *RedisFactory, err error) {
	redisFactory = &RedisFactory{
		config: config,
		logger: logger,
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) (err error) {
			return redisFactory.Close(ctx)
		},
	})

	return
}

func (f *RedisFactory) GetClient() (rdb *redis.Client, err error) {
	f.once.Do(func() {
		addr := f.config.Redis.Example.Addr
		if addr == "" {
			err = fmt.Errorf("redis addr is empty")
			return
		}

		option := &redis.Options{
			Addr: addr,
		}
		password := f.config.Redis.Example.Password
		if password != "" {
			option.Password = password
		}
		db := f.config.Redis.Example.Db
		if db != -1 {
			option.DB = db
		}

		f.rdb = redis.NewClient(option)
		f.err = f.rdb.Ping(context.Background()).Err()
	})

	return f.rdb, f.err
}

func (f *RedisFactory) Close(_ context.Context) error {
	if f.rdb != nil {
		return f.rdb.Close()
	}
	return nil
}
