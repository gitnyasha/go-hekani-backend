// Code generated by sqlc. DO NOT EDIT.
// source: article_category.sql

package db

import (
	"context"
)

const createArticleCategory = `-- name: CreateArticleCategory :one
INSERT INTO article_categories (
  name
) VALUES (
  $1
)
RETURNING id, name, created_at, updated_at
`

type CreateArticleCategoryParams struct {
	Name  string `json:"name"`
}

func (q *Queries) CreateArticleCategory(ctx context.Context, arg CreateArticleCategoryParams) (ArticleCategory, error) {
	row := q.db.QueryRowContext(ctx, createArticleCategory, arg.Name)
	var i ArticleCategory
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteArticleCategory = `-- name: DeleteArticleCategory :exec
DELETE FROM article_categories
WHERE id = $1
`

func (q *Queries) DeleteArticleCategory(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteArticleCategory, id)
	return err
}

const getArticleCategory = `-- name: GetArticleCategory :one
SELECT id, name, created_at, updated_at FROM article_categories
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetArticleCategory(ctx context.Context, id int64) (ArticleCategory, error) {
	row := q.db.QueryRowContext(ctx, getArticleCategory, id)
	var i ArticleCategory
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listArticleCategories = `-- name: ListArticleCategories :many
SELECT id, name, created_at, updated_at FROM article_categories
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListArticleCategoriesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListArticleCategories(ctx context.Context, arg ListArticleCategoriesParams) ([]ArticleCategory, error) {
	rows, err := q.db.QueryContext(ctx, listArticleCategories, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ArticleCategory
	for rows.Next() {
		var i ArticleCategory
		if err := rows.Scan(
			&i.ID,
			&i.Name,
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

const updateArticleCategory = `-- name: UpdateArticleCategory :one
UPDATE article_categories
SET name = $1, updated_at = NOW()
WHERE id = $2
RETURNING id, name, created_at, updated_at
`

type UpdateArticleCategoryParams struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
}

func (q *Queries) UpdateArticleCategory(ctx context.Context, arg UpdateArticleCategoryParams) (ArticleCategory, error) {
	row := q.db.QueryRowContext(ctx, updateArticleCategory, arg.Name, arg.ID)
	var i ArticleCategory
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
