package main

import (
	"context"
	"fmt"

	"github.com/brandonyeoxg/project52/project2/chain"
)

func main() {
	ctx := context.Background()

	err := chain.DoSequential(ctx, stage1(), stage2())
	if err != nil {
		panic("not supposed to have panic")
	}

	err = chain.DoSequential(ctx,
		stage1(),
		stage2(),
		chain.NewConcurrent(chain.Chain(
			stage3(),
			stage4(),
			stage5()),
		),
		stage6(),
		stage7(),
	)
	if err != nil {
		panic("not supposed to panic")
	}

}

func stage1() chain.ChainFn {
	return func(ctx context.Context) error {
		fmt.Println("Do stuff 1")
		return nil
	}
}

func stage2() chain.ChainFn {
	return func(ctx context.Context) error {
		fmt.Println("Do stuff 2")
		return nil
	}
}

func stage3() chain.ChainFn {
	return func(ctx context.Context) error {
		fmt.Println("Do stuff 3")
		return nil
	}
}

func stage4() chain.ChainFn {
	return func(ctx context.Context) error {
		fmt.Println("Do stuff 4")
		return nil
	}
}

func stage5() chain.ChainFn {
	return func(ctx context.Context) error {
		fmt.Println("do stuff 5")
		return nil
	}
}
func stage6() chain.ChainFn {
	return func(ctx context.Context) error {
		fmt.Println("do stuff 6")
		return nil
	}
}

func stage7() chain.ChainFn {
	return func(ctx context.Context) error {
		fmt.Println("Do stuff 7")
		return nil
	}
}
