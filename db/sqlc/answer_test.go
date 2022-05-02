package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAnswer(t *testing.T) {
	arg := CreateAnswerParams{
		UserID:     1,
		QuestionID: 1,
		Title:      "Test Answer",
	}

	answer, err := testQueries.CreateAnswer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, answer)

	require.Equal(t, arg.UserID, answer.UserID)
	require.Equal(t, arg.QuestionID, answer.QuestionID)
	require.Equal(t, arg.Title, answer.Title)

	require.NotZero(t, answer.ID)
	require.NotZero(t, answer.CreatedAt)
}
