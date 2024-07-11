package mtg

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

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

func (c *Client) GetCards(pageSize int) ([]models.Card, error) {
	req, err := http.NewRequest(http.MethodGet, MTGDomanin+"cards", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Add("pageSize", strconv.Itoa(pageSize))
	req.URL.RawQuery = q.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var cards models.Cards
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	err = json.Unmarshal(body, &cards)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return cards.Cards, nil
}
