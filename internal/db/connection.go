package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // do not delete. Required for connection to the db/
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "123678"
	dbname   = "postgres"
)

func ConnectToDatabase() (*sql.DB, error) {
	// connection string
	psqlconn := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`, host, port, user, password, dbname)
	// open database
	db, err := sql.Open("postgres", psqlconn)

	return db, err
}
