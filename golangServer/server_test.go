// main.go
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/RangerMauve/ipld-prolly-indexer/indexer"
	blocks "github.com/ipfs/go-block-format"
	carBlockstore "github.com/ipld/go-car/v2/blockstore"
	"github.com/stretchr/testify/require"
)

func copy(src string, dst string) error {
	// Read all content of src to data, may cause OOM for a large file.
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	// Write data to dst
	return ioutil.WriteFile(dst, data, 0644)
}

func Test_saveDB22(t *testing.T) {
	err := copy("./db002.car", "./db002_tmp.car")
	require.NoError(t, err)

	{
		rblockstore, err := carBlockstore.OpenReadOnly("./db002_tmp.car", carBlockstore.UseWholeCIDs(true))
		require.NoError(t, err)
		roots, err := rblockstore.Roots()
		require.NoError(t, err)

		rwblockstore, err := carBlockstore.OpenReadWrite("./db002_tmp.car", roots, carBlockstore.UseWholeCIDs(true))
		require.NoError(t, err)
		err = rwblockstore.Put(context.Background(), blocks.NewBlock([]byte{1, 2, 3, 4, 5}))
		require.NoError(t, err)

		err = rwblockstore.Finalize()
		require.NoError(t, err)
	}

	{
		rblockstore, err := carBlockstore.OpenReadOnly("./db002_tmp.car", carBlockstore.UseWholeCIDs(true))
		require.NoError(t, err)
		_, err = rblockstore.Roots()
		require.NoError(t, err)

	}
}

func Test_saveDB(t *testing.T) {
	ctx := context.Background()
	db, err := indexer.NewMemoryDatabase()
	require.NoError(t, err)
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
	db, ok := loadDBs["xasd"]
	require.True(t, ok)

	err = db.GetBlockstore().Put(context.Background(), blocks.NewBlock([]byte{1, 2, 3, 4, 5}))
	require.NoError(t, err)

	err = saveDB(ctx, filepath, dbs)
	require.NoError(t, err)
	{
		_, err = loadDb(ctx, filepath)
		require.NoError(t, err)

	}
}
