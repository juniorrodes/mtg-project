package search

import (
	"github.com/juniorrodes/mtg-project/pkg/api/search/cards"
	"github.com/juniorrodes/mtg-project/pkg/router"
)

func Search(r *router.Router, client cards.MtgClient) {
	searcher := cards.NewCardSearcher(client)
	r.Post("/api/search", searcher.RenderCardView)
}
