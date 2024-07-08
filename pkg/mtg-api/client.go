package mtg

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"

	"github.com/juniorrodes/mtg-project/pkg/mtg-api/models"
)

const (
	MTGDomanin = "https://api.magicthegathering.io/v1/"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	return &Client{
		httpClient: httpClient,
	}
}

func (c *Client) GetCards() ([]models.Card, error) {
	req, err := http.NewRequest(http.MethodGet, MTGDomanin+"cards", nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var cards []models.Card
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &cards)
	if err != nil {
		return nil, err
	}
	return cards, nil
}
