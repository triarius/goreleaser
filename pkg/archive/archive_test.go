package archive

import (
	"io"
	"os"
	"testing"

	"github.com/triarius/goreleaser/pkg/config"
	"github.com/stretchr/testify/require"
)

func TestArchive(t *testing.T) {
	folder := t.TempDir()
	empty, err := os.Create(folder + "/empty.txt")
	require.NoError(t, err)
	require.NoError(t, empty.Close())
	require.NoError(t, os.Mkdir(folder+"/folder-inside", 0o755))

	for _, format := range []string{"tar.gz", "zip", "gz", "tar.xz", "tar"} {
		format := format
		t.Run(format, func(t *testing.T) {
			archive, err := New(io.Discard, format)
			require.NoError(t, err)
			t.Cleanup(func() {
				require.NoError(t, archive.Close())
			})
			require.NoError(t, archive.Add(config.File{
				Source:      empty.Name(),
				Destination: "empty.txt",
			}))
			require.Error(t, archive.Add(config.File{
				Source:      empty.Name() + "_nope",
				Destination: "dont.txt",
			}))
		})
	}

	t.Run("7z", func(t *testing.T) {
		_, err := New(io.Discard, "7z")
		require.EqualError(t, err, "invalid archive format: 7z")
	})
}
