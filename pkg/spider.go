package pkg

import (
	"time"
)

type SpiderInfo struct {
	Mode string
	Name string

	Concurrency   int
	Interval      time.Duration
	RetryMaxTimes uint8
	Timeout       time.Duration
	Username      string
	Password      string
}

type Spider interface {
	Crawler
	GetName() string
	SetName(string)
}

type NewSpider func(Spider) (Spider, error)
