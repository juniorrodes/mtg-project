package controller

import (
	"net/http"

	"github.com/juniorrodes/mtg-project/components"
	"github.com/juniorrodes/mtg-project/pkg/mtg-api/models"
)

type mtgClient interface {
	GetCards(int) ([]models.Card, error)
}

type CardSearchController struct {
	mtgClient mtgClient
}

func NewCardSearchController(mtgClient mtgClient) *CardSearchController {
	return &CardSearchController{
		mtgClient: mtgClient,
	}
}

func (c *CardSearchController) Search(w http.ResponseWriter, r *http.Request) {
	cards, err := c.mtgClient.GetCards(10)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if err = components.CardView(cards).Render(r.Context(), w); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}
