package memory

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"golang.org/x/time/rate"
	"runtime"
	"time"
)

func (s *Scheduler) Request(ctx pkg.Context, request pkg.Request) (response pkg.Response, err error) {
	spider := s.GetSpider()
	defer func() {
		s.stateRequest.Set()
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
	spider := s.GetSpider()
	if fn != nil {
		fn(ctx, response, err)
	} else {
		s.logger.Warn("nil ErrBack")
	}
	spider.IncRequestError()
}

func (s *Scheduler) handleRequest(ctx context.Context) {
	spider := s.GetSpider()
	if ctx == nil {
		ctx = context.Background()
	}

	slot := "*"
	value, _ := s.RequestSlotLoad(slot)
	requestSlot := value.(*rate.Limiter)

	for request := range s.requestChan {
		slot = request.GetSlot()
		if slot == "" {
			slot = "*"
		}
		slotValue, ok := s.RequestSlotLoad(slot)
		if !ok {
			concurrency := uint8(1)
			if request.GetConcurrency() != nil {
				concurrency = *request.GetConcurrency()
			}
			if concurrency < 1 {
				concurrency = 1
			}
			requestSlot = rate.NewLimiter(rate.Every(request.GetInterval()/time.Duration(concurrency)), int(concurrency))
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
				s.stateRequest.Out()
				return
			}

			if request.GetCallBack() == nil {
				err = errors.New("nil CallBack")
				s.logger.Error(err)

				s.handleError(c, response, err, request.GetErrBack())
				s.stateRequest.Out()
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

				s.stateMethod.In()
				err = request.GetCallBack()(ctx, response)
				s.stateMethod.Out()
				s.stateRequest.Out()
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
		s.stateRequest.Set()
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

	s.stateRequest.In()
	s.requestChan <- request

	return
}

func (s *Scheduler) YieldExtra(ctx pkg.Context, extra any) (err error) {
	defer func() {
		s.stateRequest.Set()
	}()

	s.stateRequest.In()
	return
}
