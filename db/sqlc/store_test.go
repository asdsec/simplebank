package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)
	tAccount1 := createRandomAccount(t)
	tAccount2 := createRandomAccount(t)
	fmt.Println(">> before:", tAccount1.Balance, tAccount2.Balance)

	// run a concurrent transfer transaction
	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			ctx := context.Background()
			result, err := store.TransferTx(ctx, TransferTxParams{
				FromAccountID: tAccount1.ID,
				ToAccountID:   tAccount2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	// check results
	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, tAccount1.ID, transfer.FromAccountID)
		require.Equal(t, tAccount2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)
		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// check entries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, tAccount1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)
		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, tAccount2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)
		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// check accounts
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, tAccount1.ID, fromAccount.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, tAccount2.ID, toAccount.ID)

		// check accounts' balance
		fmt.Println(">> tx:", fromAccount.Balance, toAccount.Balance)
		diff1 := tAccount1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - tAccount2.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0)

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	// check the final updated balances
	updatedAccount1, err := testQueries.GetAccount(context.Background(), tAccount1.ID)
	require.NoError(t, err)

	updatedAccount2, err := testQueries.GetAccount(context.Background(), tAccount2.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)
	require.Equal(t, tAccount1.Balance-int64(n)*amount, updatedAccount1.Balance)
	require.Equal(t, tAccount2.Balance+int64(n)*amount, updatedAccount2.Balance)
}

func TestTransferTxDeadlock(t *testing.T) {
	store := NewStore(testDB)
	tAccount1 := createRandomAccount(t)
	tAccount2 := createRandomAccount(t)
	fmt.Println(">> before:", tAccount1.Balance, tAccount2.Balance)

	// run a concurrent transfer transaction
	n := 10
	amount := int64(10)
	errs := make(chan error)

	for i := 0; i < n; i++ {

		fromAccountID := tAccount1.ID
		toAccountID := tAccount2.ID

		if i%2 == 1 {
			fromAccountID = tAccount2.ID
			toAccountID = tAccount1.ID
		}

		go func() {
			ctx := context.Background()
			_, err := store.TransferTx(ctx, TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountID:   toAccountID,
				Amount:        amount,
			})

			errs <- err
		}()
	}

	// check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

	}

	// check the final updated balances
	updatedAccount1, err := testQueries.GetAccount(context.Background(), tAccount1.ID)
	require.NoError(t, err)

	updatedAccount2, err := testQueries.GetAccount(context.Background(), tAccount2.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)
	require.Equal(t, tAccount1.Balance, updatedAccount1.Balance)
	require.Equal(t, tAccount2.Balance, updatedAccount2.Balance)
}
