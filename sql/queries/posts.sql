-- name: CreatePost :exec
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
    gen_random_uuid(),
    Now(),
    Now(),
    $1,
    $2,
    $3,
    $4,
    $5
);
-- name: GetPosts :many
SELECT * FROM posts LIMIT $1;
