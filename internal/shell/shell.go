package shell

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/caarlos0/log"
	"github.com/triarius/goreleaser/internal/gio"
	"github.com/triarius/goreleaser/internal/logext"
	"github.com/triarius/goreleaser/pkg/context"
)

// Run a shell command with given arguments and envs
func Run(ctx *context.Context, dir string, command, env []string, output bool) error {
	fields := log.Fields{
		"cmd": command,
		"env": env,
	}

	/* #nosec */
	cmd := exec.CommandContext(ctx, command[0], command[1:]...)
	cmd.Env = env

	var b bytes.Buffer
	w := gio.Safe(&b)

	cmd.Stderr = io.MultiWriter(logext.NewConditionalWriter(fields, logext.Error, output), w)
	cmd.Stdout = io.MultiWriter(logext.NewConditionalWriter(fields, logext.Info, output), w)

	if dir != "" {
		cmd.Dir = dir
	}

	log.WithFields(fields).Debug("running")
	if err := cmd.Run(); err != nil {
		log.WithFields(fields).WithError(err).Debug("failed")
		return fmt.Errorf("failed to run '%s': %w", strings.Join(command, " "), err)
	}

	return nil
}
