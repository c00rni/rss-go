-- name: FollowFeed :one
INSERT INTO feedFollowed (feed_id, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $4)
RETURNING *;


