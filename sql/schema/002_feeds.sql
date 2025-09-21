-- +goose Up
CREATE TABLE feeds(
    id uuid primary key,
    created_at timestamp NOT NULL ,
    updated_at timestamp NOT NULL ,
    name VARCHAR NOT NULL UNIQUE,
    url VARCHAR NOT NULL UNIQUE,
    user_id uuid NOT NULL,
    CONSTRAINT "fk_user_id" FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;