package memory

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	crawlerContext "github.com/lizongying/go-crawler/pkg/context"
	"golang.org/x/time/rate"
	"net/http"
	"reflect"
	"time"
)

func (s *Scheduler) handleRequest(ctx pkg.Context) {
	slot := "*"
	value, _ := s.spider.RequestSlotLoad(slot)
	requestSlot := value.(*rate.Limiter)

out:
	for request := range s.requestChan {
		select {
		case <-ctx.GetTask().GetContext().Done():
			s.logger.Error(ctx.GetTask().GetContext().Err())
			break out
		default:
			ctx = request.GetContext()
			slot = request.GetSlot()
			if slot == "" {
				slot = "*"
			}
			slotValue, ok := s.spider.RequestSlotLoad(slot)
			if !ok {
				concurrency := uint8(1)
				if request.GetConcurrency() != nil {
					concurrency = *request.GetConcurrency()
				}
				if concurrency < 1 {
					concurrency = 1
				}
				requestSlot = rate.NewLimiter(rate.Every(request.GetInterval()/time.Duration(concurrency)), int(concurrency))
				s.spider.RequestSlotStore(slot, requestSlot)
			}

			requestSlot = slotValue.(*rate.Limiter)

			if err := requestSlot.Wait(ctx.GetTask().GetContext()); err != nil {
				s.logger.Error(err, time.Now(), ctx)
			}
			ctx.GetRequest().WithStatus(pkg.RequestStatusRunning)
			s.crawler.GetSignal().RequestChanged(request)
			go func(request pkg.Request) {
				c := request.GetContext()
				var err error

				var response pkg.Response
				response, err = s.Request(c, request.GetRequest())
				if err != nil {
					ctx.GetRequest().WithStatus(pkg.RequestStatusFailure).WithStopReason(err.Error())
					s.crawler.GetSignal().RequestChanged(request)
					s.task.RequestOut()
					return
				}

				go func(ctx pkg.Context, response pkg.Response) {
					defer func() {
						if r := recover(); r != nil {
							s.logger.Error(r)
							err = errors.New("panic")
							s.HandleError(ctx, response, err, request.GetErrBack())
						}
						s.task.MethodOut()
						s.task.RequestOut()
					}()

					s.task.MethodIn()
					if err = s.spider.CallBack(request.GetCallBack())(ctx, response); err != nil {
						s.logger.Error(err)
						s.HandleError(ctx, response, err, request.GetErrBack())
						ctx.GetRequest().WithStatus(pkg.RequestStatusFailure).WithStopReason(err.Error())
						s.crawler.GetSignal().RequestChanged(request)
						return
					}
					ctx.GetRequest().WithStatus(pkg.RequestStatusSuccess)
					s.crawler.GetSignal().RequestChanged(request)
				}(c, response)
			}(request)
		}
	}

	return
}

func (s *Scheduler) YieldRequest(ctx pkg.Context, request pkg.Request) (err error) {
	if len(s.requestChan) >= defaultRequestMax {
		err = errors.New("exceeded the maximum number of requests")
		s.logger.Error(err)
		return
	}

	requestCtx := ctx.GetRequest()
	if requestCtx != nil {
		// add referrer to request
		if requestCtx.GetReferrer() != "" {
			request.SetReferrer(requestCtx.GetReferrer())
		}

		// add cookies to request
		if len(requestCtx.GetCookies()) > 0 {
			for k, v := range requestCtx.GetCookies() {
				request.AddCookie(&http.Cookie{
					Name:  k,
					Value: v,
				})
			}
		}
	}

	ctx = new(crawlerContext.Context).
		WithCrawler(ctx.GetCrawler()).
		WithSpider(ctx.GetSpider()).
		WithJob(ctx.GetJob()).
		WithTask(ctx.GetTask()).
		WithRequest(new(crawlerContext.Request).
			WithContext(context.Background()).
			WithId(s.crawler.NextId()).
			WithStatus(pkg.RequestStatusPending))

	request.WithContext(ctx)
	s.crawler.GetSignal().RequestChanged(request)
	s.requestChan <- request
	ctx.GetTask().RequestIn()
	return
}

func (s *Scheduler) YieldExtra(c pkg.Context, extra any) (err error) {
	extraValue := reflect.ValueOf(extra)
	if extraValue.Kind() != reflect.Ptr || extraValue.IsNil() {
		err = errors.New("extra must be a non-null pointer")
		return
	}

	name := extraValue.Elem().Type().Name()
	extraChan, ok := s.extraChanMap.LoadOrStore(name, func() chan any {
		extraChan := make(chan any, defaultRequestMax)
		extraChan <- extra
		return extraChan
	}())
	if ok {
		extraChan.(chan any) <- extra
	}

	c.GetTask().RequestIn()
	return
}

func (s *Scheduler) GetExtra(ctx pkg.Context, extra any) (err error) {
	defer func() {
		s.task.RequestOut()
	}()

	extraValue := reflect.ValueOf(extra)
	if extraValue.Kind() != reflect.Ptr || extraValue.IsNil() {
		err = errors.New("extra must be a non-null pointer")
		return
	}

	name := extraValue.Elem().Type().Name()

	c, cancel := context.WithTimeout(context.Background(), time.Duration(s.config.CloseReasonQueueTimeout())*time.Second)
	defer cancel()

	resultChan := make(chan struct{})
	go func() {
		extraChan, ok := s.extraChanMap.Load(name)
		if ok {
			extra = <-extraChan.(chan any)
			resultChan <- struct{}{}
		}
	}()

	select {
	case <-resultChan:
		return
	case <-c.Done():
		close(resultChan)
		err = pkg.ErrQueueTimeout
		return
	}
}
