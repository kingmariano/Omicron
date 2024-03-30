-- +goose Up
CREATE TABLE Commands (
    id SERIAL PRIMARY KEY,
    command_name TEXT NOT NULL UNIQUE,
    instruction TEXT NOT NULL
);


-- +goose Down
DROP TABLE Commands;