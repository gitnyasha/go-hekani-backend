-- name: CreateUserArticleRelationship :one
INSERT INTO user_article_relationships (
  follower_id, followed_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteUserArticleRelationship :exec
DELETE FROM user_article_relationships
WHERE follower_id = $1;