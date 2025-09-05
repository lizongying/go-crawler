package mock_servers

import "github.com/lizongying/go-crawler/pkg"

var NewRoutes = []pkg.NewRoute{
	NewRouteBadGateway,
	NewRouteBig5,
	NewRouteBrotli,
	NewRouteCookie,
	NewRouteDeflate,
	NewRouteFile,
	NewRouteGb2312,
	NewRouteGb18030,
	NewRouteGbk,
	NewRouteGet,
	NewRouteGzip,
	NewRouteHello,
	NewRouteHtml,
	NewRouteHttpAuth,
	NewRouteInternalServerError,
	NewRouteOk,
	NewRoutePost,
	NewRouteRateLimiter,
	NewRouteRedirect,
	NewRouteRobotsTxt,
	NewRouteTimeout,
}
