package static

import (
	"net/http"

	"github.com/juniorrodes/mtg-project/pkg/router"
	"github.com/tdewolff/minify/v2"
)


func StaticContentProvider(r *router.Router) {
    fileServer := http.FileServer(http.Dir("./static"))

    minifier := minify.New()

    r.Get("/static/", minifier.Middleware(http.StripPrefix("/static/", fileServer)).ServeHTTP)
}
