package db

import (
	"context"
	"database/sql"
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
	for i := 0; i < 2; i++ {
		createRandomTransfer(t)
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

func TestUpdateTransferToId(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}

	account1, err := testQueries.GetAccount(context.Background(), 1)
	require.NoError(t, err)
	account2, err := testQueries.GetAccount(context.Background(), 2)
	require.NoError(t, err)

	arg := UpdateTransferByToIdParams{
		ToAccountID:   account2.ID,
		FromAccountID: account1.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.UpdateTransferByToId(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, arg.Amount, transfer.Amount)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, account1.ID, transfer.FromAccountID)
}

func TestDeleteTransfer(t *testing.T) {
	createRandomAccount(t)

	transfer := createRandomTransfer(t)

	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)

	require.NoError(t, err)

	trf, err := testQueries.GetTransfer(context.Background(), transfer.ID)

	require.Error(t, err)
	require.Empty(t, trf)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

func TestDeleteTransferFailure(t *testing.T) {
	err := testQueries.DeleteTransfer(context.Background(), 0)

	require.Empty(t, err)
}

func TestGetTransfer(t *testing.T) {
	createRandomTransfer(t)

	transfer, err := testQueries.GetTransfer(context.Background(), util.RandomInt(1, 2))

	require.NoError(t, err)
	require.NotEmpty(t, transfer)
}

func TestGetTransfers(t *testing.T) {
	for i := 0; i < 6; i++ {
		createRandomTransfer(t)
	}

	arg := GetTransfersParams{
		Offset: 3,
		Limit:  3,
	}
	transfers, err := testQueries.GetTransfers(context.Background(), arg)

	require.NoError(t, err)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}

func TestGetTransfersByFromId(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}

	arg := GetTransfersByFromIdParams{
		FromAccountID: 1,
		Offset:        3,
		Limit:         3,
	}
	transfers, err := testQueries.GetTransfersByFromId(context.Background(), arg)

	require.NoError(t, err)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}

func TestGetTransfersByToId(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}

	arg := GetTransfersByToIdParams{
		ToAccountID: 1,
		Offset:      3,
		Limit:       3,
	}
	transfers, err := testQueries.GetTransfersByToId(context.Background(), arg)

	require.NoError(t, err)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
