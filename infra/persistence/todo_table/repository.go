package todo_table

import (
	"context"
	"grpc-golang/adapter"
	"grpc-golang/domain"

	"gorm.io/gorm"
)

func NewRepository() adapter.TodoRepository {
	return &repository{}
}

type repository struct {}

func (r *repository) Get(ctx context.Context, db *gorm.DB, id domain.TodoId) (*domain.Todo, error) {
	var e Entity

	if err := db.Where("id = ?", id.String()).First(&e).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	return e.ToDomain(), nil
}

func (r *repository) Insert(ctx context.Context, db *gorm.DB, item *domain.Todo) error {
	if err := db.Create(entityFrom(item)).Error; err != nil {
		return err
	}

	return nil
}
