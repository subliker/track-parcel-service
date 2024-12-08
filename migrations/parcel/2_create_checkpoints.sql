-- +goose Up
CREATE TABLE IF NOT EXISTS checkpoints (
    id SERIAL PRIMARY KEY,
    time TIMESTAMP NOT NULL,
    place VARCHAR(255),
    description VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    parcel_track_number VARCHAR(25) NOT NULL,
    parcel_status PARCEL_STATUS NOT NULL,
    FOREIGN KEY (parcel_track_number) REFERENCES parcels(track_number) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS checkpoints;