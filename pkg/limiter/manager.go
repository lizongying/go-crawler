package limiter

import (
	"github.com/go-redis/redis_rate/v10"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/db"
	"sync"
)

type Manager struct {
	redisFactory *db.RedisFactory
	limiters     sync.Map
}

func (m *Manager) GetLimiter(key string) (limiter pkg.Limiter, ok bool) {
	if v, o := m.limiters.Load(key); o {
		limiter = v.(pkg.Limiter)
		ok = true
	}
	return
}

func (m *Manager) SetLimiter(limiterType pkg.LimitType, key string, ratePerHour int, concurrency int) (limiter pkg.Limiter) {
	if key == "" {
		key = "*"
	}

	if concurrency < 1 {
		concurrency = 1
	}

	val, ok := m.limiters.Load(key)
	if ok {
		limiter = val.(pkg.Limiter)
		limiter.SetBurst(concurrency)
		limiter.SetRatePerHour(ratePerHour)
		return
	}

	switch limiterType {
	case pkg.LimitSingle:
		limiter = NewSingleLimiter(key, ratePerHour, concurrency)
	case pkg.LimitCluster:
		if rdb, err := m.redisFactory.GetClient(); err == nil {
			limiter = NewClusterLimiter(redis_rate.NewLimiter(rdb), key, ratePerHour, concurrency)
		}
	default:
		limiter = NewSingleLimiter(key, ratePerHour, concurrency)
	}

	m.limiters.Store(key, limiter)
	return
}

func NewManager(redisFactory *db.RedisFactory) (*Manager, error) {
	return &Manager{
		redisFactory: redisFactory,
	}, nil
}
