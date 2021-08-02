package main

import (
	"net"
	"os"

	"com.github.Kinoshita0623/practice-grpc/app/service"
	pb "com.github.Kinoshita0623/practice-grpc/pb/cat"
	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":19003")

	if err != nil {
		os.Exit(1)
	}

	server := grpc.NewServer()
	catService := &service.MyCatService{}
	pb.RegisterCatServer(server, catService)
	server.Serve(listenPort)

}
