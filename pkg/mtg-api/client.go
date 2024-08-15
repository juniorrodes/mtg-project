package mtg

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"slices"
	"strings"

	"github.com/juniorrodes/mtg-project/pkg/models"
)

const (
	MTGDomanin = "https://api.magicthegathering.io/v1/"
	AND_OP     = "and"
	OR_OP      = "or"
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

func (c *Client) GetCards(ctx context.Context, params models.SearchParameters) ([]models.Card, error) {
	logger := ctx.Value("logger").(*log.Logger)
	req, err := http.NewRequest(http.MethodGet, MTGDomanin+"cards", nil)
	if err != nil {
		logger.Println("failed with: ", err.Error())
		return nil, err
	}
	setQueryParams(req.URL, params)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var cards models.Cards
	err = json.NewDecoder(resp.Body).Decode(&cards)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	removeDuplicates(&cards)
	removeCardsWithNoForeignNames(&cards)

	return cards.Cards, nil
}

func removeCardsWithNoForeignNames(cards *models.Cards) {
	cards.Cards = slices.DeleteFunc(cards.Cards, func(card models.Card) bool {
		return len(card.ForeignNames) < 1
	})
}

func removeDuplicates(cards *models.Cards) {
	for _, c := range cards.Cards {
		for _, variation := range c.Variations {
			cards.Cards = slices.DeleteFunc(cards.Cards, func(card models.Card) bool {
				return card.Id == variation
			})
		}
	}
}

func setQueryParams(requestURL *url.URL, params models.SearchParameters) {
	q := requestURL.Query()
	q.Add("pageSize", params.PageSize.Value)
	q.Add("name", params.Name.Value)
	if params.Color != nil {
		if params.Color.Operand == OR_OP {
			q.Add("colors", strings.Join(params.Color.Value, "|"))
		} else {
			q.Add("colors", strings.Join(params.Color.Value, ","))
		}
	}

	requestURL.RawQuery = q.Encode()

}
