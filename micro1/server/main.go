package main

import (
	"fmt"
	"log"
	"net"

	con_ "github.com/bav-demo/micro1/config"
	route "github.com/bav-demo/micro1/internal/routes"
	initConnect "github.com/bav-demo/micro1/server/init"
)

func main() {
	// lấy config
	config := con_.Config{}
	con, _ := config.NewConfig()

	// khởi tạo listen với port
	fmt.Printf("Go gRPC server in port %v!", con.Port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", con.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// connection
	initConnect.GetInstance()

	// khỏi tạo Route Register Service gRPC
	c := &route.Route{}
	g := c.NewRoute()

	if err := g.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
