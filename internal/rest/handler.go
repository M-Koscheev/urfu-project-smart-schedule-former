package rest

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	_ "github.com/M-Koscheev/urfu-project-smart-schedule-former/docs"
	"github.com/M-Koscheev/urfu-project-smart-schedule-former/internal/app"
	"github.com/M-Koscheev/urfu-project-smart-schedule-former/internal/model"

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
	router.ServeFiles("/docs/*filepath", http.Dir("docs"))

	router.GET("/api/v1/knowledge/:id", h.GetKnowledge)
	router.GET("/api/v1/technology/:id", h.GetTechnology)
	router.GET("/api/v1/competency/:id", h.GetCompetency)
	router.GET("/api/v1/profession/:id", h.GetProfession)
	router.GET("/api/v1/project/:id", h.GetProject)
	router.GET("/api/v1/organization/:id", h.GetOrganization)
	router.GET("/api/v1/educationalProgram/:id", h.GetEducationalProgram)
	router.GET("/api/v1/discipline/:id", h.GetDiscipline)
	router.GET("/api/v1/course/:id", h.GetCourse)
	router.GET("/api/v1/portfolio/:id", h.GetPortfolio)
	router.GET("/api/v1/student/:id", h.GetStudent)
	router.GET("/api/v1/trajectory/:id", h.GetTrajectory)

	router.POST("/api/v1/knowledge/", h.PostKnowledge)
	router.POST("/api/v1/technology/", h.PostTechnology)
	// router.POST("/api/competencies/v1/", h.PostCompetency)

	return h
}

// GetKnowledge return knowledge by it`s id
//
// @Summary      Show knowledge
// @Description  get single knowledge by ID
// @Tags         knowledge
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Knowledge ID"
// @Success      200  {object}  model.GetKnowledgeResponse
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /api/v1/knowledge/{id} [get]
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

// GetTechnology return technology by it`s id
//
// @Summary      Show technology
// @Description  get single technology by ID
// @Tags         technology
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Technology ID"
// @Success      200  {object}  model.GetTechnologyResponse
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /api/v1/technology/{id} [get]
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

// GetCompetency return competency by it`s id
//
// @Summary      Show competency
// @Description  get single competency by ID
// @Tags         competency
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Competency ID"
// @Success      200  {object}  model.GetCompetencyResponse
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /api/v1/competency/{id} [get]
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

// GetProfession return profession by it`s id
//
// @Summary      Show profession
// @Description  get single profession by ID
// @Tags         profession
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Profession ID"
// @Success      200  {object}  model.GetProfessionResponse
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /api/v1/profession/{id} [get]
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

// GetProject return project by it`s id
//
// @Summary      Show project
// @Description  get single project by ID
// @Tags         project
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Project ID"
// @Success      200  {object}  model.GetProjectResponse
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /api/v1/project/{id} [get]
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

// GetOrganization return organization by it`s id
//
// @Summary      Show organization
// @Description  get single organization by ID
// @Tags         project
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Organization ID"
// @Success      200  {object}  model.GetOrganizationResponse
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /api/v1/organization/{id} [get]
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

// GetEducationalProgram return educational program by it`s id
//
// @Summary      Show educational program
// @Description  get single educational program by ID
// @Tags         educational program
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Educational program ID"
// @Success      200  {object}  model.GetEducationalProgramResponse
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /api/v1/educationalProgram/{id} [get]
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

// GetDiscipline return discipline by it`s id
//
// @Summary      Show discipline
// @Description  get single discipline by ID
// @Tags         discipline
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Discipline ID"
// @Success      200  {object}  model.GetDisciplineResponse
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /api/v1/discipline/{id} [get]
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

// GetCourse return course by it`s id
//
// @Summary      Show course
// @Description  get single course by ID
// @Tags         course
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Course ID"
// @Success      200  {object}  model.GetCourseResponse
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /api/v1/course/{id} [get]
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

// GetPortfolio return portfolio by it`s id
//
// @Summary      Show portfolio
// @Description  get single portfolio by ID
// @Tags         portfolio
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Portfolio ID"
// @Success      200  {object}  model.GetPortfolioResponse
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /api/v1/portfolio/{id} [get]
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

// GetStudent return student by it`s id
//
// @Summary      Show student
// @Description  get single student by ID
// @Tags         student
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Student ID"
// @Success      200  {object}  model.GetStudentResponse
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /api/v1/student/{id} [get]
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

// GetTrajectory return trajectory by it`s id
//
// @Summary      Show trajectory
// @Description  get single trajectory by ID
// @Tags         trajectory
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Trajectory ID"
// @Success      200  {object}  model.GetTrajectoryResponse
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /api/v1/trajectory/{id} [get]
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

// PostKnowledge return knowledge by it`s id
//
// @Summary      Post knowledge
// @Description  post single knowledge (or check if it is already exists)
// @Tags         knowledge
// @Accept       json
// @Produce      json
// @Param        knowledgeTitle   body      string  true  "Knowledge title"
// @Success      200  {object}  model.GetKnowledgeResponse
// @Failure      400
// @Failure      500
// @Router       /api/v1/knowledge/ [post]
func (h *Handler) PostKnowledge(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostKnowledgeRequest
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			slog.Error("Bad Request. Wrong Type provided for field "+unmarshalErr.Field, err)
		} else {
			slog.Error("Bad Request "+err.Error(), err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.App.PostKnowledge(req.Title)
	if err != nil {
		slog.Error("error adding value to the knowledge table", err)
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
	w.WriteHeader(http.StatusOK)
}

// PostTechnology return technology by it`s id
//
// @Summary      Post technology
// @Description  post single technology (or check if it is already exists)
// @Tags         technology
// @Accept       json
// @Produce      json
// @Param        technologyTitle   body      string  true  "Technology title"
// @Success      200  {object}  model.GetTechnologyResponse
// @Failure      400
// @Failure      500
// @Router       /api/v1/technology/ [post]
func (h *Handler) PostTechnology(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostTechnologyRequest
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			slog.Error("Bad Request. Wrong Type provided for field "+unmarshalErr.Field, err)
		} else {
			slog.Error("Bad Request "+err.Error(), err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.App.PostTechnology(req.Title)
	if err != nil {
		slog.Error("error adding value to the technology table", err)
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
