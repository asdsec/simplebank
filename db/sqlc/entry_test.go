package db

import (
	"context"
	"testing"
	"time"

	"github.com/sametdmr/simplebank/utility"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Account) Entry {	
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount: utility.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	tAccount := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: tAccount.ID,
		Amount: utility.RandomMoney(),
	}

	result, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, arg.AccountID, result.AccountID)
	require.Equal(t, arg.Amount, result.Amount)
	require.NotZero(t, result.ID)
	require.NotZero(t, result.CreatedAt)
}

func TestGetEntry(t *testing.T) {
	tAccount := createRandomAccount(t)
	tEntry := createRandomEntry(t, tAccount)

	result, err := testQueries.GetEntry(context.Background(), tEntry.ID)

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, tEntry.ID, result.ID)
	require.Equal(t, tEntry.AccountID, result.AccountID)
	require.Equal(t, tEntry.Amount, result.Amount)
	require.WithinDuration(t, tEntry.CreatedAt, result.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	tAccount := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, tAccount)
	}
	arg := ListEntriesParams{
		AccountID: tAccount.ID,
		Limit: 5,
		Offset: 5,
	}

	result, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Len(t, result, 5)
	for _, entry := range result {
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
	}
}