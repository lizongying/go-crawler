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

func SelfSigned() {
	// 生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("无法生成私钥：", err)
		return
	}

	// 创建自签名证书模板
	template := x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "localhost"},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0), // 有效期一年
		BasicConstraintsValid: true,
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}

	// 添加IP地址和域名到证书模板中
	template.IPAddresses = append(template.IPAddresses, net.ParseIP("127.0.0.1"))
	template.IPAddresses = append(template.IPAddresses, net.ParseIP("::1"))
	template.DNSNames = append(template.DNSNames, "localhost")

	// 使用模板生成自签名证书
	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		fmt.Println("无法生成证书：", err)
		return
	}

	// 将证书保存到文件
	serverCertBlock := pem.Block{Type: "CERTIFICATE", Bytes: derBytes}
	for _, i := range []string{
		"static/tls/server_self.crt",
		"static/tls/server_self_crt.pem",
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

	// 将私钥保存到文件
	serverKeyBlock := pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
	for _, i := range []string{
		"static/tls/server_self.key",
		"static/tls/server_self_key.pem",
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
