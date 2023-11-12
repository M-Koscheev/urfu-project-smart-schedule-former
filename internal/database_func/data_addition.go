package database_func

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func GetKnowledgeList() string {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)

	checkError(err)

	// close database at the end of the method
	defer db.Close()

	err = db.Ping()
	checkError(err)

	fmt.Println("Connected to the database")

	rows, err := db.Query("select * from competencies")
	if err != nil {
		panic(err)
	}

	var titles string

	for rows.Next() {
		var t string
		var k string
		if err := rows.Scan(&t, &k); err != nil {
			log.Fatal(err)
		}
		titles += t + ": " + k + "\n"
	}

	return titles
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
