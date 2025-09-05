package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"math/big"
	"net"
	"os"
	"time"
)

func SelfServer(hostnames []string) {
	hostname := "localhost"
	if len(hostnames) > 0 {
		hostname = hostnames[0]
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Panicln("Unable to generate certificate key.", err)
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(4),
		Subject:      pkix.Name{CommonName: hostname},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(1, 0, 0), // Valid for 1 year
		IsCA:         false,
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}

	template.IPAddresses = append(template.IPAddresses, net.ParseIP("127.0.0.1"), net.ParseIP("::1"))
	template.DNSNames = append(template.DNSNames, "localhost")

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		log.Panicln("Unable to generate certificate.", err)
	}

	certBlock := pem.Block{Type: "CERTIFICATE", Bytes: derBytes}
	for _, i := range []string{
		"static/tls/server_self.crt",
		"static/tls/server_self_crt.pem",
	} {
		_ = save(i, &certBlock)
	}

	keyBlock := pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
	for _, i := range []string{
		"static/tls/server_self.key",
		"static/tls/server_self_key.pem",
	} {
		_ = save(i, &keyBlock)
	}
}

func SelfClient() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Panicln("Unable to generate certificate key.", err)
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(5),
		Subject:      pkix.Name{CommonName: "client"},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(1, 0, 0), // Valid for 1 year
		IsCA:         false,
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		log.Panicln("Unable to generate certificate.", err)
	}

	certBlock := pem.Block{Type: "CERTIFICATE", Bytes: derBytes}
	for _, i := range []string{
		"static/tls/client_self.crt",
		"static/tls/client_self_crt.pem",
	} {
		_ = save(i, &certBlock)
	}

	keyBlock := pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
	for _, i := range []string{
		"static/tls/client_self.key",
		"static/tls/client_self_key.pem",
	} {
		_ = save(i, &keyBlock)
	}
}

func save(path string, block *pem.Block) (err error) {
	var out *os.File
	out, err = os.Create(path)
	if err != nil {
		log.Println("Failed:", path, err)
		return
	}
	_ = pem.Encode(out, block)
	_ = out.Close()
	log.Println("Success:", path)
	return
}
