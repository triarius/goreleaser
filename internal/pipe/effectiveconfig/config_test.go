package effectiveconfig

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/triarius/goreleaser/internal/testlib"
	"github.com/triarius/goreleaser/pkg/config"
	"github.com/triarius/goreleaser/pkg/context"
	"github.com/stretchr/testify/require"
)

func TestPipeDescription(t *testing.T) {
	require.NotEmpty(t, Pipe{}.String())
}

func TestRun(t *testing.T) {
	folder := testlib.Mktmp(t)
	dist := filepath.Join(folder, "dist")
	require.NoError(t, os.Mkdir(dist, 0o755))
	ctx := context.New(
		config.Project{
			Dist: dist,
		},
	)
	require.NoError(t, Pipe{}.Run(ctx))
	bts, err := os.ReadFile(filepath.Join(dist, "config.yaml"))
	require.NoError(t, err)
	require.NotEmpty(t, string(bts))
}
