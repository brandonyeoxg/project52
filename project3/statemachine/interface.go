package statemachine

import (
	"context"
)

type State[T any] func(ctx context.Context, arg T) (T, State[T], error)

func Run[T any](ctx context.Context, state State[T], arg T) (T, error) {
	currentState := state
	for currentState != nil {
		newArg, newState, err := currentState(ctx, arg)
		if err != nil {
			return arg, err
		}
		if newState == nil {
			return arg, nil
		}
		currentState = newState
		arg = newArg
	}
	return arg, nil
}
