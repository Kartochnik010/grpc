package main

import (
	"context"
	"demo/proto"
	"io"
	"log"
	"time"
)

func callSayHelloBidirectionalStream(client proto.GreetServiceClient, names *proto.NamesList) {
	log.Println("Bidirectional streaming started")

	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Println("Bistream failed to start", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println("Failed to get response while bistreaming:", err)
				break
			}
			log.Println("Server:", message)
		}
		close(waitc)
	}()

	for _, name := range names.Name {
		req := &proto.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil && err != io.EOF {
			log.Println("Failed to send request through bistream", err)
		}
		time.Sleep(time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Println("Bistream closed")
}
