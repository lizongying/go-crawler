package spider

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/internal"
	"time"
)

func (s *BaseSpider) handleItem(_ context.Context) {
	itemConcurrencyChanLen := 0
	for item := range s.itemChan {
		itemDelay := s.itemDelay
		if itemDelay > 0 {
			s.itemTimer.Reset(itemDelay)
		}

		if s.itemConcurrencyNew != s.itemConcurrency {
			itemConcurrencyChanLen = s.itemConcurrencyNew - s.itemConcurrency + len(s.itemConcurrencyChan)
			s.itemConcurrencyChan = make(chan struct{}, s.itemConcurrencyNew)
			for i := 0; i < itemConcurrencyChanLen; i++ {
				s.itemConcurrencyChan <- struct{}{}
			}
			s.itemConcurrency = s.itemConcurrencyNew
		}

		<-s.itemConcurrencyChan
		s.Logger.Info(cap(s.itemConcurrencyChan), len(s.itemConcurrencyChan), "id:", item.Id)
		go func(itemConcurrencyChan chan struct{}, item *internal.Item) {
			defer func() {
				if itemConcurrencyChan != s.itemConcurrencyChan && itemConcurrencyChanLen < 0 {
					itemConcurrencyChanLen++
				} else {
					s.itemConcurrencyChan <- struct{}{}
				}
			}()

			ctx := context.Background()

			for _, v := range s.SortedPipelines() {
				e := v.ProcessItem(ctx, item)
				if errors.Is(e, internal.BreakErr) {
					break
				}
			}
			for _, v := range s.SortedMiddlewares() {
				e := v.ProcessItem(ctx, item)
				if errors.Is(e, internal.BreakErr) {
					break
				}
			}
		}(s.itemConcurrencyChan, item)

		if itemDelay > 0 {
			<-s.itemTimer.C
		}
	}

	return
}

func (s *BaseSpider) YieldItem(item *internal.Item) (err error) {
	if len(s.itemChan) == cap(s.itemChan) {
		err = errors.New("itemChan max limit")
		s.Logger.Error(err)
		return
	}
	s.itemChan <- item

	return
}

func (s *BaseSpider) SetItemDelay(itemDelay time.Duration) {
	s.itemDelay = itemDelay
}

func (s *BaseSpider) SetItemConcurrency(itemConcurrency int) {
	if s.itemConcurrency == itemConcurrency {
		return
	}

	if itemConcurrency < 1 {
		itemConcurrency = 1
	}

	s.itemConcurrencyNew = itemConcurrency
}
