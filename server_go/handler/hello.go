package handler

import (
	pb "../../proto"
	"context"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
	"fmt"
)

func SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Logger.Infof("gRPC New Request: %v", in.String())
	fmt.Printf("gRPC New Request: %v¥n", in.String())
	return &pb.HelloReply{Message: "Hello, " + in.Name + "!"}, nil
}
