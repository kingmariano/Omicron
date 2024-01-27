-- +goose Up
CREATE TABLE IF NOT EXISTS Unregisteredusers(
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    whatsapp_number TEXT NOT NULL,
    display_name TEXT NOT NULL
);
-- +goose Down
DROP TABLE Unregisteredusers;