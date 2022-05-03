package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/gitnyasha/go-hekani-backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomArticle(t *testing.T) Article {
	user := createRandomUser(t)
	category := createRandomCategory(t)

	arg := CreateArticleParams{
		UserID:            int32(user.ID),
		ArticleCategoryID: int32(category.ID),
		Title:             util.RandomString(),
		Link:              util.RandomString(),
		Image:             util.RandomString(),
	}

	article, err := testQueries.CreateArticle(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, article)

	require.Equal(t, arg.UserID, article.UserID)
	require.Equal(t, arg.ArticleCategoryID, article.ArticleCategoryID)
	require.Equal(t, arg.Title, article.Title)

	require.NotZero(t, article.ID)
	require.NotZero(t, article.CreatedAt)

	return article
}

// create Article
func TestCreateArticle(t *testing.T) {
	createRandomArticle(t)
}

func TestGetArticle(t *testing.T) {
	article1 := createRandomArticle(t)
	article2, err := testQueries.GetArticle(context.Background(), article1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, article2)

	require.Equal(t, article1.ID, article2.ID)
	require.Equal(t, article1.Title, article2.Title)
	require.Equal(t, article1.CreatedAt, article2.CreatedAt)
	require.Equal(t, article1.UpdatedAt, article2.UpdatedAt)
	require.WithinDuration(t, article1.CreatedAt, article2.CreatedAt, time.Second)
}

// delete article
func TestDeleteArticle(t *testing.T) {
	article1 := createRandomArticle(t)
	err := testQueries.DeleteArticle(context.Background(), article1.ID)
	require.NoError(t, err)

	article2, err := testQueries.GetArticle(context.Background(), article1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, article2)
}

// list article_
func TestListArticle(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomArticle(t)
	}

	arg := ListArticlesParams{
		Limit:  5,
		Offset: 5,
	}

	article, err := testQueries.ListArticles(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, article, 5)

	for _, article := range article {
		require.NotEmpty(t, article)
	}

}
