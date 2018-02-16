package main

import (
	res "./handler"
    pb "../proto"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"google.golang.org/grpc"
	"fmt"
    "context"
    "net"
)

var port = ":5000"

type server struct{}

func main() {
	fmt.Println("start - main")

	// Echoインスタンス生成
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())

	// -- gRPCサーバ起動
	fmt.Println("Server is starting...")
	e.Logger.Info("(logger)gRPC Server is starting...")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		e.Logger.Errorf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)

	fmt.Println("end   - main")
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return res.SayHello(ctx, in)
}

