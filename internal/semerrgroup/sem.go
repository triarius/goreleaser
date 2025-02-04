// Package semerrgroup wraps a error group with a semaphore with configurable
// size, so you can control the number of tasks being executed simultaneously.
package semerrgroup

import (
	"sync"

	"github.com/triarius/goreleaser/internal/pipe"
	"golang.org/x/sync/errgroup"
)

// Group is the Semphore ErrorGroup itself.
type Group interface {
	Go(func() error)
	Wait() error
}

// New returns a new Group of a given size.
func New(size int) Group {
	var g errgroup.Group
	g.SetLimit(size)
	return &g
}

var _ Group = &skipAwareGroup{}

// NewSkipAware returns a new Group of a given size and aware of pipe skips.
func NewSkipAware(g Group) Group {
	return &skipAwareGroup{g: g}
}

type skipAwareGroup struct {
	g        Group
	skipErr  error
	skipOnce sync.Once
}

// Go execs runs `fn` and saves the result if no error has been encountered.
func (s *skipAwareGroup) Go(fn func() error) {
	s.g.Go(func() error {
		err := fn()
		// if the err is a skip, set it for later, but return nil for now so the
		// group proceeds.
		if pipe.IsSkip(err) {
			s.skipOnce.Do(func() {
				s.skipErr = err
			})
			return nil
		}
		return err
	})
}

// Wait waits for Go to complete and returns the first error encountered.
func (s *skipAwareGroup) Wait() error {
	// if we got a "real error", return it, otherwise return skipErr or nil.
	if err := s.g.Wait(); err != nil {
		return err
	}
	return s.skipErr
}
