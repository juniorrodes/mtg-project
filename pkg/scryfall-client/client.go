package scryfallclient

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
)

const (
	ScryfallDomanin = "https://api.scryfall.com/"
)

type Client struct {
	logger     *log.Logger
	httpClient *http.Client
}

func NewClient(logger *log.Logger) *Client {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	return &Client{
		logger:     logger,
		httpClient: httpClient,
	}
}

func baseRequest(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, ScryfallDomanin+path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "mtg-project/0.1")

	return req, nil
}
