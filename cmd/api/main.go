package main

import (
	"flag"
	"fmt"
	"github.com/lizongying/go-crawler/pkg/statistics"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	portPtr := flag.Int("port", 50051, "")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *portPtr))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	statistics.NewServer(s)

	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
