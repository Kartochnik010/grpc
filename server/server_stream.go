package main

import (
	"demo/proto"
	"log"
	"time"
)

func (h *helloServer) SayHelloServerStreaming(req *proto.NamesList, stream proto.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Got request with names: %v", req.Name)
	for _, name := range req.Name {
		res := &proto.HelloResponse{
			Message: "Hi, " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}

		time.Sleep(time.Second)
	}
	return nil
}
