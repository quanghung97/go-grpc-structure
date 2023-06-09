package main

import (
	"fmt"
	"log"
	"net"

	con_ "github.com/bav-demo/config"
	route "github.com/bav-demo/internal/routes"
	initConnect "github.com/bav-demo/server/init"
)

func main() {
	// lấy config
	config := con_.Config{}
	con, _ := config.NewConfig()

	// khởi tạo listen với port
	log.Println("Go gRPC server in port!", con.Port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", con.Port))
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
