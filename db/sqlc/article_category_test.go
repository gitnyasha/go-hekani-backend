package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/gitnyasha/go-hekani-backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) ArticleCategory {
	title := util.RandomString()
	article_category, err := testQueries.CreateArticleCategory(context.Background(), title)
	require.NoError(t, err)
	require.NotEmpty(t, article_category)

	require.Equal(t, title, article_category.Name)

	require.NotZero(t, article_category.ID)
	require.NotZero(t, article_category.CreatedAt)
	require.NotZero(t, article_category.UpdatedAt)

	return article_category
}

func TestCreateArticleCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetArticleCategory(t *testing.T) {
	article_category1 := createRandomCategory(t)
	article_category2, err := testQueries.GetArticleCategory(context.Background(), article_category1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, article_category2)

	require.Equal(t, article_category1.ID, article_category2.ID)
	require.Equal(t, article_category1.Name, article_category2.Name)
	require.Equal(t, article_category1.CreatedAt, article_category2.CreatedAt)
	require.Equal(t, article_category1.UpdatedAt, article_category2.UpdatedAt)
	require.WithinDuration(t, article_category1.CreatedAt, article_category2.CreatedAt, time.Second)
}

// update article_category
func TestUpdateArticleCategory(t *testing.T) {
	article_category1 := createRandomCategory(t)

	arg := UpdateArticleCategoryParams{
		ID:   article_category1.ID,
		Name: util.RandomString(),
	}

	article_category2, err := testQueries.UpdateArticleCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, article_category2)

	require.Equal(t, arg.ID, article_category2.ID)
	require.Equal(t, arg.Name, article_category2.Name)
	require.Equal(t, article_category1.CreatedAt, article_category2.CreatedAt)
	require.NotEqual(t, article_category1.UpdatedAt, article_category2.UpdatedAt)
	require.WithinDuration(t, article_category1.CreatedAt, article_category2.UpdatedAt, time.Second)
}

// delete article_category
func TestDeleteArticleCategory(t *testing.T) {
	article_category1 := createRandomCategory(t)
	err := testQueries.DeleteArticleCategory(context.Background(), article_category1.ID)
	require.NoError(t, err)

	article_category2, err := testQueries.GetArticleCategory(context.Background(), article_category1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, article_category2)
}

// list article_categories
func TestListArticleCategories(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCategory(t)
	}

	arg := ListArticleCategoriesParams{
		Limit:  5,
		Offset: 5,
	}

	article_categories, err := testQueries.ListArticleCategories(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, article_categories, 5)

	for _, article_category := range article_categories {
		require.NotEmpty(t, article_category)
	}

}
