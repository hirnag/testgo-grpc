package handler

import (
	pb "../../proto"
	"context"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
)

func SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Logger.Infof("New Request: %v", in.String())
	return &pb.HelloReply{Message: "Hello, " + in.Name + "!"}, nil
}
