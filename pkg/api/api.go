package api

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/config"
	"go.uber.org/fx"
	"net/http"
)

type Api struct {
	key         string
	srv         *http.Server
	mux         *http.ServeMux
	routes      map[string]struct{}
	middlewares []func(next http.Handler) http.Handler
	logger      pkg.Logger
}

func (a *Api) Run() (err error) {
	go func() {
		a.logger.Infof("Starting api at http://%s\n", a.srv.Addr)
		a.logger.Info("api routes", a.GetRoutes())
		if err = a.srv.ListenAndServe(); err != nil {
			if err.Error() == "http: Server closed" {
				return
			}

			a.logger.Error(err)
			return
		}
	}()

	return
}

func (a *Api) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		a.logger.Info("Received request:", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func (a *Api) keyAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-API-Key") != a.key {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (a *Api) AddRoutes(routes ...pkg.Route) {
	for _, route := range routes {
		handler := http.Handler(route)
		for _, v := range a.middlewares {
			handler = v(handler)
		}
		a.mux.Handle(route.Pattern(), handler)
		a.routes[route.Pattern()] = struct{}{}
	}
}

func (a *Api) GetRoutes() (routes []string) {
	for route := range a.routes {
		routes = append(routes, route)
	}
	return
}

func NewApi(lc fx.Lifecycle, config *config.Config, logger pkg.Logger) (a *Api) {
	host := config.ApiHost()
	key := config.ApiKey()
	if host == nil {
		err := errors.New("nil host")
		logger.Error(err)
		return
	}

	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    host.Host,
		Handler: mux,
	}
	a = &Api{
		key:    key,
		srv:    srv,
		mux:    mux,
		routes: make(map[string]struct{}),
		logger: logger,
	}
	a.middlewares = []func(next http.Handler) http.Handler{
		a.loggingMiddleware,
		a.keyAuthMiddleware,
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return
}
