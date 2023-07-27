package memory

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"golang.org/x/time/rate"
	"runtime"
	"time"
)

func (s *Scheduler) Request(ctx context.Context, request pkg.Request) (response pkg.Response, err error) {
	defer func() {
		s.stateRequest.Set()
	}()

	if request == nil {
		err = errors.New("nil request")
		return
	}

	s.logger.DebugF("request: %+v", request)

	if ctx == nil {
		ctx = context.Background()
	}

	response, err = s.Download(ctx, request)
	if err != nil {
		if errors.Is(err, pkg.ErrIgnoreRequest) {
			s.logger.Info(err)
			return
		}

		s.logger.Error(err)
		if request != nil {
			ctx = request.Context()
		}
		s.handleError(ctx, response, err, request.GetErrBack())
		return
	}

	s.logger.DebugF("request %+v", request)

	return
}

func (s *Scheduler) handleError(ctx context.Context, response pkg.Response, err error, fn func(context.Context, pkg.Response, error)) {
	if fn != nil {
		fn(ctx, response, err)
	} else {
		s.logger.Warn("nil ErrBack")
	}
	if errors.Is(err, pkg.ErrIgnoreRequest) {
		s.crawler.GetStats().IncRequestIgnore()
	} else {
		s.crawler.GetStats().IncRequestError()
	}
}

func (s *Scheduler) handleRequest(ctx context.Context) {
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
			response, e := s.Request(ctx, request)
			if errors.Is(e, pkg.ErrIgnoreRequest) {
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

				s.handleError(request.Context(), response, err, request.GetErrBack())
				s.stateRequest.Out()
				return
			}

			go func(response pkg.Response) {
				defer func() {
					if r := recover(); r != nil {
						buf := make([]byte, 1<<16)
						runtime.Stack(buf, true)
						err = errors.New(string(buf))
						s.logger.Error(err)
						s.handleError(response.Context(), response, err, request.GetErrBack())
					}
				}()

				s.stateMethod.In()
				err = request.GetCallBack()(response.Context(), response)
				s.stateMethod.Out()
				s.stateRequest.Out()
				if e != nil {
					s.logger.Error(err)
					s.handleError(response.Context(), response, err, request.GetErrBack())
					return
				}
			}(response)
		}(request)
	}

	return
}

func (s *Scheduler) YieldRequest(ctx context.Context, request pkg.Request) (err error) {
	defer func() {
		s.stateRequest.Set()
	}()

	if len(s.requestChan) >= defaultRequestMax {
		err = errors.New("requestChan max limit")
		s.logger.Error(err)
		return
	}

	if meta, ok := ctx.Value("meta").(pkg.Meta); ok {
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
	}

	s.stateRequest.In()
	s.requestChan <- request

	return
}
