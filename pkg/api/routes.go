package api

import (
	"log"
	"net/http"

	"github.com/juniorrodes/mtg-project/pkg/controller"
	"github.com/juniorrodes/mtg-project/pkg/router"
	scryfallclient "github.com/juniorrodes/mtg-project/pkg/scryfall-client"
)

func Routes(r *router.Router, logger *log.Logger) {
	log.Println("Initializing routes ...")

	scryfallClient := scryfallclient.NewClient(logger)

	cardsController := controller.NewCardsController(logger, scryfallClient)
	homeController := controller.NewHomeController(logger)
	fs := http.FileServer(http.Dir("static"))

	r.Get("/", homeController.RenderHomePage)
	r.Get("/static/", http.StripPrefix("/static/", fs).ServeHTTP)
	r.Get("/search", cardsController.Search)
}
