package utils

import (
	"encoding/json"
	"io"
	"log"
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
	var err error
	resp, err := http.Get("https://ipinfo.io")
	if err != nil {
		log.Println(err)
		return
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	ipBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var result struct {
		Ip string `json:"ip"`
	}
	if err = json.Unmarshal(ipBytes, &result); err != nil {
		log.Println(err)
		return
	}

	ip = result.Ip
	return
}
