package api

import (
	"github.com/juniorrodes/mtg-project/pkg/api/search"
	"github.com/juniorrodes/mtg-project/pkg/api/static"
	"github.com/juniorrodes/mtg-project/pkg/mtg-api"
	"github.com/juniorrodes/mtg-project/pkg/router"
)

func Routes(r *router.Router) {
	client := mtg.NewClient()
	search.Search(r, client)
	static.StaticContentProvider(r)
	Index(r)
}
