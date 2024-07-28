package chain

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChainFn_Do(t *testing.T) {
	var isCalled bool
	generator := func() ChainFn {
		return func(ctx context.Context) error {
			isCalled = true
			return nil
		}
	}

	chainFn := generator()
	err := chainFn.Do(context.Background())
	assert.Nil(t, err)
	assert.True(t, isCalled)
}

func TestChain(t *testing.T) {
	var testChainFn ChainFn
	testChainFn = func(ctx context.Context) error {
		return nil
	}
	type args struct {
		doers []Doer
	}
	tests := []struct {
		name string
		args args
		want []Doer
	}{
		{
			name: "no doers",
			args: args{
				doers: nil,
			},
			want: nil,
		},
		{
			name: "one doer",
			args: args{
				doers: []Doer{testChainFn},
			},
			want: []Doer{testChainFn},
		},
		{
			name: "multiple doers",
			args: args{
				doers: []Doer{testChainFn, testChainFn, testChainFn},
			},
			want: []Doer{testChainFn, testChainFn, testChainFn},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doers := Chain(tt.args.doers...)
			assert.Len(t, doers, len(tt.args.doers))
		})
	}
}
