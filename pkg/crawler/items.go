package crawler

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"reflect"
	"time"
)

func (c *Crawler) handleItem(ctx context.Context) {
	itemConcurrencyChanLen := 0
	for item := range c.itemChan {
		itemDelay := c.itemDelay
		if itemDelay > 0 {
			c.itemTimer.Reset(itemDelay)
		}

		if c.itemConcurrencyNew != c.itemConcurrency {
			itemConcurrencyChanLen = c.itemConcurrencyNew - c.itemConcurrency + len(c.itemConcurrencyChan)
			c.itemConcurrencyChan = make(chan struct{}, c.itemConcurrencyNew)
			for i := 0; i < itemConcurrencyChanLen; i++ {
				c.itemConcurrencyChan <- struct{}{}
			}
			c.itemConcurrency = c.itemConcurrencyNew
		}

		<-c.itemConcurrencyChan
		c.logger.Debug(cap(c.itemConcurrencyChan), len(c.itemConcurrencyChan), "id:", item.GetId())
		go func(itemConcurrencyChan chan struct{}, item pkg.Item) {
			defer func() {
				if itemConcurrencyChan != c.itemConcurrencyChan && itemConcurrencyChanLen < 0 {
					itemConcurrencyChanLen++
				} else {
					c.itemConcurrencyChan <- struct{}{}
				}
				<-c.itemActiveChan
			}()

			err := c.Export(ctx, item)
			if err != nil {
				c.logger.Error(err)
			}
		}(c.itemConcurrencyChan, item)

		if itemDelay > 0 {
			<-c.itemTimer.C
		}
	}

	return
}

func (c *Crawler) YieldItem(ctx context.Context, item pkg.Item) (err error) {
	data := item.GetData()
	if data == nil {
		err = errors.New("nil data")
		c.logger.Error(err)
		return
	}

	if reflect.ValueOf(data).Kind() != reflect.Ptr {
		err = errors.New("data should be ptr")
		c.logger.Error(err)
		return
	}

	if len(c.itemChan) == cap(c.itemChan) {
		err = errors.New("itemChan max limit")
		c.logger.Error(err)
		return
	}

	// add referer to item
	referer := ctx.Value("referer")
	if referer != nil {
		item.SetReferer(referer.(string))
	}

	c.itemActiveChan <- struct{}{}
	c.itemChan <- item

	return
}

func (c *Crawler) SetItemDelay(itemDelay time.Duration) (spider pkg.Crawler) {
	c.itemDelay = itemDelay
	return c
}

func (c *Crawler) SetItemConcurrency(itemConcurrency int) pkg.Crawler {
	if c.itemConcurrency == itemConcurrency {
		return c
	}

	if itemConcurrency < 1 {
		itemConcurrency = 1
	}

	c.itemConcurrencyNew = itemConcurrency
	return c
}
