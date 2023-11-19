package app

import (
	"database/sql"
	"log/slog"
	"net/http"
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
	// test data
	err := app.AddKnowledge("14")
	if err != nil {
		slog.Error("error adding element 14 to the knowledge table", err)
	}
}
