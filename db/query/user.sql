-- name: CreateUser :one
INSERT INTO users (
  name, email, hashed_password, bio, birth, image
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :exec
UPDATE users
SET name = $1, bio = $2, birth = $3, updated_at = NOW()
WHERE id = $4;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;