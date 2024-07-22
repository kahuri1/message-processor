-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists messages
(
    message_id   SERIAL PRIMARY KEY,
    sender_id    INT REFERENCES users (user_id) ON DELETE CASCADE,
    recipient_id INT REFERENCES users (user_id) ON DELETE CASCADE,
    content      TEXT NOT NULL,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    processed_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists messages;
-- +goose StatementEnd
