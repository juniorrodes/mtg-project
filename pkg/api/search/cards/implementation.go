package cards

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/juniorrodes/mtg-project/components"
	"github.com/juniorrodes/mtg-project/pkg/models"
)

type MtgClient interface {
	GetCards(ctx context.Context, params models.SearchParameters) ([]models.Card, error)
}

type cardSearcher struct {
	mtgClient MtgClient
}

func NewCardSearcher(client MtgClient) *cardSearcher {
	return &cardSearcher{
		mtgClient: client,
	}
}

func (c *cardSearcher) RenderCardView(w http.ResponseWriter, r *http.Request) {
	logger := r.Context().Value("logger").(*log.Logger)

	requestBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Println("failed with: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		components.FailState(err.Error()).Render(r.Context(), w)
	}
	var params models.SearchParameters
	err = json.Unmarshal(requestBytes, &params)

	cards, err := c.mtgClient.GetCards(r.Context(), params)
	if err != nil {
		logger.Println("failed with: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		components.FailState(err.Error()).Render(r.Context(), w)
	}

	components.CardView(cards).Render(r.Context(), w)
}

func RenderSearchCards(w http.ResponseWriter)
