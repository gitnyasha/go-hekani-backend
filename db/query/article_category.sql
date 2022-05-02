
-- name: CreateArticleCategory :one
INSERT INTO article_categories (
  name
) VALUES (
  $1
)
RETURNING *;

-- name: GetArticleCategory :one
SELECT * FROM article_categories
WHERE id = $1 LIMIT 1;

-- name: ListArticleCategories :many
SELECT * FROM article_categories
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateArticleCategory :exec
UPDATE article_categories
SET name = $1, updated_at = NOW()
WHERE id = $2;

-- name: DeleteArticleCategory :exec
DELETE FROM article_categories
WHERE id = $1;