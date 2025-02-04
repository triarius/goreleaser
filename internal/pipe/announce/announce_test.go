package announce

import (
	"testing"

	"github.com/triarius/goreleaser/pkg/config"
	"github.com/triarius/goreleaser/pkg/context"
	"github.com/stretchr/testify/require"
)

func TestDescription(t *testing.T) {
	require.NotEmpty(t, Pipe{}.String())
}

func TestAnnounce(t *testing.T) {
	ctx := context.New(config.Project{
		Announce: config.Announce{
			Twitter: config.Twitter{
				Enabled: true,
			},
		},
	})
	require.Error(t, Pipe{}.Run(ctx))
}

func TestAnnounceAllDisabled(t *testing.T) {
	ctx := context.New(config.Project{})
	require.NoError(t, Pipe{}.Run(ctx))
}

func TestSkip(t *testing.T) {
	t.Run("skip", func(t *testing.T) {
		ctx := context.New(config.Project{})
		ctx.SkipAnnounce = true
		require.True(t, Pipe{}.Skip(ctx))
	})

	t.Run("skip on patches", func(t *testing.T) {
		ctx := context.New(config.Project{
			Announce: config.Announce{
				Skip: "{{gt .Patch 0}}",
			},
		})
		ctx.Semver.Patch = 1
		require.True(t, Pipe{}.Skip(ctx))
	})

	t.Run("skip on invalid template", func(t *testing.T) {
		ctx := context.New(config.Project{
			Announce: config.Announce{
				Skip: "{{if eq .Patch 123}",
			},
		})
		ctx.Semver.Patch = 1
		require.True(t, Pipe{}.Skip(ctx))
	})

	t.Run("dont skip", func(t *testing.T) {
		require.False(t, Pipe{}.Skip(context.New(config.Project{})))
	})

	t.Run("dont skip based on template", func(t *testing.T) {
		ctx := context.New(config.Project{
			Announce: config.Announce{
				Skip: "{{gt .Patch 0}}",
			},
		})
		ctx.Semver.Patch = 0
		require.False(t, Pipe{}.Skip(ctx))
	})
}
