package api

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_createNewDbFile(t *testing.T) {
	tmpDir, err := os.MkdirTemp(os.TempDir(), "*")
	require.NoError(t, err)
	fmt.Println(tmpDir)
	_, err = createNewDbFile(context.Background(), tmpDir, "db001")
	require.NoError(t, err)
}
