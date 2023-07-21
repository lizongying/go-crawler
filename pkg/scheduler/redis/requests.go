package redis

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	request2 "github.com/lizongying/go-crawler/pkg/request"
	"golang.org/x/time/rate"
	"runtime"
	"time"
)

func (s *Scheduler) Request(ctx context.Context, request pkg.Request) (response pkg.Response, err error) {
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

	s.logger.DebugF("request %+v", request.GetRequest())

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
	value, _ := s.requestSlots.Load(slot)
	requestSlot := value.(*rate.Limiter)

	for {
		req, err := s.redis.BLPop(ctx, 0, s.requestKey).Result()
		if err != nil {
			s.logger.Warn(err)
			continue
		}
		if len(req) == 0 {
			err = errors.New("req is empty")
			s.logger.Warn(err)
			continue
		}
		//s.logger.DebugF("request: %s", req[1])
		var requestJson request2.RequestJson
		err = json.Unmarshal([]byte(req[1]), &requestJson)
		if err != nil {
			s.logger.Warn(err)
			continue
		}

		requestJson.SetCallBacks(s.crawler.GetSpider().GetCallBacks())
		requestJson.SetErrBacks(s.crawler.GetSpider().GetErrBacks())
		request, err := requestJson.ToRequest()
		s.logger.DebugF("request: %+v", request)
		if err != nil {
			s.logger.Warn(err)
			continue
		}
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

		err = requestSlot.Wait(ctx)
		if err != nil {
			s.logger.Error(err)
		}
		go func(request pkg.Request) {
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

			if request.GetCallBack() == nil {
				err = errors.New("nil CallBack")
				s.logger.Error(err)

				s.handleError(request.Context(), response, err, request.GetErrBack())
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

				err = request.GetCallBack()(response.Context(), response)
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
	l, err := s.redis.LLen(ctx, s.requestKey).Result()
	if int(l) >= defaultRequestMax {
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

	s.requestActiveChan <- struct{}{}
	bs, err := request.Marshal()
	s.logger.Debug("request:", string(bs))
	if err != nil {
		s.logger.Error(err)
		return
	}
	err = s.redis.RPush(ctx, s.requestKey, bs).Err()
	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

func (s *Scheduler) YieldExtra(ctx context.Context, extra any) (err error) {
	l, err := s.redis.LLen(ctx, s.requestKey).Result()
	if int(l) >= defaultRequestMax {
		err = errors.New("requestChan max limit")
		s.logger.Error(err)
		return
	}

	s.requestActiveChan <- struct{}{}
	bs, err := json.Marshal(extra)
	if err != nil {
		s.logger.Error(err)
		return
	}
	s.redis.RPush(ctx, s.requestKey, bs)

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
