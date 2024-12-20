-- +goose Up
CREATE TABLE IF NOT EXISTS managers (
    id SERIAL PRIMARY KEY,
    telegram_id BIGINT NOT NULL UNIQUE,
    full_name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20),
    email VARCHAR(320) NOT NULL,
    company VARCHAR(255),
    api_token VARCHAR(24) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_managers_telegram_id ON managers(telegram_id);
CREATE INDEX IF NOT EXISTS idx_managers_api_token ON managers(api_token);

-- +goose Down
DROP TABLE IF EXISTS managers;