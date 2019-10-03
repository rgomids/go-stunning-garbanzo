package handlers

import (
	"encoding/json"
	"go-stunning-garbanzo/models"
	"io/ioutil"
	"log"
	"net/http"

	wsh "github.com/RafaelGomides/go-wsh"

	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
)

// AddCardHTTP ...
func AddCardHTTP(rsw http.ResponseWriter, req *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("[ERRO] Adding Card: %v\n", err)
			http.Error(rsw, "Error adding card", http.StatusInternalServerError)
		}
	}()

	cardRaw, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		return
	}

	var card models.Card
	err = json.Unmarshal(cardRaw, &card)
	if err != nil {
		return
	}

	cardID, err := models.CreateCard(&card)
	if err != nil {
		return
	}

	rsw.Write([]byte(cardID))
}

// AddCardWS ...
func AddCardWS(message *wsh.EventMessage) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("[ERRO] Adding Card: %v\n", err)
			message.Client.SendMessage(&wsh.EventMessage{Event: "INTERNAL_SERVER_ERROR"})
		}
	}()
	var card models.Card

	err = mapstructure.Decode(message.Data, &card)
	if err != nil {
		return
	}

	cardID, err := models.CreateCard(&card)
	if err != nil {
		return
	}

	message.Client.SendBroadcast(&wsh.EventMessage{Event: "ADD_CARD_SUCCESSFUL", Data: cardID})
}

// GetCardHTTP ...
func GetCardHTTP(rsw http.ResponseWriter, req *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("[ERRO] Getting Card: %v\n", err)
			http.Error(rsw, "Error getting card", http.StatusInternalServerError)
		}
	}()

	cardID := mux.Vars(req)["id"]
	if cardID == "" {
		http.Error(rsw, "Card ID is null", http.StatusBadRequest)
		return
	}

	card, err := models.GetCard(cardID)
	if err != nil {
		return
	}

	cardJSON, err := json.Marshal(card)
	if err != nil {
		return
	}

	rsw.Write(cardJSON)
}

// GetCardWS ...
func GetCardWS(message *wsh.EventMessage) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("[ERRO] Getting Card: %v\n", err)
			message.Client.SendMessage(&wsh.EventMessage{Event: "INTERNAL_SERVER_ERROR"})
		}
	}()

	cardID := message.Data.(string)
	if cardID == "" {
		message.Client.SendMessage(&wsh.EventMessage{Event: "BAD_REQUEST", Data: "Card ID is null"})
		return
	}

	card, err := models.GetCard(cardID)
	if err != nil {
		return
	}

	message.Client.SendMessage(&wsh.EventMessage{Event: "GET_CARD_SUCCESSFUL", Data: card})

}

// GetAllCardsHTTP ...
func GetAllCardsHTTP(rsw http.ResponseWriter, req *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("[ERRO] Getting Cards: %v\n", err)
			http.Error(rsw, "Error getting cards", http.StatusInternalServerError)
		}
	}()

	cards, err := models.GetCards()
	if err != nil {
		return
	}

	if len(cards) == 0 {
		log.Printf("[WARN] Cards not found")
		http.NotFound(rsw, req)
		return
	}

	cardsJSON, err := json.Marshal(cards)
	if err != nil {
		return
	}

	rsw.Write(cardsJSON)
}

// GetAllCardsWS ...
func GetAllCardsWS(message *wsh.EventMessage) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("[ERRO] Getting Cards: %v\n", err)
			message.Client.SendMessage(&wsh.EventMessage{Event: "INTERNAL_SERVER_ERROR"})
		}
	}()

	cards, err := models.GetCards()
	if err != nil {
		return
	}

	if len(cards) == 0 {
		log.Printf("[WARN] Cards not found")
		message.Client.SendMessage(&wsh.EventMessage{Event: "NOT_FOUND"})
	}

	message.Client.SendMessage(&wsh.EventMessage{Event: "GET_CARDS_SUCCESSFUL", Data: cards})
}

// UpdateCardHTTP ...
func UpdateCardHTTP(rsw http.ResponseWriter, req *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("[ERRO] Updating Cards: %v\n", err)
			http.Error(rsw, "Error updating cards", http.StatusInternalServerError)
		}
	}()

	cardRaw, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		return
	}

	var cardOld models.Card
	err = json.Unmarshal(cardRaw, &cardOld)
	if err != nil {
		return
	}

	cardUpdated, err := models.UpdateCard(&cardOld)
	if err != nil {
		return
	}

	rsw.Write([]byte(cardUpdated))
}

// UpdateCardWS ...
func UpdateCardWS(message *wsh.EventMessage) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("[ERRO] Updating Cards: %v\n", err)
			message.Client.SendMessage(&wsh.EventMessage{Event: "INTERNAL_SERVER_ERROR"})
		}
	}()

	var cardOld models.Card
	err = mapstructure.Decode(message.Data, &cardOld)
	if err != nil {
		return
	}

	cardUpdated, err := models.UpdateCard(&cardOld)
	if err != nil {
		return
	}

	message.Client.SendBroadcast(&wsh.EventMessage{Event: "UPDATE_CARD_SUCCESSFUL", Data: cardUpdated})
}

// DeleteCardHTTP ...
func DeleteCardHTTP(rsw http.ResponseWriter, req *http.Request) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("[ERRO] Deleting Cards: %v\n", err)
			http.Error(rsw, "Error deleting cards", http.StatusInternalServerError)
		}
	}()

	cardID := mux.Vars(req)["id"]
	if cardID == "" {
		http.Error(rsw, "Card ID is null", http.StatusBadRequest)
		return
	}

	cardID, err = models.DeleteCard(cardID)
	if err != nil {
		return
	}

	rsw.Write([]byte(cardID))
}

// DeleteCardWS ...
func DeleteCardWS(message *wsh.EventMessage) {
	var err error
	defer func() {
		if err != nil {
			log.Printf("[ERRO] Deleting Cards: %v\n", err)
			message.Client.SendMessage(&wsh.EventMessage{Event: "INTERNAL_SERVER_ERROR"})
		}
	}()

	cardID, err := models.DeleteCard(message.Data.(string))
	if err != nil {
		return
	}

	message.Client.SendBroadcast(&wsh.EventMessage{Event: "DELETE_CARD_SUCCESSFUL", Data: cardID})
}
