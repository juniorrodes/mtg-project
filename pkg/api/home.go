package api

import (
	"net/http"

	"github.com/juniorrodes/mtg-project/components"
)

type HomeController struct {
}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (h *HomeController) Index(w http.ResponseWriter, r *http.Request) {
	err := components.Home().Render(r.Context(), w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
}
