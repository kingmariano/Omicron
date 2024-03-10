-- +goose Up
CREATE TABLE IF NOT EXISTS UnvalidatedUsers(
    id UUID PRIMARY KEY UNIQUE,
    created_at TIMESTAMP NOT NULL,
    whatsapp_number  VARCHAR(20) NOT NULL UNIQUE,
    display_name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS ValidatedUsers(
    userid UUID PRIMARY KEY,
    updated_at TIMESTAMP NOT NULL,
    whatsapp_number VARCHAR(20) NOT NULL REFERENCES UnvalidatedUsers(whatsapp_number),
    display_name TEXT NOT NULL,
    password TEXT NOT NULL,
    email TEXT NOT NULL,
    apikey VARCHAR(64) NOT NULL,
    updated_username BOOLEAN DEFAULT FALSE,
    updated_password BOOLEAN DEFAULT FALSE,
    loggedIn BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (userid) REFERENCES UnvalidatedUsers(id)
);
-- +goose Down
DROP TABLE ValidatedUsers;
DROP TABLE UnvalidatedUsers;