package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	crawlerContext "github.com/lizongying/go-crawler/pkg/context"
	request2 "github.com/lizongying/go-crawler/pkg/request"
	"github.com/redis/go-redis/v9"
	"net/http"
	"reflect"
	"time"
)

func (s *Scheduler) handleRequest(ctx pkg.Context) {
	slot := "*"
	limiter, _ := s.spider.Limiter(slot)

out:
	for {
		select {
		case <-ctx.GetTask().GetContext().Done():
			s.logger.Error(ctx.GetTask().GetContext().Err())
			break out
		default:
			var req string
			var err error
			if s.enablePriorityQueue {
				r, e := s.redis.Do(ctx.GetTask().GetContext(), "EVALSHA", s.requestKeySha, 1, s.requestKey, s.batch).Result()
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
				r, e := s.redis.BLPop(ctx.GetTask().GetContext(), 0, s.requestKey).Result()
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

			if err = s.redis.ZRem(ctx.GetTask().GetContext(), s.requestKey, req).Err(); err != nil {
				s.logger.Warn(err)
				continue
			}

			s.logger.Debugf("request: %s", req)
			request := new(request2.Request)
			if err = request.Unmarshal([]byte(req)); err != nil {
				s.logger.Warn(err)
				continue
			}

			ctx = request.Context
			s.logger.Debugf("request: %+v", request)
			if err != nil {
				s.logger.Warn(err)
				continue
			}
			slot = request.GetSlot()
			if slot == "" {
				slot = "*"
			}
			var ok bool
			limiter, ok = s.spider.Limiter(slot)
			if !ok {
				concurrency := uint8(1)
				if request.GetConcurrency() != nil {
					concurrency = *request.GetConcurrency()
				}
				interval := request.GetInterval()
				limiter = s.spider.SetRequestRate(slot, interval, int(concurrency))
			}

			if err = limiter.Wait(ctx.GetTask().GetContext()); err != nil {
				s.logger.Error(err)
			}
			ctx.GetRequest().WithStatus(pkg.RequestStatusRunning)
			s.crawler.GetSignal().RequestChanged(request)
			go func(c pkg.Context, request pkg.Request) {
				var response pkg.Response
				response, err = s.Request(c, request)
				if err != nil {
					ctx.GetRequest().WithStatus(pkg.RequestStatusFailure).WithStopReason(err.Error())
					s.crawler.GetSignal().RequestChanged(request)
					s.task.RequestOut()
					return
				}

				go func(ctx pkg.Context, response pkg.Response) {
					defer func() {
						if r := recover(); r != nil {
							s.logger.Error(r)
							err = errors.New("panic")
							s.HandleError(ctx, response, err, request.GetErrBack())
						}
						s.task.MethodOut()
						s.task.RequestOut()
					}()

					s.task.MethodIn()
					callback, _ := s.spider.CallBack(request.GetCallBack())
					if err = callback(ctx, response); err != nil {
						s.logger.Error(err)
						s.HandleError(ctx, response, err, request.GetErrBack())
						ctx.GetRequest().WithStatus(pkg.RequestStatusFailure).WithStopReason(err.Error())
						s.crawler.GetSignal().RequestChanged(request)
						return
					}
					ctx.GetRequest().WithStatus(pkg.RequestStatusSuccess)
					s.crawler.GetSignal().RequestChanged(request)
				}(c, response)
			}(ctx, request)
		}
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

	ctx = ctx.CloneTask().
		WithRequest(new(crawlerContext.Request).
			WithContext(ctx.GetTask().GetContext()).
			WithId(s.crawler.NextId()).
			WithStatus(pkg.RequestStatusPending))

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
			s.crawler.GetSignal().RequestChanged(request)
			s.task.RequestIn()
		}
	} else {
		err = s.redis.RPush(c, s.requestKey, bs).Err()
		s.crawler.GetSignal().RequestChanged(request)
		s.task.RequestIn()
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
			s.task.RequestIn()
		}
	} else {
		extraKey := fmt.Sprintf("%s:%s:extra:%s", s.config.GetBotName(), name, spider.Name())
		if err = s.redis.RPush(ctx, extraKey, bs).Err(); err != nil {
			s.logger.Error(err)
			return
		}

		s.task.RequestIn()
	}

	return
}

func (s *Scheduler) GetExtra(c pkg.Context, extra any) (err error) {
	defer func() {
		s.task.RequestOut()
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
