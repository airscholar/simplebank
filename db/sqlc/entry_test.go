package db

import (
	"context"
	"testing"

	"github.com/airscholar/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, min, max int64) Entry {
	arg := CreateEntryParams{
		AccountID: util.RandomInt(min, max),
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	return entry
}
func TestCreateEntry(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	createRandomEntry(t, 1, 10)
}

func TestCreateEntryFailure(t *testing.T) {
	arg := CreateEntryParams{
		AccountID: util.RandomInt(1000, 10000),
		Amount:    util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.Error(t, err)
	require.Empty(t, entry)
}
