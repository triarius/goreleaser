package errhandler

import (
	"github.com/caarlos0/log"
	"github.com/triarius/goreleaser/internal/middleware"
	"github.com/triarius/goreleaser/internal/pipe"
	"github.com/triarius/goreleaser/pkg/context"
)

// Handle handles an action error, ignoring and logging pipe skipped
// errors.
func Handle(action middleware.Action) middleware.Action {
	return func(ctx *context.Context) error {
		err := action(ctx)
		if err == nil {
			return nil
		}
		if pipe.IsSkip(err) {
			log.WithField("reason", err.Error()).Warn("pipe skipped")
			return nil
		}
		return err
	}
}
