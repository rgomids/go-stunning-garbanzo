package middleware

import (
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[INFO] REQUEST RECEIVED FROM %s - URL REQUESTED %s", r.RemoteAddr, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

// Middleware ...
func Middleware(next http.Handler) http.Handler {
	return loggingMiddleware(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		}),
	)
}
