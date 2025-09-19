package controller

import (
	"log"
	"net/http"

	"github.com/juniorrodes/mtg-project/components/forms"
	"github.com/juniorrodes/mtg-project/components/layout"
)

type HomeController struct {
	logger *log.Logger
}

func NewHomeController(logger *log.Logger) *HomeController {
	return &HomeController{
		logger: logger,
	}
}

func (h *HomeController) RenderHomePage(w http.ResponseWriter, r *http.Request) {
	if err := layout.Layout(forms.NameSearch()).Render(r.Context(), w); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to render page"))
	}
}
