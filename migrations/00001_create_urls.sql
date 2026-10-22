-- +goose Up
CREATE TABLE urls (
    token TEXT PRIMARY KEY NOT NULL,
    original_url TEXT NOT NULL,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE urls;
