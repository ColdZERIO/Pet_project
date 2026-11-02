CREATE SCHEMA todo_app;

CREATE TABLE todo_app.users (
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    name VARCHAR (100) NOT NULL CHECK (char_length(name) BETWEEN 3 AND 100),
    phone_number VARCHAR(15) CHECK (
        phone_number ~ '^\+[0-9]+$'
        AND
        char_length(phone_number) BETWEEN 10 AND 15
    )
);

CREATE TABLE todo_app.tasks (
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    title VARCHAR(100) NOT NULL CHECK ( char_length(title) BETWEEN 1 AND 100),
    description VARCHAR(1000) CHECK ( char_length(description) BETWEEN 1 AND 1000),
    is_completed BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL,
    completed_at TIMESTAMP,

    CHECK (
        (is_completed=FALSE AND completed_at IS NULL)
        OR
        (is_completed=TRUE AND completed_at IS NOT NULL AND completed_at >= created_at)
    ),

    author_user_id INTEGER NOT NULL REFERENCES todo_app.users(id)
);