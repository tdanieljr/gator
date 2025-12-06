-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    gen_random_uuid(),
    Now(),
    Now(),
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;
-- name: ResetFeeds :exec
DELETE FROM feeds;
-- name: GetFeedByUrl :one
SELECT * FROM feeds where url = $1;
-- name: MarkFeedFetched :exec
UPDATE feeds SET last_fetched_at = Now() where id = $1;
-- name: GetNextFeedToFetch :one
SELECT * from feeds ORDER BY last_fetched_at ASC NULLS FIRST LIMIT 1;
