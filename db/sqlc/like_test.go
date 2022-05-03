package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomLike(t *testing.T) Like {
	user := createRandomUser(t)
	article := createRandomArticle(t)

	arg := CreateLikeParams{
		UserID:    int32(user.ID),
		ArticleID: int32(article.ID),
	}

	like, err := testQueries.CreateLike(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, like)

	require.Equal(t, arg.UserID, like.UserID)
	require.Equal(t, arg.ArticleID, like.ArticleID)

	require.NotZero(t, like.ID)
	require.NotZero(t, like.CreatedAt)

	return like
}

// create Like
func TestCreateLike(t *testing.T) {
	createRandomLike(t)
}
