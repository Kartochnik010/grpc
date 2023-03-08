package main

import (
	"demo/proto"
	"io"
	"log"
)

func SayHelloBidirectionalStream(stream proto.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Println("Client:", req.Name)
		res := &proto.HelloResponse{
			Message: "Hi, " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
