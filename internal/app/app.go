package app

import (
	"database/sql"
)

type App struct {
	db *sql.DB
}

func New(db *sql.DB) *App {
	return &App{db: db}
}

func (app *App) AddKnowledge(elem string) error {
	if _, err := app.db.Exec(`INSERT INTO knowledge(knowledge_pk) VALUES ($1) ON CONFLICT DO NOTHING`, elem); err != nil {
		return err
	}

	return nil
}

func (app *App) GetAllKnowledges() ([]string, error) {
	sqlKnow, err := app.db.Query(`SELECT knowledge_pk FROM knowledge`)
	if err != nil {
		return nil, err
	}

	temp := ""
	var strKnow []string
	for sqlKnow.Next() {
		if err = sqlKnow.Scan(&temp); err != nil {
			return nil, err
		}
		strKnow = append(strKnow, temp)
	}
	return strKnow, nil
}
