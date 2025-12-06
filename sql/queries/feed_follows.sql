-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
      gen_random_uuid(),
      Now(),
      Now(),
      $1,
      $2
    )
    RETURNING *
)
SELECT
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
INNER JOIN feeds ON feeds.id = inserted_feed_follow.feed_id
INNER JOIN users ON users.id = inserted_feed_follow.user_id;

-- name: GetFeedFollowsForUser :many
WITH user_feed_follow AS (
    SELECT * FROM feed_follows where feed_follows.user_id = $1
)
SELECT
    user_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM user_feed_follow
INNER JOIN feeds ON feeds.id = user_feed_follow.feed_id
INNER JOIN users ON users.id = user_feed_follow.user_id;

-- name: Unfollow :exec
DELETE FROM feed_follows
  WHERE user_id = $1 
  AND feed_id = $2;
