-- +goose Up
ALTER TABLE IF EXISTS users
ADD api_key VARCHAR(64) NOT NULL UNIQUE DEFAULT encode(sha256(random()::text::bytea), 'hex');
-- +goose Down
ALTER TABLE IF EXISTS users
DROP IF EXISTS api_key;
