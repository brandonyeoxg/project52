package chain

import (
	"context"

	"golang.org/x/sync/errgroup"
)

const (
	defaultMaxConcurrency = 11
)

type ConcurrentOption func(c *Concurrent)

type Concurrent struct {
	Doers []Doer

	maxConcurrency int
	earlyExit      bool
}

func NewConcurrent(doers []Doer, opts ...ConcurrentOption) Concurrent {
	var concurrent Concurrent
	concurrent.Doers = doers
	concurrent.maxConcurrency = defaultMaxConcurrency
	concurrent.earlyExit = true

	for _, opt := range opts {
		opt(&concurrent)
	}

	return concurrent
}

func (c Concurrent) Do(ctx context.Context) error {
	group := new(errgroup.Group)
	if c.earlyExit {
		group, ctx = errgroup.WithContext(ctx)
	}
	group.SetLimit(c.maxConcurrency)

	for _, doer := range c.Doers {
		group.Go(func() error {
			return doer.Do(ctx)
		})
	}
	return group.Wait()
}

func DoConcurrent(ctx context.Context, doers ...Doer) error {
	c := Concurrent{Doers: doers}
	c.maxConcurrency = defaultMaxConcurrency
	c.earlyExit = true
	return c.Do(ctx)
}

func WithMaxConcurrency(max int) ConcurrentOption {
	return func(c *Concurrent) {
		c.maxConcurrency = max
	}
}

func WithEarlyExit(isEarlyExit bool) ConcurrentOption {
	return func(c *Concurrent) {
		c.earlyExit = isEarlyExit
	}
}
