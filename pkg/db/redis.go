package db

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

func NewRedis(config *config.Config, logger pkg.Logger, lc fx.Lifecycle) (rdb *redis.Client, err error) {
	if !config.RedisEnable {
		logger.Debug("Redis Disable")
		return
	}

	addr := config.Redis.Example.Addr
	if addr == "" {
		logger.Warn("addr is empty")
		return
	}

	option := &redis.Options{
		Addr: addr,
	}

	password := config.Redis.Example.Password
	if password != "" {
		option.Password = password
	}
	db := config.Redis.Example.Db
	if db != -1 {
		option.DB = db
	}

	rdb = redis.NewClient(option)

	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) (err error) {
			if rdb == nil {
				return
			}
			return
		},
	})
	return
}
