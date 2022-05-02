-- name: CreateAnswer :one
INSERT INTO answers (
  user_id, question_id, title
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetAnswer :one
SELECT * FROM answers
WHERE id = $1 LIMIT 1;

-- name: ListAnswers :many
SELECT * FROM answers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateAnswer :exec
UPDATE answers
SET title = $1, updated_at = NOW()
WHERE id = $2;

-- name: DeleteAnswer :exec
DELETE FROM answers
WHERE id = $1;