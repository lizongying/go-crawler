package spider

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"golang.org/x/time/rate"
	"runtime"
	"time"
)

func (s *BaseSpider) Request(ctx context.Context, request *pkg.Request) (response *pkg.Response, err error) {
	// TODO limit
	s.Logger.Debug("request", utils.JsonStr(request))

	if ctx == nil {
		ctx = context.Background()
	}

	if request.TimeoutAll != 0 {
		c, cancel := context.WithTimeout(ctx, request.TimeoutAll)
		ctx = c
		defer cancel()
	}

	for _, v := range s.SortedMiddlewares() {
		s.Logger.Debug("middleware", v.GetName())
		_, r, e := v.ProcessRequest(ctx, request)
		if e != nil {
			s.Logger.Error(e)
			if errors.Is(e, pkg.ErrIgnoreRequest) {
				err = e
				return
			}
			continue
		}
		if r != nil {
			response = r
			break
		}
	}

	if response != nil {
		if request.Request != nil {
			response.Request = request
		}
	}

	for _, v := range s.SortedMiddlewares() {
		_, _, e := v.ProcessResponse(ctx, response)
		if e != nil {
			s.Logger.Error(e)
			if errors.Is(e, pkg.BreakErr) {
				break
			}
			continue
		}
	}

	if response == nil {
		err = errors.New("nil response")
		s.Logger.Error(err)
		return
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

			var response *pkg.Response
			for _, v := range s.SortedMiddlewares() {
				_, r, e := v.ProcessRequest(ctx, request)
				if e != nil {
					s.Logger.Error(e)
					if errors.Is(e, pkg.ErrIgnoreRequest) {
						if request.ErrBack != nil {
							request.ErrBack(ctx, response, e)
						} else {
							e = errors.New("nil ErrBack")
							s.Logger.Warning(e)
						}
						return
					}
					continue
				}
				if r != nil {
					response = r
					break
				}
			}

			if response != nil {
				if request.Request != nil {
					response.Request = request
				}
			}

			for _, v := range s.SortedMiddlewares() {
				_, _, e := v.ProcessResponse(ctx, response)
				if e != nil {
					s.Logger.Error(e)
					if errors.Is(e, pkg.BreakErr) {
						break
					}
					continue
				}
			}

			if response == nil {
				e := errors.New("nil response")
				s.Logger.Error(e)
				if request.ErrBack != nil {
					e = errors.New("nil ErrBack")
					s.Logger.Warning(e)
				} else {
					request.ErrBack(ctx, response, e)
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
					if request.ErrBack == nil {
						e = errors.New("nil ErrBack")
						s.Logger.Error(e)
						return
					}
					request.ErrBack(ctx, response, e)
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

func (s *BaseSpider) SetRate(slot string, interval time.Duration, concurrency int) {
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
		return
	}

	limiter := slotValue.(*rate.Limiter)
	limiter.SetBurst(concurrency)
	limiter.SetLimit(rate.Every(interval / time.Duration(concurrency)))
}
