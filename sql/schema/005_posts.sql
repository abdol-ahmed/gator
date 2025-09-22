-- +goose Up
CREATE TABLE posts(
    id uuid primary key,
    created_at timestamp NOT NULL ,
    updated_at timestamp NOT NULL ,
    title VARCHAR,
    url VARCHAR NOT NULL UNIQUE,
    description VARCHAR,
    published_at timestamp,
    feed_id uuid NOT NULL,
    CONSTRAINT "fk_feed_id" FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;