package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq" // do not delete. Required for connection to the db/
)

func CreateConnection() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		return nil, errors.New("unable to load .env file")
	}

	host, okHost := os.LookupEnv("DB_HOST")
	port, okPort := os.LookupEnv("DB_PORT")
	user, okUser := os.LookupEnv("DB_USER")
	password, okPassword := os.LookupEnv("DB_PASSWORD")
	dbname, okDBName := os.LookupEnv("DB_NAME")
	if !okHost || !okPort || !okUser || !okPassword || !okDBName {
		return nil, errors.New("unable to get environment variables")
	}

	// connection string
	psqlconn := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`, host, port, user, password, dbname)
	// open database
	conn, err := sql.Open("postgres", psqlconn)

	if connErr := conn.Ping(); connErr != nil {
		return nil, connErr
	}
	return conn, err
}
