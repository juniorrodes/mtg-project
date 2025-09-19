package scryfallclient

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/juniorrodes/mtg-project/pkg/scryfall-client/models"
)

func (c *Client) CardsSearch(searchParameters models.CardSearchParams) (*models.List[models.Cards], error) {
	req, err := baseRequest(http.MethodGet, "cards/search", nil)
	if err != nil {
		return nil, err
	}
	log.Println("Search parameters are:", searchParameters)
	query := req.URL.Query()
	query.Add("q", searchParameters.Query)
	req.URL.RawQuery = query.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	log.Println("Calling scryfall with: ", req.URL.String())

	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest && resp.StatusCode <= 599 {
		var errModel models.ScryfallError
		err = json.NewDecoder(resp.Body).Decode(&errModel)
		if err != nil {
			return nil, err
		}
		return nil, errModel
	}

	var cards models.List[models.Cards]
	err = json.NewDecoder(resp.Body).Decode(&cards)
	if err != nil {
		return nil, err
	}

	return &cards, nil
}
