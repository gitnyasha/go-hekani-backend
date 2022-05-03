package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomVote(t *testing.T) Vote {
	user := createRandomUser(t)
	answer := createRandomAnswer(t)

	arg := CreateVoteParams{
		UserID:   int32(user.ID),
		AnswerID: int32(answer.ID),
	}

	vote, err := testQueries.CreateVote(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, vote)

	require.Equal(t, arg.UserID, vote.UserID)
	require.Equal(t, arg.AnswerID, vote.AnswerID)

	require.NotZero(t, vote.ID)
	require.NotZero(t, vote.CreatedAt)

	return vote
}

// create Vote
func TestCreateVote(t *testing.T) {
	createRandomVote(t)
}
