package middlewares

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

type RedirectMiddleware struct {
	pkg.UnimplementedMiddleware
	logger            pkg.Logger
	redirectHttpCodes []int
	redirectMaxTimes  uint8
}

func (m *RedirectMiddleware) SpiderStart(_ context.Context, spider pkg.Spider) (err error) {
	m.redirectHttpCodes = []int{http.StatusMovedPermanently, http.StatusFound}
	m.redirectMaxTimes = spider.GetConfig().GetRedirectMaxTimes()
	return
}

func (m *RedirectMiddleware) ProcessRequest(c *pkg.Context) (err error) {
	m.logger.Debug("enter ProcessRequest")
	defer func() {
		m.logger.Debug("exit ProcessRequest")
	}()

	err = c.NextRequest()

	request := c.Request
	ctx := request.Context()
	if request.RedirectMaxTimes != nil {
		ctx = context.WithValue(ctx, "redirect_max_times", *request.RedirectMaxTimes)
	}
	request.Request = request.WithContext(ctx)
	return
}

//func (m *RedirectMiddleware) ProcessResponse(c *pkg.Context) (err error) {
//	err = c.NextResponse()
//
//	response := c.Response
//	request := c.Request
//	m.logger.Debug("after response")
//
//	redirectMaxTimes := m.redirectMaxTimes
//	if request.RedirectMaxTimes != nil {
//		redirectMaxTimes = *request.RedirectMaxTimes
//	}
//
//	m.logger.Info("StatusCode", response.StatusCode)
//	m.logger.Info("redirectHttpCodes", m.redirectHttpCodes)
//	if redirectMaxTimes > 0 && (utils.InSlice(response.StatusCode, m.redirectHttpCodes)) {
//		if request.RedirectTimes < redirectMaxTimes {
//			request.RedirectTimes++
//			m.logger.Info(request.UniqueKey, "redirect times:", request.RedirectTimes, "SpendTime:", request.SpendTime)
//			err = c.FirstRequest()
//			return
//		}
//
//		err = errors.New("redirect max times")
//		m.logger.Error(request.UniqueKey, err, request.RedirectTimes, request.RedirectMaxTimes)
//		return
//	}
//
//	return
//}

func (m *RedirectMiddleware) FromCrawler(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(RedirectMiddleware).FromCrawler(spider)
	}
	m.logger = spider.GetLogger()
	return m
}
