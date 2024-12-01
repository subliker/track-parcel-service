-- +goose Up
CREATE TABLE IF NOT EXISTS subscriptions (
    id SERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    parcel_track_number VARCHAR(25) NOT NULL,
    FOREIGN KEY (parcel_track_number) REFERENCES parcels(track_number) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS subscriptions;