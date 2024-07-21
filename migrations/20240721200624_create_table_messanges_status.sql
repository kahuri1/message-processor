-- +goose Up
-- +goose StatementBegin
CREATE TABLE if not exists message_status (
    status_id SERIAL PRIMARY KEY,
    message_id INT REFERENCES messages(message_id) ON DELETE CASCADE,
    user_id INT REFERENCES users(user_id) ON DELETE CASCADE,
    read_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists message_status;
-- +goose StatementEnd
