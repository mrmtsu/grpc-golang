package handler

import (
	"context"
	"time"

	"grpc-golang/adapter"
	"grpc-golang/domain"
	"grpc-golang/proto/pb"
)

func NewTodoQuery(dbFactory adapter.DB, todoRepo adapter.TodoRepository) pb.TodoQueryServer {
	return &TodoQuery{
		dbFactory: dbFactory,
		todoRepo: todoRepo,
  }
}

type TodoQuery struct {
	dbFactory adapter.DB
	todoRepo adapter.TodoRepository
}

func (q *TodoQuery) Get(ctx context.Context, req *pb.TodoGetRuest) (*pb.TodoGetResponse, error) {
	db := q.dbFactory(ctx)

	todo, err := q.todoRepo.Get(ctx, db, domain.TodoId(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.TodoGetResponse{
		Item: TodoFrom(todo),
	}, nil
}

func NewTodoCommand(dbFactory adapter.DB, todoRepo adapter.TodoRepository) pb.TodoCommandServer {
	return &TodoCommand{
		dbFactory: dbFactory,
		todoRepo: todoRepo,
	}
}

type TodoCommand struct {
	dbFactory adapter.DB
	todoRepo adapter.TodoRepository
}

func (c *TodoCommand) Create(ctx context.Context, req *pb.TodoCreateRuest) (*pb.TodoCreateResponse, error) {
	db := c.dbFactory(ctx)

	item := domain.NewTodo(req.Title, req.Body, time.Now())

	if err := c.todoRepo.Insert(ctx, db, item); err != nil {
		return nil, err
	}

	return &pb.TodoCreateResponse{
		Id: item.Id.String(),
	}, nil
}
