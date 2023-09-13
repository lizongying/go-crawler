package main

import (
	"crypto/tls"
	"flag"
	"net/http"
)

func main() {
	filterPtr := flag.String("f", "", "-f filter")
	replacePtr := flag.Bool("r", false, "-r replace")
	proxyPtr := flag.String("p", "", "-p proxy")
	flag.Parse()

	//*replacePtr = true
	//*filterPtr = "baidu"
	p, err := NewProxy(*filterPtr, *proxyPtr, *replacePtr)
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
