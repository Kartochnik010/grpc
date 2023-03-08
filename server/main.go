package main

import (
	"demo/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	proto.GreetServiceServer
}

func main() {

	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterGreetServiceServer(grpcServer, &helloServer{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}
