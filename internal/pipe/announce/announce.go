// Package announce contains the announcing pipe.
package announce

import (
	"fmt"

	"github.com/caarlos0/log"
	"github.com/triarius/goreleaser/internal/middleware/errhandler"
	"github.com/triarius/goreleaser/internal/middleware/logging"
	"github.com/triarius/goreleaser/internal/middleware/skip"
	"github.com/triarius/goreleaser/internal/pipe/discord"
	"github.com/triarius/goreleaser/internal/pipe/linkedin"
	"github.com/triarius/goreleaser/internal/pipe/mattermost"
	"github.com/triarius/goreleaser/internal/pipe/reddit"
	"github.com/triarius/goreleaser/internal/pipe/slack"
	"github.com/triarius/goreleaser/internal/pipe/smtp"
	"github.com/triarius/goreleaser/internal/pipe/teams"
	"github.com/triarius/goreleaser/internal/pipe/telegram"
	"github.com/triarius/goreleaser/internal/pipe/twitter"
	"github.com/triarius/goreleaser/internal/pipe/webhook"
	"github.com/triarius/goreleaser/internal/tmpl"
	"github.com/triarius/goreleaser/pkg/context"
)

// Announcer should be implemented by pipes that want to announce releases.
type Announcer interface {
	fmt.Stringer
	Announce(ctx *context.Context) error
}

// nolint: gochecknoglobals
var announcers = []Announcer{
	// XXX: keep asc sorting
	discord.Pipe{},
	linkedin.Pipe{},
	mattermost.Pipe{},
	reddit.Pipe{},
	slack.Pipe{},
	smtp.Pipe{},
	teams.Pipe{},
	telegram.Pipe{},
	twitter.Pipe{},
	webhook.Pipe{},
}

// Pipe that announces releases.
type Pipe struct{}

func (Pipe) String() string { return "announcing" }

func (Pipe) Skip(ctx *context.Context) bool {
	if ctx.SkipAnnounce {
		return true
	}
	if ctx.Config.Announce.Skip == "" {
		return false
	}
	skip, err := tmpl.New(ctx).Apply(ctx.Config.Announce.Skip)
	if err != nil {
		log.Error("invalid announce.skip template, will skip the announcing step")
		return true
	}
	log.Debugf("announce.skip evaluated from %q to %q", ctx.Config.Announce.Skip, skip)
	return skip == "true"
}

// Run the pipe.
func (Pipe) Run(ctx *context.Context) error {
	for _, announcer := range announcers {
		if err := skip.Maybe(
			announcer,
			logging.PadLog(
				announcer.String(),
				errhandler.Handle(announcer.Announce),
			),
		)(ctx); err != nil {
			return fmt.Errorf("%s: failed to announce release: %w", announcer.String(), err)
		}
	}
	return nil
}
