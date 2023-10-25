package utils

import (
	"fmt"
	"io"
	"net"
	"net/http"
)

func LanIp() (ip string) {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, iface := range interfaces {
			var addr []net.Addr
			addr, err = iface.Addrs()
			if err != nil {
				continue
			}

			for _, a := range addr {
				if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
					ip = ipNet.IP.String()
					break
				}
			}
		}
	}
	return
}

func InternetIp() (ip string) {
	resp, err := http.Get("https://api64.ipify.org?format=text")
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if err == nil {
		_, _ = fmt.Fscanf(resp.Body, "%s", &ip)
	}
	return
}
