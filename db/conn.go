package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Conn() (*sql.DB, *Queries, error) {

	if err := godotenv.Load(); err != nil {
		fmt.Println("Arquivo .env não encontrado, usando variáveis de ambiente.")
	}

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
