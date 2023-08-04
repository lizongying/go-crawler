package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	request2 "github.com/lizongying/go-crawler/pkg/request"
	"github.com/redis/go-redis/v9"
	"golang.org/x/time/rate"
	"reflect"
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

	s.logger.DebugF("request %+v", request.GetRequest())

	return
}

func (s *Scheduler) handleError(ctx pkg.Context, response pkg.Response, err error, fn pkg.ErrBack) {
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

	for {
		var req string
		var err error
		if s.enablePriorityQueue {
			r, e := s.redis.Do(ctx, "EVALSHA", s.requestKeySha, 1, s.requestKey, s.batch).Result()
			if e != nil {
				s.logger.Warn(e)
				continue
			}
			rs, ok := r.([]interface{})
			if !ok {
				err = errors.New("req is empty")
				s.logger.Warn(err)
				time.Sleep(1 * time.Second)
				continue
			}
			if len(rs) == 0 {
				err = errors.New("req is empty")
				s.logger.Debug(err)
				time.Sleep(1 * time.Second)
				continue
			}
			for _, v := range rs {
				req = v.(string)
				break
			}
			s.logger.Debug("req", req)
		} else {
			r, e := s.redis.BLPop(ctx, 0, s.requestKey).Result()
			if e != nil {
				s.logger.Warn(e)
				continue
			}
			if len(r) == 0 {
				err = errors.New("req is empty")
				s.logger.Warn(err)
				time.Sleep(1 * time.Second)
				continue
			}
			req = r[1]
		}

		err = s.redis.ZRem(ctx, s.requestKey, req).Err()
		if err != nil {
			s.logger.Warn(err)
			continue
		}

		s.logger.DebugF("request: %s", req)
		var requestJson request2.RequestJson
		err = json.Unmarshal([]byte(req), &requestJson)
		if err != nil {
			s.logger.Warn(err)
			continue
		}

		requestJson.SetCallBacks(s.GetSpider().GetCallBacks())
		requestJson.SetErrBacks(s.GetSpider().GetErrBacks())
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

		err = requestSlot.Wait(ctx)
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

	c := context.Background()
	c, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	var l int64
	if s.enablePriorityQueue {
		l, err = s.redis.ZCard(c, s.requestKey).Result()
	} else {
		l, err = s.redis.LLen(c, s.requestKey).Result()
	}
	if err != nil {
		s.logger.Error(err)
		return
	}

	if int(l) >= defaultRequestMax {
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

	bs, err := request.Marshal()
	s.logger.Debug("request:", string(bs))
	if err != nil {
		s.logger.Error(err)
		return
	}

	c = context.Background()
	c, cancel = context.WithTimeout(c, 10*time.Second)
	defer cancel()

	if s.enablePriorityQueue {
		z := redis.Z{
			Score:  float64(request.GetPriority()),
			Member: bs,
		}
		var res int64
		res, err = s.redis.ZAdd(c, s.requestKey, z).Result()
		if res == 1 {
			s.stateRequest.In()
		}
	} else {
		err = s.redis.RPush(c, s.requestKey, bs).Err()
		s.stateRequest.In()
	}

	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

func (s *Scheduler) YieldExtra(ctx pkg.Context, extra any) (err error) {
	defer func() {
		s.stateRequest.Set()
	}()

	spider := s.GetSpider()

	extraValue := reflect.ValueOf(extra)
	if extraValue.Kind() != reflect.Ptr || extraValue.IsNil() {
		err = errors.New("extra must be a non-null pointer")
		return
	}

	name := extraValue.Elem().Type().Name()

	bs, err := json.Marshal(extra)
	if err != nil {
		s.logger.Error(err)
		return
	}

	if s.enablePriorityQueue {
		extraKey := fmt.Sprintf("%s:%s:extra:%s:priority", s.config.GetBotName(), name, spider.GetName())
		z := redis.Z{
			Score:  float64(time.Now().Unix() - 1000000000),
			Member: bs,
		}
		var res int64
		res, err = s.redis.ZAdd(context.Background(), extraKey, z).Result()
		if err != nil {
			s.logger.Error(err)
			return
		}
		if res == 1 {
			s.stateRequest.In()
		}
	} else {
		extraKey := fmt.Sprintf("%s:%s:extra:%s", s.config.GetBotName(), name, spider.GetName())
		err = s.redis.RPush(context.Background(), extraKey, bs).Err()
		if err != nil {
			s.logger.Error(err)
			return
		}
		s.stateRequest.In()
	}

	return
}
