package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"github.com/lizongying/go-crawler/static"
	"math/big"
	"net"
	"os"
	"time"
)

func CreateCa() (caPrivateKey *rsa.PrivateKey, caCert *x509.Certificate, err error) {
	// 生成CA私钥
	caPrivateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("无法生成CA私钥：", err)
		return
	}

	// 保存ca私钥到文件
	caKeyBlock := pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(caPrivateKey)}
	for _, i := range []string{
		"static/tls/ca.key",
		"static/tls/ca_key.pem",
	} {
		var out *os.File
		out, err = os.Create(i)
		if err != nil {
			fmt.Println("无法创建服务器私钥文件：", err)
			return
		}
		_ = pem.Encode(out, &caKeyBlock)
		_ = out.Close()
		fmt.Println("CA私钥已保存到", i)
	}

	// 创建自签名CA证书模板
	caCert = &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: "GO-MITM",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0), // 有效期10年
		BasicConstraintsValid: true,
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}

	// 使用CA私钥自签名CA证书
	var caCertDER []byte
	caCertDER, err = x509.CreateCertificate(rand.Reader, caCert, caCert, &caPrivateKey.PublicKey, caPrivateKey)
	if err != nil {
		fmt.Println("无法生成CA证书：", err)
		return
	}

	// 保存CA证书到文件
	caCertBlock := pem.Block{Type: "CERTIFICATE", Bytes: caCertDER}
	for _, i := range []string{
		"static/tls/ca.crt",
		"static/tls/ca_crt.pem",
	} {
		var out *os.File
		out, err = os.Create(i)
		if err != nil {
			fmt.Println("无法创建服务器证书文件：", err)
			return
		}
		_ = pem.Encode(out, &caCertBlock)
		_ = out.Close()
		fmt.Println("CA证书已保存到", i)
	}
	return
}

func CaSigned(ca bool, ip []string, hostname []string) {
	var caPrivateKey *rsa.PrivateKey
	var caCert *x509.Certificate
	var err error

	if ca {
		caPrivateKey, caCert, err = CreateCa()
		if err != nil {
			fmt.Println("create ca error")
			return
		}
	} else {
		block, _ := pem.Decode(static.CaKey)
		if block == nil {
			return
		}
		caPrivateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return
		}

		block, _ = pem.Decode(static.CaCert)
		if block == nil {
			return
		}
		caCert, err = x509.ParseCertificate(block.Bytes)
		if err != nil {
			return
		}
	}

	// 生成服务器私钥
	serverPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("无法生成服务器私钥：", err)
		return
	}

	// 保存服务器私钥到文件
	serverKeyBlock := pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(serverPrivateKey)}
	for _, i := range []string{
		"static/tls/server.key",
		"static/tls/server_key.pem",
	} {
		var out *os.File
		out, err = os.Create(i)
		if err != nil {
			fmt.Println("无法创建服务器私钥文件：", err)
			return
		}
		_ = pem.Encode(out, &serverKeyBlock)
		_ = out.Close()
		fmt.Println("服务器私钥已保存到", i)
	}

	// 创建服务器证书模板
	serverCert := &x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject:      pkix.Name{CommonName: "localhost"},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(1, 0, 0), // 有效期一年
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}

	// 添加IP地址和域名到证书模板中
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
	if len(hostname) > 0 {
		for _, v := range hostname {
			serverCert.DNSNames = append(serverCert.DNSNames, v)
		}
	}

	// 使用CA证书签发服务器证书
	serverCertDER, err := x509.CreateCertificate(rand.Reader, serverCert, caCert, &serverPrivateKey.PublicKey, caPrivateKey)
	if err != nil {
		fmt.Println("无法生成服务器证书：", err)
		return
	}

	// 保存服务器证书到文件
	serverCertBlock := pem.Block{Type: "CERTIFICATE", Bytes: serverCertDER}
	for _, i := range []string{
		"static/tls/server.crt",
		"static/tls/server_crt.pem",
	} {
		var out *os.File
		out, err = os.Create(i)
		if err != nil {
			fmt.Println("无法创建服务器证书文件：", err)
			return
		}
		_ = pem.Encode(out, &serverCertBlock)
		_ = out.Close()
		fmt.Println("服务器证书已保存到", i)
	}
}
