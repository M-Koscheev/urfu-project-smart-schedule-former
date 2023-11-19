package app

import (
	"database/sql"
)

func (app *App) AddData(table string, needed string, column string) (sql.Result, error) {
	contains, errCont := app.CheckIfContains(table, needed, column)
	if contains || (errCont != nil && errCont != sql.ErrNoRows) {
		return nil, errCont
	}

	res, errAdd := app.db.Exec(`INSERT INTO knowledge(knowledge_pk) VALUES ($1)`, needed)

	return res, errAdd
}
