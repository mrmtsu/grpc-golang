package domain

import (
	"time"
	"github.com/google/uuid"
)

type Todo struct {
	Id        TodoId
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTodo(title string, body string, now time.Time) *Todo {
	return &Todo{
		Id: TodoId(uuid.New().String()),
		Title: title,
		Body: body,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
