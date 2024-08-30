-- name: CreatePost :one
INSERT INTO post (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetUserPosts :many
SELECT * FROM post AS p
WHERE p.feed_id IN (
    SELECT feed_id FROM feedFollowed AS f
    WHERE f.user_id = $1);
