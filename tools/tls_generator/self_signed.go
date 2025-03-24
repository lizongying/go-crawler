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

func SelfSigned(hostnames []string) {
	hostname := "localhost"
	if len(hostnames) > 0 {
		hostname = hostnames[0]
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Panicln("Unable to generate private key.", err)
	}

	template := x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: hostname},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0),
		BasicConstraintsValid: true,
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}

	template.IPAddresses = append(template.IPAddresses, net.ParseIP("127.0.0.1"), net.ParseIP("::1"))
	template.DNSNames = append(template.DNSNames, "localhost")

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		log.Panicln("Unable to generate certificate.", err)
	}

	serverCertBlock := pem.Block{Type: "CERTIFICATE", Bytes: derBytes}
	for _, i := range []string{
		"static/tls/server_self.crt",
		"static/tls/server_self_crt.pem",
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

	serverKeyBlock := pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
	for _, i := range []string{
		"static/tls/server_self.key",
		"static/tls/server_self_key.pem",
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
}
