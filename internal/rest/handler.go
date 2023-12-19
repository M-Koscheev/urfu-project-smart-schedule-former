package rest

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/M-Koscheev/urfu-project-smart-schedule-former/internal/app"
	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
)

const QueryKeyKnowledge = "knowledge"

type Handler struct {
	App    *app.App
	Router *httprouter.Router
}

func New(app *app.App) *Handler {
	router := httprouter.New()
	h := &Handler{App: app, Router: router}
	router.GET("/api/knowledge/v1/:id", h.GetKnowledge)
	router.GET("/api/technology/v1/:id", h.GetTechnology)
	router.GET("/api/competency/v1/:id", h.GetCompetency)
	router.GET("/api/profession/v1/:id", h.GetProfession)
	router.GET("/api/project/v1/:id", h.GetProject)
	router.GET("/api/organization/v1/:id", h.GetOrganization)
	router.GET("/api/educationalProgram/v1/:id", h.GetEducationalProgram)
	router.GET("/api/discipline/v1/:id", h.GetDiscipline)
	router.GET("/api/course/v1/:id", h.GetCourse)
	router.GET("/api/portfolio/v1/:id", h.GetPortfolio)
	router.GET("/api/student/v1/:id", h.GetStudent)
	router.GET("/api/trajectory/v1/:id", h.GetTrajectory)

	router.POST("/api/knowledge/v1/", h.PostKnowledge)
	router.POST("/api/technology/v1/", h.PostTechnology)
	// router.POST("/api/competencies/v1/", h.PostCompetency)

	return h
}

func (h *Handler) GetKnowledge(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := uuid.FromString(params.ByName("id"))
	if err != nil {
		slog.Error("wrong id format", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.App.GetKnowledgeByIndex(id)
	if errors.Is(err, sql.ErrNoRows) {
		slog.Info("no row with such id was found", err)
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		slog.Info("error getting knowledge by id", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		slog.Error("error converting data to JSON format", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-Type", "application/json")
	w.Write(respJSON)
}

func (h *Handler) GetTechnology(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := uuid.FromString(params.ByName("id"))
	if err != nil {
		slog.Error("wrong id format", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.App.GetTechnolgyById(id)
	if errors.Is(err, sql.ErrNoRows) {
		slog.Info("no row with such id was found", err)
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		slog.Info("error getting technology by id", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		slog.Error("error converting data to JSON format", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-Type", "application/json")
	w.Write(respJSON)
}

func (h *Handler) GetCompetency(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := uuid.FromString(params.ByName("id"))
	if err != nil {
		slog.Error("wrong id format", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.App.GetCompetencyById(id)
	if errors.Is(err, sql.ErrNoRows) {
		slog.Info("no row with such id was found", err)
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		slog.Info("error getting competency by id", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		slog.Error("error converting data to JSON format", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-Type", "application/json")
	w.Write(respJSON)
}

func (h *Handler) GetProfession(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := uuid.FromString(params.ByName("id"))
	if err != nil {
		slog.Error("wrong id format", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.App.GetProfessionById(id)
	if errors.Is(err, sql.ErrNoRows) {
		slog.Info("no row with such id was found", err)
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		slog.Info("error getting profession by id", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		slog.Error("error converting data to JSON format", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-Type", "application/json")
	w.Write(respJSON)
}

func (h *Handler) GetProject(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := uuid.FromString(params.ByName("id"))
	if err != nil {
		slog.Error("wrong id format", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.App.GetProjectById(id)
	if errors.Is(err, sql.ErrNoRows) {
		slog.Info("no row with such id was found", err)
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		slog.Info("error getting project by id", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		slog.Error("error converting data to JSON format", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-Type", "application/json")
	w.Write(respJSON)
}

func (h *Handler) GetOrganization(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := uuid.FromString(params.ByName("id"))
	if err != nil {
		slog.Error("wrong id format", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.App.GetOrganizationById(id)
	if errors.Is(err, sql.ErrNoRows) {
		slog.Info("no row with such id was found", err)
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		slog.Info("error getting project by id", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		slog.Error("error converting data to JSON format", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-Type", "application/json")
	w.Write(respJSON)
}

func (h *Handler) GetEducationalProgram(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := uuid.FromString(params.ByName("id"))
	if err != nil {
		slog.Error("wrong id format", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.App.GetEducationalProgramById(id)
	if errors.Is(err, sql.ErrNoRows) {
		slog.Info("no row with such id was found", err)
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		slog.Info("error getting project by id", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		slog.Error("error converting data to JSON format", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-Type", "application/json")
	w.Write(respJSON)
}

func (h *Handler) GetDiscipline(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := uuid.FromString(params.ByName("id"))
	if err != nil {
		slog.Error("wrong id format", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.App.GetDisciplineById(id)
	if errors.Is(err, sql.ErrNoRows) {
		slog.Info("no row with such id was found", err)
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		slog.Info("error getting project by id", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		slog.Error("error converting data to JSON format", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-Type", "application/json")
	w.Write(respJSON)
}

func (h *Handler) GetCourse(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := uuid.FromString(params.ByName("id"))
	if err != nil {
		slog.Error("wrong id format", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.App.GetCourseById(id)
	if errors.Is(err, sql.ErrNoRows) {
		slog.Info("no row with such id was found", err)
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		slog.Info("error getting project by id", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		slog.Error("error converting data to JSON format", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-Type", "application/json")
	w.Write(respJSON)
}

func (h *Handler) GetPortfolio(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := uuid.FromString(params.ByName("id"))
	if err != nil {
		slog.Error("wrong id format", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.App.GetPortfolioById(id)
	if errors.Is(err, sql.ErrNoRows) {
		slog.Info("no row with such id was found", err)
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		slog.Info("error getting project by id", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		slog.Error("error converting data to JSON format", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-Type", "application/json")
	w.Write(respJSON)
}

func (h *Handler) GetStudent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := uuid.FromString(params.ByName("id"))
	if err != nil {
		slog.Error("wrong id format", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.App.GetStudentById(id)
	if errors.Is(err, sql.ErrNoRows) {
		slog.Info("no row with such id was found", err)
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		slog.Info("error getting project by id", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		slog.Error("error converting data to JSON format", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-Type", "application/json")
	w.Write(respJSON)
}

func (h *Handler) GetTrajectory(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := uuid.FromString(params.ByName("id"))
	if err != nil {
		slog.Error("wrong id format", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.App.GetTrajectoryById(id)
	if errors.Is(err, sql.ErrNoRows) {
		slog.Info("no row with such id was found", err)
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		slog.Info("error getting project by id", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		slog.Error("error converting data to JSON format", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-Type", "application/json")
	w.Write(respJSON)
}

func (h *Handler) PostKnowledge(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	fmt.Println(string(data))
	title := params.ByName("title")
	_, err = h.App.PostKnowledge(title)
	if err != nil {
		slog.Error("error adding value to the knowledge table", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) PostTechnology(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	fmt.Println(data)
	title := params.ByName("title")
	_, err = h.App.PostTechnology(title)
	if err != nil {
		slog.Error("error adding value to the technology table", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// func (h *Handler) PostCompetency(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
// 	title := params.ByName("title")
// 	skills := params.ByName("skills")
// 	technology := params.ByName("technology")

// 	err := h.App.PostCompetency(title, skills, technology)
// 	if err != nil {
// 		slog.Error("error adding value to the technology table", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// }
