package controller

import (
	"context"
	"log"
	"net/http"

	cardView "github.com/juniorrodes/mtg-project/components/cards"
	"github.com/juniorrodes/mtg-project/pkg/scryfall-client/models"
	scryfallModels "github.com/juniorrodes/mtg-project/pkg/scryfall-client/models"
)

type searcher interface {
	CardsSearch(searchParameters scryfallModels.CardSearchParams) (*scryfallModels.List[scryfallModels.Cards], error)
}

type CardsController struct {
	logger *log.Logger
	searcher
}

func NewCardsController(logger *log.Logger, cardSearcher searcher) *CardsController {
	return &CardsController{
		logger:   logger,
		searcher: cardSearcher,
	}
}

func (c *CardsController) Search(w http.ResponseWriter, req *http.Request) {
	c.logger.Printf("query is: %s", req.URL.RawQuery)
	searchParams := scryfallModels.CardSearchParams{
		Query: req.URL.Query().Get("q"),
	}

	cards, err := c.searcher.CardsSearch(searchParams)
	if err != nil {
		customErr, ok := err.(models.ScryfallError)
		if ok {
			w.WriteHeader(customErr.Status)
			cardView.CardError(customErr).Render(req.Context(), w)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	cardView.CardView().Render(context.WithValue(req.Context(), "cards", cards.Data), w)
}
