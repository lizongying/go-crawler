package kafka

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/items"
	"reflect"
)

func (s *Scheduler) handleItem(ctx context.Context) {
	itemConcurrencyChanLen := 0
	for itemWithContext := range s.itemWithContextChan {
		itemDelay := s.GetItemDelay()
		if itemDelay > 0 {
			s.itemTimer.Reset(itemDelay)
		}

		if s.ItemConcurrencyNew() != s.ItemConcurrency() {
			itemConcurrencyChanLen = s.ItemConcurrencyNew() - s.ItemConcurrency() + len(s.itemConcurrencyChan)
			s.itemConcurrencyChan = make(chan struct{}, s.ItemConcurrencyNew())
			for i := 0; i < itemConcurrencyChanLen; i++ {
				s.itemConcurrencyChan <- struct{}{}
			}
			s.SetItemConcurrencyRaw(s.ItemConcurrencyNew())
		}

		<-s.itemConcurrencyChan
		s.logger.Debug(cap(s.itemConcurrencyChan), len(s.itemConcurrencyChan), "id:", itemWithContext.Id())
		go func(itemConcurrencyChan chan struct{}, itemWithContext pkg.ItemWithContext) {
			defer func() {
				if itemConcurrencyChan != s.itemConcurrencyChan && itemConcurrencyChanLen < 0 {
					itemConcurrencyChanLen++
				} else {
					s.itemConcurrencyChan <- struct{}{}
				}
				s.Spider().StateItem().Out()
			}()

			err := s.Export(itemWithContext)
			if err != nil {
				s.logger.Error(err)
			}
		}(s.itemConcurrencyChan, itemWithContext)

		if itemDelay > 0 {
			<-s.itemTimer.C
		}
	}

	return
}

func (s *Scheduler) YieldItem(ctx pkg.Context, item pkg.Item) (err error) {
	data := item.Data()
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

	if len(s.itemWithContextChan) == cap(s.itemWithContextChan) {
		err = errors.New("itemChan max limit")
		s.logger.Error(err)
		return
	}

	// add referrer to item
	referrer := ctx.GetMeta().Referrer
	if referrer != "" {
		item.SetReferrer(referrer)
	}

	s.Spider().StateItem().In()
	s.itemWithContextChan <- &items.ItemWithContext{
		Context: ctx,
		Item:    item,
	}

	return
}
