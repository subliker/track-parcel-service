-- +goose Up
CREATE TABLE IF NOT EXISTS checkpoints (
    id SERIAL PRIMARY KEY,
    time TIMESTAMP NOT NULL,
    place VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    FOREIGN KEY (parcel_track_number) REFERENCES parcels(track_number) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS checkpoints;