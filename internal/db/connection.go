package db

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // do not delete. Required for connection to the db/
	"github.com/pressly/goose"
)

func CreateConnection() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		slog.Warn("possible error loading environment variables")
	}

	psqlconn := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`, host, port, user, password, dbname)
	conn, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	} else if err = conn.Ping(); err != nil {
		return nil, err
	}

	goose.SetDialect("postgres")
	if err = goose.Up(conn, "migrations/"); err != nil {
		return nil, err
	}
	slog.Info("database connection was created")

	return conn, nil
}
