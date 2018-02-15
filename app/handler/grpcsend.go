package handler

import (
	"github.com/labstack/echo"
	"fmt"
	"google.golang.org/grpc"
	pb "../pb"
	"golang.org/x/net/context"
	"net/http"
)

func Send() echo.HandlerFunc {
	return func(c echo.Context) error {
		sendParam := c.Param("sendParam")

		host := "app"
		port := "5000"

		serverAddr := fmt.Sprintf("%s:%s", host, port)

		for {
			conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
			if err != nil {
				fmt.Println(err)
			}

			defer conn.Close()

			client := pb.NewTestClient(conn)
			feature, err := client.Proto(context.Background(), &pb.TestRequest{
				Name: sendParam,
			})
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(feature)
		}
		return c.String(http.StatusOK, sendParam + " を送信しました")
	}
}