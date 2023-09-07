package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"github.com/lizongying/go-crawler/static"
	"math/big"
	"net"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type Server struct {
	rootCert     *x509.Certificate
	rootKey      *rsa.PrivateKey
	privateKey   *rsa.PrivateKey
	listener     net.Listener
	serialNumber int64
}

func (s *Server) CreateFakeCertificateByDomain(domain string) (ca []byte, err error) {
	atomic.AddInt64(&s.serialNumber, 1)
	serverTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(s.serialNumber),
		Subject: pkix.Name{
			CommonName: domain,
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().AddDate(1, 0, 0),
		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageDataEncipherment,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}
	ca, err = x509.CreateCertificate(rand.Reader, serverTemplate, s.rootCert, &s.privateKey.PublicKey, s.rootKey)
	return
}

func NewServer() (s *Server, err error) {
	s = new(Server)
	// ca.cert
	block, _ := pem.Decode(static.CaCert)
	if block == nil {
		return
	}
	s.rootCert, err = x509.ParseCertificate(block.Bytes)
	if err != nil {
		return
	}
	// ca.key
	block, _ = pem.Decode(static.CaKey)
	if block == nil {
		return
	}
	if err != nil {
		return
	}
	s.rootKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	// server.key
	s.privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return
	}
	s.listener, err = net.Listen("tcp", ":"+strconv.Itoa(proxyPort+1))
	if err != nil {
		return
	}
	return
}

func CreateFakeHttpsWebSite(domain string, s *Server, successFun func()) {
	cert, err := s.CreateFakeCertificateByDomain(domain)
	if err != nil {
		return
	}

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	// fakeServer
	go func() {
		srv := &http.Server{
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("request header", r.Header)
				_, _ = w.Write([]byte(`
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Title</title>
</head>
<body>
  hello, world!
</body>
</html>
`))
			}),
		}

		tlsListener := tls.NewListener(s.listener, &tls.Config{
			Certificates: []tls.Certificate{{
				PrivateKey:  s.privateKey,
				Certificate: [][]byte{cert},
			}},
		})

		// Signal that server is open for business.
		waitGroup.Done()
		_ = srv.Serve(tlsListener)
	}()

	waitGroup.Wait()
	successFun()
}
