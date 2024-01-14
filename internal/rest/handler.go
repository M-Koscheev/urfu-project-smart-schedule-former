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
	router.POST("/api/v1/knowledgeCompetency/", h.PostCompetency)
	router.POST("/api/v1/competencyProfession/", h.PostCompetency)
	router.POST("/api/v1/project/", h.PostCompetency)
	router.POST("/api/v1/organization/", h.PostCompetency)
	router.POST("/api/v1/educationalProgram/", h.PostCompetency)
	router.POST("/api/v1/discipline/", h.PostCompetency)
	router.POST("/api/v1/course/", h.PostCompetency)
	router.POST("/api/v1/courseCompetency/", h.PostCompetency)
	router.POST("/api/v1/portfolio/", h.PostCompetency)
	router.POST("/api/v1/projectPortfolio/", h.PostCompetency)
	router.POST("/api/v1/projectPortfolioCompetency/", h.PostCompetency)
	router.POST("/api/v1/studyGroup/", h.PostCompetency)
	router.POST("/api/v1/student/", h.PostCompetency)
	router.POST("/api/v1/trajectory/", h.PostCompetency)

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
// @Tags         organization
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

// PostKnowledge
//
// @Summary      Post knowledge
// @Description  post single knowledge
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
		slog.Error("error adding record to the knowledge table", err)
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

// PostTechnology
//
// @Summary      Post technology
// @Description  post single technology
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
		slog.Error("error adding record to the technology table", err)
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

// PostCompetency
//
// @Summary      Post competency
// @Description  post single competency
// @Tags         competency
// @Accept       json
// @Produce      json
// @Param        competencyTitle   body      string  true  "Competency title"
// @Param        competencySkills   body      string  true  "Competency skills"
// @Param        competencyMainTechnologyId   body      string  true  "Competency main technology id"
// @Success      200  {object}  model.GetCompetencyResponse
// @Failure      400
// @Failure      500
// @Router       /api/v1/competency/ [post]
func (h *Handler) PostCompetency(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.GetCompetencyResponse
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

	resp, err := h.App.PostCompetency(req.Title, req.Skills, req.MainTechnologyId)
	if err != nil {
		slog.Error("error adding record to the competency table", err)
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

// PostKnowledgeCompetency
//
// @Summary      Post knowledge-competency connection
// @Description  post single knowledge-competency connection
// @Tags         knowledgeCompetency
// @Accept       json
// @Produce      json
// @Param        knowledgeId   body      string  true  "Knowledge id"
// @Param        competencyId   body      string  true  "Competency id"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /api/v1/knowledgeCompetency/ [post]
func (h *Handler) PostKnowledgeCompetency(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostKnowledgeCompetencyResponse
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

	err = h.App.PostKnowledgeCompetency(req.KnowledgeId, req.CompetencyId)
	if err != nil {
		slog.Error("error adding record to the knowledge_competency table", err)
		w.WriteHeader(http.StatusInternalServerError)
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
// @Param        professionTitle   body      string  true  "Profession title"
// @Success      200  {object}  model.GetProfessionResponse
// @Failure      400
// @Failure      500
// @Router       /api/v1/profession/ [post]
func (h *Handler) PostProfession(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.GetProfessionResponse
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

	resp, err := h.App.PostProfession(req.Title, req.Description)
	if err != nil {
		slog.Error("error adding record to the profession table", err)
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

// PostCompetencyProfession
//
// @Summary      Post competency-profession connection
// @Description  post single competency-profession connection
// @Tags         competencyProfession
// @Accept       json
// @Produce      json
// @Param        competencyId   body      string  true  "Competency id"
// @Param        professionId   body      string  true  "Profession id"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /api/v1/competencyProfession/ [post]
func (h *Handler) PostCompetencyProfession(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostCompetencyProfessionResponse
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

	err = h.App.PostKnowledgeCompetency(req.CompetencyId, req.ProfessionId)
	if err != nil {
		slog.Error("error adding record to the competency_profession table", err)
		w.WriteHeader(http.StatusInternalServerError)
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
// @Param        projectTitle   body      string  true  "Project title"
// @Param        projectDescription   body      string  true  "Project description"
// @Param        projectResult   body      string  true  "Project result"
// @Param        projectLifeScenario   body      string  true  "Project life scenario"
// @Param        projectMainTechnologyId   body      string  true  "Project main technology id"
// @Success      200  {object}  model.PostProjectResponse
// @Failure      400
// @Failure      500
// @Router       /api/v1/project/ [post]
func (h *Handler) PostProject(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostProjectResponse
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

	resp, err := h.App.PostProject(req.Title, req.Description, req.Result, req.LifeScenario, req.MainTechnologyId)
	if err != nil {
		slog.Error("error adding record to the project table", err)
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

// PostOrganization
//
// @Summary      Post organization
// @Description  post single organization
// @Tags         organization
// @Accept       json
// @Produce      json
// @Param        organizationTitle   body      string  true  "Organization title"
// @Success      200  {object}  model.GetOrganizationResponse
// @Failure      400
// @Failure      500
// @Router       /api/v1/organization/ [post]
func (h *Handler) PostOrganization(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.GetOrganizationResponse
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

	resp, err := h.App.PostOrganization(req.Title)
	if err != nil {
		slog.Error("error adding record to the organization table", err)
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

// PostEducationalProgram
//
// @Summary      Post educational program
// @Description  post single educational program
// @Tags         educational program
// @Accept       json
// @Produce      json
// @Param        eduacationalProgramTitle   body      string  true  "Educational program title"
// @Param        eduacationalProgramDescription   body      string  true  "Educational program description"
// @Param        eduacationalProgramOrganizationId   body      string  true  "Educational program organization id"
// @Success      200  {object}  model.PostEducationalProgramResponse
// @Failure      400
// @Failure      500
// @Router       /api/v1/educationalProgram/ [post]
func (h *Handler) PostEducationalProgram(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostEducationalProgramResponse
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

	resp, err := h.App.PostEducationalProgram(req.Title, req.Description, req.OrganizationId)
	if err != nil {
		slog.Error("error adding record to the educational program table", err)
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

// PostDiscipline
//
// @Summary      Post discipline
// @Description  post single discipline
// @Tags         discipline
// @Accept       json
// @Produce      json
// @Param        disciplineTitle   body      string  true  "Discipline title"
// @Param        disciplineDescription   body      string  true  "Discipline description"
// @Param        disciplineEducationalProgramId   body      string  true  "Discipline educational program id"
// @Success      200  {object}  model.PostDisciplineResponse
// @Failure      400
// @Failure      500
// @Router       /api/v1/discipline/ [post]
func (h *Handler) PostDiscipline(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostDisciplineResponse
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

	resp, err := h.App.PostDiscipline(req.Title, req.Description, req.EducationalProgramId)
	if err != nil {
		slog.Error("error adding record to the discipline table", err)
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

// PostCourse
//
// @Summary      Post course
// @Description  post single course
// @Tags         course
// @Accept       json
// @Produce      json
// @Param        courseTitle   body      string  true  "Course title"
// @Param        courseDescription   body      string  true  "Course description"
// @Param        courseTeacher   body      string  true  "Course teacher"
// @Param        courseDisciplineId   body      string  true  "Course discipline id"
// @Success      200  {object}  model.PostCourseResponse
// @Failure      400
// @Failure      500
// @Router       /api/v1/course/ [post]
func (h *Handler) PostCourse(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostCourseResponse
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

	resp, err := h.App.PostCourse(req.Title, req.Description, req.Teacher, req.DisciplineId)
	if err != nil {
		slog.Error("error adding record to the course table", err)
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

// PostCourseCompetency
//
// @Summary      Post course-competency connection
// @Description  post single course-competency connection
// @Tags         courseCompetency
// @Accept       json
// @Produce      json
// @Param        courseId   body      string  true  "Course id"
// @Param        competencyId   body      string  true  "Competency id"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /api/v1/courseCompetency/ [post]
func (h *Handler) PostCourseCompetency(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostCourseCompetencyResponse
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

	err = h.App.PostCourseCompetency(req.CourseId, req.CompetencyId)
	if err != nil {
		slog.Error("error adding record to the course_competency table", err)
		w.WriteHeader(http.StatusInternalServerError)
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
// @Success      200  {object}  model.PostPortfolioResponse
// @Failure      500
// @Router       /api/v1/portfolio/ [post]
func (h *Handler) PostPortfolio(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()
	resp, err := h.App.PostPortfolio()
	if err != nil {
		slog.Error("error adding record to the portfolio table", err)
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

// PostProjectPortolio
//
// @Summary      Post project-portfolio connection
// @Description  post single project-portfolio connection
// @Tags         projectPortfolio
// @Accept       json
// @Produce      json
// @Param        competencyId   body      string  true  "Project id"
// @Param        portfolioId   body      string  true  "Portfolio id"
// @Param        teamRole   body      string  true  "Team role id"
// @Param        semester   body      string  true  "semester"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /api/v1/projectPortfolio/ [post]
func (h *Handler) PostProjectPortfolioResponse(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostProjectPortfolioResponse
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

	resp, err := h.App.PostProjectPortolio(req.ProjectId, req.PortfolioId, req.TeamRole, req.Semester)
	if err != nil {
		slog.Error("error adding record to the project_portfolio table", err)
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

// PostProjectPortfolioCompetency
//
// @Summary      Post project-portfolio-competency connection
// @Description  post single project-portfolio-competency connection
// @Tags         projectPortfolioCompetency
// @Accept       json
// @Produce      json
// @Param        competencyId   body      string  true  "Project id"
// @Param        portfolioId   body      string  true  "Portfolio id"
// @Param        competencyId   body      string  true  "Competency id"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /api/v1/projectPortfolioCompetency/ [post]
func (h *Handler) PostProjectPortfolioCompetency(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostProjectPortfolioCompetencyResponse
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

	err = h.App.PostProjectPortfolioCompetency(req.ProjectId, req.PortfolioId, req.CompetencyId)
	if err != nil {
		slog.Error("error adding record to the project_portfolio_competency table", err)
		w.WriteHeader(http.StatusInternalServerError)
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
// @Param        courseId   body      string  true  "Course id"
// @Param        studentId   body      string  true  "Student id"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /api/v1/studyGroup/ [post]
func (h *Handler) PostStudyGroup(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostStudyGroupResponse
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

	err = h.App.PostStudyGroup(req.CourseId, req.StudentId)
	if err != nil {
		slog.Error("error adding record to the study_groups table", err)
		w.WriteHeader(http.StatusInternalServerError)
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
// @Param        studentFullName   body      string  true  "Student`s full name"
// @Param        studentAdmitionDate   body      string  true  "Student`s admition date"
// @Param        studentPortfolioId   body      string  true  "Student`s portfolio id"
// @Success      200 {object} model.PostStudentResponse
// @Failure      400
// @Failure      500
// @Router       /api/v1/student/ [post]
func (h *Handler) PostStudent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostStudentResponse
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

	resp, err := h.App.PostStudent(req.FullName, req.Admition, req.PortfolioId)
	if err != nil {
		slog.Error("error adding record to the students table", err)
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

// PostTrajectory
//
// @Summary      Post student`s archive course
// @Description  post single student`s archive course
// @Tags         trajectory
// @Accept       json
// @Produce      json
// @Param        studentId   body      string  true  "Student`s id"
// @Param        trajectorySemester   body      string  true  "Trajectory semester"
// @Param        courseId   body      string  true  "Course id"
// @Success      200 {object} model.PostTrajectoryResponse
// @Failure      400
// @Failure      500
// @Router       /api/v1/trajectory/ [post]
func (h *Handler) PostTrajectoryResponse(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	defer r.Body.Close()

	var req model.PostTrajectoryResponse
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

	resp, err := h.App.PostTrajectory(req.Semester, req.StudentId, req.CourseId)
	if err != nil {
		slog.Error("error adding record to the trajectory table", err)
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
