package db

import (
	"context"
	"testing"

	"github.com/airscholar/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	arg := CreateTransferParams{
		FromAccountID: util.RandomInt(1, 5),
		ToAccountID:   util.RandomInt(6, 10),
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	return transfer
}
func TestCreateTransfer(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}
	createRandomTransfer(t)
}

func TestCreateTransferFailure(t *testing.T) {
	arg := CreateTransferParams{
		FromAccountID: 0,
		ToAccountID:   0,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.Error(t, err)
	require.Empty(t, transfer)
}

func TestUpdateTransferFromId(t *testing.T) {
	for i := 0; i < 2; i++ {
		createRandomAccount(t)
	}

	account1, err := testQueries.GetAccount(context.Background(), 1)
	require.NoError(t, err)
	account2, err := testQueries.GetAccount(context.Background(), 2)
	require.NoError(t, err)

	arg := UpdateTransferByFromIdParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.UpdateTransferByFromId(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, arg.Amount, transfer.Amount)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, account1.ID, transfer.FromAccountID)
}