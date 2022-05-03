package db

import (
	"context"
	"testing"

	"github.com/gitnyasha/go-hekani-backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomReply(t *testing.T) Reply {
	user := createRandomUser(t)
	article := createRandomArticle(t)

	arg := CreateReplyParams{
		UserID:    int32(user.ID),
		ArticleID: int32(article.ID),
		Title:     util.RandomString(),
	}

	reply, err := testQueries.CreateReply(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, reply)

	require.Equal(t, arg.UserID, reply.UserID)
	require.Equal(t, arg.ArticleID, reply.ArticleID)
	require.Equal(t, arg.Title, reply.Title)

	require.NotZero(t, reply.ID)
	require.NotZero(t, reply.CreatedAt)

	return reply
}

// create Reply
func TestCreateReply(t *testing.T) {
	createRandomReply(t)
}
