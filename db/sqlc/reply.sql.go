// Code generated by sqlc. DO NOT EDIT.
// source: reply.sql

package db

import (
	"context"
)

const createReply = `-- name: CreateReply :one
INSERT INTO replies (
  user_id, article_id, title
) VALUES (
  $1, $2, $3
)
RETURNING id, title, article_id, user_id, created_at, updated_at
`

type CreateReplyParams struct {
	UserID    int32  `json:"user_id"`
	ArticleID int32  `json:"article_id"`
	Title     string `json:"title"`
}

func (q *Queries) CreateReply(ctx context.Context, arg CreateReplyParams) (Reply, error) {
	row := q.db.QueryRowContext(ctx, createReply, arg.UserID, arg.ArticleID, arg.Title)
	var i Reply
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.ArticleID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteReply = `-- name: DeleteReply :exec
DELETE FROM replies
WHERE id = $1
`

func (q *Queries) DeleteReply(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteReply, id)
	return err
}

const getReply = `-- name: GetReply :one
SELECT id, title, article_id, user_id, created_at, updated_at FROM replies
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetReply(ctx context.Context, id int64) (Reply, error) {
	row := q.db.QueryRowContext(ctx, getReply, id)
	var i Reply
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.ArticleID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listReplies = `-- name: ListReplies :many
SELECT id, title, article_id, user_id, created_at, updated_at FROM replies
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListRepliesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListReplies(ctx context.Context, arg ListRepliesParams) ([]Reply, error) {
	rows, err := q.db.QueryContext(ctx, listReplies, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Reply{}
	for rows.Next() {
		var i Reply
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.ArticleID,
			&i.UserID,
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

const updateReply = `-- name: UpdateReply :one
UPDATE replies
SET title = $1, updated_at = NOW()
WHERE id = $2
RETURNING id, title, article_id, user_id, created_at, updated_at
`

type UpdateReplyParams struct {
	Title string `json:"title"`
	ID    int64  `json:"id"`
}

func (q *Queries) UpdateReply(ctx context.Context, arg UpdateReplyParams) (Reply, error) {
	row := q.db.QueryRowContext(ctx, updateReply, arg.Title, arg.ID)
	var i Reply
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.ArticleID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
