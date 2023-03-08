package main

import (
	"context"
	"demo/proto"
	"log"
	"time"
)

func callHelloBidirectionalStream(client proto.GreetServiceClient, names *proto.NamesList) {
	log.Println("Client streaming started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalln("Client streaming error:", err)
	}
	for _, name := range names.Name {
		req := &proto.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Println("Client streaming failed to send message:", err)
		}

		log.Println("Sent request with name:", req.Name)
		time.Sleep(time.Second)
	}
	res, err := stream.CloseAndRecv()
	log.Println("Client streaming closed")
	if err != nil {
		log.Println("Error closing client stream:", err)
	}
	log.Println("Server:", res.Message)
}
