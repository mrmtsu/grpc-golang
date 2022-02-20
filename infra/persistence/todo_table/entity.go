package todo_table

import (
	"grpc-golang/domain"
	"time"
)

func (e *Entity) TableName() string {
	return "todos"
}

type Entity struct {
	Id        string    `gorm:"column:id;primary_key"`
	Title     string    `gorm:"column:title"`
	Body      string    `gorm:"column:body"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func entityFrom(d *domain.Todo) *Entity {
	return &Entity{
		Id:        d.Id.String(),
		Title:     d.Title,
		Body:      d.Body,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

func (e *Entity) ToDomain() *domain.Todo {
	return &domain.Todo{
		Id:                  domain.TodoId(e.Id),
		Title: e.Title,
		Body: e.Body,
		CreatedAt:           e.CreatedAt,
		UpdatedAt:           e.UpdatedAt,
	}
}
