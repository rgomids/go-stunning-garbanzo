package middleware

import (
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

// HTTP ...
func HTTP(next http.Handler) http.Handler {
	return loggingMiddleware(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Do stuff here
			// ...
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
		}),
	)
}
