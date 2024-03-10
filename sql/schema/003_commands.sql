-- +goose Up
CREATE TABLE Commands (
    id SERIAL PRIMARY KEY,
    command_name TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL
);


-- +goose Down
DROP TABLE Commands;