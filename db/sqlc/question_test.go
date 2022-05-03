package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomQuestion(t *testing.T) Question {
	user := createRandomUser(t)
	question_category := createRandomQuestionCategory(t)

	arg := CreateQuestionParams{
		Title:              "Test Question",
		UserID:             int32(user.ID),
		QuestionCategoryID: int32(question_category.ID),
	}

	question, err := testQueries.CreateQuestion(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, question)

	require.Equal(t, arg.Title, question.Title)

	require.NotZero(t, question.ID)
	require.NotZero(t, question.CreatedAt)
	return question
}

// create question
func TestCreateQuestion(t *testing.T) {
	createRandomQuestion(t)
}

// get question
func TestGetQuestion(t *testing.T) {
	question := createRandomQuestion(t)

	question, err := testQueries.GetQuestion(context.Background(), question.ID)
	require.NoError(t, err)
	require.NotEmpty(t, question)

	require.Equal(t, question.ID, question.ID)
	require.Equal(t, question.Title, question.Title)
}
