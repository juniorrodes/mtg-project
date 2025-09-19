package main

import (
	"log"
	"os"

	"github.com/juniorrodes/mtg-project/pkg/api"
	"github.com/juniorrodes/mtg-project/pkg/router"
)

func main() {
	logger := log.New(os.Stdout, "[mtg-app]", log.Ldate|log.Lshortfile|log.Ltime)

	router := router.NewRouter(logger)
	api.Routes(router, logger)

	router.ListenAndServe(":8080")
}
