-- +goose Up
CREATE TABLE users(
    id uuid primary key,
    created_at timestamp NOT NULL ,
    updated_at timestamp NOT NULL ,
    name VARCHAR NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE users;