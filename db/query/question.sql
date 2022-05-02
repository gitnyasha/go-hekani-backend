-- name: CreateQuestion :one
INSERT INTO questions (
  title, user_id, question_category_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetQuestion :one
SELECT * FROM questions
WHERE id = $1 LIMIT 1;

-- name: ListQuestions :many
SELECT * FROM questions
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateQuestion :exec
UPDATE questions
SET title = $1, updated_at = NOW()
WHERE id = $2;

-- name: DeleteQuestion :exec
DELETE FROM questions
WHERE id = $1;