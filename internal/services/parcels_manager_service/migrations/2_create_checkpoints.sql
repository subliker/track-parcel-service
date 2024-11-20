-- +goose Up
CREATE TABLE IF NOT EXISTS checkpoints (
    id SERIAL PRIMARY KEY,
    time TIMESTAMP NOT NULL,
    place VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    status PARCEL_STATUS NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS checkpoints;