package main

import (
	"log"
	"net"
	"os"

	"grpc-golang/di"
)

func main() {
	grpcServer := di.ResolveAPIHandler()

	grpcPort := "3000"
	log.Printf("listening grpc on port %s", grpcPort)
	li, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to open grpc listener: %+v", err)
	}
	grpcServer.Serve(li)

	log.Println("unexpected end of process")
	os.Exit(1)
}
