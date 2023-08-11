package pkg

import (
	"context"
	"golang.org/x/time/rate"
	"sync"
	"time"
)

type SchedulerType string

const (
	SchedulerUnknown SchedulerType = ""
	SchedulerMemory  SchedulerType = "memory"
	SchedulerRedis   SchedulerType = "redis"
	SchedulerKafka   SchedulerType = "kafka"
)

type Scheduler interface {
	GetDownloader() Downloader
	SetDownloader(Downloader)
	GetExporter() Exporter
	SetExporter(Exporter)
	SetItemDelay(time.Duration)
	SetItemConcurrency(int)
	SetRequestRate(string, time.Duration, int)
	YieldItem(Context, Item) error
	Request(Context, Request) (Response, error)
	YieldRequest(Context, Request) error
	YieldExtra(Context, any) error
	StartScheduler(context.Context) error
	StopScheduler(context.Context) error
	GetSpider() Spider
	SetSpider(spider Spider)
	Interval() time.Duration
	SetInterval(time.Duration)

	Downloader
	Exporter
}

type UnimplementedScheduler struct {
	itemConcurrency    int
	itemConcurrencyNew int
	itemDelay          time.Duration
	requestSlots       sync.Map

	Downloader
	Exporter

	spider Spider
}

func (s *UnimplementedScheduler) GetSpider() Spider {
	return s.spider
}
func (s *UnimplementedScheduler) SetSpider(spider Spider) {
	s.spider = spider
}
func (s *UnimplementedScheduler) GetDownloader() Downloader {
	return s.Downloader
}
func (s *UnimplementedScheduler) SetDownloader(downloader Downloader) {
	s.Downloader = downloader
}
func (s *UnimplementedScheduler) GetExporter() Exporter {
	return s.Exporter
}
func (s *UnimplementedScheduler) SetExporter(exporter Exporter) {
	s.Exporter = exporter
}
func (s *UnimplementedScheduler) GetItemDelay() time.Duration {
	return s.itemDelay
}
func (s *UnimplementedScheduler) SetItemDelay(itemDelay time.Duration) {
	s.itemDelay = itemDelay
}
func (s *UnimplementedScheduler) GetItemConcurrencyNew() int {
	return s.itemConcurrencyNew
}
func (s *UnimplementedScheduler) SetItemConcurrencyNew(itemConcurrency int) {
	s.itemConcurrencyNew = itemConcurrency
}
func (s *UnimplementedScheduler) GetItemConcurrency() int {
	return s.itemConcurrency
}
func (s *UnimplementedScheduler) SetItemConcurrencyRaw(itemConcurrency int) {
	s.itemConcurrency = itemConcurrency
}
func (s *UnimplementedScheduler) SetItemConcurrency(itemConcurrency int) {
	if s.itemConcurrency == itemConcurrency {
		return
	}

	if itemConcurrency < 1 {
		itemConcurrency = 1
	}

	s.itemConcurrencyNew = itemConcurrency
}
func (s *UnimplementedScheduler) RequestSlotLoad(slot string) (value any, ok bool) {
	return s.requestSlots.Load(slot)
}
func (s *UnimplementedScheduler) RequestSlotStore(slot string, value any) {
	s.requestSlots.Store(slot, value)
}
func (s *UnimplementedScheduler) SetRequestRate(slot string, interval time.Duration, concurrency int) {
	if slot == "" {
		slot = "*"
	}

	if concurrency < 1 {
		concurrency = 1
	}

	slotValue, ok := s.requestSlots.Load(slot)
	if !ok {
		requestSlot := rate.NewLimiter(rate.Every(interval/time.Duration(concurrency)), concurrency)
		s.requestSlots.Store(slot, requestSlot)
		return
	}

	limiter := slotValue.(*rate.Limiter)
	limiter.SetBurst(concurrency)
	limiter.SetLimit(rate.Every(interval / time.Duration(concurrency)))

	return
}
