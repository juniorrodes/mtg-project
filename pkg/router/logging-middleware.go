package router

import (
	"context"
	"log"
	"net/http"
)

func LoggerInjector(handler http.HandlerFunc, logger *log.Logger) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler(w, r.WithContext(context.WithValue(r.Context(), "logger", logger)))
	})
}
