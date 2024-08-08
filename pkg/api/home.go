package api

import (
	"net/http"

	"github.com/juniorrodes/mtg-project/components"
	"github.com/juniorrodes/mtg-project/pkg/router"
)

func Index(r *router.Router) {
	r.Get("/", renderHomePage)
}

func renderHomePage(w http.ResponseWriter, r *http.Request) {
    err := components.Home().Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
}
