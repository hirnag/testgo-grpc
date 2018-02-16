package handler

import (
	pb "../../proto"
	"context"
	"log"
)


func SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("New Request: %v", in.String())
	return &pb.HelloReply{Message: "Hello, " + in.Name + "!"}, nil
}
