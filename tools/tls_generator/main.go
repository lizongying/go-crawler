package main

import (
	"flag"
	"strings"
)

func main() {
	selfSignedPtr := flag.Bool("s", false, "Self-Signed")
	caPtr := flag.Bool("c", false, "Create CA")
	ipPtr := flag.String("i", "", "IP")
	hostnamePtr := flag.String("n", "", "HostName")
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
		SelfServer(hostnames)
		SelfClient()
	} else {
		CaServer(*caPtr, ip, hostnames)
		CaClient(*caPtr)
	}
}
