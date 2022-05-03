-- name: CreateArticle :one
INSERT INTO Articles (
  title, link, image, article_category_id, user_id
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetArticle :one
SELECT * 
FROM articles
WHERE id = $1 
LIMIT 1;

-- name: ListArticles :many
SELECT * FROM articles
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteArticle :exec
DELETE FROM articles
WHERE id = $1;