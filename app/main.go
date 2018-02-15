package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "./handler"
    "google.golang.org/grpc"
    pb "./pb"
    "fmt"
    "golang.org/x/net/context"
    "net"
    "strconv"
	"strings"
)

func GRPCMiddleware(grpcServer *grpc.Server) func(next echo.HandlerFunc) echo.HandlerFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            r := c.Request()
            w := c.Response()
            if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
                grpcServer.ServeHTTP(w, r)
                return nil
            }
            if err := next(c); err != nil {
                c.Error(err)
            }
            return nil
        }
    }
}

type testServer struct {
}

func (p *testServer) Response(c context.Context, request *pb.TestRequest) (*pb.TestResponse, error) {
    fmt.Println("Request recieved.")
    fmt.Println(request.Name)

    response := pb.TestResponse{
        Value: request.Name,
    }

    return &response, nil
}

func main() {
	fmt.Println("start - main")
    // Echoインスタンス生成
    e := echo.New()

    // 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // gRPCサービス起動
    fmt.Println("Server is starting...")
    e.Logger.Info("(logger)gRPC Server is starting...")

    port := 5000
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
    if err != nil {
        e.Logger.Errorf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    //pb.RegisterTestServer(grpcServer, &testServer{})
    grpcServer.Serve(lis)

    e.Use(GRPCMiddleware(grpcServer))

    // ルーティング
    e.GET("/", handler.Hello())
    e.GET("/hello", handler.Hello())
    e.GET("/calc/:calcValue", handler.Clac())
    e.GET("/send/:sendParam", handler.Send())

    // サーバー起動
    e.Start(":" + strconv.Itoa(port))
    fmt.Println("end   - main")
}

