package chain

import (
	"context"
)

type Sequential struct {
	Doers []Doer
}

func NewSequential(doers []Doer) Sequential {
	return Sequential{
		Doers: doers,
	}
}

func (s Sequential) Do(ctx context.Context) error {
	for _, chain := range s.Doers {
		if err := chain.Do(ctx); err != nil {
			return err
		}
	}
	return nil
}

func DoSequential(ctx context.Context, doers ...Doer) error {
	s := Sequential{Doers: doers}
	return s.Do(ctx)
}
