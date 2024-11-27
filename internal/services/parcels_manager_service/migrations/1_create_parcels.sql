-- +goose Up
CREATE TYPE PARCEL_STATUS AS ENUM('UNKNOWN', 'PENDING', 'IN_TRANSIT', 'DELIVERED', 'CANCELED');

CREATE TABLE IF NOT EXISTS parcels (
    id SERIAL PRIMARY KEY,
    track_number VARCHAR(25) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    manager_id BIGINT NOT NULL,
    recipient VARCHAR(255) NOT NULL,
    arrival_address VARCHAR(255) NOT NULL,
    forecast_date TIMESTAMP NOT NULL,
    description VARCHAR(255) NOT NULL,
    status PARCEL_STATUS NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_parcels_track_number ON parcels(track_number);

-- +goose Down
DROP TABLE IF EXISTS parcels;

DROP TYPE IF EXISTS PARCEL_STATUS;