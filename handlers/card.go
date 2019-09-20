package handlers

import (
	"encoding/json"
	"fmt"
	"go-stunning-garbanzo/models"
	"go-stunning-garbanzo/server"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

func addCard(card *models.Card) (string, error) {
	return models.CreateCard(card)
}

// AddCardHTTP ...
func AddCardHTTP(rsw http.ResponseWriter, req *http.Request) {
	cardRaw, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(rsw, err.Error(), http.StatusInternalServerError)
		return
	}

	var card models.Card
	err = json.Unmarshal(cardRaw, &card)
	if err != nil {
		http.Error(rsw, err.Error(), http.StatusInternalServerError)
		return
	}

	cardID, err := addCard(&card)
	if err != nil {
		http.Error(rsw, err.Error(), http.StatusInternalServerError)
		return
	}

	rsw.Write([]byte(cardID))
}

// AddCardWS ...
func AddCardWS(message *server.EventMessage) {
	var card models.Card
	err := mapstructure.Decode(message.Data, &card)
	if err != nil {
		message.Client.SendMessage(&server.EventMessage{Event: "INTERNAL_SERVER_ERROR", Data: nil})
		return
	}

	cardID, err := addCard(&card)
	if err != nil {
		message.Client.SendMessage(&server.EventMessage{Event: "INTERNAL_SERVER_ERROR", Data: nil})
		return
	}

	message.Client.SendMessage(&server.EventMessage{Event: "ADD_CARD_SUCCESSFUL", Data: cardID})
}

func getCard(card *models.Card) (string, error) {
	return models.CreateCard(card)
}

// GetCardHTTP ...
func GetCardHTTP(rsw http.ResponseWriter, req *http.Request) {}

// GetCardWS ...
func GetCardWS(message *server.EventMessage) {}

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
		message.Client.SendMessage(&server.EventMessage{Event: "GET_CARDS_SUCCESSFUL", Data: cards})
	} else {
		log.Printf("[ERRO] Ao tentar buscar cards: %v", err)
		message.Client.SendMessage(&server.EventMessage{Event: "INTERNAL_SERVER_ERROR", Data: nil})
	}
}

func updateCard(card *models.Card) (string, error) {
	return models.CreateCard(card)
}

// UpdateCardHTTP ...
func UpdateCardHTTP(rsw http.ResponseWriter, req *http.Request) {}

// UpdateCardWS ...
func UpdateCardWS(message *server.EventMessage) {}

func deleteCard(card *models.Card) (string, error) {
	return models.CreateCard(card)
}

// DeleteCardHTTP ...
func DeleteCardHTTP(rsw http.ResponseWriter, req *http.Request) {}

// DeleteCardWS ...
func DeleteCardWS(message *server.EventMessage) {}
