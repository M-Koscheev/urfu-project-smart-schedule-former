package cmd

import (
	"log/slog"

	"github.com/M-Koscheev/urfu-project-smart-schedule-former/internal/app"
	"github.com/M-Koscheev/urfu-project-smart-schedule-former/internal/db"
)

// cmd - control panel, so there is no program logic here/
func StartApp() {
	db, err := db.CreateConnection()
	if err != nil {
		slog.Error("unable to connect to the database", err)
	}
	defer db.Close()

	service := app.New(db)
	if err := service.Run(); err != nil {
		slog.Error("unable to start app", slog.String("err", err.Error()))
	}
}
