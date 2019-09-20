package handlers

import (
	"encoding/json"
	"fmt"
	"go-stunning-garbanzo/models"
	"go-stunning-garbanzo/server"
	"log"
	"net/http"
)

// GetAllCardsHTTP ...
func GetAllCardsHTTP(rsw http.ResponseWriter, req *http.Request) {
	if cards, err := getCards(); err == nil {
		if len(cards) == 0 {
			log.Printf("[WARN] Cards não encontrados")
			http.NotFound(rsw, req)
		}
		cardsJSON, err := json.Marshal(cards)
		if err != nil {
			log.Printf("[ERRO] Ao tentar parsear a estrutura: %v", err)
			http.Error(rsw, err.Error(), http.StatusInternalServerError)
			return
		}
		rsw.Write(cardsJSON)
	} else {
		log.Printf("[ERRO] Ao tentar buscar cards: %v", err)
		http.Error(rsw, err.Error(), http.StatusInternalServerError)
	}
}

// GetAllCardsWS ...
func GetAllCardsWS(message *server.EventMessage) {
	if cards, err := getCards(); err == nil {
		if len(cards) == 0 {
			log.Printf("[WARN] Cards não encontrados")
			message.Client.SendMessage(&server.EventMessage{Event: "NOT_FOUND", Data: nil})
		}
		message.Client.SendMessage(&server.EventMessage{Event: "GET_CARDS_SUCCESFUL", Data: cards})
	} else {
		log.Printf("[ERRO] Ao tentar buscar cards: %v", err)
		message.Client.SendMessage(&server.EventMessage{Event: "INTERNAL_SERVER_ERROR", Data: nil})
	}
}

func getCards() ([]*models.Card, error) {
	cardsRaw, err := models.GetCards()
	if err != nil {
		return nil, fmt.Errorf("Erro ao tentar coletar cards: %v", err)
	}
	if len(cardsRaw) == 0 {
		return nil, nil
	}
	return cardsRaw, nil
}
