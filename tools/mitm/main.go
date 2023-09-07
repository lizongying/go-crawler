package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"time"
)

func handleHTTP(w http.ResponseWriter, req *http.Request) {
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	_, _ = io.Copy(w, resp.Body)
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer func() {
		_ = destination.Close()
	}()
	defer func() {
		_ = source.Close()
	}()
	_, _ = io.Copy(destination, source)
}

const proxyPort = 8082

func main() {
	s, err := NewServer()
	if err != nil {
		return
	}

	fmt.Println("Proxy server listening on port", proxyPort)
	addr := "127.0.0.1:" + strconv.Itoa(proxyPort+1)
	srv := &http.Server{
		Addr: ":" + strconv.Itoa(proxyPort),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodConnect {
				CreateFakeHttpsWebSite(r.URL.Hostname(), s, func() {
					srvSocket, err := net.DialTimeout("tcp", addr, time.Minute)
					if err != nil {
						http.Error(w, err.Error(), http.StatusServiceUnavailable)
						return
					}

					w.WriteHeader(http.StatusOK)

					hijacker, ok := w.(http.Hijacker)
					if !ok {
						http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
						return
					}

					cltSocket, _, err := hijacker.Hijack()
					if err != nil {
						http.Error(w, err.Error(), http.StatusServiceUnavailable)
					}

					go transfer(srvSocket, cltSocket)
					go transfer(cltSocket, srvSocket)
				})
			} else {
				handleHTTP(w, r)
			}
		}),
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}
	_ = srv.ListenAndServe()
}
