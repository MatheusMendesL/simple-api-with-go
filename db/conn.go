package db

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func Conn() (*sql.DB, *Queries, error) {
	dsn := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		return nil, nil, err
	}

	queries := New(db)

	return db, queries, nil
}
