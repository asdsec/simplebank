package db

import (
	"context"
	"testing"
	"time"

	"github.com/sametdmr/simplebank/utility"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := utility.HashPassword(utility.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       utility.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       utility.RandomOwner(),
		Email:          utility.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	hashedPassword, err := utility.HashPassword(utility.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       utility.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       utility.RandomOwner(),
		Email:          utility.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)
}

func TestGetUser(t *testing.T) {
	tUser := createRandomUser(t)

	result, err := testQueries.GetUser(context.Background(), tUser.Username)

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, tUser.Username, result.Username)
	require.Equal(t, tUser.HashedPassword, result.HashedPassword)
	require.Equal(t, tUser.FullName, result.FullName)
	require.Equal(t, tUser.Email, result.Email)
	require.WithinDuration(t, tUser.PasswordChangedAt, result.PasswordChangedAt, time.Second)
	require.WithinDuration(t, tUser.CreatedAt, result.CreatedAt, time.Second)
}
