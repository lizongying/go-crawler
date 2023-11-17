package main

import (
	"flag"
)

func main() {
	selfSignedPtr := flag.Bool("s", false, "-s Self-Signed")
	caPtr := flag.Bool("c", false, "-c Create CA")
	flag.Parse()

	if *selfSignedPtr {
		SelfSigned()
	} else {
		CaSigned(*caPtr)
	}
}
