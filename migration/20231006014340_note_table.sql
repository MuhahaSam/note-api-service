-- +goose Up
-- +goose StatementBegin
create table note (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    title text,
    author text not null,
    text text,
    created_at timestamp not null default now(),
    updated_at timestamp,
    deleted_at timestamp
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
drop table note;

-- +goose StatementEnd