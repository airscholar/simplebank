package db

import (
	"context"
	"database/sql"
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

func TestUpdateEntry(t *testing.T) {
	for i := 0; i < 3; i++ {
		createRandomAccount(t)
	}

	for i := 0; i < 3; i++ {
		createRandomEntry(t, 1, 3)
	}

	entry1, err := testQueries.GetAccount(context.Background(), util.RandomInt(1, 3))
	require.NoError(t, err)

	arg := UpdateEntryParams{
		ID:     entry1.ID,
		Amount: util.RandomMoney(),
	}

	updatedEntry, err := testQueries.UpdateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry1)
	require.NotEmpty(t, updatedEntry)
	require.EqualValues(t, arg.Amount, updatedEntry.Amount)
}

func TestUpdateEntryFailure(t *testing.T) {
	arg := UpdateEntryParams{
		ID:     util.RandomInt(0, 0),
		Amount: util.RandomMoney(),
	}

	updatedEntry, err := testQueries.UpdateEntry(context.Background(), arg)

	require.Error(t, err)
	require.Empty(t, updatedEntry)
}

func TestDeleteEntry(t *testing.T) {
	createRandomAccount(t)

	entry := createRandomEntry(t, 1, 1)

	err := testQueries.DeleteEntry(context.Background(), entry.ID)

	require.NoError(t, err)

	entry, err = testQueries.GetEntry(context.Background(), entry.ID)

	require.Empty(t, entry)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

func TestDeleteEntryFailure(t *testing.T) {
	err := testQueries.DeleteEntry(context.Background(), -10)

	require.Empty(t, err)
}
