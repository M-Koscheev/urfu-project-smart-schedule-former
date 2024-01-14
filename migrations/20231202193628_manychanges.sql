-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE EXTENSION "uuid-ossp";

CREATE TABLE knowledge ( -- Знания (составляющая часть компетенций) (есть в таблице, состявляемой аналитиками)
    knowledge_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR NOT NULL UNIQUE
);

CREATE TABLE technologies ( -- Ключевые технологии
    technology_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR NOT NULL UNIQUE
);

CREATE TABLE competencies ( -- Компетенции (есть в таблице, состявляемой аналитиками)
    competency_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR NOT NULL,
    skills VARCHAR,
    main_technology_id UUID REFERENCES technologies(technology_id) DEFAULT uuid_nil()
);

CREATE TABLE knowledge_competency ( -- Связь между компетенцией и составляющим ее знанием
    knowledge_id UUID REFERENCES knowledge(knowledge_id) ON DELETE CASCADE,
    competency_id UUID REFERENCES competencies(competency_id) ON DELETE CASCADE,
    PRIMARY KEY (knowledge_id, competency_id)
);

CREATE TABLE professions ( -- Список доступных для выбора профессий.
    profession_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR NOT NULL,
    description VARCHAR
);  

CREATE TABLE competency_profession ( -- Связь между профессией и требующейся для нее компетенции
    competency_id UUID REFERENCES competencies(competency_id) ON DELETE CASCADE,
    profession_id UUID REFERENCES professions(profession_id) ON DELETE CASCADE,
    PRIMARY KEY (profession_id, competency_id)
);

CREATE TABLE projects ( -- Учебный проект по проектной деятельности в конкретном семесте у конкретной команды.
    project_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR NOT NULL,
    description VARCHAR NOT NUll,
    result VARCHAR,
    life_scenario VARCHAR,
    main_technology_id UUID REFERENCES technologies(technology_id) DEFAULT uuid_nil()
);

-- институт/университет/компания
CREATE TABLE organizations ( 
    organization_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR
);

-- Например: Программная инженерия или Прикладная информатика?
CREATE TABLE educational_programs (
    educational_program_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR NOT NULL,
    description VARCHAR,
    organizations_id UUID NOT NULL REFERENCES organizations(organization_id) ON DELETE CASCADE
);

-- Например: Физкультура или Программирование(скажем в 3 семестре)?
CREATE TABLE disciplines (
    discipline_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR NOT NULL,
    description VARCHAR,
    educational_program_id UUID NOT NULL REFERENCES educational_programs(educational_program_id) ON DELETE CASCADE
);

-- Например в разделе Программирование (в 3 семестре): Go от geekbrains, Java от УрФУ, Kotlin от ИТМО?
CREATE TABLE courses (
    course_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR NOT NULL,
    description VARCHAR,
    teacher VARCHAR, 
    discipline_id UUID NOT NULL REFERENCES disciplines(discipline_id) ON DELETE CASCADE
);

CREATE TABLE portfolios (  
    portfolio_id UUID PRIMARY KEY DEFAULT gen_random_uuid()
);

CREATE TABLE project_portfolio ( -- Связь между проектом и портофлио проектов конкретного студента.
    project_id UUID REFERENCES projects(project_id) ON DELETE CASCADE,
    portfolio_id UUID REFERENCES portfolios(portfolio_id) ON DELETE CASCADE,
    team_role VARCHAR,
    semester SMALLINT CHECK (semester > 0),
    PRIMARY KEY (project_id, portfolio_id)
);

CREATE TABLE project_portfolio_competencies (
    competency_id UUID REFERENCES competencies(competency_id) ON DELETE CASCADE,
    project_id UUID REFERENCES projects(project_id) ON DELETE CASCADE,
    portfolio_id UUID REFERENCES portfolios(portfolio_id) ON DELETE CASCADE,
    FOREIGN KEY (project_id, portfolio_id) REFERENCES project_portfolio(project_id, portfolio_id),
    PRIMARY KEY (competency_id, project_id, portfolio_id)
)


CREATE TABLE students ( -- Информация о конкретном студенте.
    student_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    full_name VARCHAR NOT NULL,
    portfolio_id UUID REFERENCES portfolios(portfolio_id),
    -- semester SMALLINT NOT NULL CHECK (semester > 0) 
    admition DATE NOT NULL
);

CREATE TABLE study_groups ( -- Учебная группа конкретного человека в конкретном семестре (Например РИ-220942)
    -- study_group_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    -- semester SMALLINT CHECK (semester > 0), 
    course_id UUID REFERENCES courses(course_id) ON DELETE CASCADE,
    student_id UUID REFERENCES students(student_id) ON DELETE CASCADE,
    PRIMARY KEY (course_id, student_id)
);

CREATE TABLE trajectories ( -- Траектория конкретного студента в конкретном семестре. (архивные данные)
    trajectory_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    student_id UUID NOT NULL REFERENCES students(student_id),
    course_id UUID REFERENCES courses(course_id)
    semester SMALLINT CHECK (semester > 0)
);

CREATE TABLE course_competencies ( -- список компетенции у траектории
    course_id UUID REFERENCES courses(course_id) ON DELETE CASCADE,
    competency_id UUID REFERENCES competencies(competency_id) ON DELETE CASCADE,
    PRIMARY KEY (trajectory_id, competency_id)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE course_competencies;
DROP TABLE trajectories;
DROP TABLE study_groups;
DROP TABLE students;
DROP TABLE project_portfolio_competencies;
DROP TABLE project_portfolio;
DROP TABLE portfolios;
DROP TABLE courses;
DROP TABLE disciplines;
DROP TABLE educational_programs;
DROP TABLE organizations;
DROP TABLE projects;
DROP TABLE competency_profession;
DROP TABLE professions;
DROP TABLE professions;
DROP TABLE knowledge_competency;
DROP TABLE competencies;
DROP TABLE technologies;
DROP TABLE knowledge;
