// Package defaults make the list of Defaulter implementations available
// so projects extending GoReleaser are able to use it, namely, GoDownloader.
package defaults

import (
	"fmt"

	"github.com/triarius/goreleaser/internal/pipe/archive"
	"github.com/triarius/goreleaser/internal/pipe/artifactory"
	"github.com/triarius/goreleaser/internal/pipe/aur"
	"github.com/triarius/goreleaser/internal/pipe/blob"
	"github.com/triarius/goreleaser/internal/pipe/brew"
	"github.com/triarius/goreleaser/internal/pipe/build"
	"github.com/triarius/goreleaser/internal/pipe/checksums"
	"github.com/triarius/goreleaser/internal/pipe/discord"
	"github.com/triarius/goreleaser/internal/pipe/docker"
	"github.com/triarius/goreleaser/internal/pipe/gomod"
	"github.com/triarius/goreleaser/internal/pipe/krew"
	"github.com/triarius/goreleaser/internal/pipe/linkedin"
	"github.com/triarius/goreleaser/internal/pipe/mattermost"
	"github.com/triarius/goreleaser/internal/pipe/milestone"
	"github.com/triarius/goreleaser/internal/pipe/nfpm"
	"github.com/triarius/goreleaser/internal/pipe/project"
	"github.com/triarius/goreleaser/internal/pipe/reddit"
	"github.com/triarius/goreleaser/internal/pipe/release"
	"github.com/triarius/goreleaser/internal/pipe/sbom"
	"github.com/triarius/goreleaser/internal/pipe/scoop"
	"github.com/triarius/goreleaser/internal/pipe/sign"
	"github.com/triarius/goreleaser/internal/pipe/slack"
	"github.com/triarius/goreleaser/internal/pipe/smtp"
	"github.com/triarius/goreleaser/internal/pipe/snapcraft"
	"github.com/triarius/goreleaser/internal/pipe/snapshot"
	"github.com/triarius/goreleaser/internal/pipe/sourcearchive"
	"github.com/triarius/goreleaser/internal/pipe/teams"
	"github.com/triarius/goreleaser/internal/pipe/telegram"
	"github.com/triarius/goreleaser/internal/pipe/twitter"
	"github.com/triarius/goreleaser/internal/pipe/universalbinary"
	"github.com/triarius/goreleaser/internal/pipe/webhook"
	"github.com/triarius/goreleaser/pkg/context"
)

// Defaulter can be implemented by a Piper to set default values for its
// configuration.
type Defaulter interface {
	fmt.Stringer

	// Default sets the configuration defaults
	Default(ctx *context.Context) error
}

// Defaulters is the list of defaulters.
// nolint: gochecknoglobals
var Defaulters = []Defaulter{
	snapshot.Pipe{},
	release.Pipe{},
	project.Pipe{},
	gomod.Pipe{},
	build.Pipe{},
	universalbinary.Pipe{},
	sourcearchive.Pipe{},
	archive.Pipe{},
	nfpm.Pipe{},
	snapcraft.Pipe{},
	checksums.Pipe{},
	sign.Pipe{},
	sign.DockerPipe{},
	sbom.Pipe{},
	docker.Pipe{},
	docker.ManifestPipe{},
	artifactory.Pipe{},
	blob.Pipe{},
	aur.Pipe{},
	brew.Pipe{},
	krew.Pipe{},
	scoop.Pipe{},
	discord.Pipe{},
	reddit.Pipe{},
	slack.Pipe{},
	teams.Pipe{},
	twitter.Pipe{},
	smtp.Pipe{},
	mattermost.Pipe{},
	milestone.Pipe{},
	linkedin.Pipe{},
	telegram.Pipe{},
	webhook.Pipe{},
}
