package memory

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	crawlerContext "github.com/lizongying/go-crawler/pkg/context"
	"github.com/lizongying/go-crawler/pkg/utils"
	"golang.org/x/time/rate"
	"net/http"
	"reflect"
	"runtime"
	"time"
)

func (s *Scheduler) Request(ctx pkg.Context, request pkg.Request) (response pkg.Response, err error) {
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
	ctx.GetTask().ReadyRequest()
	return
}

func (s *Scheduler) handleRequest(ctx pkg.Context) {
	slot := "*"
	value, _ := s.spider.RequestSlotLoad(slot)
	requestSlot := value.(*rate.Limiter)

	for request := range s.requestChan {
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

		err := requestSlot.Wait(ctx.GetTaskContext())
		if err != nil {
			s.logger.Error(err, time.Now(), ctx)
		}
		go func(request pkg.Request) {
			c := request.GetContext()

			response, e := s.Request(c, request.GetRequest())
			if e != nil {
				s.task.StopRequest()
				return
			}

			go func(ctx pkg.Context, response pkg.Response) {
				defer func() {
					if r := recover(); r != nil {
						buf := make([]byte, 1<<16)
						runtime.Stack(buf, true)
						err = errors.New(string(buf))
						//s.logger.Error(err)
						s.HandleError(ctx, response, err, request.GetErrBack())
					}
				}()

				if err = s.spider.CallBack(request.GetCallBack())(ctx, response); err != nil {
					s.logger.Error(err)
					s.HandleError(ctx, response, err, request.GetErrBack())
				}
				s.task.StopRequest()
			}(c, response)
		}(request)
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

	c := new(crawlerContext.Context).
		WithCrawler(ctx.GetCrawler()).
		WithSpider(ctx.GetSpider()).
		WithSchedule(ctx.GetSchedule()).
		WithTask(ctx.GetTask()).
		WithRequest(new(crawlerContext.Request).
			WithContext(context.Background()).
			WithId(utils.UUIDV1WithoutHyphens()).
			WithStatus(pkg.RequestStatusPending).
			WithStartTime(time.Now()))
	s.crawler.GetSignal().RequestStarted(c)

	request.WithContext(c)
	s.requestChan <- request

	ctx.GetTask().StartRequest()
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

	c.GetTask().StartRequest()
	return
}

func (s *Scheduler) GetExtra(_ pkg.Context, extra any) (err error) {
	defer func() {
		s.task.StopRequest()
	}()

	extraValue := reflect.ValueOf(extra)
	if extraValue.Kind() != reflect.Ptr || extraValue.IsNil() {
		err = errors.New("extra must be a non-null pointer")
		return
	}

	name := extraValue.Elem().Type().Name()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.config.CloseReasonQueueTimeout())*time.Second)
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
	case <-ctx.Done():
		close(resultChan)
		err = pkg.ErrQueueTimeout
		return
	}
}
