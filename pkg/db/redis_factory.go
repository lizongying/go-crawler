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
	config sync.Map
	logger pkg.Logger

	clients sync.Map
}

func (f *RedisFactory) GetClient(name string) (rdb *redis.Client, err error) {
	if v, ok := f.clients.Load(name); ok {
		return v.(*redis.Client), nil
	}

	c, ok := f.config.Load(name)
	if !ok {
		return nil, fmt.Errorf("redis config %s not found", name)
	}

	conf := c.(pkg.Redis)

	addr := conf.Addr
	if addr == "" {
		err = fmt.Errorf("redis addr is empty")
		return
	}

	option := &redis.Options{
		Addr: addr,
	}
	password := conf.Password
	if password != "" {
		option.Password = password
	}
	db := conf.Db
	if db != -1 {
		option.DB = db
	}

	rdb = redis.NewClient(option)
	if err = rdb.Ping(context.Background()).Err(); err != nil {
		f.logger.Error(err)
		return
	}

	actual, loaded := f.clients.LoadOrStore(name, rdb)
	if loaded {
		_ = rdb.Close()
	}

	return actual.(*redis.Client), nil
}

func (f *RedisFactory) Close(_ context.Context) error {
	f.clients.Range(func(key, value interface{}) bool {
		_ = value.(*redis.Client).Close()
		return true
	})
	return nil
}

func NewRedisFactory(config *config.Config, logger pkg.Logger, lc fx.Lifecycle) (redisFactory *RedisFactory, err error) {
	redisFactory = &RedisFactory{
		logger: logger,
	}
	for _, i := range config.RedisList {
		redisFactory.config.Store(i.Name, i)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) (err error) {
			return redisFactory.Close(ctx)
		},
	})

	return
}
