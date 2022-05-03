package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/gitnyasha/go-hekani-backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomAnswer(t *testing.T) Answer {
	user := createRandomUser(t)
	question := createRandomQuestion(t)

	arg := CreateAnswerParams{
		UserID:     int32(user.ID),
		QuestionID: int32(question.ID),
		Title:      util.RandomString(),
	}

	answer, err := testQueries.CreateAnswer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, answer)

	require.Equal(t, arg.UserID, answer.UserID)
	require.Equal(t, arg.QuestionID, answer.QuestionID)
	require.Equal(t, arg.Title, answer.Title)

	require.NotZero(t, answer.ID)
	require.NotZero(t, answer.CreatedAt)

	return answer
}

// create answer
func TestCreateAnswer(t *testing.T) {
	createRandomAnswer(t)
}

func TestGetAnswer(t *testing.T) {
	answer1 := createRandomAnswer(t)
	answer2, err := testQueries.GetAnswer(context.Background(), answer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, answer2)

	require.Equal(t, answer1.ID, answer2.ID)
	require.Equal(t, answer1.Title, answer2.Title)
	require.Equal(t, answer1.CreatedAt, answer2.CreatedAt)
	require.Equal(t, answer1.UpdatedAt, answer2.UpdatedAt)
	require.WithinDuration(t, answer1.CreatedAt, answer2.CreatedAt, time.Second)
}

// delete answer
func TestDeleteAnswer(t *testing.T) {
	answer1 := createRandomAnswer(t)
	err := testQueries.DeleteAnswer(context.Background(), answer1.ID)
	require.NoError(t, err)

	answer2, err := testQueries.GetAnswer(context.Background(), answer1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, answer2)
}

// list answer_
func TestListAnswer(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAnswer(t)
	}

	arg := ListAnswersParams{
		Limit:  5,
		Offset: 5,
	}

	answer, err := testQueries.ListAnswers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, answer, 5)

	for _, answer := range answer {
		require.NotEmpty(t, answer)
	}

}
