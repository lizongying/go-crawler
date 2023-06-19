package scheduler

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"reflect"
	"time"
)

func (s *Scheduler) handleItem(ctx context.Context) {
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
		s.logger.Debug(cap(s.itemConcurrencyChan), len(s.itemConcurrencyChan), "id:", item.GetId())
		go func(itemConcurrencyChan chan struct{}, item pkg.Item) {
			defer func() {
				if itemConcurrencyChan != s.itemConcurrencyChan && itemConcurrencyChanLen < 0 {
					itemConcurrencyChanLen++
				} else {
					s.itemConcurrencyChan <- struct{}{}
				}
				<-s.itemActiveChan
			}()

			err := s.Export(ctx, item)
			if err != nil {
				s.logger.Error(err)
			}
		}(s.itemConcurrencyChan, item)

		if itemDelay > 0 {
			<-s.itemTimer.C
		}
	}

	return
}

func (s *Scheduler) YieldItem(ctx context.Context, item pkg.Item) (err error) {
	data := item.GetData()
	if data == nil {
		err = errors.New("nil data")
		s.logger.Error(err)
		return
	}

	if reflect.ValueOf(data).Kind() != reflect.Ptr {
		err = errors.New("data should be ptr")
		s.logger.Error(err)
		return
	}

	if len(s.itemChan) == cap(s.itemChan) {
		err = errors.New("itemChan max limit")
		s.logger.Error(err)
		return
	}

	// add referer to item
	referer := ctx.Value("referer")
	if referer != nil {
		item.SetReferer(referer.(string))
	}

	s.itemActiveChan <- struct{}{}
	s.itemChan <- item

	return
}

func (s *Scheduler) SetItemDelay(itemDelay time.Duration) {
	s.itemDelay = itemDelay
}

func (s *Scheduler) SetItemConcurrency(itemConcurrency int) {
	if s.itemConcurrency == itemConcurrency {
		return
	}

	if itemConcurrency < 1 {
		itemConcurrency = 1
	}

	s.itemConcurrencyNew = itemConcurrency
}
