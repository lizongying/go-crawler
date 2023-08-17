package memory

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"golang.org/x/time/rate"
	"reflect"
	"runtime"
	"time"
)

func (s *Scheduler) Request(ctx pkg.Context, request pkg.Request) (response pkg.Response, err error) {
	spider := s.Spider()
	defer func() {
		s.Spider().StateRequest().Set()
	}()

	if request == nil {
		err = errors.New("nil request")
		return
	}

	s.logger.DebugF("request: %+v", request)

	response, err = s.Download(ctx, request)
	if err != nil {
		if errors.Is(err, pkg.ErrIgnoreRequest) {
			s.logger.Info(err)
			spider.IncRequestIgnore()
			return
		}

		s.logger.Error(err)
		s.handleError(ctx, response, err, request.GetErrBack())
		return
	}

	s.logger.DebugF("request %+v", request)

	return
}

func (s *Scheduler) handleError(ctx pkg.Context, response pkg.Response, err error, fn func(pkg.Context, pkg.Response, error)) {
	spider := s.Spider()
	if fn != nil {
		fn(ctx, response, err)
	} else {
		s.logger.Warn("nil ErrBack")
	}
	spider.IncRequestError()
}

func (s *Scheduler) handleRequest(ctx context.Context) {
	spider := s.Spider()
	if ctx == nil {
		ctx = context.Background()
	}

	slot := "*"
	value, _ := s.RequestSlotLoad(slot)
	requestSlot := value.(*rate.Limiter)

	for request := range s.requestChan {
		slot = request.Slot()
		if slot == "" {
			slot = "*"
		}
		slotValue, ok := s.RequestSlotLoad(slot)
		if !ok {
			concurrency := uint8(1)
			if request.Concurrency() != nil {
				concurrency = *request.Concurrency()
			}
			if concurrency < 1 {
				concurrency = 1
			}
			requestSlot = rate.NewLimiter(rate.Every(request.Interval()/time.Duration(concurrency)), int(concurrency))
			s.RequestSlotStore(slot, requestSlot)
		}

		requestSlot = slotValue.(*rate.Limiter)

		err := requestSlot.Wait(ctx)
		if err != nil {
			s.logger.Error(err)
		}
		go func(request pkg.Request) {
			c := pkg.Context{}
			response, e := s.Request(c, request)
			if errors.Is(e, pkg.ErrIgnoreRequest) {
				s.logger.Info(err)
				spider.IncRequestIgnore()
				return
			}

			if e != nil {
				err = e
				s.logger.Error(err)
				s.Spider().StateRequest().Out()
				return
			}

			if request.GetCallBack() == nil {
				err = errors.New("nil CallBack")
				s.logger.Error(err)

				s.handleError(c, response, err, request.GetErrBack())
				s.Spider().StateRequest().Out()
				return
			}

			go func(ctx pkg.Context, response pkg.Response) {
				defer func() {
					if r := recover(); r != nil {
						buf := make([]byte, 1<<16)
						runtime.Stack(buf, true)
						err = errors.New(string(buf))
						s.logger.Error(err)
						s.handleError(ctx, response, err, request.GetErrBack())
					}
				}()

				s.Spider().StateMethod().In()
				err = request.GetCallBack()(ctx, response)
				s.Spider().StateMethod().Out()
				s.Spider().StateRequest().Out()
				if e != nil {
					s.logger.Error(err)
					s.handleError(ctx, response, err, request.GetErrBack())
					return
				}
			}(c, response)
		}(request)
	}

	return
}

func (s *Scheduler) YieldRequest(ctx pkg.Context, request pkg.Request) (err error) {
	defer func() {
		s.Spider().StateRequest().Set()
	}()

	if len(s.requestChan) >= defaultRequestMax {
		err = errors.New("exceeded the maximum number of requests")
		s.logger.Error(err)
		return
	}

	meta := ctx.Meta

	// add referrer to request
	if meta.Referrer != nil {
		request.SetReferrer(meta.Referrer.String())
	}

	// add cookies to request
	if len(meta.Cookies) > 0 {
		for _, cookie := range meta.Cookies {
			request.AddCookie(cookie)
		}
	}

	s.Spider().StateRequest().In()
	s.requestChan <- request

	return
}

func (s *Scheduler) YieldExtra(extra any) (err error) {
	defer func() {
		s.Spider().StateRequest().In()
		s.Spider().StateRequest().Set()
	}()

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

	return
}

func (s *Scheduler) GetExtra(extra any) (err error) {
	defer func() {
		s.Spider().StateRequest().Out()
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
