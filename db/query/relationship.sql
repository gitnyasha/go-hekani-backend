-- name: CreateRelationship :one
INSERT INTO relationships (
  follower_id, followed_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteRelationship :exec
DELETE FROM relationships
WHERE follower_id = $1;