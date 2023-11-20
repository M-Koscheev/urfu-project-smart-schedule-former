package app

import (
	"database/sql"
	"log/slog"
	"net/http"
	"strings"
)

type App struct {
	db *sql.DB
}

func New(db *sql.DB) *App {
	return &App{db: db}
}

func (app *App) Run() error {
	http.HandleFunc("/", app.greet)
	err := http.ListenAndServe(":8080", nil)
	return err
}

func (app *App) greet(w http.ResponseWriter, r *http.Request) {
	err := app.AddKnowledge("14")
	if err != nil {
		slog.Error("error adding element 14 to the knowledge table", err)
	}
}

func (app *App) AddKnowledge(input string) error {
	knowledgeList := strings.Split(input, ", ")
	for _, elem := range knowledgeList {
		_, err := app.AddData("knowledge", elem, "knowledge_pk")
		if err != nil {
			return err
		}
	}

	return nil
}

func (app *App) CheckIfContains(table string, needed string, column string) (bool, error) {
	var data bool

	err := app.db.QueryRow(`SELECT exists (SELECT 1 FROM knowledge WHERE $1 = $2 LIMIT 1)`, column, needed).Scan(&data)

	return data, err
}

func (app *App) AddData(table string, needed string, column string) (sql.Result, error) {
	contains, errCont := app.CheckIfContains(table, needed, column)
	if contains || (errCont != nil && errCont != sql.ErrNoRows) {
		return nil, errCont
	}

	res, errAdd := app.db.Exec(`INSERT INTO knowledge(knowledge_pk) VALUES ($1)`, needed)

	return res, errAdd
}
