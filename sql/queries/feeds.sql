-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeedByName :one
SELECT * FROM feeds WHERE name = $1;

-- name: GetFeedByURL :one
SELECT * FROM feeds WHERE url = $1;

-- name: GetFeedByUser :many
SELECT * FROM feeds WHERE user_id = $1;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedsWithUsers :many
SELECT sqlc.embed(f), sqlc.embed(u) FROM feeds f INNER JOIN users u ON u.id = f.user_id;

-- name: MarkFeedFetched :one
UPDATE feeds
SET last_fetched_at = NOW() AT TIME ZONE 'utc', updated_at = NOW() AT TIME ZONE 'utc'
WHERE id = $1
RETURNING *;

-- name: GetNextFeedToFetch :one
SELECT *
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST LIMIT 1;