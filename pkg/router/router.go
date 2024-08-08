package router

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type ErrorModel struct {
	Motive       string `json:"motive"`
	ErrorMessage string `json:"message"`
}

type Router struct {
	mux    *http.ServeMux
	logger *log.Logger
}

func NewRouter() *Router {
	mux := http.NewServeMux()
	logger := log.New(os.Stdout, "[mtg-api]", log.LUTC|log.Lshortfile|log.Lmicroseconds)
	return &Router{
		mux:    mux,
		logger: logger,
	}
}

func (r *Router) Get(path string, handler http.HandlerFunc) {
	path = fmt.Sprintf("%s %s", http.MethodGet, path)

	r.mux.HandleFunc(path, LoggerInjector(handler, r.logger))
}

func (r *Router) Post(path string, handler http.HandlerFunc) {
	path = fmt.Sprintf("%s %s", http.MethodPost, path)
	r.mux.HandleFunc(path, LoggerInjector(handler, r.logger))
}

func (r *Router) ListenAndServe(addr string) error {
	server := &http.Server{
		Addr:    addr,
		Handler: r.mux,
	}

	log.Printf("Server listening on %s", addr)

	return server.ListenAndServe()
}
