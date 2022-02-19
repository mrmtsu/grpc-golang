package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"golang.org/x/sync/errgroup"
	"grpc-golang/di"
)

func main() {
	grpcServer := di.ResolveAPIHandler()
	grpcWebServer := grpcweb.WrapServer(
		grpcServer,
		grpcweb.WithOriginFunc(func(origin string) bool {
			return true
		}),
		grpcweb.WithAllowNonRootResource(true),
	)

	grpcPort := "3000"
	grpcWebPort := "3001"

	eg := &errgroup.Group{}
	eg.Go(func() error {
		log.Printf("listening grpc on port %s", grpcPort)
		li, err := net.Listen("tcp", ":"+grpcPort)
		if err != nil {
			log.Fatalf("failed to open grpc listener: %+v", err)
		}
		return grpcServer.Serve(li)
	})

	eg.Go(func() error {
		log.Printf("listening grpc web on port %s", grpcWebPort)
		li, err := net.Listen("tcp", ":"+grpcWebPort)
		if err != nil {
			log.Fatalf("failed to open grpc web listener: %+v", err)
		}
		return http.Serve(li, grpcWebServer)
	})
	log.Println(eg.Wait())

	log.Println("unexpected end of process")
	os.Exit(1)
}
