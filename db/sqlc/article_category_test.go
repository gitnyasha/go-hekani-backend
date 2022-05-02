package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateArticleCategory(t *testing.T) {
	arg := CreateArticleCategoryParams{
		Name: "Sport",
	}

	article_category, err := testQueries.CreateArticleCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, article_category)

	require.Equal(t, arg.Name, article_category.Name)

	require.NotZero(t, article_category.ID)
	require.NotZero(t, article_category.CreatedAt)
	require.NotZero(t, article_category.UpdatedAt)
}
