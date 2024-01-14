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

		resp.MainTechnologyId = technology.Id
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
		return resp, errors.New("incorrect admition date")
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

func (app *App) PostKnowledge(knowledge string) (model.GetKnowledgeResponse, error) {
	var resp model.GetKnowledgeResponse
	if knowledge == "" {
		return resp, errors.New("empty title")
	}

	resp.Id = uuid.Nil
	resp.Title = knowledge
	knowledgeData := app.db.QueryRow(`INSERT INTO knowledge (title) VALUES ($1) ON CONFLICT (title) DO UPDATE SET title = excluded.title RETURNING knowledge_id`, knowledge)
	if err := knowledgeData.Scan(&resp.Id); err != nil {
		return resp, err
	}

	return resp, nil
}

func (app *App) PostTechnology(technology string) (model.GetTechnologyResponse, error) {
	var resp model.GetTechnologyResponse
	if technology == "" {
		return resp, errors.New("empty title")
	}

	resp.Id = uuid.Nil
	resp.Title = technology
	technologyData := app.db.QueryRow(`INSERT INTO technology (title) VALUES ($1) ON CONFLICT (title) DO UPDATE SET title = excluded.title RETURNING technology_id `, technology)
	if err := technologyData.Scan(&resp.Id); err != nil {
		return resp, err
	}
	return resp, nil
}

//

func (app *App) PostCompetency(comptency string, skills string, technologyId uuid.UUID) (model.GetCompetencyResponse, error) {
	var resp model.GetCompetencyResponse
	if comptency == "" {
		return resp, errors.New("empty title")
	}

	resp.MainTechnologyId = technologyId
	resp.Title = comptency
	resp.Skills = skills
	resp.Id = uuid.Nil
	competencyData := app.db.QueryRow(`INSERT INTO competencies (title, skills, main_technology_id) VALUES ($1, $2, $3)
								ON CONFLICT (title) DO UPDATE SET title = excluded.title RETURNING competency_id`, comptency, skills, technologyId)
	if err := competencyData.Scan(&resp.Id); err != nil {
		return resp, err
	}

	return resp, nil
}

func (app *App) PostKnowledgeCompetency(knowledgeId uuid.UUID, competencyId uuid.UUID) error {
	if knowledgeId == uuid.Nil {
		return errors.New("empty knowledge id")
	}
	if competencyId == uuid.Nil {
		return errors.New("empty competency id")
	}

	_, err := app.db.Exec(`INSERT INTO knowledge_competency (knowledge_id, competency_id) VALUES ($1, $2)`, knowledgeId, competencyId)
	return err
}

func (app *App) PostProfession(profession string, description string) (model.GetProfessionResponse, error) {
	var resp model.GetProfessionResponse
	if profession == "" {
		return resp, errors.New("empty title")
	}

	resp.Id = uuid.Nil
	resp.Title = profession
	resp.Description = description
	professionData := app.db.QueryRow(`INSERT INTO professions (title, description) VALUES ($1, $2)
				 ON CONFLICT (title) DO UPDATE SET title = excluded.title RETURNING profession_id `, profession, description)
	if err := professionData.Scan(&resp.Id); err != nil {
		return resp, err
	}
	return resp, nil
}

func (app *App) PostCompetencyProfession(competencyId uuid.UUID, professionId uuid.UUID) error {
	if professionId == uuid.Nil {
		return errors.New("empty profession id")
	}
	if competencyId == uuid.Nil {
		return errors.New("empty competency id")
	}

	_, err := app.db.Exec(`INSERT INTO competency_profession (competency_id, profession_id) VALUES ($1, $2)`, competencyId, professionId)
	return err
}

func (app *App) PostProject(project string, description string, result string, lifeScenario string, technologyId uuid.UUID) (model.PostProjectResponse, error) {
	var resp model.PostProjectResponse
	if project == "" {
		return resp, errors.New("empty title")
	}

	resp.Id = uuid.Nil
	resp.Title = project
	resp.Description = description
	resp.MainTechnologyId = technologyId
	projectData := app.db.QueryRow(`INSERT INTO projects (title, description, result, life_scenario, main_technology_id) VALUES ($1, $2, $3, $4, $5)
				 ON CONFLICT (title) DO UPDATE SET title = excluded.title RETURNING project_id `, project, description, result, lifeScenario, technologyId)
	if err := projectData.Scan(&resp.Id); err != nil {
		return resp, err
	}
	return resp, nil
}

func (app *App) PostOrganization(organization string) (model.GetOrganizationResponse, error) {
	var resp model.GetOrganizationResponse
	if organization == "" {
		return resp, errors.New("empty title")
	}

	resp.Id = uuid.Nil
	resp.Title = organization
	organizationData := app.db.QueryRow(`INSERT INTO organizations (title) VALUES ($1) ON CONFLICT (title)
							 DO UPDATE SET title = excluded.title RETURNING organization_id `, organization)
	if err := organizationData.Scan(&resp.Id); err != nil {
		return resp, err
	}
	return resp, nil
}

func (app *App) PostEducationalProgram(educationalProgram string, description string, organizationId uuid.UUID) (model.PostEducationalProgramResponse, error) {
	var resp model.PostEducationalProgramResponse
	if educationalProgram == "" {
		return resp, errors.New("empty title")
	}

	resp.Id = uuid.Nil
	resp.Title = educationalProgram
	resp.Description = description
	resp.OrganizationId = organizationId
	educationalProgramData := app.db.QueryRow(`INSERT INTO educational_programs (title, description, organizations_id) VALUES ($1, $2, $3) ON CONFLICT (title)
							 	DO UPDATE SET title = excluded.title RETURNING educational_program_id `, educationalProgram, description, organizationId)
	if err := educationalProgramData.Scan(&resp.Id); err != nil {
		return resp, err
	}
	return resp, nil
}

func (app *App) PostDiscipline(discipline string, description string, educationalProgramId uuid.UUID) (model.PostDisciplineResponse, error) {
	var resp model.PostDisciplineResponse
	if discipline == "" {
		return resp, errors.New("empty title")
	}

	resp.Id = uuid.Nil
	resp.Title = discipline
	resp.Description = description
	resp.EducationalProgramId = educationalProgramId
	disciplineData := app.db.QueryRow(`INSERT INTO disciplines (title, description, educational_program_id) VALUES ($1, $2, $3) ON CONFLICT (title)
							 	DO UPDATE SET title = excluded.title RETURNING discipline_id `, discipline, description, educationalProgramId)
	if err := disciplineData.Scan(&resp.Id); err != nil {
		return resp, err
	}
	return resp, nil
}

func (app *App) PostCourse(course string, description string, teacher string, disciplineId uuid.UUID) (model.PostCourseResponse, error) {
	var resp model.PostCourseResponse
	if course == "" {
		return resp, errors.New("empty title")
	}

	courseData := app.db.QueryRow(`INSERT INTO courses (title, description, teacher, discipline_id) VALUES ($1, $2, $3, $4) ON CONFLICT (title)
							 	DO UPDATE SET title = excluded.title RETURNING course_id, title, description, teacher, discipline_id`, course, description, teacher, disciplineId)
	if err := courseData.Scan(&resp.Id, &resp.Title, &resp.Description, &resp.Teacher, &resp.DisciplineId); err != nil {
		return resp, err
	}
	return resp, nil
}

func (app *App) PostCourseCompetency(courseId uuid.UUID, competencyId uuid.UUID) error {
	if courseId == uuid.Nil {
		return errors.New("empty course id")
	}
	if competencyId == uuid.Nil {
		return errors.New("empty competency id")
	}

	_, err := app.db.Exec(`INSERT INTO course_competency (course_id, competency_id) VALUES ($1, $2)`, courseId, competencyId)
	return err
}

func (app *App) PostPortfolio() (model.PostPortfolioResponse, error) {
	var resp model.PostPortfolioResponse
	portfolioData := app.db.QueryRow(`INSERT INTO portfolios (portfolio_id) VALUES (DEFAULT) RETURNING portfolio_id`)
	if err := portfolioData.Scan(&resp.Id); err != nil {
		return resp, err
	}
	return resp, nil
}

func (app *App) PostProjectPortolio(projectId uuid.UUID, portfolioId uuid.UUID, teamRole string, semester uint8) (model.PostProjectPortfolioResponse, error) {
	var resp model.PostProjectPortfolioResponse
	if projectId == uuid.Nil {
		return resp, errors.New("empty project id")
	}
	if portfolioId == uuid.Nil {
		return resp, errors.New("empty portfolio id")
	}

	projectPortfolioData := app.db.QueryRow(`INSERT INTO project_portfolio (project_id, portfolio_id, team_role, semester) VALUES ($1, $2, $3, $4)
								RETURNING project_id, portfolio_id, team_role, semester`, projectId, portfolioId, teamRole, semester)
	if err := projectPortfolioData.Scan(&resp.PortfolioId, resp.PortfolioId, resp.TeamRole, resp.Semester); err != nil {
		return resp, err
	}
	return resp, nil
}

func (app *App) PostProjectPortfolioCompetency(projectId uuid.UUID, portfolioId uuid.UUID, competencyId uuid.UUID) error {
	if projectId == uuid.Nil {
		return errors.New("empty project id")
	}
	if portfolioId == uuid.Nil {
		return errors.New("empty portfolio id")
	}
	if competencyId == uuid.Nil {
		return errors.New("empty competency id")
	}

	_, err := app.db.Exec(`INSERT INTO project_portfolio_competencies (competency_id, project_id, portfolio_id) VALUES ($1, $2, $3)`, competencyId, projectId, portfolioId)
	return err
}

func (app *App) PostStudyGroup(courseId uuid.UUID, studentId uuid.UUID) error {
	if studentId == uuid.Nil {
		return errors.New("empty student id")
	}
	if courseId == uuid.Nil {
		return errors.New("empty course id")
	}

	_, err := app.db.Exec(`INSERT INTO study_groups (course_id, student_id) VALUES ($1, $2)`, courseId, studentId)
	return err
}

func (app *App) PostStudent(fullName string, admition time.Time, portfolioId uuid.UUID) (model.PostStudentResponse, error) {
	var resp model.PostStudentResponse
	if fullName == "" {
		return resp, errors.New("empty student name")
	}

	//if empty time given than current time will be set
	if admition.IsZero() {
		admition = time.Now()
	}

	//if zero portfolio is given than new portfolio id will be generated
	if portfolioId == uuid.Nil {
		portfolio, err := app.PostPortfolio()
		if err != nil {
			return resp, err
		}
		portfolioId = portfolio.Id
	}

	studentData := app.db.QueryRow(`INSERT INTO students (full_name, portfolio_id, admition) VALUES ($1, $2, $3)
								RETURNING student_id, full_name, portfolio_id, admition`, fullName, portfolioId, admition)
	if err := studentData.Scan(&resp.Id, &resp.FullName, &resp.PortfolioId, &resp.Admition); err != nil {
		return resp, err
	}
	return resp, nil
}

func (app *App) PostTrajectory(semester uint8, studentId uuid.UUID, courseId uuid.UUID) (model.PostTrajectoryResponse, error) {
	var resp model.PostTrajectoryResponse
	if semester <= 0 {
		return resp, errors.New("semester must be positive")
	}

	if studentId == uuid.Nil {
		return resp, errors.New("empty student id")
	}

	if courseId == uuid.Nil {
		return resp, errors.New("empty course id")
	}

	trajectoryData := app.db.QueryRow(`INSERT INTO trajectories (student_id, course_id, semester) VALUES ($1, $2, $3)
								RETURNING trajectory_id, student_id, course_id, semester`, studentId, courseId, semester)
	if err := trajectoryData.Scan(&resp.Id, &resp.StudentId, &resp.CourseId, &resp.Semester); err != nil {
		return resp, err
	}
	return resp, nil
}
