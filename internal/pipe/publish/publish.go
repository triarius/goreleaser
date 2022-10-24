// Package publish contains the publishing pipe.
package publish

import (
	"fmt"

	"github.com/triarius/goreleaser/internal/middleware/errhandler"
	"github.com/triarius/goreleaser/internal/middleware/logging"
	"github.com/triarius/goreleaser/internal/middleware/skip"
	"github.com/triarius/goreleaser/internal/pipe/artifactory"
	"github.com/triarius/goreleaser/internal/pipe/aur"
	"github.com/triarius/goreleaser/internal/pipe/blob"
	"github.com/triarius/goreleaser/internal/pipe/brew"
	"github.com/triarius/goreleaser/internal/pipe/custompublishers"
	"github.com/triarius/goreleaser/internal/pipe/docker"
	"github.com/triarius/goreleaser/internal/pipe/krew"
	"github.com/triarius/goreleaser/internal/pipe/milestone"
	"github.com/triarius/goreleaser/internal/pipe/release"
	"github.com/triarius/goreleaser/internal/pipe/scoop"
	"github.com/triarius/goreleaser/internal/pipe/sign"
	"github.com/triarius/goreleaser/internal/pipe/snapcraft"
	"github.com/triarius/goreleaser/internal/pipe/upload"
	"github.com/triarius/goreleaser/pkg/context"
)

// Publisher should be implemented by pipes that want to publish artifacts.
type Publisher interface {
	fmt.Stringer

	// Default sets the configuration defaults
	Publish(ctx *context.Context) error
}

// nolint: gochecknoglobals
var publishers = []Publisher{
	blob.Pipe{},
	upload.Pipe{},
	artifactory.Pipe{},
	custompublishers.Pipe{},
	docker.Pipe{},
	docker.ManifestPipe{},
	sign.DockerPipe{},
	snapcraft.Pipe{},
	// This should be one of the last steps
	release.Pipe{},
	// brew et al use the release URL, so, they should be last
	brew.Pipe{},
	aur.Pipe{},
	krew.Pipe{},
	scoop.Pipe{},
	milestone.Pipe{},
}

// Pipe that publishes artifacts.
type Pipe struct{}

func (Pipe) String() string                 { return "publishing" }
func (Pipe) Skip(ctx *context.Context) bool { return ctx.SkipPublish }

func (Pipe) Run(ctx *context.Context) error {
	for _, publisher := range publishers {
		if err := skip.Maybe(
			publisher,
			logging.PadLog(
				publisher.String(),
				errhandler.Handle(publisher.Publish),
			),
		)(ctx); err != nil {
			return fmt.Errorf("%s: failed to publish artifacts: %w", publisher.String(), err)
		}
	}
	return nil
}
