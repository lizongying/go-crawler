package mock_servers

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
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

var versionMap = map[uint16]string{
	0x0301: "TLS 1.0",
	0x0302: "TLS 1.1",
	0x0303: "TLS 1.2",
	0x0304: "TLS 1.3",
}

var cipherSuiteMap = map[uint16]string{
	0x0005: "TLS_RSA_WITH_RC4_128_SHA",
	0x000a: "TLS_RSA_WITH_3DES_EDE_CBC_SHA",
	0x002f: "TLS_RSA_WITH_AES_128_CBC_SHA",
	0x0035: "TLS_RSA_WITH_AES_256_CBC_SHA",
	0x003c: "TLS_RSA_WITH_AES_128_CBC_SHA256",
	0x009c: "TLS_RSA_WITH_AES_128_GCM_SHA256",
	0x009d: "TLS_RSA_WITH_AES_256_GCM_SHA384",
	0xc007: "TLS_ECDHE_ECDSA_WITH_RC4_128_SHA",
	0xc009: "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA",
	0xc00a: "TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA",
	0xc011: "TLS_ECDHE_RSA_WITH_RC4_128_SHA",
	0xc012: "TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA",
	0xc013: "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA",
	0xc014: "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA",
	0xc023: "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256",
	0xc027: "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256",
	0xc02f: "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
	0xc02b: "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
	0xc030: "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
	0xc02c: "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
	0xcca8: "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256",
	0xcca9: "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256",
	0x1301: "TLS_AES_128_GCM_SHA256",
	0x1302: "TLS_AES_256_GCM_SHA384",
	0x1303: "TLS_CHACHA20_POLY1305_SHA256",
}

type HttpServer struct {
	url        *url.URL
	enableJa3  bool
	clientAuth uint8
	srv        *http.Server
	logger     pkg.Logger

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
	h.logger.Info("mock server at", h.url.String())
	h.srv.Handler = h.mux
	listener, e := net.Listen("tcp", h.url.Host)
	if e != nil {
		err = e
		h.logger.Error(err)
		return
	}
	go func() {
		if h.url.Scheme == "https" {
			var cer tls.Certificate
			cer, err = tls.X509KeyPair(static.ServerCert, static.ServerKey)
			if err != nil {
				h.logger.Error(err)
				return
			}

			conf := &tls.Config{
				GetCertificate: func(info *tls.ClientHelloInfo) (*tls.Certificate, error) {
					if h.enableJa3 {
						h.logger.Debug("ja3", convertToJA3(info))
					}
					return &cer, nil
				},
				VerifyConnection: func(state tls.ConnectionState) error {
					h.logger.Debug("Version:", versionMap[state.Version], "CipherSuite:", cipherSuiteMap[state.CipherSuite])
					return nil
				},
			}

			caCertPool := x509.NewCertPool()
			caCertPool.AppendCertsFromPEM(static.CaCert)
			conf.ClientCAs = caCertPool
			//conf.ClientAuth = tls.RequireAndVerifyClientCert
			if h.clientAuth != 0 {
				conf.ClientAuth = tls.ClientAuthType(h.clientAuth)
			}

			tlsListener := tls.NewListener(listener, conf)
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
		h.logger.Info("mock route added:", route.Pattern())
	}
}

func (h *HttpServer) GetRoutes() (routes []string) {
	for route := range h.routes {
		routes = append(routes, route)
	}
	return
}

func NewHttpServer(lc fx.Lifecycle, config *config.Config, logger pkg.Logger) (httpServer pkg.MockServer) {
	mockServer := config.MockServerHost()
	if mockServer == nil {
		err := errors.New("nil mockServer")
		logger.Error(err)
		return
	}

	srv := &http.Server{}
	httpServer = &HttpServer{
		url:        mockServer,
		enableJa3:  config.GetEnableJa3(),
		clientAuth: config.MockServerClientAuth(),
		srv:        srv,
		logger:     logger,
		mux:        http.NewServeMux(),
		routes:     make(map[string]struct{}),
	}
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return
}
