package devServer

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/lizongying/go-crawler/static"
	"go.uber.org/fx"
	"net"
	"net/http"
	"net/url"
	"strings"
)

type HttpServer struct {
	url       *url.URL
	enableJa3 bool
	srv       *http.Server
	logger    pkg.Logger

	mux    *http.ServeMux
	routes map[string]struct{}
}

func convertToJA3(info *tls.ClientHelloInfo) string {
	var ja3Builder strings.Builder

	ja3Builder.WriteString(fmt.Sprintf("%d,", info.SupportedVersions[0]))

	cipherSuites := make([]string, len(info.CipherSuites))
	for i, suite := range info.CipherSuites {
		cipherSuites[i] = fmt.Sprintf("%d", suite)
	}
	ja3Builder.WriteString(strings.Join(cipherSuites, "-"))
	ja3Builder.WriteString(",")

	//extensions := make([]string, len(info.Extensions))
	//for i, ext := range info.Extensions {
	//	extensions[i] = fmt.Sprintf("%04x", ext)
	//}
	//ja3Builder.WriteString(strings.Join(extensions, "-"))
	ja3Builder.WriteString(",")

	curves := make([]string, len(info.SupportedCurves))
	for i, curve := range info.SupportedCurves {
		curves[i] = fmt.Sprintf("%d", curve)
	}
	ja3Builder.WriteString(strings.Join(curves, "-"))
	ja3Builder.WriteString(",")

	points := make([]string, len(info.SupportedPoints))
	for i, point := range info.SupportedPoints {
		points[i] = fmt.Sprintf("%d", point)
	}
	ja3Builder.WriteString(strings.Join(points, "-"))

	return ja3Builder.String()
}

func (h *HttpServer) Run() (err error) {
	h.logger.Info("Starting dev server at", h.url.String())
	h.srv.Handler = h.mux
	listener, e := net.Listen("tcp", h.url.Host)
	if e != nil {
		err = e
		h.logger.Error(err)
		return
	}
	go func() {
		if h.url.Scheme == "https" {
			cer, e := tls.X509KeyPair(static.Cert, static.Key)
			if e != nil {
				err = e
				h.logger.Error(err)
				return
			}
			tlsListener := tls.NewListener(listener, &tls.Config{
				GetCertificate: func(info *tls.ClientHelloInfo) (*tls.Certificate, error) {
					if h.enableJa3 {
						h.logger.Info("ja3", convertToJA3(info))
					}
					return &cer, nil
				},
				VerifyConnection: func(state tls.ConnectionState) error {
					h.logger.Info("Version:", state.Version)
					h.logger.Info("CipherSuite:", state.CipherSuite)
					return nil
				},
			})
			err = h.srv.Serve(tlsListener)
		} else {
			err = h.srv.Serve(listener)
		}
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

func NewHttpServer(lc fx.Lifecycle, config *config.Config, logger pkg.Logger) (httpServer pkg.DevServer) {
	devServer, err := config.GetDevServer()
	if err != nil {
		logger.Error(err)
		return
	}
	srv := &http.Server{}
	httpServer = &HttpServer{
		url:       devServer,
		enableJa3: config.GetEnableJa3(),
		srv:       srv,
		logger:    logger,
		mux:       http.NewServeMux(),
		routes:    make(map[string]struct{}),
	}
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return
}
