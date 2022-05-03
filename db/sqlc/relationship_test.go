package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomRelationship(t *testing.T) Relationship {
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)

	arg := CreateRelationshipParams{
		FollowerID: int32(user1.ID),
		FollowedID: int32(user2.ID),
	}

	relationship, err := testQueries.CreateRelationship(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, relationship)
	require.Equal(t, arg.FollowedID, relationship.FollowedID)
	require.Equal(t, arg.FollowerID, relationship.FollowerID)
	require.NotZero(t, relationship.ID)
	require.NotZero(t, relationship.CreatedAt)
	return relationship
}

func TestCreateRelationship(t *testing.T) {
	createRandomRelationship(t)
}
