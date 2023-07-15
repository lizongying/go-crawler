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

func (m *RedirectMiddleware) ProcessRequest(_ context.Context, request pkg.Request) (err error) {
	if request.GetRedirectMaxTimes() != nil {
		ctx := context.WithValue(request.Context(), "redirect_max_times", *request.GetRedirectMaxTimes())
		request.WithContext(ctx)
	}
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
//		if request.GetRedirectTimes() < redirectMaxTimes {
//			request.SetRedirectTimes(request.GetRedirectTimes()+1)
//			m.logger.Info(request.GetUniqueKey(), "redirect times:", request.GetRedirectTimes(), "SpendTime:", request.GetSpendTime())
//			err = c.FirstRequest()
//			return
//		}
//
//		err = errors.New("redirect max times")
//		m.logger.Error(request.GetUniqueKey(), err, request.GetRedirectTimes(), request.RedirectMaxTimes)
//		return
//	}
//
//	return
//}

func (m *RedirectMiddleware) FromCrawler(crawler pkg.Crawler) pkg.Middleware {
	if m == nil {
		return new(RedirectMiddleware).FromCrawler(crawler)
	}

	m.logger = crawler.GetLogger()
	m.redirectHttpCodes = []int{http.StatusMovedPermanently, http.StatusFound}
	m.redirectMaxTimes = crawler.GetConfig().GetRedirectMaxTimes()
	return m
}
