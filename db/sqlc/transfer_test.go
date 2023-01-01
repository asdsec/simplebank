package db

import (
	"context"
	"testing"
	"time"

	"github.com/sametdmr/simplebank/utility"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, account1, account2 Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID: account2.ID,
		Amount: utility.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func Test(t *testing.T) {
	tAccount1 := createRandomAccount(t)
	tAccount2 := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: tAccount1.ID,
		ToAccountID: tAccount2.ID,
		Amount: utility.RandomMoney(),
	}

	result, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, arg.FromAccountID, result.FromAccountID)
	require.Equal(t, arg.ToAccountID, result.ToAccountID)
	require.Equal(t, arg.Amount, result.Amount)
	require.NotZero(t, result.ID)
	require.NotZero(t, result.CreatedAt)
}

func TestGetTransfer(t *testing.T) {
	tAccount1 := createRandomAccount(t)
	tAccount2 := createRandomAccount(t)
	tTransfer := createRandomTransfer(t, tAccount1, tAccount2)

	result, err := testQueries.GetTransfer(context.Background(), tTransfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, tTransfer.ID, result.ID)
	require.Equal(t, tTransfer.Amount, result.Amount)
	require.Equal(t, tTransfer.FromAccountID, result.FromAccountID)
	require.Equal(t, tTransfer.ToAccountID, result.ToAccountID)
	require.WithinDuration(t, tTransfer.CreatedAt, result.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	tAccount1 := createRandomAccount(t)
	tAccount2 := createRandomAccount(t)
	for i := 0; i < 5; i++ {
		createRandomTransfer(t, tAccount1, tAccount2)
		createRandomTransfer(t, tAccount2, tAccount1)
	}
	arg := ListTransfersParams{
		FromAccountID: tAccount1.ID,
		ToAccountID: tAccount1.ID,
		Limit: 5,
		Offset: 5,
	}

	result, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Len(t, result, 5)
	for _, transfer := range result {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == tAccount1.ID || transfer.ToAccountID == tAccount1.ID)
	}
}