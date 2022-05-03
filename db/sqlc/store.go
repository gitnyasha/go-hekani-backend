package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("%s; %s", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type RelationshipTxParams struct {
	FollowerID int32 `json:"follower_id"`
	FollowedID int32 `json:"followed_id"`
}

type RelationshipTxResult struct {
	Relationship Relationship `json:"follow"`
	Follower     User         `json:"follower_id"`
	Followed     User         `json:"followed_id"`
}

func (store *Store) RelationshipTx(ctx context.Context, arg RelationshipTxParams) (RelationshipTxResult, error) {
	var result RelationshipTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Relationship, err = q.CreateRelationship(ctx, CreateRelationshipParams{
			FollowerID: arg.FollowerID,
			FollowedID: arg.FollowedID,
		})

		if err != nil {
			return err
		}
		return nil
	})

	return result, err
}
