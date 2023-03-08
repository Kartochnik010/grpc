package main

import (
	"demo/proto"
	"io"
	"log"
)

func (h *helloServer) SayHelloClientStreaming(stream proto.GreetService_SayHelloClientStreamingServer) error {
	var names []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&proto.MessagesList{Message: names})
			break
		}
		if err != nil {
			return err
		}
		names = append(names, req.Name)
		log.Println("Client:", req.Name)
	}

	return nil
}
