package mtg

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/juniorrodes/mtg-project/pkg/mtg-api/models"
)

const (
	MTGDomanin = "https://api.magicthegathering.io/v1/"
	AND_OP     = "and"
	OR_OP      = "or"
)

type nameParam struct {
	Value string `form:"value"`
}

type colorParam struct {
	Value   []string `form:"value"`
	Operand string   `form:"operand"`
}

type searchParameters struct {
	Name  *nameParam  `form:"name"`
	Color *colorParam `form:"color"`
}

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

func (c *Client) GetCards(ctx context.Context, requestBytes []byte, pageSize int) ([]models.Card, error) {
	logger := ctx.Value("logger").(*log.Logger)
	req, err := http.NewRequest(http.MethodGet, MTGDomanin+"cards", nil)
	if err != nil {
		logger.Println("failed with: ", err.Error())
		return nil, err
	}
	var requestBody searchParameters
	err = json.Unmarshal(requestBytes, &requestBody)
	if err != nil {
		logger.Println("failed with: ", err.Error())
		return nil, err
	}

	q := req.URL.Query()
	q.Add("pageSize", strconv.Itoa(pageSize))
	q.Add("name", requestBody.Name.Value)
	if requestBody.Color != nil {
		if requestBody.Color.Operand == OR_OP {
			q.Add("colors", strings.Join(requestBody.Color.Value, "|"))
		} else {
			q.Add("colors", strings.Join(requestBody.Color.Value, ","))
		}
	}

	req.URL.RawQuery = q.Encode()

	logger.Println(req.URL.RawQuery)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var cards models.Cards
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
	logger.Println(cards)
	return cards.Cards, nil
}
