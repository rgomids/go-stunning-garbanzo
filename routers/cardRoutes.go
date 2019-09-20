package routers

import (
	"go-stunning-garbanzo/handlers"
	"go-stunning-garbanzo/server"

	"github.com/gorilla/mux"
)

func cardRoutesHTTP(r *mux.Router) {
	r.HandleFunc("/api/card", handlers.AddCardHTTP).Methods("POST")
	r.HandleFunc("/api/card/{id}", handlers.GetCardHTTP).Methods("GET")
	r.HandleFunc("/api/cards", handlers.GetAllCardsHTTP).Methods("GET")
	r.HandleFunc("/api/card/{id}", handlers.UpdateCardHTTP).Methods("PUT")
	r.HandleFunc("/api/card/{id}", handlers.DeleteCardHTTP).Methods("DELETE")
}

func cardRoutesWS(h *server.EventHub) {
	h.AddHandler("ADD_CARD", handlers.AddCardWS)
	h.AddHandler("GET_CARD", handlers.GetCardWS)
	h.AddHandler("GET_CARDS", handlers.GetAllCardsWS)
	h.AddHandler("UPDATE_CARD", handlers.UpdateCardWS)
	h.AddHandler("DELETE_CARD", handlers.DeleteCardWS)
}
