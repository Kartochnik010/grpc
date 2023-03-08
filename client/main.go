package main

import (
	"demo/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

type helloClient struct {
	proto.GreetServiceClient
}

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := &helloClient{proto.NewGreetServiceClient(conn)}

	names := &proto.NamesList{
		Name: []string{"Ziyash", "Ilyas", "Amir", "Uali", "Ailin", "Daniya", "Ailin"},
	}

	callSayHello(client)
	// callSayHelloClientStreaming(client, names)
	// callSayHelloServerStream(client, names)
	callHelloBidirectionalStream(client, names)
}
