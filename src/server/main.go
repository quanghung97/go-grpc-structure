package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/quanghung97/grpc-docker/src/pb/echopb"

	"google.golang.org/grpc"
)

type server struct {
	echopb.EchoServiceServer
}

// unary API type
// implement interface
// con trỏ server khởi tạo mới, mỗi lần request thì sẽ tạo vùng nhớ khác nhau
func (*server) Echo(ctx context.Context, req *echopb.EchoRequest) (*echopb.EchoResponse, error) {
	log.Println("echo called...")
	// nội dung của response đều cố định được tham chiếu bằng địa chỉ
	resp := &echopb.EchoResponse{
		Message: req.GetMessage(),
	}

	return resp, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:9090")

	if err != nil {
		log.Fatalf("err listen %v", err)
	}

	// certFile := "ssl/server.crt"
	// keyFile := "ssl/server.key"

	// creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
	// if sslErr != nil {
	// 	log.Fatalf("create creds ssl err %v\n", sslErr)
	// 	return
	// }
	// opts := grpc.Creds(creds)

	// s := grpc.NewServer(opts)

	s := grpc.NewServer()

	echopb.RegisterEchoServiceServer(s, &server{})

	fmt.Println("echo is running 9090")

	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("err serve %v", err)
	}
}
