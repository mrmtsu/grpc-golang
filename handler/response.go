package handler

import (
	"grpc-golang/domain"
	"grpc-golang/proto/pb"
)

func TodoFrom(todo *domain.Todo) *pb.TodoGetResponse_Todo {
	return &pb.TodoGetResponse_Todo{
		Id: todo.Id.String(),
		Title: todo.Title,
		Body: todo.Body,
	}
}
