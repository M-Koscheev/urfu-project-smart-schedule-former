package app

import (
	"database/sql"
	"fmt"
)

type App struct {
	db *sql.DB
}

func New(db *sql.DB) *App {
	return &App{db: db}
}

func (app *App) AddKnowledge(elem string) error {
	var contains bool

	errCont := app.db.QueryRow(`SELECT exists (SELECT 1 FROM knowledge WHERE $1 = $2 LIMIT 1)`, "knowledge_pk", elem).Scan(&contains)
	if contains || (errCont != nil && errCont != sql.ErrNoRows) {
		return errCont
	}

	_, errAdd := app.db.Exec(`INSERT INTO knowledge(knowledge_pk) VALUES ($1)`, elem)
	if errAdd != nil {
		return errAdd
	}

	return nil
}

func (app *App) GetKnowledges() (string, error) {
	knowledges, errKnow := app.db.Query(`SELECT * FROM knowledge`)
	if errKnow != nil {
		return "", errKnow
	}
	output := ""
	temp := ""
	for knowledges.Next() {
		errRes := knowledges.Scan(&temp)
		if errRes != nil {
			return output, errRes
		}
		output += fmt.Sprint(temp, ", ")
	}
	return output, nil
}
