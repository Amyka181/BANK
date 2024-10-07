-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS bank.users (
    id integer not null unique,
    balance integer not null
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS bank.users;
-- +goose StatementEnd
