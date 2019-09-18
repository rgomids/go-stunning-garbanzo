package routers

import (
	"go-stunning-garbanzo/handlers"

	"github.com/gorilla/mux"
)

func cardRoutes(r *mux.Router) {
	r.HandleFunc("/api/cards", handlers.GetAllCards).Methods("GET")
}
