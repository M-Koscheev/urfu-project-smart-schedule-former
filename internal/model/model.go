package model

import (
	uuid "github.com/satori/go.uuid"
)

type GetKnowledgeResponse struct {
	Id    uuid.UUID `json:"knowledgeId,omitempty"`
	Title string    `json:"knowledgeTitle"`
}

type GetTechnologyResponse struct {
	Id    uuid.UUID `json:"technologyId,omitempty"`
	Title string    `json:"technologyTitle"`
}

type GetCompetencyResponse struct {
	Id             uuid.UUID `json:"competencyId,omitempty"`
	Title          string    `json:"competencyTitle"`
	Skills         string    `json:"competencySkills,omitempty"`
	MainTechnology string    `json:"competencyMainTechnology,omitempty"`
	Knowledge      []string  `json:"competencyKnowledge,omitempty"`
}

type GetProfessionResponse struct {
	Id           uuid.UUID `json:"professionId,omitempty"`
	Title        string    `json:"professionTitle"`
	Description  string    `json:"professionDescription,omitempty"`
	Competencies []string  `json:"professionCompetencies,omitempty"`
}

type GetProjectResponse struct {
	Id             uuid.UUID `json:"projectId,omitempty"`
	Title          string    `json:"projectTitle"`
	Description    string    `json:"projectDescription,omitempty"`
	Result         string    `json:"projectResult,omitempty"`
	LifeScenario   string    `json:"projectLifeScenarion,omitempty"`
	MainTechnology string    `json:"projectMainTechnology,omitempty"`
}

type GetOrganizationResponse struct {
	Id    uuid.UUID `json:"organizationId,omitempty"`
	Title string    `json:"organizationTitle"`
}

type GetEducationalProgramResponse struct {
	Id           uuid.UUID `json:"educationalProgramId,omitempty"`
	Title        string    `json:"educationalProgramTitle"`
	Description  string    `json:"educationalProgramDexcription,omitempty"`
	Organization string    `json:"educationalProgramOrganization,omitempty"`
}

type GetDisciplineResponse struct {
	Id                 uuid.UUID `json:"disciplineId,omitempty"`
	Title              string    `json:"disciplineTitle"`
	Description        string    `json:"disciplineDescription,omitempty"`
	EducationalProgram string    `json:"disciplineEducationalProgram,omitempty"`
}

type GetCourseResponse struct {
	Id           uuid.UUID `json:"courseId,omitempty"`
	Title        string    `json:"courseTitle"`
	Description  string    `json:"courseDescription,omitempty"`
	Teacher      string    `json:"courseTeacher,omitempty"`
	Discipline   string    `json:"courseDiscipline,omitempty"`
	Competencies []string  `json:"courseCompetencies,omitempty"`
}

type GetPersonalProjectResponse struct {
	// Id             uuid.UUID `json:"personalProjectId,omitempty"`
	Title          string   `json:"personalProjectTitle"`
	Description    string   `json:"personalProjectDescription,omitempty"`
	Result         string   `json:"personalProjectResult,omitempty"`
	LifeScenario   string   `json:"personalProjectLifeScenarion,omitempty"`
	MainTechnology string   `json:"personalProjectMainTechnology,omitempty"`
	TeamRole       string   `json:"personalProjectTeamRole"`
	Semester       string   `json:"personalProjectSemester"`
	Competencies   []string `json:"personalProjectCompetencies,omitempty"`
}

type GetPortfolioResponse struct {
	Id       uuid.UUID                    `json:"portfolioId,omitempty"`
	Projects []GetPersonalProjectResponse `json:"portfolioProjects,omitempty"`
}

type GetStudentResponse struct {
	Id       uuid.UUID `json:"studentId,omitempty"`
	FullName string    `json:"studentFullName"`
	// Admition  time.Time            `json:"studentAdmitionDate"`
	Portfolio GetPortfolioResponse `json:"studentPortfolio,omitempty"`
	Semester  uint8                `json:"studentSemester"`
}

type GetStudyGroupsResponse struct {
	// Id       uuid.UUID          `json:"studyGroupId,omitempty"`
	// Title    string `json:"studyGroupTitle"`
	// Semester        uint8    `json:"studyGroupSemester"`
	Courses []string `json:"studyGroupCourse"`
	// StudentFullName string   `json:"studyGroupStudent"`
}

type GetTrajectoryResponse struct {
	Id       uuid.UUID `json:"trajectoryId,omitempty"`
	Student  string    `json:"trajectoryStudent"`
	Semester uint8     `json:"trajectorySemester"`
	Course   string    `json:"trajectoryCourse"`
}

type PostKnowledgeRequest struct {
	Title string `json:"knowledgeTitle"`
}

type PostTechnologyRequest struct {
	Title string `json:"technologyTitle"`
}
