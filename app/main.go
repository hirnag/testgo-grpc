package main

import (
	res "./handler"
    pb "../proto"
	"google.golang.org/grpc"
    "context"
    "net"
	"log"
)

var port = ":5000"

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return res.SayHello(ctx, in)
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("boo")
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)

}

