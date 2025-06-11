-- +goose Up
create table chat (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    usernames TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

-- +goose Down
drop table chat;