// Code generated by sqlc. DO NOT EDIT.
// source: relationship.sql

package db

import (
	"context"
)

const createRelationship = `-- name: CreateRelationship :one
INSERT INTO relationships (
  follower_id, followed_id
) VALUES (
  $1, $2
)
RETURNING id, follower_id, followed_id, created_at, updated_at
`

type CreateRelationshipParams struct {
	FollowerID int32 `json:"follower_id"`
	FollowedID int32 `json:"followed_id"`
}

func (q *Queries) CreateRelationship(ctx context.Context, arg CreateRelationshipParams) (Relationship, error) {
	row := q.db.QueryRowContext(ctx, createRelationship, arg.FollowerID, arg.FollowedID)
	var i Relationship
	err := row.Scan(
		&i.ID,
		&i.FollowerID,
		&i.FollowedID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteRelationship = `-- name: DeleteRelationship :exec
DELETE FROM relationships
WHERE follower_id = $1
`

func (q *Queries) DeleteRelationship(ctx context.Context, followerID int32) error {
	_, err := q.db.ExecContext(ctx, deleteRelationship, followerID)
	return err
}