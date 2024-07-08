package main

import (
	"github.com/juniorrodes/mtg-project/pkg/api"
	"github.com/juniorrodes/mtg-project/pkg/router"
)

func main() {
	router := router.NewRouter()
	api.Routes(router)

	router.ListenAndServe(":8080")
}
