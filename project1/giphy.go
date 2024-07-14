package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type giphy struct {
	apiKey string
}

func newGiphy(apiKey string) *giphy {
	return &giphy{
		apiKey: apiKey,
	}
}

func (g giphy) MemeIDs(tags ...string) ([]string, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.giphy.com/v1/gifs/search", nil)
	if err != nil {
		return nil, fmt.Errorf("forming request err: %w", err)
	}
	q := req.URL.Query()
	q.Add("api_key", g.apiKey)
	q.Add("q", strings.Join(tags, " "))
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending request err: %w", err)
	}
	defer res.Body.Close()

	type responseBody struct {
		Data []struct {
			ID string `json:"id"`
		} `json:"data"`
	}

	var body responseBody

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return nil, fmt.Errorf("decoding err: %w", err)
	}

	memeCount := len(body.Data)
	memeIDs := make([]string, memeCount)

	for i, meme := range body.Data {
		memeIDs[i] = meme.ID
	}

	return memeIDs, nil
}

func (g giphy) MemeDownload(id string) ([]byte, error) {
	memeFile := id + "." + "gif"
	url := url.URL{
		Scheme: "https",
		Host:   "i.giphy.com",
		Path:   memeFile,
	}
	res, err := http.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("get request err: %w for URL: %s", err, url.String())
	}
	defer res.Body.Close()

	var bbuf bytes.Buffer
	if _, err := io.Copy(&bbuf, res.Body); err != nil {
		return nil, fmt.Errorf("copying response body err: %w", err)
	}
	return bbuf.Bytes(), nil
}
