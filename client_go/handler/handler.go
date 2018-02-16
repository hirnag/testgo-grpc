package handler

import (
    pb "../../proto"
    "net/http"
    "github.com/labstack/echo"
    "google.golang.org/grpc"
    "github.com/labstack/echo/middleware"
    "context"
    "fmt"
)

func Hello() echo.HandlerFunc {
    return func(c echo.Context) error {     //c をいじって Request, Responseを色々する 
        return c.String(http.StatusOK, "Hello World")
    }
}
func Clac() echo.HandlerFunc {
    return func(c echo.Context) error {
        calcValue := c.Param("calcValue")
        return c.String(http.StatusOK, "result:" + calcValue)
    }
}
func Send() echo.HandlerFunc {
    return func(c echo.Context) error {
        sendValue := c.Param("sendValue")

        e := echo.New()
        e.Use(middleware.Logger())

        conn, err := grpc.Dial("server_go:5000", grpc.WithInsecure())
        if err != nil {
            e.Logger.Fatalf("Connection error: %v", err)
        }
        defer conn.Close()

        cli := pb.NewGreeterClient(conn)

        res, err := cli.SayHello(context.Background(), &pb.HelloRequest{Name: sendValue})
        if err != nil {
            e.Logger.Fatalf("RPC error: %v", err)
        }

        fmt.Printf("gRPC Greeting: %s¥n", res.Message)
        return c.String(http.StatusOK, res.Message)
    }
}
