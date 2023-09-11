package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"os"
	"time"
)

func CaSigned() {
	// 生成CA私钥
	caPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("无法生成CA私钥：", err)
		return
	}

	// 创建自签名CA证书模板
	caTemplate := x509.Certificate{
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
	caCertDER, err := x509.CreateCertificate(rand.Reader, &caTemplate, &caTemplate, &caPrivateKey.PublicKey, caPrivateKey)
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
		out, err := os.Create(i)
		if err != nil {
			fmt.Println("无法创建服务器证书文件：", err)
			return
		}
		_ = pem.Encode(out, &caCertBlock)
		_ = out.Close()
		fmt.Println("CA证书已保存到", i)
	}

	// 保存ca私钥到文件
	caKeyBlock := pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(caPrivateKey)}
	for _, i := range []string{
		"static/tls/ca.key",
		"static/tls/ca_key.pem",
	} {
		out, err := os.Create(i)
		if err != nil {
			fmt.Println("无法创建服务器私钥文件：", err)
			return
		}
		_ = pem.Encode(out, &caKeyBlock)
		_ = out.Close()
		fmt.Println("CA私钥已保存到", i)
	}

	// 生成服务器私钥
	serverPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("无法生成服务器私钥：", err)
		return
	}

	// 创建服务器证书模板
	serverTemplate := x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject:      pkix.Name{CommonName: "localhost"},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(1, 0, 0), // 有效期一年
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}

	// 添加IP地址和域名到证书模板中
	serverTemplate.IPAddresses = append(serverTemplate.IPAddresses, net.ParseIP("127.0.0.1"))
	serverTemplate.IPAddresses = append(serverTemplate.IPAddresses, net.ParseIP("::1"))
	serverTemplate.DNSNames = append(serverTemplate.DNSNames, "localhost")

	// 使用CA证书签发服务器证书
	serverCertDER, err := x509.CreateCertificate(rand.Reader, &serverTemplate, &caTemplate, &serverPrivateKey.PublicKey, caPrivateKey)
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
		out, err := os.Create(i)
		if err != nil {
			fmt.Println("无法创建服务器证书文件：", err)
			return
		}
		_ = pem.Encode(out, &serverCertBlock)
		_ = out.Close()
		fmt.Println("服务器证书已保存到", i)
	}

	// 保存服务器私钥到文件
	serverKeyBlock := pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(serverPrivateKey)}
	for _, i := range []string{
		"static/tls/server.key",
		"static/tls/server_key.pem",
	} {
		out, err := os.Create(i)
		if err != nil {
			fmt.Println("无法创建服务器私钥文件：", err)
			return
		}
		_ = pem.Encode(out, &serverKeyBlock)
		_ = out.Close()
		fmt.Println("服务器私钥已保存到", i)
	}
}
