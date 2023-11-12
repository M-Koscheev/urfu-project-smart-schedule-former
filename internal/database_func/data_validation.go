package database_func

import (
	"database/sql"
	"fmt"
)

func CheckIfContains[T comparable](db *sql.DB, table string, needed T, column string) bool {
	req := fmt.Sprintf("SELECT %s FROM %s WHERE %s = ANY(%s))", column, table, column, needed)
	row := db.QueryRow(req)
	var data T
	err := row.Scan(&data)
	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		panic(err)
	}

	return true
}
