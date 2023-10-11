package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	context2 "github.com/lizongying/go-crawler/pkg/context"
	request2 "github.com/lizongying/go-crawler/pkg/request"
	"github.com/redis/go-redis/v9"
	"golang.org/x/time/rate"
	"reflect"
	"runtime"
	"time"
)

func (s *Scheduler) Request(ctx pkg.Context, request pkg.Request) (response pkg.Response, err error) {
	defer func() {
		s.Spider().StateRequest().Set()
	}()

	if request == nil {
		err = errors.New("nil request")
		return
	}

	s.logger.Debugf("request: %+v", request)

	response, err = s.Download(ctx, request)
	if err != nil {
		if errors.Is(err, pkg.ErrIgnoreRequest) {
			s.logger.Info(err)
			err = nil
			return
		}

		s.HandleError(ctx, response, err, request.ErrBack())
		return
	}

	s.logger.Debugf("request %+v", request)
	return
}

func (s *Scheduler) handleRequest(ctx context.Context) {
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

		s.logger.Debugf("request: %s", req)
		var requestJsonWithContext request2.JsonWithContext
		requestJsonWithContext.ContextJson = new(context2.Json)
		requestJsonWithContext.RequestJson = new(request2.Json)
		err = json.Unmarshal([]byte(req), &requestJsonWithContext)
		if err != nil {
			s.logger.Warn(err)
			continue
		}

		request, err := requestJsonWithContext.ToRequest()
		c := requestJsonWithContext.ToContext()
		s.logger.Debugf("request: %+v", request)
		if err != nil {
			s.logger.Warn(err)
			continue
		}
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

		err = requestSlot.Wait(ctx)
		if err != nil {
			s.logger.Error(err)
		}
		go func(c pkg.Context, request pkg.Request) {
			response, e := s.Request(c, request)
			if e != nil {
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
						s.HandleError(ctx, response, err, request.ErrBack())
					}
				}()

				s.Spider().StateMethod().In()
				if err = s.Spider().CallBack(request.CallBack())(ctx, response); err != nil {
					s.logger.Error(err)
					s.HandleError(ctx, response, err, request.ErrBack())
				}
				s.Spider().StateMethod().Out()
				s.Spider().StateRequest().Out()
			}(c, response)
		}(c, request)
	}

	return
}

func (s *Scheduler) YieldRequest(ctx pkg.Context, request pkg.Request) (err error) {
	defer func() {
		s.Spider().StateRequest().Set()
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

	meta := ctx.Meta()

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

	bs, err := (&request2.WithContext{
		Context: ctx,
		Request: request,
	}).MarshalWithContext()
	s.logger.Info("request with context:", string(bs))
	if err != nil {
		s.logger.Error(err)
		return
	}

	c = context.Background()
	c, cancel = context.WithTimeout(c, 10*time.Second)
	defer cancel()

	if s.enablePriorityQueue {
		z := redis.Z{
			Score:  float64(request.Priority()),
			Member: bs,
		}
		var res int64
		res, err = s.redis.ZAdd(c, s.requestKey, z).Result()
		if res == 1 {
			s.Spider().StateRequest().In()
		}
	} else {
		err = s.redis.RPush(c, s.requestKey, bs).Err()
		s.Spider().StateRequest().In()
	}

	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

func (s *Scheduler) YieldExtra(extra any) (err error) {
	defer func() {
		s.Spider().StateRequest().Set()
	}()

	spider := s.Spider()

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

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if s.enablePriorityQueue {
		extraKey := fmt.Sprintf("%s:%s:extra:%s:priority", s.config.GetBotName(), name, spider.Name())
		z := redis.Z{
			Score:  float64(time.Now().Unix() - 1000000000),
			Member: bs,
		}
		var res int64
		res, err = s.redis.ZAdd(ctx, extraKey, z).Result()
		if err != nil {
			s.logger.Error(err)
			return
		}

		if res == 1 {
			s.Spider().StateRequest().In()
		}
	} else {
		extraKey := fmt.Sprintf("%s:%s:extra:%s", s.config.GetBotName(), name, spider.Name())
		if err = s.redis.RPush(ctx, extraKey, bs).Err(); err != nil {
			s.logger.Error(err)
			return
		}

		s.Spider().StateRequest().In()
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
		var msg []byte
		if s.enablePriorityQueue {
			key := fmt.Sprintf("%s:%s:extra:%s:priority", s.config.GetBotName(), s.Spider().Name(), name)
			r, e := s.redis.Do(ctx, "EVALSHA", s.requestKeySha, 1, key, 1).Result()
			if e != nil {
				err = e
				s.logger.Error(e)
				return
			}
			rs, ok := r.([]interface{})
			if !ok {
				err = errors.New("msg error")
				s.logger.Error(err)
				return
			}
			if len(rs) == 0 {
				err = errors.New("msg error")
				s.logger.Error(err)
				return
			}
			for _, v := range rs {
				msg = []byte(v.(string))
				break
			}
		} else {
			key := fmt.Sprintf("%s:%s:extra:%s", s.config.GetBotName(), s.Spider().Name(), name)
			r, e := s.redis.BLPop(ctx, 0, key).Result()
			if e != nil {
				err = e
				s.logger.Error(e)
				return
			}
			if len(r) == 0 {
				err = errors.New("msg error")
				s.logger.Error(err)
				return
			}
			msg = []byte(r[1])
		}

		err = json.Unmarshal(msg, extra)
		if err != nil {
			s.logger.Error(err)
			return
		}

		resultChan <- struct{}{}
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
