package limiter

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"golang.org/x/time/rate"
	"math"
	"time"
)

type SingleLimiter struct {
	key         string
	limiter     *rate.Limiter
	ratePerHour int
	concurrency int
}

func (r *SingleLimiter) Burst() (concurrency int) {
	concurrency = r.concurrency
	return
}

func (r *SingleLimiter) SetBurst(concurrency int) pkg.Limiter {
	if concurrency <= 0 {
		concurrency = 1
	}
	if concurrency != r.concurrency {
		r.concurrency = concurrency
		r.limiter.SetBurst(concurrency)
	}

	return r
}

func (r *SingleLimiter) RatePerHour() (ratePerHour int) {
	ratePerHour = r.ratePerHour
	return
}

func (r *SingleLimiter) SetRatePerHour(ratePerHour int) pkg.Limiter {
	if ratePerHour <= 0 {
		ratePerHour = math.MaxInt32
	}

	if ratePerHour != r.ratePerHour {
		r.ratePerHour = ratePerHour
		interval := time.Hour / time.Duration(ratePerHour)
		r.limiter.SetLimit(rate.Every(interval))
	}
	return r
}

func (r *SingleLimiter) Key() string {
	return r.key
}

func (r *SingleLimiter) Wait(ctx context.Context) (err error) {
	return r.limiter.Wait(ctx)
}

func NewSingleLimiter(key string, ratePerHour int, concurrency int) *SingleLimiter {
	if ratePerHour <= 0 {
		ratePerHour = math.MaxInt32
	}

	if concurrency < 1 {
		concurrency = 1
	}

	interval := time.Hour / time.Duration(ratePerHour)
	limiter := rate.NewLimiter(rate.Every(interval), concurrency)

	return &SingleLimiter{
		key,
		limiter,
		ratePerHour,
		concurrency,
	}
}
