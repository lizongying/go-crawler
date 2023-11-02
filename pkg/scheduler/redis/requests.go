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

	for {
		var req string
		var err error
		if s.enablePriorityQueue {
			r, e := s.redis.Do(ctx.GetTaskContext(), "EVALSHA", s.requestKeySha, 1, s.requestKey, s.batch).Result()
			if e != nil {
				s.logger.Warn(e)
				time.Sleep(1 * time.Second)
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
			r, e := s.redis.BLPop(ctx.GetTaskContext(), 0, s.requestKey).Result()
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

		err = s.redis.ZRem(ctx.GetTaskContext(), s.requestKey, req).Err()
		if err != nil {
			s.logger.Warn(err)
			continue
		}

		s.logger.Debugf("request: %s", req)
		request := new(request2.Request)
		if err = request.Unmarshal([]byte(req)); err != nil {
			s.logger.Warn(err)
			continue
		}

		c := request.Context
		s.logger.Debugf("request: %+v", request)
		if err != nil {
			s.logger.Warn(err)
			continue
		}
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

		err = requestSlot.Wait(ctx.GetTaskContext())
		if err != nil {
			s.logger.Error(err)
		}
		go func(c pkg.Context, request pkg.Request) {
			response, e := s.Request(c, request)
			if e != nil {
				s.task.StopRequest()
				return
			}

			go func(ctx pkg.Context, response pkg.Response) {
				defer func() {
					if r := recover(); r != nil {
						s.logger.Error(r)
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
		}(c, request)
	}
}

func (s *Scheduler) YieldRequest(ctx pkg.Context, request pkg.Request) (err error) {
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

	// add referrer to request
	if ctx.GetRequest().GetReferrer() != "" {
		request.SetReferrer(ctx.GetRequest().GetReferrer())
	}

	// add cookies to request
	if len(ctx.GetRequest().GetCookies()) > 0 {
		for k, v := range ctx.GetRequest().GetCookies() {
			request.AddCookie(&http.Cookie{
				Name:  k,
				Value: v,
			})
		}
	}

	request.WithContext(ctx)
	bs, err := request.Marshal()
	if err != nil {
		s.logger.Error(err)
		return
	}

	s.logger.Info("request:", string(bs))

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
			ctx.GetTask().StartRequest()
		}
	} else {
		err = s.redis.RPush(c, s.requestKey, bs).Err()
		ctx.GetTask().StartRequest()
	}

	if err != nil {
		s.logger.Error(err)
		return
	}

	return
}

func (s *Scheduler) YieldExtra(c pkg.Context, extra any) (err error) {
	spider := s.spider

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
			c.GetTask().StartRequest()
		}
	} else {
		extraKey := fmt.Sprintf("%s:%s:extra:%s", s.config.GetBotName(), name, spider.Name())
		if err = s.redis.RPush(ctx, extraKey, bs).Err(); err != nil {
			s.logger.Error(err)
			return
		}

		c.GetTask().StartRequest()
	}

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
		var msg []byte
		if s.enablePriorityQueue {
			key := fmt.Sprintf("%s:%s:extra:%s:priority", s.config.GetBotName(), s.spider.Name(), name)
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
			key := fmt.Sprintf("%s:%s:extra:%s", s.config.GetBotName(), s.spider.Name(), name)
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
