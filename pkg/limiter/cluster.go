package limiter

import (
	"github.com/go-redis/redis_rate/v10"
	"github.com/lizongying/go-crawler/pkg"
	"golang.org/x/net/context"
	"math"
	"time"
)

type ClusterLimiter struct {
	key         string
	limit       redis_rate.Limit
	limiter     *redis_rate.Limiter
	ratePerHour int
	concurrency int
}

func (r *ClusterLimiter) Burst() (concurrency int) {
	concurrency = r.concurrency
	return
}

func (r *ClusterLimiter) SetBurst(concurrency int) pkg.Limiter {
	if concurrency <= 0 {
		concurrency = 1
	}
	if concurrency != r.concurrency {
		r.concurrency = concurrency
		r.limit.Burst = concurrency
	}
	return r
}

func (r *ClusterLimiter) RatePerHour() (ratePerHour int) {
	ratePerHour = r.ratePerHour
	return
}

func (r *ClusterLimiter) SetRatePerHour(ratePerHour int) pkg.Limiter {
	if ratePerHour <= 0 {
		ratePerHour = math.MaxInt32
	}

	if ratePerHour != r.ratePerHour {
		r.ratePerHour = ratePerHour
		r.limit.Rate = ratePerHour
	}
	return r
}

func (r *ClusterLimiter) Key() string {
	return r.key
}

func (r *ClusterLimiter) Wait(ctx context.Context) (err error) {
	for {
		var res *redis_rate.Result
		res, err = r.limiter.Allow(ctx, r.key, r.limit)
		if err != nil {
			return err
		}

		if res.Allowed > 0 {
			return nil
		}

		waitDuration := res.RetryAfter
		if waitDuration > 0 {
			select {
			case <-time.After(waitDuration):
				continue
			case <-ctx.Done():
				return ctx.Err()
			}
		} else {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				return nil
			}
		}
	}
}

func NewClusterLimiter(limiter *redis_rate.Limiter, key string, ratePerHour int, concurrency int) *ClusterLimiter {
	if ratePerHour <= 0 {
		ratePerHour = math.MaxInt32
	}

	if concurrency < 1 {
		concurrency = 1
	}

	limit := redis_rate.Limit{
		Rate:   ratePerHour,
		Period: time.Hour,
		Burst:  concurrency,
	}

	return &ClusterLimiter{
		key,
		limit,
		limiter,
		ratePerHour,
		concurrency,
	}
}
