package persistence

import (
	"context"
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"grpc-golang/adapter"
)

func NewStandardRawDB(connect string) *sql.DB {
	db, err := sql.Open("mysql", connect)
	if err != nil {
		panic(fmt.Errorf("failed to initialize db: %w", err))
	}
	return db
}

func NewDB(rawDB *sql.DB, conf adapter.DBConfig) adapter.DB {
	gormConf := &gorm.Config{
		DisableAutomaticPing: true,
	}

	conf.Apply(rawDB)
	conn, err := gorm.Open(
		mysql.New(mysql.Config{Conn: rawDB}),
		gormConf,
	)
	if err != nil {
		panic(fmt.Errorf("failed to initialize db: %w", err))
	}

	return func(ctx context.Context) *gorm.DB {
		return conn.WithContext(ctx)
	}
}
