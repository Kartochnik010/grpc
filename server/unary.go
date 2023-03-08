package main

import (
	"context"
	"demo/proto"
)

func (h *helloServer) SayHello(ctx context.Context, req *proto.NoParam) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{
		Message: "Hello",
	}, nil
}
