package router

import (
	"fmt"
	"log"
	"net/http"
)

type Router struct {
	logger *log.Logger
	mux    *http.ServeMux
}

func NewRouter(logger *log.Logger) *Router {
	mux := http.NewServeMux()
	return &Router{
		logger: logger,
		mux:    mux,
	}
}

func (r *Router) Get(path string, handler http.HandlerFunc) {
	path = fmt.Sprintf("%s %s", http.MethodGet, path)
	r.mux.HandleFunc(path, handler)
}

func (r *Router) Post(path string, handler http.HandlerFunc) {
	path = fmt.Sprintf("%s %s", http.MethodPost, path)
	r.mux.HandleFunc(path, handler)
}

func (r *Router) ListenAndServe(addr string) error {
	server := &http.Server{
		Addr:    addr,
		Handler: r.mux,
	}

	r.logger.Printf("Server listening on %s", addr)

	return server.ListenAndServe()
}
