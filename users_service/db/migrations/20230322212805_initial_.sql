-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id  serial      not null,
    fio text not null,
    phone text not null,
    age int not null,
    level int not null default 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
