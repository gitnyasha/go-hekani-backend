// Code generated by sqlc. DO NOT EDIT.
// source: answer.sql

package db

import (
	"context"
)

const createAnswer = `-- name: CreateAnswer :one
INSERT INTO answers (
  user_id, question_id, title
) VALUES (
  $1, $2, $3
)
RETURNING id, user_id, question_id, title, created_at, updated_at
`

type CreateAnswerParams struct {
	UserID     int32  `json:"user_id"`
	QuestionID int32  `json:"question_id"`
	Title      string `json:"title"`
}

func (q *Queries) CreateAnswer(ctx context.Context, arg CreateAnswerParams) (Answer, error) {
	row := q.db.QueryRowContext(ctx, createAnswer, arg.UserID, arg.QuestionID, arg.Title)
	var i Answer
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.QuestionID,
		&i.Title,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteAnswer = `-- name: DeleteAnswer :exec
DELETE FROM answers
WHERE id = $1
`

func (q *Queries) DeleteAnswer(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAnswer, id)
	return err
}

const getAnswer = `-- name: GetAnswer :one
SELECT id, user_id, question_id, title, created_at, updated_at FROM answers
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAnswer(ctx context.Context, id int64) (Answer, error) {
	row := q.db.QueryRowContext(ctx, getAnswer, id)
	var i Answer
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.QuestionID,
		&i.Title,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAnswers = `-- name: ListAnswers :many
SELECT id, user_id, question_id, title, created_at, updated_at FROM answers
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListAnswersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAnswers(ctx context.Context, arg ListAnswersParams) ([]Answer, error) {
	rows, err := q.db.QueryContext(ctx, listAnswers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Answer{}
	for rows.Next() {
		var i Answer
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.QuestionID,
			&i.Title,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAnswer = `-- name: UpdateAnswer :one
UPDATE answers
SET title = $1, updated_at = NOW()
WHERE id = $2
RETURNING id, user_id, question_id, title, created_at, updated_at
`

type UpdateAnswerParams struct {
	Title string `json:"title"`
	ID    int64  `json:"id"`
}

func (q *Queries) UpdateAnswer(ctx context.Context, arg UpdateAnswerParams) (Answer, error) {
	row := q.db.QueryRowContext(ctx, updateAnswer, arg.Title, arg.ID)
	var i Answer
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.QuestionID,
		&i.Title,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
