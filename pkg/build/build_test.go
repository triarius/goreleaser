package build

import (
	"testing"

	"github.com/triarius/goreleaser/pkg/config"
	"github.com/triarius/goreleaser/pkg/context"
	"github.com/stretchr/testify/require"
)

type dummy struct{}

func (*dummy) WithDefaults(build config.Build) (config.Build, error) {
	return build, nil
}

func (*dummy) Build(ctx *context.Context, build config.Build, options Options) error {
	return nil
}

func TestRegisterAndGet(t *testing.T) {
	builder := &dummy{}
	Register("dummy", builder)
	require.Equal(t, builder, For("dummy"))
}
