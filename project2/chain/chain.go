package chain

import (
	"context"
)

type Doer interface {
	Do(ctx context.Context) error
}

type ChainFn func(ctx context.Context) error

func (fn ChainFn) Do(ctx context.Context) error {
	return fn(ctx)
}

func Chain(doers ...Doer) []Doer {
	return doers
}
