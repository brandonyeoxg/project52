package chain

import (
	"context"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConcurrent(t *testing.T) {
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
	c := NewConcurrent(Chain(testChainFn0, testChainFn1, testChainFn2))
	assert.Len(t, c.Doers, 3)
}

func TestConcurrent_Do(t *testing.T) {
	var mut sync.Mutex
	generator := func(called *int) ChainFn {
		return func(ctx context.Context) error {
			mut.Lock()
			defer mut.Unlock()
			*called++
			return nil
		}
	}

	t.Run("no doers", func(t *testing.T) {
		c := NewConcurrent(nil)
		assert.Nil(t, c.Do(context.Background()))
	})

	t.Run("1 doer", func(t *testing.T) {
		var called int
		testChainFn := generator(&called)

		c := NewConcurrent(Chain(testChainFn))
		assert.Nil(t, c.Do(context.Background()))
		assert.Equal(t, called, 1)
	})

	t.Run("multiple doers", func(t *testing.T) {
		var called int
		testChainFn := generator(&called)

		c := NewConcurrent(Chain(testChainFn, testChainFn, testChainFn))
		assert.Nil(t, c.Do(context.Background()))
		assert.Equal(t, called, 3)
	})

	t.Run("multiple doers one error for early exit", func(t *testing.T) {
		var called int
		testChainFn := generator(&called)

		var errChainFn ChainFn = func(ctx context.Context) error {
			return assert.AnError
		}

		c := NewConcurrent(Chain(testChainFn, testChainFn, errChainFn, testChainFn))
		assert.NotNil(t, c.Do(context.Background()))
		assert.LessOrEqual(t, called, 3)
	})

	t.Run("multiple doers one error for no early exit", func(t *testing.T) {
		var called int
		testChainFn := generator(&called)

		var errChainFn ChainFn = func(ctx context.Context) error {
			return assert.AnError
		}

		c := NewConcurrent(Chain(testChainFn, testChainFn, errChainFn, testChainFn), WithEarlyExit(false))
		assert.NotNil(t, c.Do(context.Background()))
		assert.Equal(t, called, 3)
	})
}

func TestDoConcurrent(t *testing.T) {
	var mut sync.Mutex
	generator := func(called *int) ChainFn {
		return func(ctx context.Context) error {
			mut.Lock()
			defer mut.Unlock()
			*called++
			return nil
		}
	}

	t.Run("no doers", func(t *testing.T) {
		assert.Nil(t, DoConcurrent(context.Background()))
	})

	t.Run("1 doer", func(t *testing.T) {
		var called int
		testChainFn := generator(&called)

		assert.Nil(t, DoConcurrent(context.Background(), testChainFn))
		assert.Equal(t, called, 1)
	})

	t.Run("multiple doers", func(t *testing.T) {
		var called int
		testChainFn := generator(&called)

		assert.Nil(t, DoConcurrent(context.Background(), testChainFn, testChainFn, testChainFn))
		assert.Equal(t, called, 3)
	})

	t.Run("multiple doers one error for early exit", func(t *testing.T) {
		var called int
		testChainFn := generator(&called)

		var errChainFn ChainFn = func(ctx context.Context) error {
			return assert.AnError
		}

		assert.NotNil(t, DoConcurrent(context.Background(), testChainFn, testChainFn, errChainFn, testChainFn))
		assert.LessOrEqual(t, called, 3)
	})
}

func TestWithMaxConcurrency(t *testing.T) {
	maxConcurrency := 77
	c := NewConcurrent(nil, WithMaxConcurrency(maxConcurrency))
	assert.Equal(t, maxConcurrency, c.maxConcurrency)
}

func TestWithEarlyExit(t *testing.T) {
	earlyExit := false
	c := NewConcurrent(nil, WithEarlyExit(false))
	assert.Equal(t, earlyExit, c.earlyExit)
}
