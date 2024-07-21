-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists statistics (
    stat_id SERIAL PRIMARY KEY,
    total_messages INT NOT NULL DEFAULT 0,
    processed_messages INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists statistics;
-- +goose StatementEnd
