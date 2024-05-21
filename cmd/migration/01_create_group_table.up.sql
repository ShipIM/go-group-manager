CREATE TYPE study_grade AS ENUM ('bachelor', 'master', 'postgraduate');

CREATE TABLE IF NOT EXISTS _group (
    name VARCHAR PRIMARY KEY,
    course INTEGER NOT NULL,
    grade study_grade NOT NULL
    CONSTRAINT valid_course CHECK (course > 0 AND course < 5)
);
