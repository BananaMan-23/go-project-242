package pathSize

import (
	"github.com/stretchr/testify/require"
	"path/filepath"
	"testing"
)

func TestGetSizeFile(t *testing.T) {
	path := filepath.Join("testdata", "file1.txt")

	got, err := GetSize(path)
	require.NoError(t, err)

	want := "3B\t" + path
	require.Equal(t, want, got)
}

func TestGetSizeDirFirstLevel(t *testing.T) {
	path := filepath.Join("testdata")

	got, err := GetSize(path)
	require.NoError(t, err)

	want := "12B\t" + path
	require.Equal(t, want, got)
}
