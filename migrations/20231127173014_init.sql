-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- DROP TABLE IF EXISTS trajectories, students, project_portfolio, portfolios, study_groups, courses, disciplines, 
-- educational_programs, teachers, projects, professions, knowledge_competence, competencies, knowledge;

CREATE TABLE IF NOT EXISTS knowledge ( -- Знания (составляющая часть компетенций) (есть в таблице, состявляемой аналитиками)
    knowledge_id BIGSERIAL PRIMARY KEY,
    knowledge TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS competencies ( -- Компетенции (есть в таблице, состявляемой аналитиками)
    competency_id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    experience TEXT,
    skills TEXT
);

CREATE TABLE IF NOT EXISTS knowledge_competence ( -- Связь между компетенцией и составляющим ее знанием
    knowledge_id BIGSERIAL REFERENCES knowledge(knowledge_id) ON DELETE CASCADE,
    competency_id BIGSERIAL REFERENCES competencies(competency_id) ON DELETE CASCADE,
    PRIMARY KEY (knowledge_id, competency_id)
);

CREATE TABLE IF NOT EXISTS professions ( -- Список доступных для выбора профессий.
    profession_id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT
);  

CREATE TABLE IF NOT EXISTS competence_profession ( -- Связь между профессией и требующейся для нее компетенции
    competency_id BIGSERIAL REFERENCES competencies(competency_id) ON DELETE CASCADE,
    profession_id BIGSERIAL REFERENCES professions(profession_id) ON DELETE CASCADE,
    PRIMARY KEY (profession_id, competency_id)
);

CREATE TABLE IF NOT EXISTS projects ( -- Учебный проект по проектной деятельности в конкретном семесте у конкретной команды.
    project_id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    result TEXT,
    life_scenario TEXT,
    main_technology_id BIGSERIAL REFERENCES knowledge(knowledge_id) -- ключевая технология == знание (knowledge)?
);

-- Например: Программная инженерия или Прикладная информатика?
CREATE TABLE IF NOT EXISTS educational_programs ( -- как связана с дисциплиной?
    educational_program_id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    belonging TEXT -- Что за поле?
);

-- Например: Физкультура или Программирование(скажем в 3 семестре)?
CREATE TABLE IF NOT EXISTS disciplines (
    discipline_id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    educational_program_id BIGSERIAL NOT NULL REFERENCES educational_programs(educational_program_id) ON DELETE CASCADE -- так?
);

CREATE TABLE IF NOT EXISTS teachers ( -- Список учителей
    teacher_id BIGSERIAL PRIMARY KEY,
    full_name TEXT NOT NULL
);

-- Например в разделе Программирование (в 3 семестре): Go от geekbrains, Java от УрФУ, Kotlin от ИТМО?
CREATE TABLE IF NOT EXISTS courses (
    course_id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    duration SMALLINT CHECK (duration > 0), -- в семестрах/курсах/месяцах? (предположу, что в семестрах)
    teacher_id BIGSERIAL REFERENCES teachers(teacher_id) ON DELETE RESTRICT, 
    discipline_id BIGSERIAL NOT NULL REFERENCES disciplines(discipline_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS study_groups ( -- Учебная группа (Например РИ-220942)
    study_group_id BIGSERIAL PRIMARY KEY,
    term SMALLINT CHECK (term > 0), 
    title TEXT NOT NULL, -- заменить на число?
    course_id BIGSERIAL REFERENCES courses(course_id) -- Не вся группа обучается на одном и том же курсе. (непонятно)
);

-- Портфолио проектов отдельного ученика в конкретном семестре. 
-- Возможно стоит убрать семестр, и сделать портфолио на период всего обучения?
-- А семестр указывать в каждой связи прокт-портфолио (таким образом на одном проекте смогут быть ученики разных семестров)
CREATE TABLE IF NOT EXISTS portfolios ( -- портфолио проектов конкретного ученика? 
    portfolio_id BIGSERIAL PRIMARY KEY, -- в конкретном семестре или на протяжении всего обучения?
    term SMALLINT CHECK (term > 0)
);

CREATE TABLE IF NOT EXISTS project_portfolio ( -- Связь между проектом и портофлио проектов конкретного студента.
    project_id BIGSERIAL REFERENCES projects(project_id) ON DELETE CASCADE,
    portfolio_id BIGSERIAL REFERENCES portfolios(portfolio_id) ON DELETE CASCADE,
    team_role TEXT, -- переставил столбец из портфолио. Могу создать отдельную таблицу, перечисляющую все роли.
    PRIMARY KEY (project_id, portfolio_id)
);

CREATE TABLE IF NOT EXISTS students ( -- Информация о конкретном студенте.
    student_id BIGSERIAL PRIMARY KEY,
    full_name TEXT NOT NULL,
    study_group_id BIGSERIAL NOT NULL REFERENCES study_groups(study_group_id) ON DELETE CASCADE,
    portfolio_id BIGSERIAL REFERENCES portfolios(portfolio_id)
);

CREATE TABLE IF NOT EXISTS trajectories ( -- Траектория конкретного студента в конкретном семестре.
    trajectory_id BIGSERIAL PRIMARY KEY,
    student_id BIGSERIAL NOT NULL REFERENCES students(student_id),
    term SMALLINT NOT NULL CHECK (term > 0),
    course_id BIGSERIAL REFERENCES courses(course_id), -- Опять же не понятно, как связана с курсов
    competency_id BIGSERIAL REFERENCES competencies(competency_id) -- Вероятно имеются ввиду компетенции, и стоит добавить исвязующую таблицу. 
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE trajectories, students, project_portfolio, portfolios, study_groups, courses, disciplines, 
educational_programs, teachers, projects, professions, knowledge_competence, competencies, knowledge;