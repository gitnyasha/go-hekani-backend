-- name: CreateQuestionCategory :one
INSERT INTO question_categories (
  name
) VALUES (
  $1
)
RETURNING *;

-- name: GetQuestionCategory :one
SELECT * FROM question_categories
WHERE id = $1 LIMIT 1;

-- name: ListQuestionCategories :many
SELECT * FROM question_categories
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateQuestionCategory :one
UPDATE question_categories
SET name = $1, updated_at = NOW()
WHERE id = $2
RETURNING *;

-- name: DeleteQuestionCategory :exec
DELETE FROM question_categories
WHERE id = $1;