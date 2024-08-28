-- name: FollowFeed :one
INSERT INTO feedFollowed (feed_id, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: Unfollow :exec
DELETE FROM feedFollowed
WHERE feed_id = $1
AND user_id = $2;

-- name: GetUserFeeds :many
SELECT *
FROM feeds
WHERE id IN (SELECT feed_id from feedFollowed as f where f.user_id = $1);
