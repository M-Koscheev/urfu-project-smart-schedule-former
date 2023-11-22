package cmd

import (
	"log/slog"

	"github.com/M-Koscheev/urfu-project-smart-schedule-former/internal/app"
	"github.com/M-Koscheev/urfu-project-smart-schedule-former/internal/db"
	"github.com/M-Koscheev/urfu-project-smart-schedule-former/internal/rest"
)

// cmd - control panel, so there is no program logic here/
func StartApp() {
	db, err := db.CreateConnection()
	if err != nil {
		slog.Error("unable to connect to the database", err)
	}
	defer db.Close()

	service := app.New(db)
	handler := rest.New(service)
	if err = handler.Run(); err != nil {
		slog.Error("an error occurred during the execution of the program", err)
	}
}
