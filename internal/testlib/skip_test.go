package testlib

import (
	"testing"

	"github.com/triarius/goreleaser/internal/pipe"
)

func TestAssertSkipped(t *testing.T) {
	AssertSkipped(t, pipe.Skip("skip"))
}
