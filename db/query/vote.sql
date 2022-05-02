-- name: CreateVote :one
INSERT INTO votes (
  user_id, answer_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetVotes :many
SELECT * FROM votes
WHERE answer_id = $1;

-- name: DeleteVote :exec
DELETE FROM votes
WHERE id = $1;