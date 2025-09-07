package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	crawlerContext "github.com/lizongying/go-crawler/pkg/context"
	request2 "github.com/lizongying/go-crawler/pkg/request"
	"github.com/segmentio/kafka-go"
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
			req, err := s.kafkaReader.FetchMessage(ctx.GetTask().GetContext())
			if err != nil {
				s.logger.Warn(err)
				continue
			}
			if len(req.Value) == 0 {
				err = errors.New("req is empty")
				s.logger.Warn(err)
				continue
			}

			request := new(request2.Request)
			if err = request.Unmarshal(req.Value); err != nil {
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
	requestCtx := ctx.GetRequest()
	if requestCtx != nil {
		// add referrer to request
		if requestCtx.GetReferrer() != "" {
			request.SetReferrer(requestCtx.GetReferrer())
		}

		// add cookies to request
		if len(requestCtx.GetCookies()) > 0 {
			for k, v := range requestCtx.GetCookies() {
				request.AddCookie(&http.Cookie{
					Name:  k,
					Value: v,
				})
			}
		}
	}

	ctx = ctx.CloneTask().
		WithRequest(new(crawlerContext.Request).
			WithContext(ctx.GetTask().GetContext()).
			WithId(s.crawler.NextId()).
			WithStatus(pkg.RequestStatusPending))

	request.WithContext(ctx)

	bs, err := request.Marshal()
	s.logger.Info("request with context:", string(bs))
	if err != nil {
		s.logger.Error(err)
		s.crawler.GetSignal().RequestChanged(request)
		return
	}

	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = s.kafkaWriter.WriteMessages(ctxTimeout, kafka.Message{
		Value: bs,
	}); err != nil {
		s.logger.Error(err)
		s.crawler.GetSignal().RequestChanged(request)
		return
	}

	s.crawler.GetSignal().RequestChanged(request)
	s.task.RequestIn()
	return
}

func (s *Scheduler) YieldExtra(c pkg.Context, extra any) (err error) {
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

	kafkaWriter, err := s.crawler.GetKafkaWriter(s.config.GetKafka(), fmt.Sprintf("%s-%s-extra-%s", s.config.GetBotName(), s.spider.Name(), name))
	if err != nil {
		s.logger.Error(err)
		return
	}
	if err = kafkaWriter.WriteMessages(ctx, kafka.Message{
		Value: bs,
	}); err != nil {
		s.logger.Error(err)
		return
	}

	s.task.RequestIn()
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
		var msg kafka.Message
		msg, err = kafka.NewReader(kafka.ReaderConfig{
			Brokers:  s.kafkaReader.Config().Brokers,
			MaxBytes: 10e6, // 10MB
			Topic:    fmt.Sprintf("%s-%s-extra-%s", s.config.GetBotName(), s.spider.Name(), name),
			GroupID:  s.config.GetBotName(),
		}).FetchMessage(ctx)
		if err != nil {
			s.logger.Error(err)
			return
		}

		if len(msg.Value) == 0 {
			err = errors.New("msg error")
			s.logger.Error(err)
			return
		}

		err = json.Unmarshal(msg.Value, extra)
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
