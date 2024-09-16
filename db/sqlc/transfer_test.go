package db

import (
	"context"
	"github.com/bolusarz/simplebank/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createTransferFunc(t *testing.T) Transfer {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	args := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, args.FromAccountID, transfer.FromAccountID)
	require.Equal(t, args.ToAccountID, transfer.ToAccountID)
	require.Equal(t, args.Amount, transfer.Amount)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createTransferFunc(t)
}

func TestGetTransfer(t *testing.T) {
	transfer := createTransferFunc(t)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer.ID, transfer2.ID)
	require.Equal(t, transfer.Amount, transfer2.Amount)
	require.Equal(t, transfer.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer.ToAccountID, transfer2.ToAccountID)
	require.NotZero(t, transfer2.ID)
	require.WithinDuration(t, transfer.CreatedAt.Time, transfer2.CreatedAt.Time, time.Second)
}

func TestListTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		args := CreateTransferParams{
			FromAccountID: account1.ID,
			ToAccountID:   account2.ID,
			Amount:        util.RandomMoney(),
		}

		_, _ = testQueries.CreateTransfer(context.Background(), args)
	}

	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Limit:         10,
		Offset:        0,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 10)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, transfer.FromAccountID, account1.ID)
		require.Equal(t, transfer.ToAccountID, account2.ID)
	}
}
