package rest

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strings"
	"time"

	_ "github.com/M-Koscheev/urfu-project-smart-schedule-former/docs"
	"github.com/M-Koscheev/urfu-project-smart-schedule-former/internal/app"
	"github.com/M-Koscheev/urfu-project-smart-schedule-former/internal/model"
	"github.com/lib/pq"

	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
)

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
	router.POST("/api/v1/competency/", h.PostCompetency)
	router.POST("/api/v1/knowledgeCompetency/", h.PostKnowledgeCompetency)
	router.POST("/api/v1/profession/", h.PostProfession)
	router.POST("/api/v1/competencyProfession/", h.PostCompetencyProfession)
	router.POST("/api/v1/project/", h.PostProject)
	router.POST("/api/v1/organization/", h.PostOrganization)
	router.POST("/api/v1/educationalProgram/", h.PostEducationalProgram)
	router.POST("/api/v1/discipline/", h.PostDiscipline)
	router.POST("/api/v1/course/", h.PostCourse)
	router.POST("/api/v1/courseCompetency/", h.PostCourseCompetency)
	router.POST("/api/v1/portfolio/", h.PostPortfolio)
	router.POST("/api/v1/projectPortfolio/", h.PostProjectPortfolio)
	router.POST("/api/v1/projectPortfolioCompetency/", h.PostProjectPortfolioCompetency)
	router.POST("/api/v1/studyGroup/", h.PostStudyGroup)
	router.POST("/api/v1/student/", h.PostStudent)
	router.POST("/api/v1/trajectory/", h.PostTrajectory)

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
// @Success      200  {object}  model.GetKnowledge
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
// @Success      200  {object}  model.GetTechnology
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
// @Success      200  {object}  model.GetCompetency
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
// @Success      200  {object}  model.GetProfession
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
// @Success      200  {object}  model.GetProject
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
// @Tags         organization
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Organization ID"
// @Success      200  {object}  model.GetOrganization
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
// @Success      200  {object}  model.GetEducationalProgram
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
// @Success      200  {object}  model.GetDiscipline
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
// @Success      200  {object}  model.GetCourse
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
// @Success      200  {object}  model.GetPortfolio
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
// @Success      200  {object}  model.GetStudent
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
// @Success      200  {object}  model.GetTrajectory
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

// PostKnowledge
//
// @Summary      Post knowledge
// @Description  post single knowledge
// @Tags         knowledge
// @Accept       json
// @Produce      json
// @Param        input   body      model.PostKnowledge  true  "Knowledge request"
// @Success      200  {object}  model.GetKnowledge
// @Failure      400
// @Failure      500
// @Failure      502
// @Router       /api/v1/knowledge/ [post]
func (h *Handler) PostKnowledge(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostKnowledge
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			slog.Error("Bad request. Wrong Type provided for field "+unmarshalErr.Field, err)
		} else {
			slog.Error("Bad request "+err.Error(), err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.App.PostKnowledge(req.Title)
	if errors.Is(err, app.ErrEmptyTitle) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("empty title"))
		return
	}
	if err != nil {
		switch e := err.(type) {
		case *pq.Error:
			w.WriteHeader(http.StatusBadRequest)
			switch e.Code {
			case "23503":
				w.Write([]byte("foreign key violation"))
			case "23505":
				w.Write([]byte("duplicate value"))
			default:
				w.Write([]byte(e.Message))
			}
		default:
			slog.Error("unknown error", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		slog.Error("error converting data to JSON format", err)
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	w.Header().Set("content-Type", "application/json")
	w.Write(respJSON)
}

// PostTechnology
//
// @Summary      Post technology
// @Description  post single technology
// @Tags         technology
// @Accept       json
// @Produce      json
// @Param        input   body      model.PostTechnology  true  "Technology request"
// @Success      200  {object}  model.GetTechnology
// @Failure      400
// @Failure      500
// @Failure      502
// @Router       /api/v1/technology/ [post]
func (h *Handler) PostTechnology(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostTechnology
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			slog.Error("Bad request. Wrong Type provided for field "+unmarshalErr.Field, err)
		} else {
			slog.Error("Bad request "+err.Error(), err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.App.PostTechnology(req.Title)
	if errors.Is(err, app.ErrEmptyTitle) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("empty title"))
		return
	}
	if err != nil {
		switch e := err.(type) {
		case *pq.Error:
			w.WriteHeader(http.StatusBadRequest)
			switch e.Code {
			case "23503":
				w.Write([]byte("foreign key violation"))
			case "23505":
				w.Write([]byte("duplicate value"))
			default:
				w.Write([]byte(e.Message))
			}
		default:
			slog.Error("unknown error", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
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

// PostCompetency
//
// @Summary      Post competency
// @Description  post single competency
// @Tags         competency
// @Accept       json
// @Produce      json
// @Param        input   body      model.PostCompetency  true  "Competency request"
// @Success      200  {object}  model.GetCompetency
// @Failure      400
// @Failure      500
// @Failure      502
// @Router       /api/v1/competency/ [post]
func (h *Handler) PostCompetency(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostCompetency
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			slog.Error("Bad request. Wrong Type provided for field "+unmarshalErr.Field, err)
		} else {
			slog.Error("Bad request "+err.Error(), err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.App.PostCompetency(req.Title, req.Skills, req.MainTechnologyId)
	if errors.Is(err, app.ErrEmptyTitle) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("empty title"))
		return
	}
	if err != nil {
		switch e := err.(type) {
		case *pq.Error:
			w.WriteHeader(http.StatusBadRequest)
			switch e.Code {
			case "23503":
				w.Write([]byte("foreign key violation"))
			case "23505":
				w.Write([]byte("duplicate value"))
			default:
				w.Write([]byte(e.Message))
			}
		default:
			slog.Error("unknown error", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
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

// PostKnowledgeCompetency
//
// @Summary      Post knowledge-competency connection
// @Description  post single knowledge-competency connection
// @Tags         knowledgeCompetency
// @Accept       json
// @Produce      json
// @Param        input   body      model.PostKnowledgeCompetency  true  "Knowledge-competency request"
// @Success      200
// @Failure      400
// @Failure      500
// @Failure      502
// @Router       /api/v1/knowledgeCompetency/ [post]
func (h *Handler) PostKnowledgeCompetency(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostKnowledgeCompetency
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			slog.Error("Bad request. Wrong Type provided for field "+unmarshalErr.Field, err)
		} else {
			slog.Error("Bad request "+err.Error(), err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.App.PostKnowledgeCompetency(req.KnowledgeId, req.CompetencyId)
	if errors.Is(err, app.ErrEmptyId) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("empty id"))
		return
	}
	if err != nil {
		switch e := err.(type) {
		case *pq.Error:
			w.WriteHeader(http.StatusBadRequest)
			switch e.Code {
			case "23503":
				w.Write([]byte("foreign key violation"))
			case "23505":
				w.Write([]byte("duplicate value"))
			default:
				w.Write([]byte(e.Message))
			}
		default:
			slog.Error("unknown error", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// PostProfession
//
// @Summary      Post profession
// @Description  post single profession
// @Tags         profession
// @Accept       json
// @Produce      json
// @Param        input   body      model.PostProfession  true  "Profession data"
// @Success      200  {object}  model.GetProfession
// @Failure      400
// @Failure      500
// @Failure      502
// @Router       /api/v1/profession/ [post]
func (h *Handler) PostProfession(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostProfession
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			slog.Error("Bad request. Wrong Type provided for field "+unmarshalErr.Field, err)
		} else {
			slog.Error("Bad request "+err.Error(), err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.App.PostProfession(req.Title, req.Description)
	if errors.Is(err, app.ErrEmptyTitle) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("empty title"))
		return
	}
	if err != nil {
		switch e := err.(type) {
		case *pq.Error:
			w.WriteHeader(http.StatusBadRequest)
			switch e.Code {
			case "23503":
				w.Write([]byte("foreign key violation"))
			case "23505":
				w.Write([]byte("duplicate value"))
			default:
				w.Write([]byte(e.Message))
			}
		default:
			slog.Error("unknown error", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
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

// PostCompetencyProfession
//
// @Summary      Post competency-profession connection
// @Description  post single competency-profession connection
// @Tags         competencyProfession
// @Accept       json
// @Produce      json
// @Param        input   body      model.PostCompetencyProfession  true  "CompetencyProfession data"
// @Success      200
// @Failure      400
// @Failure      500
// @Failure      502
// @Router       /api/v1/competencyProfession/ [post]
func (h *Handler) PostCompetencyProfession(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostCompetencyProfession
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			slog.Error("Bad request. Wrong Type provided for field "+unmarshalErr.Field, err)
		} else {
			slog.Error("Bad request "+err.Error(), err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.App.PostCompetencyProfession(req.CompetencyId, req.ProfessionId)
	if errors.Is(err, app.ErrEmptyId) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("empty id"))
		return
	}
	if err != nil {
		switch e := err.(type) {
		case *pq.Error:
			w.WriteHeader(http.StatusBadRequest)
			switch e.Code {
			case "23503":
				w.Write([]byte("foreign key violation"))
			case "23505":
				w.Write([]byte("duplicate value"))
			default:
				w.Write([]byte(e.Message))
			}
		default:
			slog.Error("unknown error", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// PostProject
//
// @Summary      Post project
// @Description  post single project
// @Tags         project
// @Accept       json
// @Produce      json
// @Param        input   body      model.PostProject  true  "Project data"
// @Success      200  {object}  model.GetProject
// @Failure      400
// @Failure      500
// @Failure      502
// @Router       /api/v1/project/ [post]
func (h *Handler) PostProject(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostProject
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			slog.Error("Bad request. Wrong Type provided for field "+unmarshalErr.Field, err)
		} else {
			slog.Error("Bad request "+err.Error(), err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.App.PostProject(req.Title, req.Description, req.Result, req.LifeScenario, req.MainTechnologyId)
	if errors.Is(err, app.ErrEmptyTitle) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("empty title"))
		return
	}
	if err != nil {
		switch e := err.(type) {
		case *pq.Error:
			w.WriteHeader(http.StatusBadRequest)
			switch e.Code {
			case "23503":
				w.Write([]byte("foreign key violation"))
			case "23505":
				w.Write([]byte("duplicate value"))
			default:
				w.Write([]byte(e.Message))
			}
		default:
			slog.Error("unknown error", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
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

// PostOrganization
//
// @Summary      Post organization
// @Description  post single organization
// @Tags         organization
// @Accept       json
// @Produce      json
// @Param        input   body      model.PostOrganization  true  "Organization data"
// @Success      200  {object}  model.GetOrganization
// @Failure      400
// @Failure      500
// @Failure      502
// @Router       /api/v1/organization/ [post]
func (h *Handler) PostOrganization(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.GetOrganization
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			slog.Error("Bad request. Wrong Type provided for field "+unmarshalErr.Field, err)
		} else {
			slog.Error("Bad request "+err.Error(), err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.App.PostOrganization(req.Title)
	if errors.Is(err, app.ErrEmptyTitle) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("empty title"))
		return
	}
	if err != nil {
		switch e := err.(type) {
		case *pq.Error:
			w.WriteHeader(http.StatusBadRequest)
			switch e.Code {
			case "23503":
				w.Write([]byte("foreign key violation"))
			case "23505":
				w.Write([]byte("duplicate value"))
			default:
				w.Write([]byte(e.Message))
			}
		default:
			slog.Error("unknown error", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
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

// PostEducationalProgram
//
// @Summary      Post educational program
// @Description  post single educational program
// @Tags         educational program
// @Accept       json
// @Produce      json
// @Param        input   body      model.PostEducationalProgram  true  "Educational program data"
// @Success      200  {object}  model.GetEducationalProgram
// @Failure      400
// @Failure      500
// @Failure      502
// @Router       /api/v1/educationalProgram/ [post]
func (h *Handler) PostEducationalProgram(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostEducationalProgram
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			slog.Error("Bad request. Wrong Type provided for field "+unmarshalErr.Field, err)
		} else {
			slog.Error("Bad request "+err.Error(), err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.App.PostEducationalProgram(req.Title, req.Description, req.OrganizationId)
	if errors.Is(err, app.ErrEmptyTitle) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("empty title"))
		return
	}
	if err != nil {
		switch e := err.(type) {
		case *pq.Error:
			w.WriteHeader(http.StatusBadRequest)
			switch e.Code {
			case "23503":
				w.Write([]byte("foreign key violation"))
			case "23505":
				w.Write([]byte("duplicate value"))
			default:
				w.Write([]byte(e.Message))
			}
		default:
			slog.Error("unknown error", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
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

// PostDiscipline
//
// @Summary      Post discipline
// @Description  post single discipline
// @Tags         discipline
// @Accept       json
// @Produce      json
// @Param        input   body      model.PostDiscipline  true  "Discipline data"
// @Success      200  {object}  model.GetDiscipline
// @Failure      400
// @Failure      500
// @Failure      502
// @Router       /api/v1/discipline/ [post]
func (h *Handler) PostDiscipline(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostDiscipline
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			slog.Error("Bad request. Wrong Type provided for field "+unmarshalErr.Field, err)
		} else {
			slog.Error("Bad request "+err.Error(), err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.App.PostDiscipline(req.Title, req.Description, req.EducationalProgramId)
	if errors.Is(err, app.ErrEmptyTitle) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("empty title"))
		return
	}
	if err != nil {
		switch e := err.(type) {
		case *pq.Error:
			w.WriteHeader(http.StatusBadRequest)
			switch e.Code {
			case "23503":
				w.Write([]byte("foreign key violation"))
			case "23505":
				w.Write([]byte("duplicate value"))
			default:
				w.Write([]byte(e.Message))
			}
		default:
			slog.Error("unknown error", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
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

// PostCourse
//
// @Summary      Post course
// @Description  post single course
// @Tags         course
// @Accept       json
// @Produce      json
// @Param        input   body      model.PostCourse  true  "Course data"
// @Success      200  {object}  model.GetCourse
// @Failure      400
// @Failure      500
// @Failure      502
// @Router       /api/v1/course/ [post]
func (h *Handler) PostCourse(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostCourse
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			slog.Error("Bad request. Wrong Type provided for field "+unmarshalErr.Field, err)
		} else {
			slog.Error("Bad request "+err.Error(), err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.App.PostCourse(req.Title, req.Description, req.Teacher, req.DisciplineId)
	if errors.Is(err, app.ErrEmptyTitle) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("empty title"))
		return
	}
	if err != nil {
		switch e := err.(type) {
		case *pq.Error:
			w.WriteHeader(http.StatusBadRequest)
			switch e.Code {
			case "23503":
				w.Write([]byte("foreign key violation"))
			case "23505":
				w.Write([]byte("duplicate value"))
			default:
				w.Write([]byte(e.Message))
			}
		default:
			slog.Error("unknown error", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
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

// PostCourseCompetency
//
// @Summary      Post course-competency connection
// @Description  post single course-competency connection
// @Tags         courseCompetency
// @Accept       json
// @Produce      json
// @Param        input   body      model.PostCourseCompetency  true  "Course-competency data"
// @Success      200
// @Failure      400
// @Failure      500
// @Failure      502
// @Router       /api/v1/courseCompetency/ [post]
func (h *Handler) PostCourseCompetency(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostCourseCompetency
	var unmarshalErr *json.UnmarshalTypeError
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			slog.Error("Bad request. Wrong Type provided for field "+unmarshalErr.Field, err)
		} else {
			slog.Error("Bad request "+err.Error(), err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.App.PostCourseCompetency(req.CourseId, req.CompetencyId)
	if errors.Is(err, app.ErrEmptyId) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("empty id"))
		return
	}
	if err != nil {
		switch e := err.(type) {
		case *pq.Error:
			w.WriteHeader(http.StatusBadRequest)
			switch e.Code {
			case "23503":
				w.Write([]byte("foreign key violation"))
			case "23505":
				w.Write([]byte("duplicate value"))
			default:
				w.Write([]byte(e.Message))
			}
		default:
			slog.Error("unknown error", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// PostPortfolio
//
// @Summary      Post portfolio
// @Description  post single portfolio
// @Tags         portfolio
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.PostPortfolio
// @Failure      500
// @Failure      502
// @Router       /api/v1/portfolio/ [post]
func (h *Handler) PostPortfolio(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()
	resp, err := h.App.PostPortfolio()
	if err != nil {
		switch e := err.(type) {
		case *pq.Error:
			w.WriteHeader(http.StatusBadRequest)
			switch e.Code {
			case "23503":
				w.Write([]byte("foreign key violation"))
			case "23505":
				w.Write([]byte("duplicate value"))
			default:
				w.Write([]byte(e.Message))
			}
		default:
			slog.Error("unknown error", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
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

// PostProjectPortolio
//
// @Summary      Post project-portfolio connection
// @Description  post single project-portfolio connection
// @Tags         projectPortfolio
// @Accept       json
// @Produce      json
// @Param        input   body      model.PostProjectPortfolio  true  "Personal project data"
// @Success      200
// @Failure      400
// @Failure      500
// @Failure      502
// @Router       /api/v1/projectPortfolio/ [post]
func (h *Handler) PostProjectPortfolio(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostProjectPortfolio
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			slog.Error("Bad request. Wrong Type provided for field "+unmarshalErr.Field, err)
		} else {
			slog.Error("Bad request "+err.Error(), err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.App.PostProjectPortolio(req.ProjectId, req.PortfolioId, req.TeamRole, req.Semester)
	if errors.Is(err, app.ErrEmptyId) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("empty id"))
		return
	}
	if err != nil {
		switch e := err.(type) {
		case *pq.Error:
			w.WriteHeader(http.StatusBadRequest)
			switch e.Code {
			case "23503":
				w.Write([]byte("foreign key violation"))
			case "23505":
				w.Write([]byte("duplicate value"))
			default:
				w.Write([]byte(e.Message))
			}
		default:
			slog.Error("unknown error", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
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

// PostProjectPortfolioCompetency
//
// @Summary      Post project-portfolio-competency connection
// @Description  post single project-portfolio-competency connection
// @Tags         projectPortfolioCompetency
// @Accept       json
// @Produce      json
// @Param        input   body      model.PostProjectPortfolioCompetency  true  "Personal project competency"
// @Success      200
// @Failure      400
// @Failure      500
// @Failure      502
// @Router       /api/v1/projectPortfolioCompetency/ [post]
func (h *Handler) PostProjectPortfolioCompetency(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostProjectPortfolioCompetency
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			slog.Error("Bad request. Wrong Type provided for field "+unmarshalErr.Field, err)
		} else {
			slog.Error("Bad request "+err.Error(), err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.App.PostProjectPortfolioCompetency(req.ProjectId, req.PortfolioId, req.CompetencyId)
	if errors.Is(err, app.ErrEmptyId) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("empty id"))
		return
	}
	if err != nil {
		switch e := err.(type) {
		case *pq.Error:
			w.WriteHeader(http.StatusBadRequest)
			switch e.Code {
			case "23503":
				w.Write([]byte("foreign key violation"))
			case "23505":
				w.Write([]byte("duplicate value"))
			default:
				w.Write([]byte(e.Message))
			}
		default:
			slog.Error("unknown error", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// PostStudyGroup
//
// @Summary      Post student`s course in current semester
// @Description  Student`s course in current semester
// @Tags         studyGroup
// @Accept       json
// @Produce      json
// @Param        input   body      model.PostStudyGroup  true  "Personal current student`s project"
// @Success      200
// @Failure      400
// @Failure      500
// @Failure      502
// @Router       /api/v1/studyGroup/ [post]
func (h *Handler) PostStudyGroup(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostStudyGroup
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			slog.Error("Bad request. Wrong Type provided for field "+unmarshalErr.Field, err)
		} else {
			slog.Error("Bad request "+err.Error(), err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.App.PostStudyGroup(req.CourseId, req.StudentId)
	if errors.Is(err, app.ErrEmptyId) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("empty id"))
		return
	}
	if err != nil {
		switch e := err.(type) {
		case *pq.Error:
			w.WriteHeader(http.StatusBadRequest)
			switch e.Code {
			case "23503":
				w.Write([]byte("foreign key violation"))
			case "23505":
				w.Write([]byte("duplicate value"))
			default:
				w.Write([]byte(e.Message))
			}
		default:
			slog.Error("unknown error", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// PostStudent
//
// @Summary      Post student
// @Description  post single student
// @Tags         student
// @Accept       json
// @Produce      json
// @Param        input   body      model.PostStudent  true  "Student`s data"
// @Success      200 {object} model.GetStudent
// @Failure      400
// @Failure      500
// @Failure      502
// @Router       /api/v1/student/ [post]
func (h *Handler) PostStudent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostStudent
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			slog.Error("Bad request. Wrong Type provided for field "+unmarshalErr.Field, err)
		} else {
			slog.Error("Bad request "+err.Error(), err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes, err := req.Admition.MarshalJSON()
	if err != nil {
		slog.Error("date marshalling error " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	dateString := strings.Trim(string(bytes), "\"")
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		slog.Error("date parsing error " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	resp, err := h.App.PostStudent(req.FullName, date, req.PortfolioId)
	if errors.Is(err, app.ErrEmptyTitle) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("empty student`s name"))
		return
	}
	if err != nil {
		switch e := err.(type) {
		case *pq.Error:
			w.WriteHeader(http.StatusBadRequest)
			switch e.Code {
			case "23503":
				w.Write([]byte("foreign key violation"))
			case "23505":
				w.Write([]byte("duplicate value"))
			default:
				w.Write([]byte(e.Message))
			}
		default:
			slog.Error("unknown error", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
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

// PostTrajectory
//
// @Summary      Post student`s archive course
// @Description  post single student`s archive course
// @Tags         trajectory
// @Accept       json
// @Produce      json
// @Param        input   body      model.PostTrajectory  true  "Trajectory`s data"
// @Success      200 {object} model.GetTrajectory
// @Failure      400
// @Failure      500
// @Failure      502
// @Router       /api/v1/trajectory/ [post]
func (h *Handler) PostTrajectory(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostTrajectory
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			slog.Error("Bad request. Wrong Type provided for field "+unmarshalErr.Field, err)
		} else {
			slog.Error("Bad request "+err.Error(), err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := h.App.PostTrajectory(req.Semester, req.StudentId, req.CourseId)
	if errors.Is(err, app.ErrEmptyId) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("empty id"))
		return
	}
	if err != nil {
		switch e := err.(type) {
		case *pq.Error:
			w.WriteHeader(http.StatusBadRequest)
			switch e.Code {
			case "23503":
				w.Write([]byte("foreign key violation"))
			case "23505":
				w.Write([]byte("duplicate value"))
			default:
				w.Write([]byte(e.Message))
			}
		default:
			slog.Error("unknown error", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
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
