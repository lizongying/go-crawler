package mock_servers

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlCookie = "/cookie"

type RouteCookie struct {
	logger pkg.Logger
}

func (h *RouteCookie) Pattern() string {
	return UrlCookie
}

func (h *RouteCookie) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("into HandlerCookie")
	defer func() {
		h.logger.Info("exit HandlerCookie")
	}()

	cookie := &http.Cookie{
		Name:  "myCookie",
		Value: "Hello, Cookie!",
	}

	http.SetCookie(w, cookie)
	cookies, err := r.Cookie("myCookie")
	if err == nil {
		_, _ = fmt.Fprintln(w, "Cookie Value:", cookies.Value)
	} else {
		_, _ = fmt.Fprintln(w, "Cookie Not Found")
	}

	_, err = fmt.Fprintln(w, "Cookie Set")
	if err != nil {
		h.logger.Error(err)
	}

	_, _ = w.Write([]byte(fmt.Sprintf("Header: %v", r.Header)))

	w.WriteHeader(http.StatusOK)
}

func NewRouteCookie(logger pkg.Logger) pkg.Route {
	return &RouteCookie{logger: logger}
}
