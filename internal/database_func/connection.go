package database_func

import (
	"database/sql"
	"fmt"
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
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// open database
	return sql.Open("postgres", psqlconn)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
