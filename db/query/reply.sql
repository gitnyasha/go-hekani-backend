-- name: CreateReply :one
INSERT INTO replies (
  user_id, article_id, title
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetReply :one
SELECT * FROM replies
WHERE id = $1 LIMIT 1;

-- name: ListReplies :many
SELECT * FROM replies
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateReply :one
UPDATE replies
SET title = $1, updated_at = NOW()
WHERE id = $2
RETURNING *;

-- name: DeleteReply :exec
DELETE FROM replies
WHERE id = $1;