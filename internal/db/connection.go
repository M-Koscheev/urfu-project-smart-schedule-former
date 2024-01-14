package db

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq" // do not delete. Required for connection to the db/
	"github.com/pressly/goose"
)

func CreateConnection() (*sql.DB, error) {
	// if err := godotenv.Load(); err != nil {
	// 	return nil, err
	// }

	slog.Info("begin connection")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		slog.Warn(fmt.Sprint("empty environment variable: host: ", host, ", port: ", port, ", user: ", user, ", password: ", password, ", dbname: ", dbname))
	}
	slog.Info("got environment variables")

	psqlconn := fmt.Sprint("postgresql://", user, ":", password, "@", host, ":", port, "/", dbname, "?sslmode=disable")
	conn, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	} else if err = conn.Ping(); err != nil {
		return nil, err
	}
	slog.Info("connection opened")

	goose.SetDialect("postgres")
	if err = goose.Up(conn, "migrations/"); err != nil {
		return nil, err
	}
	slog.Info("database tables were created")

	return conn, nil
}
