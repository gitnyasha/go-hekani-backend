-- name: CreateLike :one
INSERT INTO likes (
  user_id, article_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetLikes :many
SELECT * FROM likes
WHERE article_id = $1;

-- name: DeleteLike :exec
DELETE FROM likes
WHERE id = $1;