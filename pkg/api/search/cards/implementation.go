package cards

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/juniorrodes/mtg-project/components"
	"github.com/juniorrodes/mtg-project/pkg/mtg-api/models"
)

type MtgClient interface {
	GetCards(ctx context.Context, requestBytes []byte, pageSize int) ([]models.Card, error)
}

type cardSearcher struct {
	mtgClient MtgClient
}

func NewCardSearcher(client MtgClient) *cardSearcher {
	return &cardSearcher{
		mtgClient: client,
	}
}

func (c *cardSearcher) RenderSearchResults(w http.ResponseWriter, r *http.Request) {
	logger := r.Context().Value("logger").(*log.Logger)

	requestBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Println("failed with: ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		components.FailState(err.Error()).Render(r.Context(), w)
	}

	cards, err := c.mtgClient.GetCards(r.Context(), requestBytes, 10)
	if err != nil {
		logger.Println("failed with: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		components.FailState(err.Error()).Render(r.Context(), w)
	}

	components.CardView(cards).Render(r.Context(), w)
	w.WriteHeader(http.StatusOK)
}
