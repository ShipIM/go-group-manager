CREATE TABLE IF NOT EXISTS student (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    surname VARCHAR NOT NULL,
    patronymic VARCHAR,
    age INTEGER NOT NULL,
    group_name VARCHAR NOT NULL REFERENCES _group(name) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT valid_age CHECK (age > 0 AND age < 100)
);
