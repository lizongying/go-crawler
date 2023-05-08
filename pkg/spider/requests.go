package spider

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"golang.org/x/time/rate"
	"runtime"
	"time"
)

func (s *BaseSpider) Request(ctx context.Context, request *pkg.Request) (response *pkg.Response, err error) {
	// TODO limit
	s.Logger.DebugF("request: %+v", request)

	if ctx == nil {
		ctx = context.Background()
	}

	requestContext := pkg.Context{
		Request:     request,
		Middlewares: s.SortedMiddlewares(),
	}

	err = requestContext.FirstRequest()
	response = requestContext.Response
	if err != nil {
		s.Logger.Error(err)

		//if errors.Is(err, pkg.ErrIgnoreRequest) {
		//	s.Logger.Warning(err)
		//}

		if request.ErrBack != nil {
			request.ErrBack(ctx, response, err)
		} else {
			s.Logger.Warning("nil ErrBack")
		}
		return
	}

	if response == nil {
		err = errors.New("nil response")
		s.Logger.Error(err)
		return
	}

	if response != nil && response.Request == nil {
		if request != nil {
			response.Request = request
		}
	}

	return
}

func (s *BaseSpider) handleRequest(ctx context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}

	slot := "*"
	value, _ := s.requestSlots.Load(slot)
	requestSlot := value.(*rate.Limiter)

	for request := range s.requestChan {
		slot = request.Slot
		if slot == "" {
			slot = "*"
		}
		slotValue, ok := s.requestSlots.Load(slot)
		if !ok {
			if request.Concurrency < 1 {
				request.Concurrency = 1
			}
			requestSlot = rate.NewLimiter(rate.Every(request.Interval/time.Duration(request.Concurrency)), request.Concurrency)
			s.requestSlots.Store(slot, requestSlot)
		}

		requestSlot = slotValue.(*rate.Limiter)

		err := requestSlot.Wait(ctx)
		if err != nil {
			s.Logger.Error(err)
		}
		go func(request *pkg.Request) {
			defer func() {
				<-s.requestActiveChan
			}()

			requestContext := pkg.Context{
				Request:     request,
				Middlewares: s.SortedMiddlewares(),
			}

			err = requestContext.FirstRequest()
			response := requestContext.Response
			if err != nil {
				s.Logger.Error(err)

				if request.ErrBack != nil {
					request.ErrBack(ctx, response, err)
				} else {
					err = errors.New("nil ErrBack")
					s.Logger.Warning(err)
				}

				//if errors.Is(err, pkg.ErrIgnoreRequest) {
				//	s.Logger.Warning(err)
				//}

				return
			}

			if response != nil {
				if request.Request != nil {
					response.Request = request
				}
			}

			if response == nil {
				err = errors.New("nil response")
				s.Logger.Error(err)
				if request.ErrBack != nil {
					request.ErrBack(ctx, response, err)
				} else {
					err = errors.New("nil ErrBack")
					s.Logger.Warning(err)
				}
				return
			}

			if request.CallBack == nil {
				e := errors.New("nil CallBack")
				s.Logger.Error(e)
				return
			}

			go func() {
				defer func() {
					if e := recover(); e != nil {
						buf := make([]byte, 1<<16)
						runtime.Stack(buf, true)
						s.Logger.Error(string(buf))
					}
				}()
				e := request.CallBack(ctx, response)
				if e != nil {
					s.Logger.Error(e)
					if request.ErrBack != nil {
						request.ErrBack(ctx, response, e)
					} else {
						e = errors.New("nil ErrBack")
						s.Logger.Warning(e)
					}
					return
				}
			}()
		}(request)
	}

	return
}

func (s *BaseSpider) YieldRequest(request *pkg.Request) (err error) {
	if len(s.requestChan) == cap(s.requestChan) {
		err = errors.New("requestChan max limit")
		s.Logger.Error(err)
		return
	}

	if request.Skip {
		s.Logger.Debug("skip")
		return
	}
	s.requestActiveChan <- struct{}{}
	s.requestChan <- request

	return
}

func (s *BaseSpider) SetRequestRate(slot string, interval time.Duration, concurrency int) pkg.Spider {
	if slot == "" {
		slot = "*"
	}

	if concurrency < 1 {
		concurrency = 1
	}

	slotValue, ok := s.requestSlots.Load(slot)
	if !ok {
		requestSlot := rate.NewLimiter(rate.Every(interval/time.Duration(concurrency)), concurrency)
		s.requestSlots.Store(slot, requestSlot)
		return s
	}

	limiter := slotValue.(*rate.Limiter)
	limiter.SetBurst(concurrency)
	limiter.SetLimit(rate.Every(interval / time.Duration(concurrency)))

	return s
}
