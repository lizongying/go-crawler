package mockServers

import (
	"github.com/lizongying/go-crawler/pkg"
	"golang.org/x/text/encoding/simplifiedchinese"
	"net/http"
)

const UrlGbk = "/gbk"

type RouteGbk struct {
	logger pkg.Logger
}

func (h *RouteGbk) Pattern() string {
	return UrlGbk
}

func (h *RouteGbk) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	h.logger.Info("into HandlerGbk")
	defer func() {
		h.logger.Info("exit HandlerGbk")
	}()

	encoder := simplifiedchinese.GBK.NewEncoder()
	gbkBytes, _ := encoder.Bytes([]byte("汉字GBK"))

	w.Header().Set("Content-Type", "text/plain; charset=GBK")

	_, _ = w.Write(gbkBytes)
}

func NewRouteGbk(logger pkg.Logger) pkg.Route {
	return &RouteGbk{logger: logger}
}
