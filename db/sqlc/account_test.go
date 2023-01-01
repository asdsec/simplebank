package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/sametdmr/simplebank/utility"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner: utility.RandomOwner(),
		Balance: utility.RandomMoney(),
		Currency: utility.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner: utility.RandomOwner(),
		Balance: utility.RandomMoney(),
		Currency: utility.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestGetAccount(t *testing.T) {
	expected := createRandomAccount(t)
	
	result, err := testQueries.GetAccount(context.Background(), expected.ID)
	
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, expected.ID, result.ID)
	require.Equal(t, expected.Owner, result.Owner)
	require.Equal(t, expected.Balance, result.Balance)
	require.Equal(t, expected.Currency, result.Currency)
	require.WithinDuration(t, expected.CreatedAt, result.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	tAccount := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID: tAccount.ID,
		Balance: utility.RandomMoney(),
	}

	result, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, tAccount.ID, result.ID)
	require.Equal(t, tAccount.Owner, result.Owner)
	require.Equal(t, arg.Balance, result.Balance)
	require.Equal(t, tAccount.Currency, result.Currency)
	require.WithinDuration(t, tAccount.CreatedAt, result.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	tAccount := createRandomAccount(t)
	
	err := testQueries.DeleteAccount(context.Background(), tAccount.ID)

	require.NoError(t, err)
	account, err := testQueries.GetAccount(context.Background(), tAccount.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}
	arg := ListAccountsParams{
		Limit: 5,
		Offset: 5,
	}

	result, err := testQueries.ListAccounts(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Len(t, result, 5)
	for _, account := range result {
		require.NotEmpty(t, account)
	}	
}