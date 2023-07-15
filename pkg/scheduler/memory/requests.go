package memory

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"golang.org/x/time/rate"
	"net/http"
	"runtime"
	"time"
)

func (s *Scheduler) Request(ctx context.Context, request *pkg.Request) (response *pkg.Response, err error) {
	if request == nil {
		err = errors.New("nil request")
		return
	}

	s.logger.DebugF("request: %+v", *request)

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
		s.handleError(ctx, response, err, request.GetErrback())
		return
	}

	s.logger.DebugF("request %+v", *request)

	return
}

func (s *Scheduler) handleError(ctx context.Context, response *pkg.Response, err error, fn func(context.Context, *pkg.Response, error)) {
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
	value, _ := s.requestSlots.Load(slot)
	requestSlot := value.(*rate.Limiter)

	for request := range s.requestChan {
		slot = request.GetSlot()
		if slot == "" {
			slot = "*"
		}
		slotValue, ok := s.requestSlots.Load(slot)
		if !ok {
			concurrency := uint8(1)
			if request.GetConcurrency() != nil {
				concurrency = *request.GetConcurrency()
			}
			if concurrency < 1 {
				concurrency = 1
			}
			requestSlot = rate.NewLimiter(rate.Every(request.GetInterval()/time.Duration(concurrency)), int(concurrency))
			s.requestSlots.Store(slot, requestSlot)
		}

		requestSlot = slotValue.(*rate.Limiter)

		err := requestSlot.Wait(ctx)
		if err != nil {
			s.logger.Error(err)
		}
		go func(request *pkg.Request) {
			defer func() {
				<-s.requestActiveChan
			}()

			response, e := s.Request(ctx, request)
			if errors.Is(e, pkg.ErrIgnoreRequest) {
				return
			}

			if e != nil {
				err = e
				s.logger.Error(err)
				return
			}

			if request.GetCallback() == nil {
				err = errors.New("nil CallBack")
				s.logger.Error(err)

				s.handleError(request.Context(), response, err, request.GetErrback())
				return
			}

			go func(response *pkg.Response) {
				defer func() {
					if r := recover(); r != nil {
						buf := make([]byte, 1<<16)
						runtime.Stack(buf, true)
						err = errors.New(string(buf))
						s.logger.Error(err)
						s.handleError(response.Request.Context(), response, err, request.GetErrback())
					}
				}()

				err = request.GetCallback()(response.Request.Context(), response)
				if e != nil {
					s.logger.Error(err)
					s.handleError(response.Request.Context(), response, err, request.GetErrback())
					return
				}
			}(response)
		}(request)
	}

	return
}

func (s *Scheduler) YieldRequest(ctx context.Context, request *pkg.Request) (err error) {
	if len(s.requestChan) >= defaultRequestMax {
		err = errors.New("requestChan max limit")
		s.logger.Error(err)
		return
	}

	// add referer to request
	referer := ctx.Value("referer")
	if referer != nil {
		request.SetReferer(referer.(string))
	}

	// add cookies to request
	cookies := ctx.Value("cookies")
	if cookies != nil {
		for _, cookie := range cookies.([]*http.Cookie) {
			request.AddCookie(cookie)
		}
	}

	s.requestActiveChan <- struct{}{}
	s.requestChan <- request

	return
}

func (s *Scheduler) SetRequestRate(slot string, interval time.Duration, concurrency int) {
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

	return
}
