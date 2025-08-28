package scheduler

import (
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
func (s *UnimplementedScheduler) GetTask() pkg.Task {
	return s.task
}
func (s *UnimplementedScheduler) SetTask(task pkg.Task) {
	s.task = task
}
func (s *UnimplementedScheduler) SetLogger(logger pkg.Logger) {
	s.logger = logger
}
func (s *UnimplementedScheduler) HandleItem(ctx pkg.Context) {
out:
	for item := range s.itemChan {
		select {
		case <-ctx.GetTask().GetContext().Done():
			err := ctx.GetTask().GetContext().Err()
			s.logger.Error(err)
			item.GetContext().GetItem().
				WithStatus(pkg.ItemStatusFailure).
				WithStopReason(err.Error())
			s.crawler.GetSignal().ItemChanged(item)
			break out
		default:
			itemDelay := s.crawler.GetItemDelay()
			if itemDelay > 0 {
				s.crawler.ItemTimer().Reset(itemDelay)
			}

			<-s.crawler.ItemConcurrencyChan()
			s.logger.Debug(cap(s.crawler.ItemConcurrencyChan()), len(s.crawler.ItemConcurrencyChan()), "id:", item.Id())
			go func(item pkg.Item) {
				defer func() {
					s.crawler.ItemConcurrencyChan() <- struct{}{}
					s.task.ItemOut()
				}()

				contextItem := item.GetContext().GetItem()

				contextItem.
					WithStatus(pkg.ItemStatusRunning)
				s.crawler.GetSignal().ItemChanged(item)

				err := s.spider.Export(item)
				if err != nil {
					s.logger.Error(err)
					contextItem.
						WithStatus(pkg.ItemStatusFailure).
						WithStopReason(err.Error())
				} else {
					contextItem.
						WithStatus(pkg.ItemStatusSuccess)
				}
				s.crawler.GetSignal().ItemChanged(item)
			}(item)

			if itemDelay > 0 {
				<-s.crawler.ItemTimer().C
			}
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
			WithContext(ctx.GetTask().GetContext()).
			WithId(s.crawler.NextId()).
			WithStatus(pkg.ItemStatusPending))

	item.WithContext(c)
	s.crawler.GetSignal().ItemChanged(item)
	s.itemChan <- item
	s.task.ItemIn()
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
	if errBack, e := s.spider.ErrBack(errBackName); e != nil {
		errBack(ctx, response, e)
	}

	ctx.GetTask().IncRequestError()
}
func (s *UnimplementedScheduler) SyncRequest(ctx pkg.Context, request pkg.Request) (response pkg.Response, err error) {
	if request == nil {
		err = errors.New("nil request")
		return
	}

	s.logger.Debugf("request: %+v", request)

	response, err = s.spider.Download(ctx, request)
	if err != nil {
		if errors.Is(err, pkg.ErrIgnoreRequest) {
			s.logger.Info(err)
			err = nil
			return
		}

		s.HandleError(ctx, response, err, request.GetErrBack())
		return
	}

	s.logger.Debugf("request %+v", request)
	return
}
func (s *UnimplementedScheduler) Request(ctx pkg.Context, request pkg.Request) (response pkg.Response, err error) {
	s.task.RequestIn()
	response, err = s.SyncRequest(ctx, request)
	s.task.RequestOut()
	return
}
