-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetPostByTitle :one
SELECT * FROM posts WHERE title = $1;

-- name: GetPostByURL :one
SELECT * FROM posts WHERE url = $1;

-- name: GetPostByFeed :many
SELECT * FROM posts WHERE feed_id = $1;

-- name: GetPostsOfUser :many
SELECT p.*, f.name AS feed_name
FROM posts p
INNER JOIN feeds f ON f.id = p.feed_id
INNER JOIN feed_follows ff ON ff.feed_id = p.feed_id
WHERE f.user_id = $1
ORDER BY p.created_at DESC
LIMIT $2;