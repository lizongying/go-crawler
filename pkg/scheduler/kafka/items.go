package kafka

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"reflect"
)

func (s *Scheduler) handleItem(ctx context.Context) {
	itemConcurrencyChanLen := 0
	for item := range s.itemChan {
		itemDelay := s.GetItemDelay()
		if itemDelay > 0 {
			s.itemTimer.Reset(itemDelay)
		}

		if s.GetItemConcurrencyNew() != s.GetItemConcurrency() {
			itemConcurrencyChanLen = s.GetItemConcurrencyNew() - s.GetItemConcurrency() + len(s.itemConcurrencyChan)
			s.itemConcurrencyChan = make(chan struct{}, s.GetItemConcurrencyNew())
			for i := 0; i < itemConcurrencyChanLen; i++ {
				s.itemConcurrencyChan <- struct{}{}
			}
			s.SetItemConcurrencyRaw(s.GetItemConcurrencyNew())
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
				s.stateItem.Out()
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

	dataValue := reflect.ValueOf(data)
	if !dataValue.IsNil() && dataValue.Kind() != reflect.Ptr {
		err = errors.New("item.Data must be a pointer")
		s.logger.Error(err)
		return
	}

	if len(s.itemChan) == cap(s.itemChan) {
		err = errors.New("itemChan max limit")
		s.logger.Error(err)
		return
	}

	// add referrer to item
	if meta, ok := ctx.Value("meta").(pkg.Meta); ok {
		if meta.Referrer != nil {
			item.SetReferrer(meta.Referrer.String())
		}
	}

	s.stateItem.In()
	s.itemChan <- item

	return
}
