package app

import (
	"database/sql"
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/M-Koscheev/urfu-project-smart-schedule-former/internal/model"
)

type App struct {
	db *sql.DB
}

func New(db *sql.DB) *App {
	return &App{db: db}
}

func (app *App) GetAllKnowledges() ([]string, error) {
	sqlKnow, err := app.db.Query(`SELECT knowledge_pk FROM knowledge`)
	if err != nil {
		return nil, err
	}

	temp := ""
	var strKnow []string
	for sqlKnow.Next() {
		if err = sqlKnow.Scan(&temp); err != nil {
			return nil, err
		}
		strKnow = append(strKnow, temp)
	}
	return strKnow, nil
}

func (app *App) GetKnowledgeByIndex(id uuid.UUID) (model.GetKnowledgeResponse, error) {
	var resp model.GetKnowledgeResponse
	data := app.db.QueryRow(`SELECT knowledge_id, title FROM knowledge WHERE knowledge_id = $1`, id)
	err := data.Scan(&resp.Id, &resp.Title)
	return resp, err
}

func (app *App) getKnowledgeByCompetency(competencyId uuid.UUID) ([]string, error) {
	var knowledge []string
	rows, err := app.db.Query(`SELECT title FROM knowledge WHERE knowledge_id in 
		(SELECT knowledge_id FROM knowledge_competency WHERE competency_id = $1)`, competencyId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var title string
		if err = rows.Scan(&title); err != nil {
			return nil, err
		}

		knowledge = append(knowledge, title)
	}

	return knowledge, nil
}

func (app *App) GetTechnolgyById(id uuid.UUID) (model.GetTechnologyResponse, error) {
	var resp model.GetTechnologyResponse
	data := app.db.QueryRow(`SELECT technology_id, title FROM technologies WHERE technology_id = $1`, id)
	err := data.Scan(&resp.Id, &resp.Title)
	return resp, err
}

func (app *App) GetCompetencyById(id uuid.UUID) (model.GetCompetencyResponse, error) {
	var resp model.GetCompetencyResponse
	row := app.db.QueryRow(`SELECT competency_id, title, skills, main_technology_id FROM competencies WHERE competency_id = $1`, id)
	var mainTechnologyId uuid.UUID
	err := row.Scan(&resp.Id, &resp.Title, &resp.Skills, &mainTechnologyId)
	if err != nil {
		return resp, err
	}

	if mainTechnologyId != uuid.Nil {
		var technology model.GetTechnologyResponse
		if technology, err = app.GetTechnolgyById(mainTechnologyId); err != nil {
			return resp, err
		}

		resp.MainTechnology = technology.Title
	}

	resp.Knowledge, err = app.getKnowledgeByCompetency(id)
	return resp, err
}

func (app *App) getCompetenciesByProfession(professionId uuid.UUID) ([]string, error) {
	var competencies []string
	rows, err := app.db.Query(`SELECT title FROM competencies WHERE competency_id in 
		(SELECT competency_id FROM competency_profession WHERE profession_id = $1)`, professionId)
	if err != nil {
		return competencies, err
	}
	defer rows.Close()

	for rows.Next() {
		var title string
		if err = rows.Scan(&title); err != nil {
			return competencies, err
		}

		competencies = append(competencies, title)
	}

	return competencies, nil
}

func (app *App) GetProfessionById(id uuid.UUID) (model.GetProfessionResponse, error) {
	var resp model.GetProfessionResponse
	data := app.db.QueryRow(`SELECT profession_id, title, description FROM professions WHERE profession_id = $1`, id)
	err := data.Scan(&resp.Id, &resp.Title, &resp.Description)
	if err != nil {
		return resp, err
	}

	resp.Competencies, err = app.getCompetenciesByProfession(id)
	return resp, err
}

func (app *App) GetProjectById(id uuid.UUID) (model.GetProjectResponse, error) {
	var resp model.GetProjectResponse
	data := app.db.QueryRow(`SELECT project_id, title, description, result, life_scenario, main_technology_id FROM projects WHERE project_id = $1`, id)
	var mainTechnologyId uuid.UUID
	err := data.Scan(&resp.Id, &resp.Title, &resp.Description, &resp.Result, &resp.LifeScenario, mainTechnologyId)
	if err != nil {
		return resp, err
	}

	if mainTechnologyId != uuid.Nil {
		var technology model.GetTechnologyResponse
		if technology, err = app.GetTechnolgyById(mainTechnologyId); err != nil {
			return resp, err
		}

		resp.MainTechnology = technology.Title
	}

	return resp, nil
}

// func (app *App) GetProjectById(id uuid.UUID) (model.GetPersonalProjectResponse, error) {
// 	var resp model.GetPersonalProjectResponse
// 	data := app.db.QueryRow(`SELECT project_id, title, description, result, life_scenario, main_technology_id FROM projects WHERE project_id = $1`, id)
// }

func (app *App) GetOrganizationById(id uuid.UUID) (model.GetOrganizationResponse, error) {
	var resp model.GetOrganizationResponse
	data := app.db.QueryRow(`SELECT organization_id, title FROM organizations WHERE organization_id = $1`, id)
	err := data.Scan(&resp.Id, &resp.Title)
	return resp, err
}

func (app *App) GetEducationalProgramById(id uuid.UUID) (model.GetEducationalProgramResponse, error) {
	var resp model.GetEducationalProgramResponse
	var organizationId uuid.UUID
	data := app.db.QueryRow(`SELECT educational_program_id, title, description, organizations_id FROM educational_programs WHERE educational_program_id = $1`, id)
	err := data.Scan(&resp.Id, &resp.Title, &resp.Description, &organizationId)
	if err != nil {
		return resp, err
	}

	var organization model.GetOrganizationResponse
	if organization, err = app.GetOrganizationById(organizationId); err != nil {
		return resp, err
	}

	resp.Organization = organization.Title
	return resp, nil
}

func (app *App) GetDisciplineById(id uuid.UUID) (model.GetDisciplineResponse, error) {
	var resp model.GetDisciplineResponse
	var educationalProgramId uuid.UUID
	data := app.db.QueryRow(`SELECT discipline_id, title, description, educational_program_id FROM disciplines WHERE discipline_id = $1`, id)
	err := data.Scan(&resp.Id, &resp.Title, &resp.Description, &educationalProgramId)
	if err != nil {
		return resp, err
	}

	var educationalProgram model.GetEducationalProgramResponse
	if educationalProgram, err = app.GetEducationalProgramById(educationalProgramId); err != nil {
		return resp, err
	}

	resp.EducationalProgram = educationalProgram.Title
	return resp, nil
}

func (app *App) getCompetenciesByCourse(courseId uuid.UUID) ([]string, error) {
	var competencies []string
	rows, err := app.db.Query(`SELECT title FROM competencies WHERE competency_id in 
		(SELECT competency_id FROM course_competencies WHERE course_id = $1)`, courseId)
	if err != nil {
		return competencies, err
	}
	defer rows.Close()

	for rows.Next() {
		var title string
		if err = rows.Scan(&title); err != nil {
			return competencies, err
		}

		competencies = append(competencies, title)
	}

	return competencies, nil
}

func (app *App) GetCourseById(id uuid.UUID) (model.GetCourseResponse, error) {
	var resp model.GetCourseResponse
	var disciplineId uuid.UUID
	data := app.db.QueryRow(`SELECT course_id, title, description, teacher, discipline_id FROM disciplines WHERE discipline_id = $1`, id)
	err := data.Scan(&resp.Id, &resp.Title, &resp.Description, &resp.Teacher, &disciplineId)
	if err != nil {
		return resp, nil
	}

	if disciplineId != uuid.Nil {
		var discipline model.GetDisciplineResponse
		if discipline, err = app.GetDisciplineById(disciplineId); err != nil {
			return resp, err
		}

		resp.Discipline = discipline.Title
	}

	resp.Competencies, err = app.getCompetenciesByCourse(id)
	return resp, err
}

func (app *App) getCompetenciesByPersonalProject(portfolioId uuid.UUID, projectId uuid.UUID) ([]string, error) {
	var resp []string
	rows, err := app.db.Query(`SELECT title FROM competencies WHERE competency_id in 
		(SELECT competency_id FROM project_portfolio_competencies WHERE portfolio_id = $1 and project_id = $2)`, portfolioId, projectId)
	if err != nil {
		return resp, nil
	}

	var competency string
	for rows.Next() {
		if err = rows.Scan(&competency); err != nil {
			return resp, err
		}

		resp = append(resp, competency)
	}

	return resp, nil
}

func (app *App) GetPersonalProjectsByPortfolio(portfolioId uuid.UUID) ([]model.GetPersonalProjectResponse, error) {
	var resp []model.GetPersonalProjectResponse
	rows, err := app.db.Query(`SELECT project_id, team_role, semester FROM project_portfolio WHERE portfolio_id = $1`, portfolioId)
	if err != nil {
		return resp, nil
	}

	for rows.Next() {
		var personlaProject model.GetPersonalProjectResponse
		var projectId uuid.UUID
		if err := rows.Scan(projectId, personlaProject.TeamRole, personlaProject.Semester); err != nil {
			return resp, err
		}

		project, err := app.GetProjectById(projectId)
		if err != nil {
			return resp, err
		}

		personlaProject.Title = project.Title
		personlaProject.Description = project.Description
		personlaProject.Result = project.Result
		personlaProject.LifeScenario = project.LifeScenario
		personlaProject.MainTechnology = project.MainTechnology
		personlaProject.Competencies, err = app.getCompetenciesByPersonalProject(portfolioId, projectId)
		if err != nil {
			return resp, nil
		}

		resp = append(resp, personlaProject)
	}

	return resp, nil
}

func (app *App) GetPortfolioById(id uuid.UUID) (model.GetPortfolioResponse, error) {
	var resp model.GetPortfolioResponse
	var err error
	resp.Id = id
	resp.Projects, err = app.GetPersonalProjectsByPortfolio(id)
	if err != nil {
		return resp, nil
	}

	return resp, nil
}

func (app *App) GetStudentById(id uuid.UUID) (model.GetStudentResponse, error) {
	var resp model.GetStudentResponse
	var portfolioId uuid.UUID
	var admition time.Time
	var err error
	data := app.db.QueryRow(`SELECT student_id, full_name, portfolio_id, admition FROM students WHERE student_id = $1`, id)
	if err = data.Scan(&resp.Id, &resp.FullName, &portfolioId, &admition); err != nil {
		return resp, err
	}
	if admition.After(time.Now()) {
		return resp, errors.New("Incorrect admition date")
	}

	resp.Semester = (uint8(time.Now().Year())-uint8(admition.Year()))*2 + 1
	if time.Now().Month() > time.January && time.Now().Month() < time.September { // approximate date
		resp.Semester += 1
	}

	resp.Portfolio, err = app.GetPortfolioById(portfolioId)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (app *App) GetStudyGroupsByStudent(studentId uuid.UUID) (model.GetStudyGroupsResponse, error) {
	var resp model.GetStudyGroupsResponse
	rows, err := app.db.Query(`SELECT title FROM courses WHERE course_id in
		(SELECT course_id FROM study_groups WHERE student_id = $1)`, studentId)
	if err != nil {
		return resp, err
	}

	var course string
	for rows.Next() {
		if err = rows.Scan(&course); err != nil {
			return resp, err
		}

		resp.Courses = append(resp.Courses, course)
	}

	return resp, nil
}

func (app *App) GetTrajectoryById(trajectoryId uuid.UUID) (model.GetTrajectoryResponse, error) {
	var resp model.GetTrajectoryResponse
	var studentId uuid.UUID
	var courseId uuid.UUID
	var err error
	data := app.db.QueryRow(`SELECT student_id, course_id, semester FROM trajectories WHERE trajectory_id = $1`, trajectoryId)
	if err = data.Scan(&studentId, &courseId, resp.Semester); err != nil {
		return resp, err
	}

	studentData := app.db.QueryRow(`SELECT full_name FROM students WHERE student_id = $1`, studentId)
	if err = studentData.Scan(&resp.Student); err != nil {
		return resp, err
	}

	courseData := app.db.QueryRow(`SELECT title FROM courses WHERE course_id = $1`, courseId)
	if err = courseData.Scan(&resp.Course); err != nil {
		return resp, err
	}

	return resp, nil
}

func (app *App) PostKnowledge(knowledge string) (uuid.UUID, error) {
	if knowledge == "" {
		return uuid.Nil, errors.New("empty title")
	}

	knowledgeId := uuid.Nil
	knowledgeData := app.db.QueryRow(`INSERT INTO knowledge (title) VALUES ($1) RETURNING knowledge_id`, knowledge)
	if err := knowledgeData.Scan(&knowledgeId); err != nil {
		return uuid.Nil, err
	}

	return knowledgeId, nil
}

func (app *App) PostTechnology(technology string) (uuid.UUID, error) {
	if technology == "" {
		return uuid.Nil, errors.New("empty title")
	}

	technologyId := uuid.Nil
	technologyData := app.db.QueryRow(`INSERT INTO technology (title) VALUES ($1) RETURNING technology_id`, technology)
	if err := technologyData.Scan(&technologyId); err != nil {
		return uuid.Nil, err
	}
	return technologyId, nil
}

// func (app *App) PostCompetency(comptency string, skills string, technologyId uuid.UUID, technology string) error {
// 	if comptency == "" {
// 		return errors.New("empty title")
// 	}

// 	if technologyId == uuid.Nil && technology != "" {
// 		tech, err := app.PostTechnology(technology)
// 		if errors.Is(err, errors.)
// 	}

// 	_, err := app.db.Exec(`INSERT INTO technology (title) VALUES ($1)`, knowledge)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
