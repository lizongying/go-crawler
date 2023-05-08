package devServer

import (
	"context"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/lizongying/go-crawler/pkg/logger"
	"go.uber.org/fx"
	"net"
	"net/http"
	"strings"
)

const defaultAddr = ":8081"

type HttpServer struct {
	srv    *http.Server
	logger *logger.Logger

	mux    *http.ServeMux
	routes map[string]struct{}
}

func (h *HttpServer) Run() (err error) {
	h.srv.Handler = h.mux
	ln, err := net.Listen("tcp", h.srv.Addr)
	if err != nil {
		h.logger.Error(err)
		return
	}

	h.logger.Info("Starting dev server at", h.srv.Addr)
	go func() {
		err = h.srv.Serve(ln)
		if err != nil {
			if err.Error() == "http: Server closed" {
				return
			}
			h.logger.Error(err)
		}
	}()
	return
}

func (h *HttpServer) AddRoutes(routes ...pkg.Route) {
	for _, route := range routes {
		h.mux.Handle(route.Pattern(), route)
		h.routes[route.Pattern()] = struct{}{}
	}
}

func (h *HttpServer) GetRoutes() (routes []string) {
	for route := range h.routes {
		routes = append(routes, route)
	}
	return
}

func (h *HttpServer) GetHost() (host string) {
	host = h.srv.Addr
	if !strings.Contains(host, "http://") && !strings.Contains(host, "https://") {
		host = fmt.Sprintf("http://%s", host)
	}
	return
}

func NewHttpServer(lc fx.Lifecycle, config *config.Config, logger *logger.Logger) (httpServer *HttpServer) {
	addr := defaultAddr
	if config.DevAddr != "" {
		addr = config.DevAddr
	}
	srv := &http.Server{Addr: addr}
	httpServer = &HttpServer{
		srv:    srv,
		logger: logger,
		mux:    http.NewServeMux(),
		routes: make(map[string]struct{}),
	}
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return
}
