-- name: CreateComment :one
INSERT INTO comments (
  user_id, answer_id, title
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetComment :one
SELECT * FROM comments
WHERE id = $1 LIMIT 1;

-- name: ListComments :many
SELECT * FROM comments
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateComment :one
UPDATE comments
SET title = $1, updated_at = NOW()
WHERE id = $2
RETURNING *;

-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = $1;