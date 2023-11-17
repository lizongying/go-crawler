package main

import (
	"flag"
	"strings"
)

func main() {
	selfSignedPtr := flag.Bool("s", false, "-s Self-Signed")
	caPtr := flag.Bool("c", false, "-c Create CA")
	ipPtr := flag.String("i", "", "-i IP")
	hostnamePtr := flag.String("h", "", "-h HostName")
	flag.Parse()

	if *selfSignedPtr {
		SelfSigned()
	} else {
		var ip []string
		var hostname []string
		if *ipPtr != "" {
			ip = strings.Split(*ipPtr, ",")
		}
		if *hostnamePtr != "" {
			hostname = strings.Split(*hostnamePtr, ",")
		}
		CaSigned(*caPtr, ip, hostname)
	}
}
