package db

import (
	"context"
	"testing"
	"time"

	"github.com/gitnyasha/go-hekani-backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString())
	require.NoError(t, err)
	arg := CreateUserParams{
		Name:           util.RandomString(),
		Email:          util.RandomEmail(),
		HashedPassword: hashedPassword,
		Bio:            util.RandomString(),
		Birth:          time.Time(time.Date(int(util.RandomInt(1900, 2020)), 1, 1, 0, 0, 0, 0, time.UTC)),
		Image:          util.RandomString(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Bio, user.Bio)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	return user
}

// create user
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

// get user
func TestGetUser(t *testing.T) {
	user := createRandomUser(t)

	user2, err := testQueries.GetUser(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user.Name, user2.Name)
	require.Equal(t, user.Email, user2.Email)
	require.Equal(t, user.HashedPassword, user2.HashedPassword)
	require.Equal(t, user.Bio, user2.Bio)
	require.Equal(t, user.Birth, user2.Birth)
	require.Equal(t, user.Image, user2.Image)
	require.Equal(t, user.CreatedAt, user2.CreatedAt)
	require.Equal(t, user.UpdatedAt, user2.UpdatedAt)
}

// delete user
func TestDeleteUser(t *testing.T) {
	user := createRandomUser(t)

	err := testQueries.DeleteUser(context.Background(), user.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user.ID)
	require.Error(t, err)
	require.Empty(t, user2)
}

// list users
func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}

// update user
func TestUpdateUser(t *testing.T) {
	user := createRandomUser(t)

	arg := UpdateUserParams{
		ID:    user.ID,
		Name:  user.Name,
		Bio:   user.Bio,
		Birth: user.Birth,
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, arg.Name, user2.Name)
	require.Equal(t, arg.Bio, user2.Bio)
	require.Equal(t, arg.Birth, user2.Birth)
	require.Equal(t, user.CreatedAt, user2.CreatedAt)
	require.NotZero(t, user2.UpdatedAt)
}
