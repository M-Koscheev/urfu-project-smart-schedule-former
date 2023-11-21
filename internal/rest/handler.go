package rest

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/M-Koscheev/urfu-project-smart-schedule-former/internal/app"
)

type Handler struct {
	app *app.App
}

func New(app *app.App) *Handler {
	return &Handler{app: app}
}

func (h *Handler) Run() error {
	http.HandleFunc("/", h.greet)
	err := http.ListenAndServe(":8080", nil)
	return err
}

func (h *Handler) greet(w http.ResponseWriter, r *http.Request) {
	value := "15"
	switch r.Method {
	case http.MethodGet:
		tableTitle := r.URL.Query().Get("table")
		fmt.Println(tableTitle)
		if tableTitle == "knowledge" {
			data, err := h.app.GetKnowledges()
			fmt.Println(data)
			if err != nil {
				slog.Error("Error getting data from knowledge table", err)
			}
			http.ServeFile(w, r, data)
			w.Write([]byte(data))
		}
		return
	case http.MethodPut:
		err := h.app.AddKnowledge(value)
		if err != nil {
			slog.Error(fmt.Sprintf("error adding element %v to the knowledge table", value), err)
		}
		return
	default:
		slog.Error("Now such method implemented yet", w, http.StatusMethodNotAllowed)
	}
}
