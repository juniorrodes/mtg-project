package api

import (
	"github.com/juniorrodes/mtg-project/pkg/api/controller"
	"github.com/juniorrodes/mtg-project/pkg/mtg-api"
	"github.com/juniorrodes/mtg-project/pkg/router"
)

func Routes(r *router.Router) {
	mtgClient := mtg.NewClient()

	hc := NewHomeController()
	csc := controller.NewCardSearchController(mtgClient)

	r.Post("/api/search", csc.Search)
	r.Get("/", hc.Index)
}
