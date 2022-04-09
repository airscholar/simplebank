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
