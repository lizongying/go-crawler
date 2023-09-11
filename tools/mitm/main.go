package main

import (
	"crypto/tls"
	"flag"
	"net/http"
)

func main() {
	replacePtr := flag.Bool("r", false, "-r replace")
	proxyPtr := flag.String("p", "", "-p proxy")
	flag.Parse()

	//*replacePtr = true
	p, err := NewProxy(*proxyPtr, *replacePtr)
	if err != nil {
		return
	}

	srv := &http.Server{
		Addr:         ":8082",
		Handler:      p,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}
	_ = srv.ListenAndServe()
}
