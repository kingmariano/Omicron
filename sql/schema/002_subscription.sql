-- +goose Up
CREATE TYPE subscription_status_enum AS ENUM ('Active', 'Expired');
CREATE TYPE subscription_tier_enum AS ENUM('Basic', 'Pro', 'Free-Trial');
CREATE TABLE Subscription(
    subscription_id UUID PRIMARY KEY,
    userid UUID UNIQUE,
    expiry_date DATE NOT NULL,
    subscription_status subscription_status_enum DEFAULT 'Active',
    subscription_tier subscription_tier_enum DEFAULT 'Free-Trial',
    FOREIGN KEY (userid) REFERENCES UnvalidatedUsers(id)
);

-- +goose Down
DROP TABLE Subscription;
DROP TYPE subscription_tier_enum;
DROP TYPE subscription_status_enum;