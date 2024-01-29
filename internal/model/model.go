package model

import (
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

// First create a type alias
type JsonAdmitionDate time.Time

// Implement Marshaler and Unmarshaler interface
func (j *JsonAdmitionDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = JsonAdmitionDate(t)
	return nil
}

func (j JsonAdmitionDate) MarshalJSON() ([]byte, error) {
	return []byte("\"" + time.Time(j).Format("2006-01-02") + "\""), nil
}

// Maybe a Format function for printing your date
func (j JsonAdmitionDate) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}

type GetKnowledge struct {
	Id    uuid.UUID `json:"knowledgeId,omitempty" example:"00000000-0000-0000-0000-000000000000"`
	Title string    `json:"knowledgeTitle" example:"Название знания"`
}

type GetTechnology struct {
	Id    uuid.UUID `json:"technologyId,omitempty" example:"00000000-0000-0000-0000-000000000000"`
	Title string    `json:"technologyTitle" example:"Название технологии"`
}

type GetCompetency struct {
	Id               uuid.UUID `json:"competencyId,omitempty" example:"00000000-0000-0000-0000-000000000000"`
	Title            string    `json:"competencyTitle" example:"Название компетенции"`
	Skills           string    `json:"competencySkills,omitempty" example:"навык 1, навык 2..."`
	MainTechnologyId uuid.UUID `json:"competencyMainTechnology,omitempty" example:"00000000-0000-0000-0000-000000000000"`
	Knowledge        []string  `json:"competencyKnowledge,omitempty" example:"знание 1, знание 2..."`
}

type PostCompetency struct {
	Title            string    `json:"competencyTitle" example:"Название компетенции"`
	Skills           string    `json:"competencySkills,omitempty" example:"навык 1, навык 2..."`
	MainTechnologyId uuid.UUID `json:"competencyMainTechnology,omitempty" example:"00000000-0000-0000-0000-000000000000"`
}

type GetProfession struct {
	Id           uuid.UUID `json:"professionId,omitempty" example:"00000000-0000-0000-0000-000000000000"`
	Title        string    `json:"professionTitle" example:"Название профессии"`
	Description  string    `json:"professionDescription,omitempty" example:"Описание профессии"`
	Competencies []string  `json:"professionCompetencies,omitempty" example:"компетенция 1, компетенция 2..."`
}

type PostProfession struct {
	Title       string `json:"professionTitle" example:"Название профессии"`
	Description string `json:"professionDescription,omitempty" example:"Описание профессии"`
}

type GetProject struct {
	Id             uuid.UUID `json:"projectId,omitempty" example:"00000000-0000-0000-0000-000000000000"`
	Title          string    `json:"projectTitle" example:"Исследование зивисимости длины шерстки капибар от продолжительности их жизни"`
	Description    string    `json:"projectDescription,omitempty" example:"Курс направлен на изучение поведения капибар в дикой природе..."`
	Result         string    `json:"projectResult,omitempty" example:"Данные о зависимости длины шерстки капибар от продолжительности их жизни"`
	LifeScenario   string    `json:"projectLifeScenarion,omitempty" example:"Проект позволит подобрать идеальную длину шерстки для ваших капибар"`
	MainTechnology string    `json:"projectMainTechnology,omitempty" example:"Название технологии"`
}

type GetOrganization struct {
	Id    uuid.UUID `json:"organizationId,omitempty" example:"00000000-0000-0000-0000-000000000000"`
	Title string    `json:"organizationTitle" example:"Название организации"`
}

type PostOrganization struct {
	Title string `json:"organizationTitle" example:"Название организации"`
}

type GetEducationalProgram struct {
	Id           uuid.UUID `json:"educationalProgramId,omitempty" example:"00000000-0000-0000-0000-000000000000"`
	Title        string    `json:"educationalProgramTitle" example:"Название образовательной программы"`
	Description  string    `json:"educationalProgramDexcription,omitempty" example:"Описание образовательной программы"`
	Organization string    `json:"educationalProgramOrganization,omitempty" example:"организация, отвечающая за образовательную программу"`
}

type GetDiscipline struct {
	Id                 uuid.UUID `json:"disciplineId,omitempty" example:"00000000-0000-0000-0000-000000000000"`
	Title              string    `json:"disciplineTitle" example:"Название дисциплины"`
	Description        string    `json:"disciplineDescription,omitempty" example:"Описание дисциплины"`
	EducationalProgram string    `json:"disciplineEducationalProgram,omitempty" example:"Образовательная программа, в которую входит цисциплина"`
}

type GetCourse struct {
	Id           uuid.UUID `json:"courseId,omitempty" example:"00000000-0000-0000-0000-000000000000"`
	Title        string    `json:"courseTitle" example:"Название курса"`
	Description  string    `json:"courseDescription,omitempty"  example:"Описание курса"`
	Teacher      string    `json:"courseTeacher,omitempty" example:"Преподаватель"`
	Discipline   string    `json:"courseDiscipline,omitempty" example:"Дисциплина, к которой отностися курс"`
	Competencies []string  `json:"courseCompetencies,omitempty" example:"компетенция 1, компетенция 2..."`
}

type GetPersonalProject struct {
	Id             uuid.UUID `json:"projectId,omitempty" example:"00000000-0000-0000-0000-000000000000"`
	Title          string    `json:"projectTitle" example:"Исследование зивисимости длины шерстки капибар от продолжительности их жизни"`
	Description    string    `json:"projectDescription,omitempty" example:"Курс направлен на изучение поведения капибар в дикой природе..."`
	Result         string    `json:"projectResult,omitempty" example:"Данные о зависимости длины шерстки капибар от продолжительности их жизни"`
	LifeScenario   string    `json:"projectLifeScenarion,omitempty" example:"Проект позволит подобрать идеальную длину шерстки для ваших капибар"`
	MainTechnology string    `json:"personalProjectMainTechnology,omitempty" example:"Основная технология проекта"`
	TeamRole       string    `json:"personalProjectTeamRole" example:"Роль участника в команде"`
	Semester       string    `json:"personalProjectSemester" example:"Семестр, в котором участик работал над проектом"`
	Competencies   []string  `json:"personalProjectCompetencies,omitempty" example:"Компетенции участика в этом проекте"`
}

type GetPortfolio struct {
	Id       uuid.UUID            `json:"portfolioId,omitempty" example:"00000000-0000-0000-0000-000000000000"`
	Projects []GetPersonalProject `json:"portfolioProjects,omitempty"`
}

type GetStudent struct {
	Id        uuid.UUID    `json:"studentId,omitempty" example:"00000000-0000-0000-0000-000000000000"`
	FullName  string       `json:"studentFullName" example:"Фамилия Имя Отчество"`
	Portfolio GetPortfolio `json:"studentPortfolio,omitempty"`
	Semester  uint8        `json:"studentSemester" example:"3"`
}

type GetStudyGroups struct {
	// Id       uuid.UUID          `json:"studyGroupId,omitempty"`
	// Title    string `json:"studyGroupTitle"`
	// Semester        uint8    `json:"studyGroupSemester"`
	Courses []string `json:"studyGroupCourse" example:"Курсы, которые студент изучает"`
	// StudentFullName string   `json:"studyGroupStudent"`
}

type GetTrajectory struct {
	Id       uuid.UUID `json:"trajectoryId,omitempty" example:"00000000-0000-0000-0000-000000000000"`
	Student  string    `json:"trajectoryStudent" example:"Фамилия Имя Отчество"`
	Semester uint8     `json:"trajectorySemester" example:"3"`
	Course   string    `json:"trajectoryCourse" example:"Название курса"`
}

type PostKnowledge struct {
	Title string `json:"knowledgeTitle" example:"Название знания"`
}

type PostTechnology struct {
	Title string `json:"technologyTitle" example:"Название технологии"`
}

type PostProject struct {
	Title            string    `json:"projectTitle" example:"Исследование зивисимости длины шерстки капибар от продолжительности их жизни"`
	Description      string    `json:"projectDescription,omitempty" example:"Курс направлен на изучение поведения капибар в дикой природе..."`
	Result           string    `json:"projectResult,omitempty" example:"Данные о зависимости длины шерстки капибар от продолжительности их жизни"`
	LifeScenario     string    `json:"projectLifeScenarion,omitempty" example:"Проект позволит подобрать идеальную длину шерстки для ваших капибар"`
	MainTechnologyId uuid.UUID `json:"projectMainTechnologyId,omitempty" example:"00000000-0000-0000-0000-000000000000"`
}

type PostEducationalProgram struct {
	Title          string    `json:"educationalProgramTitle" example:"Название образовательной программы"`
	Description    string    `json:"educationalProgramDescription,omitempty" example:"Описание образовательной программы"`
	OrganizationId uuid.UUID `json:"educationalProgramOrganizationId,omitempty" example:"00000000-0000-0000-0000-000000000000"`
}

type PostDiscipline struct {
	Title                string    `json:"disciplineTitle" example:"Название дисциплины"`
	Description          string    `json:"disciplineDescription,omitempty" example:"Описание дисциплины"`
	EducationalProgramId uuid.UUID `json:"disciplineEducationalProgramId,omitempty" example:"00000000-0000-0000-0000-000000000000"`
}

type PostCourse struct {
	Title        string    `json:"courseTitle" example:"Название курса"`
	Description  string    `json:"courseDescription,omitempty" example:"Описание курса"`
	Teacher      string    `json:"courseTeacher,omitempty" example:"Фамилия Имя Отчество"`
	DisciplineId uuid.UUID `json:"courseDisciplineId,omitempty" example:"00000000-0000-0000-0000-000000000000"`
}

type PostPortfolio struct {
	Id uuid.UUID `json:"portfolioId,omitempty" example:"00000000-0000-0000-0000-000000000000"`
}

type PostProjectPortfolio struct {
	ProjectId   uuid.UUID `json:"ProjectId" example:"00000000-0000-0000-0000-000000000000"`
	PortfolioId uuid.UUID `json:"PortfolioId" example:"00000000-0000-0000-0000-000000000000"`
	TeamRole    string    `json:"TeamRole,omitempty" example:"Роль в команде"`
	Semester    uint8     `json:"projectSemester,omitempty" example:"3"`
}

type PostProjectPortfolioCompetency struct {
	ProjectId    uuid.UUID `json:"projectId" example:"00000000-0000-0000-0000-000000000000"`
	PortfolioId  uuid.UUID `json:"PortfolioId" example:"00000000-0000-0000-0000-000000000000"`
	CompetencyId uuid.UUID `json:"CompetencyId" example:"00000000-0000-0000-0000-000000000000"`
}

type PostStudent struct {
	FullName    string           `json:"studentFullName" example:"Фамилия Имя Отчество"`
	Admition    JsonAdmitionDate `json:"studentAdmitionDate" example:"2024-01-19"`
	PortfolioId uuid.UUID        `json:"studentPortfolioId,omitempty" example:"00000000-0000-0000-0000-000000000000"`
}

type PostTrajectory struct {
	StudentId uuid.UUID `json:"trajectoryStudentId" example:"00000000-0000-0000-0000-000000000000"`
	Semester  uint8     `json:"trajectorySemester" example:"3"`
	CourseId  uuid.UUID `json:"trajectoryCourseId" example:"00000000-0000-0000-0000-000000000000"`
}

type PostKnowledgeCompetency struct {
	KnowledgeId  uuid.UUID `json:"knowledgeId" example:"00000000-0000-0000-0000-000000000000"`
	CompetencyId uuid.UUID `json:"competencyId" example:"00000000-0000-0000-0000-000000000000"`
}

type PostCompetencyProfession struct {
	CompetencyId uuid.UUID `json:"competencyId" example:"00000000-0000-0000-0000-000000000000"`
	ProfessionId uuid.UUID `json:"professionId" example:"00000000-0000-0000-0000-000000000000"`
}

type PostCourseCompetency struct {
	CourseId     uuid.UUID `json:"courseId" example:"00000000-0000-0000-0000-000000000000"`
	CompetencyId uuid.UUID `json:"competencyId" example:"00000000-0000-0000-0000-000000000000"`
}

type PostStudyGroup struct {
	CourseId  uuid.UUID `json:"courseId" example:"00000000-0000-0000-0000-000000000000"`
	StudentId uuid.UUID `json:"studentId" example:"00000000-0000-0000-0000-000000000000"`
}
