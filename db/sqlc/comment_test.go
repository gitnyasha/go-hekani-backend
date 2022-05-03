package db

import (
	"context"
	"testing"

	"github.com/gitnyasha/go-hekani-backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomComment(t *testing.T) Comment {
	user := createRandomUser(t)
	answer := createRandomAnswer(t)

	arg := CreateCommentParams{
		UserID:   int32(user.ID),
		AnswerID: int32(answer.ID),
		Title:    util.RandomString(),
	}

	comment, err := testQueries.CreateComment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, comment)

	require.Equal(t, arg.UserID, comment.UserID)
	require.Equal(t, arg.AnswerID, comment.AnswerID)
	require.Equal(t, arg.Title, comment.Title)

	require.NotZero(t, comment.ID)
	require.NotZero(t, comment.CreatedAt)

	return comment
}

// create Comment
func TestCreateComment(t *testing.T) {
	createRandomComment(t)
}
