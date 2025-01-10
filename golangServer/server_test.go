// main.go
package main

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/RangerMauve/ipld-prolly-indexer/indexer"
	"github.com/stretchr/testify/require"
)

func Test_saveDB(t *testing.T) {
	ctx := context.Background()
	db, err := indexer.NewMemoryDatabase()

	dbs := map[string]*indexer.Database{
		"xasd": db,
	}
	require.NoError(t, err)

	filepath, err := os.MkdirTemp(os.TempDir(), "*")
	fmt.Println(filepath)
	require.NoError(t, err)
	saveDB(ctx, filepath, dbs)
	loadDBs, err := loadDb(ctx, filepath)
	require.NoError(t, err)
	require.Len(t, loadDBs, 1)
	_, ok := loadDBs["xasd"]
	require.True(t, ok)
	err = saveDB(ctx, filepath, dbs)
	require.NoError(t, err)
	{
		_, err = loadDb(ctx, filepath)
		require.NoError(t, err)

	}
}
