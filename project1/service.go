package main

import (
	"fmt"
	"math/rand"
)

type service struct {
	memeProvider MemeGetter
}

func newService(provider MemeGetter) *service {
	return &service{
		memeProvider: provider,
	}
}

func (s service) findMeme(tags ...string) ([]byte, error) {
	ids, err := s.memeProvider.MemeIDs(tags...)
	if err != nil {
		return nil, fmt.Errorf("get memes err: %w", err)
	}

	chosenID := ids[rand.Intn(len(ids))]

	meme, err := s.memeProvider.MemeDownload(chosenID)
	if err != nil {
		return nil, fmt.Errorf("downloading meme err: %w", err)
	}
	return meme, nil
}
