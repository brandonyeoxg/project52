package chain

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSequential(t *testing.T) {
	var (
		testChainFn0, testChainFn1, testChainFn2 ChainFn
	)

	testChainFn0 = func(ctx context.Context) error {
		return nil
	}
	testChainFn1 = func(ctx context.Context) error {
		return nil
	}
	testChainFn2 = func(ctx context.Context) error {
		return nil
	}
	s := NewSequential(Chain(testChainFn0, testChainFn1, testChainFn2))
	assert.Len(t, s.Doers, 3)
}

func TestSequential_Do(t *testing.T) {
	generator := func(called *int) ChainFn {
		return func(ctx context.Context) error {
			*called++
			return nil
		}
	}

	t.Run("no doers", func(t *testing.T) {
		s := NewSequential(nil)
		assert.Nil(t, s.Do(context.Background()))
	})

	t.Run("1 doer", func(t *testing.T) {
		var called int
		testChainFn := generator(&called)

		s := NewSequential(Chain(testChainFn))
		assert.Nil(t, s.Do(context.Background()))
		assert.Equal(t, called, 1)
	})

	t.Run("multiple doers", func(t *testing.T) {
		var called int
		testChainFn := generator(&called)

		s := NewSequential(Chain(testChainFn, testChainFn, testChainFn))
		assert.Nil(t, s.Do((context.Background())))
		assert.Equal(t, called, 3)
	})

	t.Run("multiple doers one error", func(t *testing.T) {
		var called int
		testChainFn := generator(&called)

		var errChainFn ChainFn = func(ctx context.Context) error {
			return assert.AnError
		}

		s := NewSequential(Chain(testChainFn, testChainFn, errChainFn, testChainFn))
		assert.NotNil(t, s.Do((context.Background())))
		assert.Equal(t, called, 2)
	})
}

func TestDoSequential(t *testing.T) {
	generator := func(called *int) ChainFn {
		return func(ctx context.Context) error {
			*called++
			return nil
		}
	}

	t.Run("no doers", func(t *testing.T) {
		err := DoSequential(context.Background())
		assert.Nil(t, err)
	})

	t.Run("1 doer", func(t *testing.T) {
		var called int
		testChainFn := generator(&called)

		assert.Nil(t, DoSequential(context.Background(), testChainFn))
		assert.Equal(t, called, 1)
	})

	t.Run("multiple doers", func(t *testing.T) {
		var called int
		testChainFn := generator(&called)

		assert.Nil(t, DoSequential(context.Background(), testChainFn, testChainFn, testChainFn))
		assert.Equal(t, called, 3)
	})

	t.Run("multiple doers one error", func(t *testing.T) {
		var called int
		testChainFn := generator(&called)

		var errChainFn ChainFn = func(ctx context.Context) error {
			return assert.AnError
		}

		assert.NotNil(t, DoSequential(context.Background(), testChainFn, testChainFn, errChainFn, testChainFn))
		assert.Equal(t, called, 2)
	})
}
