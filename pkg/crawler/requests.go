package crawler

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"golang.org/x/time/rate"
	"net/http"
	"runtime"
	"time"
)

func (c *Crawler) Request(ctx context.Context, request *pkg.Request) (response *pkg.Response, err error) {
	if request == nil {
		err = errors.New("nil request")
		return
	}

	c.logger.DebugF("request: %+v", *request)

	if ctx == nil {
		ctx = context.Background()
	}

	response, err = c.Download(ctx, request)
	if err != nil {
		c.logger.Error(err)
		if request != nil && request.Request != nil {
			ctx = request.Context()
		}
		c.handleError(ctx, response, err, request.ErrBack)
		return
	}

	c.logger.DebugF("request %+v", *request)

	return
}

func (c *Crawler) handleError(ctx context.Context, response *pkg.Response, err error, fn func(context.Context, *pkg.Response, error)) {
	if fn != nil {
		fn(ctx, response, err)
	} else {
		c.logger.Warn("nil ErrBack")
	}
	if errors.Is(err, pkg.ErrIgnoreRequest) {
		c.GetStats().IncRequestIgnore()
	} else {
		c.GetStats().IncRequestError()
	}
}

func (c *Crawler) handleRequest(ctx context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}

	slot := "*"
	value, _ := c.requestSlots.Load(slot)
	requestSlot := value.(*rate.Limiter)

	for request := range c.requestChan {
		slot = request.Slot
		if slot == "" {
			slot = "*"
		}
		slotValue, ok := c.requestSlots.Load(slot)
		if !ok {
			if request.Concurrency < 1 {
				request.Concurrency = 1
			}
			requestSlot = rate.NewLimiter(rate.Every(request.Interval/time.Duration(request.Concurrency)), request.Concurrency)
			c.requestSlots.Store(slot, requestSlot)
		}

		requestSlot = slotValue.(*rate.Limiter)

		err := requestSlot.Wait(ctx)
		if err != nil {
			c.logger.Error(err)
		}
		go func(request *pkg.Request) {
			defer func() {
				<-c.requestActiveChan
			}()

			response, e := c.Request(ctx, request)
			if e != nil {
				err = e
				c.logger.Error(err)
				return
			}

			if request.CallBack == nil {
				err = errors.New("nil CallBack")
				c.logger.Error(err)

				c.handleError(request.Context(), response, err, request.ErrBack)
				return
			}

			go func(response *pkg.Response) {
				defer func() {
					if r := recover(); r != nil {
						buf := make([]byte, 1<<16)
						runtime.Stack(buf, true)
						err = errors.New(string(buf))
						c.logger.Error(err)
						c.handleError(response.Request.Context(), response, err, request.ErrBack)
					}
				}()

				err = request.CallBack(response.Request.Context(), response)
				if e != nil {
					c.logger.Error(err)
					c.handleError(response.Request.Context(), response, err, request.ErrBack)
					return
				}
			}(response)
		}(request)
	}

	return
}

func (c *Crawler) YieldRequest(ctx context.Context, request *pkg.Request) (err error) {
	if len(c.requestChan) == cap(c.requestChan) {
		err = errors.New("requestChan max limit")
		c.logger.Error(err)
		return
	}

	if request.Skip {
		c.logger.Debug("skip")
		return
	}

	// add referer to request
	referer := ctx.Value("referer")
	if referer != nil {
		request.Referer = referer.(string)
	}

	// add cookies to request
	cookies := ctx.Value("cookies")
	if cookies != nil {
		request.Cookies = cookies.([]*http.Cookie)
	}

	c.requestActiveChan <- struct{}{}
	c.requestChan <- request

	return
}

func (c *Crawler) SetRequestRate(slot string, interval time.Duration, concurrency int) pkg.Crawler {
	if slot == "" {
		slot = "*"
	}

	if concurrency < 1 {
		concurrency = 1
	}

	slotValue, ok := c.requestSlots.Load(slot)
	if !ok {
		requestSlot := rate.NewLimiter(rate.Every(interval/time.Duration(concurrency)), concurrency)
		c.requestSlots.Store(slot, requestSlot)
		return c
	}

	limiter := slotValue.(*rate.Limiter)
	limiter.SetBurst(concurrency)
	limiter.SetLimit(rate.Every(interval / time.Duration(concurrency)))

	return c
}
