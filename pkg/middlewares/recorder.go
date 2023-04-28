package middlewares

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/logger"
	"net/http"
	"time"
)

type RecorderMiddleware struct {
	pkg.UnimplementedMiddleware
	logger                    *logger.Logger
	interval                  time.Duration
	timer                     *time.Timer
	chanStop                  chan struct{}
	requestWithoutFilterCount int
	requestCount              int
	responseCount             int
	statusOK                  int
	statusIgnore              int
	statusErr                 int
	spiderInfo                *pkg.SpiderInfo
}

func (m *RecorderMiddleware) GetName() string {
	return "recorder"
}

func (m *RecorderMiddleware) SpiderStart(_ context.Context, spider pkg.Spider) (err error) {
	m.spiderInfo = spider.GetInfo()
	m.chanStop = make(chan struct{})
	m.timer = time.NewTimer(m.interval)
	go m.log()
	return
}

func (m *RecorderMiddleware) ProcessRequest(c *pkg.Context) (err error) {
	r := c.Request
	m.logger.DebugF("request: %+v", r)

	m.requestWithoutFilterCount++
	m.requestCount++

	err = c.NextRequest()
	return
}

func (m *RecorderMiddleware) ProcessResponse(c *pkg.Context) (err error) {
	r := c.Response
	m.logger.Debug("response body len:", len(r.BodyBytes))

	m.responseCount++
	if r == nil {
		m.statusErr++
	} else {
		if r.StatusCode == http.StatusOK {
			m.statusOK++
		} else {
			m.statusErr++
		}
	}

	err = c.NextResponse()
	return
}

func (m *RecorderMiddleware) SpiderStop(_ context.Context) (err error) {
	m.timer.Stop()
	m.chanStop <- struct{}{}
	m.logger.Info(m.spiderInfo.Name, "requestWithoutFilterCount:", m.requestWithoutFilterCount, "requestCount:", m.requestCount, "responseCount:", m.responseCount, "statusOK:", m.statusOK, "statusErr:", m.statusErr)
	return
}

func (m *RecorderMiddleware) log() {
	for {
		m.timer.Reset(m.interval)
		select {
		case <-m.timer.C:
			m.logger.Info(m.spiderInfo.Name, "requestWithoutFilterCount:", m.requestWithoutFilterCount, "requestCount:", m.requestCount, "responseCount:", m.responseCount, "statusOK:", m.statusOK, "statusErr:", m.statusErr)
		case <-m.chanStop:
			return
		}
	}
}

func NewRecorderMiddleware(logger *logger.Logger) (m pkg.Middleware) {
	interval := time.Minute
	m = &RecorderMiddleware{
		logger:   logger,
		interval: interval,
	}
	return
}
