package spider

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"reflect"
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
		s.Logger.Debug(cap(s.itemConcurrencyChan), len(s.itemConcurrencyChan), "id:", item.GetId())
		go func(itemConcurrencyChan chan struct{}, item pkg.Item) {
			defer func() {
				if itemConcurrencyChan != s.itemConcurrencyChan && itemConcurrencyChanLen < 0 {
					itemConcurrencyChanLen++
				} else {
					s.itemConcurrencyChan <- struct{}{}
				}
				<-s.itemActiveChan
			}()

			requestContext := pkg.Context{
				Item:        item,
				Middlewares: s.SortedMiddlewares(),
			}

			err := requestContext.FirstItem()
			if err != nil {
				s.Logger.Error(err)
			}
		}(s.itemConcurrencyChan, item)

		if itemDelay > 0 {
			<-s.itemTimer.C
		}
	}

	return
}

func (s *BaseSpider) YieldItem(ctx context.Context, item pkg.Item) (err error) {
	data := item.GetData()
	if data == nil {
		err = errors.New("nil data")
		s.Logger.Error(err)
		return
	}

	if reflect.ValueOf(data).Kind() != reflect.Ptr {
		err = errors.New("data should be ptr")
		s.Logger.Error(err)
		return
	}

	if len(s.itemChan) == cap(s.itemChan) {
		err = errors.New("itemChan max limit")
		s.Logger.Error(err)
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

func (s *BaseSpider) SetItemDelay(itemDelay time.Duration) (spider pkg.Spider) {
	s.itemDelay = itemDelay
	return s
}

func (s *BaseSpider) SetItemConcurrency(itemConcurrency int) pkg.Spider {
	if s.itemConcurrency == itemConcurrency {
		return s
	}

	if itemConcurrency < 1 {
		itemConcurrency = 1
	}

	s.itemConcurrencyNew = itemConcurrency
	return s
}
