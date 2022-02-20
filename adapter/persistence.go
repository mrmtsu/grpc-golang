package adapter

import (
	"context"
	"database/sql"
	"time"

	"gorm.io/gorm"

	"grpc-golang/domain"
)

type DB func(ctx context.Context) *gorm.DB

type DBConfig struct {
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func (c DBConfig) Apply(db *sql.DB) {
	if c.MaxOpenConns != 0 {
		db.SetMaxOpenConns(c.MaxOpenConns)
	}
	if c.MaxIdleConns != 0 {
		db.SetMaxIdleConns(c.MaxIdleConns)
	}
	if c.ConnMaxLifetime != 0 {
		db.SetConnMaxLifetime(c.ConnMaxLifetime)
	}
}

type TodoRepository interface {
	Get(ctx context.Context, db *gorm.DB, id domain.TodoId) (*domain.Todo, error)
	Insert(ctx context.Context, db *gorm.DB, item *domain.Todo) error
}
