-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE knowledge ( -- Знания (составляющая часть компетенций) (есть в таблице, состявляемой аналитиками)
    knowledge_id SERIAL PRIMARY KEY,
    title TEXT NOT NULL UNIQUE
);

CREATE TABLE competencies ( -- Компетенции (есть в таблице, состявляемой аналитиками)
    competency_id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    experience TEXT,
    skills TEXT
);

CREATE TABLE knowledge_competency ( -- Связь между компетенцией и составляющим ее знанием
    knowledge_id SERIAL REFERENCES knowledge(knowledge_id) ON DELETE CASCADE,
    competency_id SERIAL REFERENCES competencies(competency_id) ON DELETE CASCADE,
    PRIMARY KEY (knowledge_id, competency_id)
);

CREATE TABLE professions ( -- Список доступных для выбора профессий.
    profession_id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT
);  

CREATE TABLE competency_profession ( -- Связь между профессией и требующейся для нее компетенции
    competency_id SERIAL REFERENCES competencies(competency_id) ON DELETE CASCADE,
    profession_id SERIAL REFERENCES professions(profession_id) ON DELETE CASCADE,
    PRIMARY KEY (profession_id, competency_id)
);

CREATE TABLE projects ( -- Учебный проект по проектной деятельности в конкретном семесте у конкретной команды.
    project_id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    result TEXT,
    life_scenario TEXT,
    main_technology_id SERIAL REFERENCES knowledge(knowledge_id) -- ключевая технология == знание (knowledge)?
);

-- Например: Программная инженерия или Прикладная информатика?
CREATE TABLE educational_programs ( -- как связана с дисциплиной?
    educational_program_id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    belonging TEXT -- Что за поле?
);

-- Например: Физкультура или Программирование(скажем в 3 семестре)?
CREATE TABLE disciplines (
    discipline_id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    educational_program_id SERIAL NOT NULL REFERENCES educational_programs(educational_program_id) ON DELETE CASCADE -- так?
);

CREATE TABLE teachers ( -- Список учителей
    teacher_id SERIAL PRIMARY KEY,
    full_name TEXT NOT NULL
);

-- Например в разделе Программирование (в 3 семестре): Go от geekbrains, Java от УрФУ, Kotlin от ИТМО?
CREATE TABLE courses (
    course_id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    -- duration SMALLINT CHECK (duration > 0), -- в семестрах
    teacher_id SERIAL REFERENCES teachers(teacher_id) ON DELETE RESTRICT, 
    discipline_id SERIAL NOT NULL REFERENCES disciplines(discipline_id) ON DELETE CASCADE
);

CREATE TABLE study_groups ( -- Учебная группа (Например РИ-220942)
    study_group_id SERIAL PRIMARY KEY,
    semester SMALLINT CHECK (semester > 0), 
    title TEXT NOT NULL, -- заменить на число?
    course_id SERIAL REFERENCES courses(course_id) -- Не вся группа обучается на одном и том же курсе. (непонятно)
);

-- Портфолио проектов отдельного ученика в конкретном семестре. 
-- Возможно стоит убрать семестр, и сделать портфолио на период всего обучения?
-- А семестр указывать в каждой связи прокт-портфолио (таким образом на одном проекте смогут быть ученики разных семестров)
CREATE TABLE portfolios ( -- портфолио проектов конкретного ученика? 
    portfolio_id SERIAL PRIMARY KEY -- в конкретном семестре или на протяжении всего обучения?
);

CREATE TABLE project_portfolio ( -- Связь между проектом и портофлио проектов конкретного студента.
    project_id SERIAL REFERENCES projects(project_id) ON DELETE CASCADE,
    portfolio_id SERIAL REFERENCES portfolios(portfolio_id) ON DELETE CASCADE,
    team_role TEXT, -- переставил столбец из портфолио. Могу создать отдельную таблицу, перечисляющую все роли.
    semester SMALLINT CHECK (semester > 0),
    -- compitence_id -- список 
    PRIMARY KEY (project_id, portfolio_id)
);

CREATE TABLE students ( -- Информация о конкретном студенте.
    student_id SERIAL PRIMARY KEY,
    full_name TEXT NOT NULL,
    study_group_id SERIAL NOT NULL REFERENCES study_groups(study_group_id) ON DELETE CASCADE,
    -- перенести в учебную группу
    portfolio_id SERIAL REFERENCES portfolios(portfolio_id)
);

CREATE TABLE trajectories ( -- Траектория конкретного студента в конкретном семестре.
    trajectory_id SERIAL PRIMARY KEY,
    student_id SERIAL NOT NULL REFERENCES students(student_id),
    semester SMALLINT NOT NULL CHECK (semester > 0),
    course_id SERIAL REFERENCES courses(course_id), -- Опять же не понятно, как связана с курсов
    competency_id SERIAL REFERENCES competencies(competency_id) -- Вероятно имеются ввиду компетенции, и стоит добавить исвязующую таблицу. 
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE trajectories;
DROP TABLE students;
DROP TABLE project_portfolio;
DROP TABLE portfolios;
DROP TABLE study_groups;
DROP TABLE courses;
DROP TABLE teachers;
DROP TABLE disciplines;
DROP TABLE educational_programs;
DROP TABLE projects;
DROP TABLE competency_profession;
DROP TABLE professions;
DROP TABLE knowledge_competency;
DROP TABLE competencies;
DROP TABLE knowledge;
