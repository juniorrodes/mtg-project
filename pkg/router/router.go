package router

import (
	"fmt"
	"log"
	"net/http"
)

type Router struct {
	mux *http.ServeMux
}

func NewRouter() *Router {
	mux := http.NewServeMux()
	return &Router{
		mux: mux,
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

	log.Printf("Server listening on %s", addr)

	return server.ListenAndServe()
}
