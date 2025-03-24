package statistics

import (
	"net/http"
)

type Getaway struct {
}

func NewGetaway(server *http.Server) (s *Getaway) {

	s = &Getaway{}
	//gwmux := runtime.NewServeMux()
	//// Register Greeter
	//err := pb.RegisterStatisticsServer(context.Background(), gwmux)
	//if err != nil {
	//	log.Fatalln("Failed to register gateway:", err)
	//}
	//
	//server.Handler = gwmux
	//
	//pb.RegisterStatisticsServer(server, s)
	return
}
