package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomUserQuestionRelation(t *testing.T) UserQuestionRelation {
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)

	arg := CreateUserQuestionRelationParams{
		FollowerID: int32(user1.ID),
		FollowedID: int32(user2.ID),
	}

	user_question_relationship, err := testQueries.CreateUserQuestionRelation(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user_question_relationship)
	require.Equal(t, arg.FollowedID, user_question_relationship.FollowedID)
	require.Equal(t, arg.FollowerID, user_question_relationship.FollowerID)
	require.NotZero(t, user_question_relationship.ID)
	require.NotZero(t, user_question_relationship.CreatedAt)
	return user_question_relationship
}

func TestCreateUserquestionRelationship(t *testing.T) {
	createRandomUserQuestionRelation(t)
}
