package middleware

import (
	"log"
	"net/http"
	"strings"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func Middleware(next http.Handler) http.Handler {
	return loggingMiddleware(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			urlSplited := strings.Split(r.URL.String(), "/")
			switch urlSplited[1] {
			case "api":
				midHTTP(w, r)
			case "ws":
				midWS(w, r)
			}
			next.ServeHTTP(w, r)
		}),
	)
}

func midHTTP(w http.ResponseWriter, r *http.Request) {

}

func midWS(w http.ResponseWriter, r *http.Request) {

}
