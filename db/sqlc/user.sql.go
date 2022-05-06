// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package db

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  name, email, hashed_password, bio, birth, image
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id, name, email, hashed_password, bio, birth, created_at, updated_at, image
`

type CreateUserParams struct {
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"hashed_password"`
	Bio            string    `json:"bio"`
	Birth          time.Time `json:"birth"`
	Image          string    `json:"image"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Name,
		arg.Email,
		arg.HashedPassword,
		arg.Bio,
		arg.Birth,
		arg.Image,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.Bio,
		&i.Birth,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Image,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, name, email, hashed_password, bio, birth, created_at, updated_at, image FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.Bio,
		&i.Birth,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Image,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, name, email, hashed_password, bio, birth, created_at, updated_at, image FROM users
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.HashedPassword,
			&i.Bio,
			&i.Birth,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Image,
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

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET name = $1, bio = $2, birth = $3, updated_at = NOW()
WHERE id = $4
RETURNING id, name, email, hashed_password, bio, birth, created_at, updated_at, image
`

type UpdateUserParams struct {
	Name  string    `json:"name"`
	Bio   string    `json:"bio"`
	Birth time.Time `json:"birth"`
	ID    int64     `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.Name,
		arg.Bio,
		arg.Birth,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.HashedPassword,
		&i.Bio,
		&i.Birth,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Image,
	)
	return i, err
}
