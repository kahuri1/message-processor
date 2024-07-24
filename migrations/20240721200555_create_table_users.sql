-- Active: 1721598613974@@127.0.0.1@5432@mydb
-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(255)  NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
-- +goose StatementEnd
