package middlewares

import (
	"context"
	"github.com/lizongying/go-crawler/internal"
	"github.com/lizongying/go-crawler/internal/logger"
	"github.com/lizongying/go-crawler/internal/utils"
	"net/http"
	"time"
)

type RecorderMiddleware struct {
	internal.UnimplementedMiddleware
	logger                    *logger.Logger
	interval                  time.Duration
	timer                     *time.Timer
	chanStop                  chan struct{}
	requestWithoutFilterCount int
	requestCount              int
	responseCount             int
	statusOK                  int
	statusErr                 int
	spiderInfo                *internal.SpiderInfo
}

func (m *RecorderMiddleware) GetName() string {
	return "recorder"
}

func (m *RecorderMiddleware) SpiderStart(_ context.Context, spider internal.Spider) (err error) {
	m.spiderInfo = spider.GetInfo()
	m.chanStop = make(chan struct{})
	m.timer = time.NewTimer(m.interval)
	go m.log()
	return
}

func (m *RecorderMiddleware) ProcessRequest(_ context.Context, r *internal.Request) (request *internal.Request, response *internal.Response, err error) {
	m.logger.Debug("request", utils.JsonStr(r))
	m.requestWithoutFilterCount++
	m.requestCount++
	return
}

func (m *RecorderMiddleware) ProcessResponse(_ context.Context, r *internal.Response) (request *internal.Request, response *internal.Response, err error) {
	m.responseCount++
	if r.StatusCode == http.StatusOK {
		m.statusOK++
	} else {
		m.statusErr++
	}
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

func NewRecorderMiddleware(logger *logger.Logger) (m internal.Middleware) {
	interval := time.Minute
	m = &RecorderMiddleware{
		logger:   logger,
		interval: interval,
	}
	return
}
