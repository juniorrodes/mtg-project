package api

import "github.com/juniorrodes/mtg-project/pkg/router"

func Routes(r *router.Router) {
	hc := NewHomeController()
	r.Get("/", hc.Index)
}
