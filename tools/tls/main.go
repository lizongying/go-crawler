package main

import (
	"flag"
)

func main() {
	selfSignedPtr := flag.Bool("s", false, "-s Self-Signed")
	flag.Parse()

	if *selfSignedPtr {
		SelfSigned()
	} else {
		CaSigned()
	}
}
