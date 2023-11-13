package scheduler

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	crawlerContext "github.com/lizongying/go-crawler/pkg/context"
	"reflect"
)

const defaultChanItemMax = 1000 * 1000

type UnimplementedScheduler struct {
	crawler  pkg.Crawler
	spider   pkg.Spider
	task     pkg.Task
	logger   pkg.Logger
	itemChan chan pkg.Item
}

func (s *UnimplementedScheduler) Init() {
	s.itemChan = make(chan pkg.Item, defaultChanItemMax)
	return
}
func (s *UnimplementedScheduler) SetCrawler(crawler pkg.Crawler) {
	s.crawler = crawler
}
func (s *UnimplementedScheduler) SetSpider(spider pkg.Spider) {
	s.spider = spider
}
func (s *UnimplementedScheduler) SetTask(task pkg.Task) {
	s.task = task
}
func (s *UnimplementedScheduler) SetLogger(logger pkg.Logger) {
	s.logger = logger
}
func (s *UnimplementedScheduler) HandleItem(_ pkg.Context) {
	for item := range s.itemChan {
		itemDelay := s.crawler.GetItemDelay()
		if itemDelay > 0 {
			s.crawler.ItemTimer().Reset(itemDelay)
		}

		<-s.crawler.ItemConcurrencyChan()
		s.logger.Debug(cap(s.crawler.ItemConcurrencyChan()), len(s.crawler.ItemConcurrencyChan()), "id:", item.Id())
		go func(item pkg.Item) {
			defer func() {
				s.crawler.ItemConcurrencyChan() <- struct{}{}
				s.task.StopItem()
			}()

			if err := s.spider.Export(item); err != nil {
				s.logger.Error(err)
			}
		}(item)

		if itemDelay > 0 {
			<-s.crawler.ItemTimer().C
		}
	}

	return
}
func (s *UnimplementedScheduler) YieldItem(ctx pkg.Context, item pkg.Item) (err error) {
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

	if len(s.itemChan) == cap(s.itemChan) {
		err = errors.New("itemChan max limit")
		s.logger.Error(err)
		return
	}

	// add referrer to item
	referrer := ctx.GetRequest().GetReferrer()
	if referrer != "" {
		item.SetReferrer(referrer)
	}

	c := new(crawlerContext.Context).
		WithCrawler(ctx.GetCrawler()).
		WithSpider(ctx.GetSpider()).
		WithJob(ctx.GetJob()).
		WithTask(ctx.GetTask()).
		WithItem(new(crawlerContext.Item).
			WithContext(context.Background()).
			WithId(s.crawler.NextId()).
			WithStatus(pkg.ItemStatusPending))
	//s.crawler.GetSignal().ItemChanged(c)

	item.WithContext(c)
	s.itemChan <- item

	s.task.StartItem()
	return
}
func (s *UnimplementedScheduler) HandleError(ctx pkg.Context, response pkg.Response, err error, errBackName string) {
	processed := false
	for _, v := range s.spider.GetMiddlewares().Middlewares() {
		next := v.ProcessError(ctx, response, err)
		if !next {
			break
		}
		processed = true
	}

	if processed {
		s.logger.Debug("error processed")
	}
	s.spider.ErrBack(errBackName)(ctx, response, err)
	ctx.GetTask().IncRequestError()
}
