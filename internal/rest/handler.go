package rest

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/M-Koscheev/urfu-project-smart-schedule-former/internal/app"
	"github.com/julienschmidt/httprouter"
)

const QueryKeyKnowledge = "knowledge"

type Handler struct {
	App    *app.App
	Router *httprouter.Router
}

func New(app *app.App) *Handler {
	router := httprouter.New()
	h := &Handler{App: app, Router: router}
	router.GET("/knowledge", h.knowledgeGET)
	router.POST("/knowledge", h.knowledgePOST)
	return h
}

func (h *Handler) knowledgeGET(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := h.App.GetAllKnowledges()
	if err != nil {
		slog.Error("Error getting data from knowledge table", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(strings.Join(data, " ")))
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) knowledgePOST(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	reqData := r.URL.Query()
	addData := reqData.Get(QueryKeyKnowledge)
	err := h.App.AddKnowledge(addData)
	if err != nil {
		slog.Error(fmt.Sprintf("error adding element %v to the knowledge table", addData), err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
