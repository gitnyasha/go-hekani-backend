-- name: CreateUserQuestionRelation :one
INSERT INTO user_question_relations (
  follower_id, followed_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteUserQuestionRelation :exec
DELETE FROM user_question_relations
WHERE follower_id = $1;