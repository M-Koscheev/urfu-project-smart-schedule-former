package app

import (
	"database/sql"
	"errors"
	"log/slog"
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

var ErrEmptyTitle = errors.New("empty title")
var ErrEmptyId = errors.New("empty id")

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

func (app *App) GetKnowledgeByIndex(id uuid.UUID) (model.GetKnowledge, error) {
	var resp model.GetKnowledge
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

func (app *App) GetTechnolgyById(id uuid.UUID) (model.GetTechnology, error) {
	var resp model.GetTechnology
	data := app.db.QueryRow(`SELECT technology_id, title FROM technologies WHERE technology_id = $1`, id)
	err := data.Scan(&resp.Id, &resp.Title)
	return resp, err
}

func (app *App) GetCompetencyById(id uuid.UUID) (model.GetCompetency, error) {
	var resp model.GetCompetency
	row := app.db.QueryRow(`SELECT competency_id, title, skills, main_technology_id FROM competencies WHERE competency_id = $1`, id)
	var mainTechnologyId uuid.UUID
	err := row.Scan(&resp.Id, &resp.Title, &resp.Skills, &mainTechnologyId)
	if err != nil {
		return resp, err
	}

	if mainTechnologyId != uuid.Nil {
		var technology model.GetTechnology
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

func (app *App) GetProfessionById(id uuid.UUID) (model.GetProfession, error) {
	var resp model.GetProfession
	data := app.db.QueryRow(`SELECT profession_id, title, description FROM professions WHERE profession_id = $1`, id)
	err := data.Scan(&resp.Id, &resp.Title, &resp.Description)
	if err != nil {
		return resp, err
	}

	resp.Competencies, err = app.getCompetenciesByProfession(id)
	return resp, err
}

func (app *App) GetProjectById(id uuid.UUID) (model.GetProject, error) {
	var resp model.GetProject
	data := app.db.QueryRow(`SELECT project_id, title, description, result, life_scenario, main_technology_id FROM projects WHERE project_id = $1`, id)
	var mainTechnologyId uuid.UUID
	err := data.Scan(&resp.Id, &resp.Title, &resp.Description, &resp.Result, &resp.LifeScenario, mainTechnologyId)
	if err != nil {
		return resp, err
	}

	if mainTechnologyId != uuid.Nil {
		var technology model.GetTechnology
		if technology, err = app.GetTechnolgyById(mainTechnologyId); err != nil {
			return resp, err
		}

		resp.MainTechnology = technology.Title
	}

	return resp, nil
}

func (app *App) GetOrganizationById(id uuid.UUID) (model.GetOrganization, error) {
	var resp model.GetOrganization
	data := app.db.QueryRow(`SELECT organization_id, title FROM organizations WHERE organization_id = $1`, id)
	err := data.Scan(&resp.Id, &resp.Title)
	return resp, err
}

func (app *App) GetEducationalProgramById(id uuid.UUID) (model.GetEducationalProgram, error) {
	var resp model.GetEducationalProgram
	var organizationId uuid.UUID
	data := app.db.QueryRow(`SELECT educational_program_id, title, description, organizations_id FROM educational_programs WHERE educational_program_id = $1`, id)
	err := data.Scan(&resp.Id, &resp.Title, &resp.Description, &organizationId)
	if err != nil {
		return resp, err
	}

	var organization model.GetOrganization
	if organization, err = app.GetOrganizationById(organizationId); err != nil {
		return resp, err
	}

	resp.Organization = organization.Title
	return resp, nil
}

func (app *App) GetDisciplineById(id uuid.UUID) (model.GetDiscipline, error) {
	var resp model.GetDiscipline
	var educationalProgramId uuid.UUID
	data := app.db.QueryRow(`SELECT discipline_id, title, description, educational_program_id FROM disciplines WHERE discipline_id = $1`, id)
	err := data.Scan(&resp.Id, &resp.Title, &resp.Description, &educationalProgramId)
	if err != nil {
		return resp, err
	}

	var educationalProgram model.GetEducationalProgram
	if educationalProgram, err = app.GetEducationalProgramById(educationalProgramId); err != nil {
		return resp, err
	}

	resp.EducationalProgram = educationalProgram.Title
	return resp, nil
}

func (app *App) getCompetenciesByCourse(courseId uuid.UUID) ([]string, error) {
	var competencies []string
	rows, err := app.db.Query(`SELECT title FROM competencies WHERE competency_id in 
		(SELECT competency_id FROM course_competency WHERE course_id = $1)`, courseId)
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

func (app *App) GetCourseById(id uuid.UUID) (model.GetCourse, error) {
	var resp model.GetCourse
	var disciplineId uuid.UUID
	data := app.db.QueryRow(`SELECT course_id, title, description, teacher, discipline_id FROM courses WHERE course_id = $1`, id)
	err := data.Scan(&resp.Id, &resp.Title, &resp.Description, &resp.Teacher, &disciplineId)
	if err != nil {
		return resp, nil
	}

	if disciplineId != uuid.Nil {
		var discipline model.GetDiscipline
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
		(SELECT competency_id FROM project_portfolio_competency WHERE portfolio_id = $1 and project_id = $2)`, portfolioId, projectId)
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

func (app *App) GetPersonalProjectsByPortfolio(portfolioId uuid.UUID) ([]model.GetPersonalProject, error) {
	var resp []model.GetPersonalProject
	rows, err := app.db.Query(`SELECT project_id, team_role, semester FROM project_portfolio WHERE portfolio_id = $1`, portfolioId)
	if err != nil {
		return resp, nil
	}

	for rows.Next() {
		var personlaProject model.GetPersonalProject
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

func (app *App) GetPortfolioById(id uuid.UUID) (model.GetPortfolio, error) {
	var resp model.GetPortfolio
	var err error
	resp.Id = id
	resp.Projects, err = app.GetPersonalProjectsByPortfolio(id)
	if err != nil {
		return resp, nil
	}

	return resp, nil
}

func (app *App) GetStudentById(id uuid.UUID) (model.GetStudent, error) {
	var resp model.GetStudent
	var portfolioId uuid.UUID
	var admition time.Time
	var err error
	data := app.db.QueryRow(`SELECT student_id, full_name, portfolio_id, admition FROM students WHERE student_id = $1`, id)
	if err = data.Scan(&resp.Id, &resp.FullName, &portfolioId, &admition); err != nil {
		return resp, err
	}
	if admition.After(time.Now()) {
		slog.Error("admition date - ", admition.String(), ", now - ", time.Now().String(), errors.New("incorrect admition date"))
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

func (app *App) GetStudyGroupsByStudent(studentId uuid.UUID) (model.GetStudyGroups, error) {
	var resp model.GetStudyGroups
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

func (app *App) GetTrajectoryById(trajectoryId uuid.UUID) (model.GetTrajectory, error) {
	var resp model.GetTrajectory
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

func (app *App) PostKnowledge(knowledge string) (model.GetKnowledge, error) {
	var resp model.GetKnowledge
	if knowledge == "" {
		return resp, ErrEmptyTitle
	}

	resp.Id = uuid.Nil
	resp.Title = knowledge
	knowledgeData := app.db.QueryRow(`INSERT INTO knowledge (title) VALUES ($1) ON CONFLICT (title) DO UPDATE SET title = excluded.title RETURNING knowledge_id`, knowledge)
	if err := knowledgeData.Scan(&resp.Id); err != nil {
		return resp, err
	}

	return resp, nil
}

func (app *App) PostTechnology(technology string) (model.GetTechnology, error) {
	var resp model.GetTechnology
	if technology == "" {
		return resp, ErrEmptyTitle
	}

	resp.Id = uuid.Nil
	resp.Title = technology
	technologyData := app.db.QueryRow(`INSERT INTO technologies (title) VALUES ($1) ON CONFLICT (title) DO UPDATE SET title = excluded.title RETURNING technology_id `, technology)
	if err := technologyData.Scan(&resp.Id); err != nil {
		return resp, err
	}
	return resp, nil
}

//

func (app *App) PostCompetency(comptency string, skills string, technologyId uuid.UUID) (model.GetCompetency, error) {
	var resp model.GetCompetency
	if comptency == "" {
		return resp, ErrEmptyTitle
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
		return ErrEmptyId
	}
	if competencyId == uuid.Nil {
		return ErrEmptyId
	}

	_, err := app.db.Exec(`INSERT INTO knowledge_competency (knowledge_id, competency_id) VALUES ($1, $2)`, knowledgeId, competencyId)
	return err
}

func (app *App) PostProfession(profession string, description string) (model.GetProfession, error) {
	var resp model.GetProfession
	if profession == "" {
		return resp, ErrEmptyTitle
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
		return ErrEmptyId
	}
	if competencyId == uuid.Nil {
		return ErrEmptyId
	}

	_, err := app.db.Exec(`INSERT INTO competency_profession (competency_id, profession_id) VALUES ($1, $2)`, competencyId, professionId)
	return err
}

func (app *App) PostProject(project string, description string, result string, lifeScenario string, technologyId uuid.UUID) (model.GetProject, error) {
	var resp model.GetProject
	if project == "" {
		return resp, ErrEmptyTitle
	}

	resp.Id = uuid.Nil
	resp.Title = project
	resp.Description = description

	technology, err := app.GetTechnolgyById(technologyId)
	if err != nil {
		return resp, err
	}
	resp.MainTechnology = technology.Title

	projectData := app.db.QueryRow(`INSERT INTO projects (title, description, result, life_scenario, main_technology_id) VALUES ($1, $2, $3, $4, $5)
				 ON CONFLICT (title) DO UPDATE SET title = excluded.title RETURNING project_id `, project, description, result, lifeScenario, technologyId)
	if err := projectData.Scan(&resp.Id); err != nil {
		return resp, err
	}
	return resp, nil
}

func (app *App) PostOrganization(organization string) (model.GetOrganization, error) {
	var resp model.GetOrganization
	if organization == "" {
		return resp, ErrEmptyTitle
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

func (app *App) PostEducationalProgram(educationalProgram string, description string, organizationId uuid.UUID) (model.GetEducationalProgram, error) {
	var resp model.GetEducationalProgram
	if educationalProgram == "" {
		return resp, ErrEmptyTitle
	}
	resp.Title = educationalProgram
	resp.Description = description
	if organizationId != uuid.Nil {
		organization, err := app.GetOrganizationById(organizationId)
		if err != nil {
			return resp, err
		}
		resp.Organization = organization.Title
	}

	educationalProgramData := app.db.QueryRow(`INSERT INTO educational_programs (title, description, organizations_id) VALUES ($1, $2, $3) ON CONFLICT (title)
							 	DO UPDATE SET title = excluded.title RETURNING educational_program_id `, educationalProgram, description, organizationId)
	if err := educationalProgramData.Scan(&resp.Id); err != nil {
		return resp, err
	}
	return resp, nil
}

func (app *App) PostDiscipline(discipline string, description string, educationalProgramId uuid.UUID) (model.GetDiscipline, error) {
	var resp model.GetDiscipline
	if discipline == "" {
		return resp, ErrEmptyTitle
	}

	resp.Title = discipline
	resp.Description = description
	if educationalProgramId != uuid.Nil {
		educationalProgram, err := app.GetEducationalProgramById(educationalProgramId)
		if err != nil {
			return resp, err
		}
		resp.EducationalProgram = educationalProgram.Title
	}

	disciplineData := app.db.QueryRow(`INSERT INTO disciplines (title, description, educational_program_id) VALUES ($1, $2, $3) ON CONFLICT (title)
							 	DO UPDATE SET title = excluded.title RETURNING discipline_id `, discipline, description, educationalProgramId)
	if err := disciplineData.Scan(&resp.Id); err != nil {
		return resp, err
	}

	return resp, nil
}

func (app *App) PostCourse(course string, description string, teacher string, disciplineId uuid.UUID) (model.GetCourse, error) {
	var resp model.GetCourse
	if course == "" {
		return resp, ErrEmptyTitle
	}
	if disciplineId != uuid.Nil {
		discipline, err := app.GetDisciplineById(disciplineId)
		if err != nil {
			return resp, err
		}
		resp.Discipline = discipline.Title
	}

	resp.Title = course
	resp.Description = description
	resp.Teacher = teacher
	courseData := app.db.QueryRow(`INSERT INTO courses (title, description, teacher, discipline_id) VALUES ($1, $2, $3, $4) ON CONFLICT (title)
							 	DO UPDATE SET title = excluded.title RETURNING course_id`, course, description, teacher, disciplineId)
	if err := courseData.Scan(&resp.Id); err != nil {
		return resp, err
	}
	return resp, nil
}

func (app *App) PostCourseCompetency(courseId uuid.UUID, competencyId uuid.UUID) error {
	if courseId == uuid.Nil {
		return ErrEmptyId
	}
	if competencyId == uuid.Nil {
		return ErrEmptyId
	}

	_, err := app.db.Exec(`INSERT INTO course_competency (course_id, competency_id) VALUES ($1, $2)`, courseId, competencyId)
	return err
}

func (app *App) PostPortfolio() (model.PostPortfolio, error) {
	var resp model.PostPortfolio
	portfolioData := app.db.QueryRow(`INSERT INTO portfolios (portfolio_id) VALUES (DEFAULT) RETURNING portfolio_id`)
	if err := portfolioData.Scan(&resp.Id); err != nil {
		return resp, err
	}
	return resp, nil
}

func (app *App) PostProjectPortolio(projectId uuid.UUID, portfolioId uuid.UUID, teamRole string, semester uint8) (model.PostProjectPortfolio, error) {
	var resp model.PostProjectPortfolio
	if projectId == uuid.Nil {
		return resp, ErrEmptyId
	}
	if portfolioId == uuid.Nil {
		return resp, ErrEmptyId
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
		return ErrEmptyId
	}
	if portfolioId == uuid.Nil {
		return ErrEmptyId
	}
	if competencyId == uuid.Nil {
		return ErrEmptyId
	}

	_, err := app.db.Exec(`INSERT INTO project_portfolio_competency (competency_id, project_id, portfolio_id) VALUES ($1, $2, $3)`, competencyId, projectId, portfolioId)
	return err
}

func (app *App) PostStudyGroup(courseId uuid.UUID, studentId uuid.UUID) error {
	if studentId == uuid.Nil {
		return ErrEmptyId
	}
	if courseId == uuid.Nil {
		return ErrEmptyId
	}

	_, err := app.db.Exec(`INSERT INTO study_groups (course_id, student_id) VALUES ($1, $2)`, courseId, studentId)
	return err
}

func (app *App) PostStudent(fullName string, admition time.Time, portfolioId uuid.UUID) (model.GetStudent, error) {
	var resp model.GetStudent
	if fullName == "" {
		return resp, ErrEmptyTitle
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

	portfolio, err := app.GetPortfolioById(portfolioId)
	if err != nil {
		return resp, err
	}
	resp.Portfolio = portfolio

	studentData := app.db.QueryRow(`INSERT INTO students (full_name, portfolio_id, admition) VALUES ($1, $2, $3)
								RETURNING student_id, full_name`, fullName, portfolioId, admition)

	var semesterChange time.Time = time.Time.AddDate(time.Now(), 0, 1, 0)
	resp.Semester = uint8(time.Now().Year()) - uint8(admition.Year()) + 1
	if time.Now().Month() > semesterChange.Month() {
		resp.Semester += 1
	} else if time.Now().Month() == semesterChange.Month() && time.Now().Day() > semesterChange.Day() {
		resp.Semester += 1
	}

	if err := studentData.Scan(&resp.Id, &resp.FullName); err != nil {
		return resp, err
	}
	return resp, nil
}

func (app *App) PostTrajectory(semester uint8, studentId uuid.UUID, courseId uuid.UUID) (model.GetTrajectory, error) {
	var resp model.GetTrajectory
	if semester <= 0 {
		return resp, errors.New("semester must be positive")
	}

	if studentId == uuid.Nil {
		return resp, ErrEmptyId
	}

	if courseId == uuid.Nil {
		return resp, ErrEmptyId
	}
	slog.Info("id check passed")

	trajectoryData := app.db.QueryRow(`INSERT INTO trajectories (student_id, course_id, semester) VALUES ($1, $2, $3)
								RETURNING trajectory_id`, studentId, courseId, semester)

	student, err := app.GetStudentById(studentId)
	if err != nil {
		return resp, nil
	}
	resp.Student = student.FullName
	resp.Semester = semester

	course, err := app.GetCourseById(courseId)
	if err != nil {
		return resp, nil
	}
	resp.Course = course.Title

	if err := trajectoryData.Scan(&resp.Id); err != nil {
		return resp, err
	}
	return resp, nil
}
