package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/gitnyasha/go-hekani-backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomQuestionCategory(t *testing.T) QuestionCategory {
	title := util.RandomString()
	question_category, err := testQueries.CreateQuestionCategory(context.Background(), title)
	require.NoError(t, err)
	require.NotEmpty(t, question_category)

	require.Equal(t, title, question_category.Name)

	require.NotZero(t, question_category.ID)
	require.NotZero(t, question_category.CreatedAt)
	require.NotZero(t, question_category.UpdatedAt)

	return question_category
}

func TestCreatequestionCategory(t *testing.T) {
	createRandomQuestionCategory(t)
}

func TestGetquestionCategory(t *testing.T) {
	question_category1 := createRandomQuestionCategory(t)
	question_category2, err := testQueries.GetQuestionCategory(context.Background(), question_category1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, question_category2)

	require.Equal(t, question_category1.ID, question_category2.ID)
	require.Equal(t, question_category1.Name, question_category2.Name)
	require.Equal(t, question_category1.CreatedAt, question_category2.CreatedAt)
	require.Equal(t, question_category1.UpdatedAt, question_category2.UpdatedAt)
	require.WithinDuration(t, question_category1.CreatedAt, question_category2.CreatedAt, time.Second)
}

// update question_category
func TestUpdatequestionCategory(t *testing.T) {
	question_category1 := createRandomQuestionCategory(t)

	arg := UpdateQuestionCategoryParams{
		ID:   question_category1.ID,
		Name: util.RandomString(),
	}

	question_category2, err := testQueries.UpdateQuestionCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, question_category2)

	require.Equal(t, arg.ID, question_category2.ID)
	require.Equal(t, arg.Name, question_category2.Name)
	require.Equal(t, question_category1.CreatedAt, question_category2.CreatedAt)
	require.NotEqual(t, question_category1.UpdatedAt, question_category2.UpdatedAt)
	require.WithinDuration(t, question_category1.CreatedAt, question_category2.UpdatedAt, time.Second)
}

// delete question_category
func TestDeletequestionCategory(t *testing.T) {
	question_category1 := createRandomQuestionCategory(t)
	err := testQueries.DeleteQuestionCategory(context.Background(), question_category1.ID)
	require.NoError(t, err)

	question_category2, err := testQueries.GetQuestionCategory(context.Background(), question_category1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, question_category2)
}

// list question_categories
func TestListQuestionCategories(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomQuestionCategory(t)
	}

	arg := ListQuestionCategoriesParams{
		Limit:  5,
		Offset: 5,
	}

	question_categories, err := testQueries.ListQuestionCategories(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, question_categories, 5)

	for _, question_category := range question_categories {
		require.NotEmpty(t, question_category)
	}

}
