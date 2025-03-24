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

	var ip []string
	var hostnames []string
	if *ipPtr != "" {
		ip = strings.Split(*ipPtr, ",")
	}
	if *hostnamePtr != "" {
		hostnames = strings.Split(*hostnamePtr, ",")
	}

	if *selfSignedPtr {
		SelfSigned(hostnames)
	} else {
		CaSigned(*caPtr, ip, hostnames)
	}
}
