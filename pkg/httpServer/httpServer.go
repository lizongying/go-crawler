package httpServer

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/lizongying/go-crawler/pkg/logger"
	"go.uber.org/fx"
	"net"
	"net/http"
)

const defaultAddr = ":8081"

type HttpServer struct {
	srv    *http.Server
	logger *logger.Logger
}

func (h *HttpServer) Run() (err error) {
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
	mux := http.NewServeMux()
	for _, route := range routes {
		mux.Handle(route.Pattern(), route)
	}
	h.srv.Handler = mux
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
	}
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return
}
