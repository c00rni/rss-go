-- +goose Up
ALTER TABLE IF EXISTS feeds
ADD last_fetched_at TIMESTAMP NULL;

-- +goose Down
ALTER TABLE IF EXISTS feeds
DROP IF EXISTS last_fetched_at;
