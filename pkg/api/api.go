package api

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/lizongying/go-crawler/pkg/utils"
	"github.com/lizongying/go-crawler/static"
	"go.uber.org/fx"
	"net/http"
)

type Api struct {
	accessKey   string
	https       bool
	srv         *http.Server
	mux         *http.ServeMux
	routes      map[string]struct{}
	middlewares []func(next http.Handler) http.Handler
	logger      pkg.Logger
}

func (a *Api) Run() (err error) {
	go func() {
		a.logger.Info("access key", a.accessKey)
		a.logger.Info("api routes", a.GetRoutes())
		if !a.https {
			a.logger.Infof("api at http://%s%s http://%s%s http://%s%s\n", "localhost", a.srv.Addr, utils.LanIp(), a.srv.Addr, utils.InternetIp(), a.srv.Addr)
			if err = a.srv.ListenAndServe(); err != nil {
				if err.Error() == "http: Server closed" {
					return
				}

				a.logger.Error(err)
				return
			}
		} else {
			a.logger.Infof("api at https://%s%s https://%s%s https://%s%s\n", "localhost", a.srv.Addr, utils.LanIp(), a.srv.Addr, utils.InternetIp(), a.srv.Addr)
			if err = a.srv.ListenAndServeTLS("", ""); err != nil {
				if err.Error() == "http: Server closed" {
					return
				}

				a.logger.Error(err)
				return
			}
		}
	}()

	return
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
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.ApiPort()),
		Handler: mux,
	}
	apiAccessKey := config.ApiAccessKey()
	if apiAccessKey == "" {
		apiAccessKey = utils.StrMd5(utils.NowStr())
	}
	https := config.ApiHttps()
	if https {
		cer, err := tls.X509KeyPair(static.ServerCert, static.ServerKey)
		if err != nil {
			a.logger.Error(err)
			return
		}

		srv.TLSConfig = &tls.Config{
			Certificates: []tls.Certificate{
				cer,
			},
		}
	}
	a = &Api{
		accessKey: apiAccessKey,
		https:     https,
		srv:       srv,
		mux:       mux,
		routes:    make(map[string]struct{}),
		logger:    logger,
	}
	a.middlewares = []func(next http.Handler) http.Handler{
		a.loggingMiddleware,
		a.keyAuthMiddleware,
		a.crossDomainMiddleware,
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return
}
