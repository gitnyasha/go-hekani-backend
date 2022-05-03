package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomUserArticleRelationship(t *testing.T) UserArticleRelationship {
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)

	arg := CreateUserArticleRelationshipParams{
		FollowerID: int32(user1.ID),
		FollowedID: int32(user2.ID),
	}

	user_article_relationship, err := testQueries.CreateUserArticleRelationship(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user_article_relationship)
	require.Equal(t, arg.FollowedID, user_article_relationship.FollowedID)
	require.Equal(t, arg.FollowerID, user_article_relationship.FollowerID)
	require.NotZero(t, user_article_relationship.ID)
	require.NotZero(t, user_article_relationship.CreatedAt)
	return user_article_relationship
}

func TestCreateUserArticleRelationship(t *testing.T) {
	createRandomUserArticleRelationship(t)
}
