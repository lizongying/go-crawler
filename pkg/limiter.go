package pkg

import (
	"context"
)

type Limiter interface {
	Burst() (concurrency int)
	SetBurst(concurrency int) Limiter
	RatePerHour() (ratePerHour int)
	SetRatePerHour(ratePerHour int) Limiter
	Wait(ctx context.Context) (err error)
}

type LimitType string

const (
	LimitUnknown LimitType = ""
	LimitSingle  LimitType = "single"
	LimitCluster LimitType = "cluster"
)
