package handler

import (
	"context"

	"grpc-golang/proto/pb"
)

func NewHelloQuery() pb.HelloQueryServer {
	return &helloQuery{}
}

type helloQuery struct {}

func (q *helloQuery) Get(ctx context.Context, req *pb.HelloGetRuest) (*pb.HelloGetResponse, error) {
	return &pb.HelloGetResponse{
		Message: "Hello, World!",
	}, nil
}
