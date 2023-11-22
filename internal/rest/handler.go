package rest

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/M-Koscheev/urfu-project-smart-schedule-former/internal/app"
)

type Handler struct {
	app *app.App
}

func New(app *app.App) *Handler {
	return &Handler{app: app}
}

func (h *Handler) Run() error {
	http.HandleFunc("/knowledge", h.knowledge)
	return http.ListenAndServe(":8080", nil)
}

func (h *Handler) knowledge(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		data, err := h.app.GetAllKnowledges()
		if err != nil {
			slog.Error("Error getting data from knowledge table", err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte(strings.Join(data, " ")))
			w.WriteHeader(http.StatusOK)
		}
	case http.MethodPost:
		reqData := r.URL.Query()
		addData := reqData.Get("knowledge")
		err := h.app.AddKnowledge(addData)
		if err != nil {
			slog.Error(fmt.Sprintf("error adding element %v to the knowledge table", addData), err)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusCreated)
		}
	default:
		slog.Warn("Now such method implemented yet", w, http.StatusMethodNotAllowed)
		w.WriteHeader(http.StatusNotImplemented)
	}
}
