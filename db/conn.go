package db

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

func Conn() (*sql.DB, *Queries, error) {

	db, err := sql.Open("postgres", "postgresql://neondb_owner:npg_r92qGLSQEMlg@ep-empty-tooth-ac925cgm-pooler.sa-east-1.aws.neon.tech/go-project?sslmode=require&channel_binding=require")
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
