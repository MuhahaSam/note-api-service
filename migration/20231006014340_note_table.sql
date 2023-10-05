-- +goose Up
-- +goose StatementBegin
create table note (
    id bigserial primary key,
    title text,
    author text not null,
    text text,
    created_at timestamp not null default now(),
    deleted_at timestamp
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
drop table note;

-- +goose StatementEnd