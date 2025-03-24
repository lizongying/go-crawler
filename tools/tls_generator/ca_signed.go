package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"github.com/lizongying/go-crawler/static"
	"log"
	"math/big"
	"net"
	"os"
	"time"
)

func CreateCa() (caPrivateKey *rsa.PrivateKey, caCert *x509.Certificate, err error) {
	caPrivateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Panicln("Unable to generate ca private key.", err)
	}

	caKeyBlock := pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(caPrivateKey)}
	for _, i := range []string{
		"static/tls/ca.key",
		"static/tls/ca_key.pem",
	} {
		var out *os.File
		out, err = os.Create(i)
		if err != nil {
			log.Panicln("Unable to generate private key file.", err)
		}
		_ = pem.Encode(out, &caKeyBlock)
		_ = out.Close()
		log.Println("The ca certificate has been saved to", i)
	}

	caCert = &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: "ZONGYING",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0), // 有效期10年
		BasicConstraintsValid: true,
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}

	var caCertDER []byte
	caCertDER, err = x509.CreateCertificate(rand.Reader, caCert, caCert, &caPrivateKey.PublicKey, caPrivateKey)
	if err != nil {
		log.Panicln("Unable to create ca certificate.", err)
	}

	caCertBlock := pem.Block{Type: "CERTIFICATE", Bytes: caCertDER}
	for _, i := range []string{
		"static/tls/ca.crt",
		"static/tls/ca_crt.pem",
	} {
		var out *os.File
		out, err = os.Create(i)
		if err != nil {
			log.Panicln("Unable to create ca certificate file.", err)
		}
		_ = pem.Encode(out, &caCertBlock)
		_ = out.Close()
		log.Println("The ca certificate has been saved to", i)
	}
	return
}

func CaSigned(ca bool, ip []string, hostnames []string) {
	hostname := "localhost"
	if len(hostnames) > 0 {
		hostname = hostnames[0]
	}

	var caPrivateKey *rsa.PrivateKey
	var caCert *x509.Certificate
	var err error

	if ca {
		caPrivateKey, caCert, err = CreateCa()
		if err != nil {
			log.Panicln("create ca error", err)
		}
	} else {
		block, _ := pem.Decode(static.CaKey)
		if block == nil {
			err = errors.New("block nil")
			log.Panicln(err)
		}
		caPrivateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			log.Panicln(err)
		}

		block, _ = pem.Decode(static.CaCert)
		if block == nil {
			err = errors.New("block nil")
			log.Panicln(err)
		}
		caCert, err = x509.ParseCertificate(block.Bytes)
		if err != nil {
			log.Panicln(err)
		}
	}

	serverPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Panicln("Unable to create server certificate key.", err)
	}

	serverKeyBlock := pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(serverPrivateKey)}
	for _, i := range []string{
		"static/tls/server.key",
		"static/tls/server_key.pem",
	} {
		var out *os.File
		out, err = os.Create(i)
		if err != nil {
			log.Panicln("Unable to create server certificate key file.", err)
		}
		_ = pem.Encode(out, &serverKeyBlock)
		_ = out.Close()
		log.Println("The server certificate key has been saved to", i)
	}

	serverCert := &x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject:      pkix.Name{CommonName: hostname},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(1, 0, 0),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}

	serverCert.IPAddresses = append(serverCert.IPAddresses,
		net.ParseIP("127.0.0.1"),
		net.ParseIP("::1"),
	)
	if len(ip) > 0 {
		for _, v := range ip {
			serverCert.IPAddresses = append(serverCert.IPAddresses,
				net.ParseIP(v),
			)
		}
	}
	serverCert.DNSNames = append(serverCert.DNSNames, "localhost")
	if len(hostnames) > 0 {
		for _, v := range hostnames {
			serverCert.DNSNames = append(serverCert.DNSNames, v)
		}
	}

	serverCertDER, err := x509.CreateCertificate(rand.Reader, serverCert, caCert, &serverPrivateKey.PublicKey, caPrivateKey)
	if err != nil {
		log.Panicln("Unable to generate server certificate.", err)
	}

	serverCertBlock := pem.Block{Type: "CERTIFICATE", Bytes: serverCertDER}
	for _, i := range []string{
		"static/tls/server.crt",
		"static/tls/server_crt.pem",
	} {
		var out *os.File
		out, err = os.Create(i)
		if err != nil {
			log.Panicln("Unable to create server certificate file.", err)
		}
		_ = pem.Encode(out, &serverCertBlock)
		_ = out.Close()
		log.Println("The server certificate has been saved to", i)
	}
}
