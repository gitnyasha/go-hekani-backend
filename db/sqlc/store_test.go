package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRelationshipTx(t *testing.T) {
	store := NewStore(testDB)

	n := 5
	errs := make(chan error)
	results := make(chan RelationshipTxResult)

	for i := 0; i < n; i++ {
		user1 := createRandomUser(t)
		user2 := createRandomUser(t)

		go func() {
			result, err := store.RelationshipTx(context.Background(), RelationshipTxParams{
				FollowerID: int32(user1.ID),
				FollowedID: int32(user2.ID),
			})

			errs <- err
			results <- result
		}()

		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		relationship := result.Relationship
		require.Equal(t, int32(user1.ID), relationship.FollowerID)
		require.Equal(t, int32(user2.ID), relationship.FollowedID)
		require.NotZero(t, relationship.ID)
		require.NotZero(t, relationship.CreatedAt)
		require.NotZero(t, relationship.UpdatedAt)
	}

}
