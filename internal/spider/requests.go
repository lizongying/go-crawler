package spider

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/internal"
	"github.com/lizongying/go-crawler/internal/utils"
	"time"
)

func (s *BaseSpider) Request(ctx context.Context, request *internal.Request) (response *internal.Response, err error) {
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
			if errors.Is(e, internal.ErrIgnoreRequest) {
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
			if errors.Is(e, internal.BreakErr) {
				break
			}
			continue
		}
	}

	return
}

func (s *BaseSpider) handleRequest(_ context.Context) {
	slotsCurrent := make(map[string]internal.RequestSlot)

	slot := "*"
	value, _ := s.requestSlots.Load(slot)
	requestSlot := value.(*internal.RequestSlot)
	slotsCurrent[slot] = *requestSlot

	for request := range s.requestChan {
		slot = request.Slot
		if slot == "" {
			slot = "*"
		}
		slotValue, ok := s.requestSlots.Load(slot)
		if !ok {
			requestSlot = new(internal.RequestSlot)
			if request.Delay > 0 {
				requestSlot.Delay = request.Delay
				requestSlot.Timer = time.NewTimer(requestSlot.Delay)
			}
			if request.Concurrency < 1 {
				request.Concurrency = 1
			}
			requestSlot.Concurrency = request.Concurrency
			requestSlot.ConcurrencyChan = make(chan struct{}, requestSlot.Concurrency)
			for i := 0; i < requestSlot.Concurrency; i++ {
				requestSlot.ConcurrencyChan <- struct{}{}
			}
			s.requestSlots.Store(slot, requestSlot)
			slotsCurrent[slot] = *requestSlot
		}

		requestSlot = slotValue.(*internal.RequestSlot)
		if requestSlot.Delay != slotsCurrent[slot].Delay {
			if requestSlot.Delay > 0 {
				requestSlot.Timer = time.NewTimer(requestSlot.Delay)
			}
		}
		if requestSlot.Concurrency != slotsCurrent[slot].Concurrency {
			requestConcurrency := requestSlot.Concurrency - slotsCurrent[slot].Concurrency + len(slotsCurrent[slot].ConcurrencyChan)
			requestSlot.ConcurrencyChan = make(chan struct{}, requestSlot.Concurrency)
			for i := 0; i < requestConcurrency; i++ {
				requestSlot.ConcurrencyChan <- struct{}{}
			}
		}
		slotsCurrent[slot] = *requestSlot

		<-requestSlot.ConcurrencyChan
		go func(requestConcurrency int, requestSlot *internal.RequestSlot, request *internal.Request) {
			defer func() {
				if requestSlot.Delay > 0 {
					<-requestSlot.Timer.C
				}
				//if requestConcurrency != requestSlot.Concurrency && requestConcurrencyChanLen < 0 {
				//	requestConcurrencyChanLen++
				//} else {
				//	requestSlot.ConcurrencyChan <- struct{}{}
				//}
			}()

			if requestSlot.Delay > 0 {
				requestSlot.Timer.Reset(requestSlot.Delay)
			}

			ctx := context.Background()

			var response *internal.Response
			for _, v := range s.SortedMiddlewares() {
				_, r, e := v.ProcessRequest(ctx, request)
				if e != nil {
					s.Logger.Error(e)
					if errors.Is(e, internal.ErrIgnoreRequest) {
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
					if errors.Is(e, internal.BreakErr) {
						break
					}
					continue
				}
			}

			if request.CallBack == nil {
				err := errors.New("nil CallBack")
				s.Logger.Error(err)
				return
			}

			err := request.CallBack(ctx, response)
			if err != nil {
				s.Logger.Error(err)
				if request.ErrBack == nil {
					err = errors.New("nil ErrBack")
					s.Logger.Error(err)
					return
				}
				request.ErrBack(ctx, response, err)
				return
			}
		}(requestSlot.Concurrency, requestSlot, request)
	}

	return
}

func (s *BaseSpider) YieldRequest(request *internal.Request) (err error) {
	if len(s.requestChan) == cap(s.requestChan) {
		err = errors.New("requestChan max limit")
		s.Logger.Error(err)
		return
	}

	if request.Skip {
		s.Logger.Debug("skip")
		return
	}

	s.requestChan <- request

	return
}

func (s *BaseSpider) SetRequestDelay(slot string, requestDelay time.Duration) {
	if slot == "" {
		slot = "*"
	}

	slotValue, ok := s.requestSlots.Load(slot)
	if !ok {
		requestSlot := &internal.RequestSlot{
			Delay: requestDelay,
		}
		s.requestSlots.Store(slot, requestSlot)
		return
	}

	requestSlot := slotValue.(*internal.RequestSlot)
	requestSlot.Delay = requestDelay
}

func (s *BaseSpider) SetRequestConcurrency(slot string, requestConcurrency int) {
	if requestConcurrency < 1 {
		requestConcurrency = 1
	}

	if slot == "" {
		slot = "*"
	}

	slotValue, ok := s.requestSlots.Load(slot)
	if !ok {
		requestSlot := &internal.RequestSlot{
			Concurrency: requestConcurrency,
		}
		s.requestSlots.Store(slot, requestSlot)
		return
	}

	requestSlot := slotValue.(*internal.RequestSlot)
	requestSlot.Concurrency = requestConcurrency
}
