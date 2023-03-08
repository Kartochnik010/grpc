package main

import (
	"context"
	"demo/proto"
	"log"
	"time"
)

func callSayHello(client proto.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &proto.NoParam{})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Println(resp.Message)
}
