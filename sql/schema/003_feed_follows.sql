-- +goose Up
CREATE TABLE feed_follows(
    id uuid primary key,
    created_at timestamp NOT NULL ,
    updated_at timestamp NOT NULL ,
    user_id uuid NOT NULL,
    CONSTRAINT "fk_user_id" FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    feed_id uuid NOT NULL,
    CONSTRAINT "fk_feed_id" FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE,
CONSTRAINT "un_user_feed_id" UNIQUE (user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;