package main

import (
	"context"
	"demo/proto"
	"io"
	"log"
)

func callSayHelloServerStream(client proto.GreetServiceClient, names *proto.NamesList) {
	log.Println("Server streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalln("could not send names:", err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("error while streaming:", err)
		}
		log.Println("Server:", message.Message)
	}
	if err := stream.CloseSend(); err != nil {
		log.Println("Error while closing server stream:", err)
	}
	log.Fatalln("Server stream closed")
}
