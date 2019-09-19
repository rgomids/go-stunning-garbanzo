package handlers

import (
	"go-stunning-garbanzo/server"
	"net/http"
)

// GetAllCardsHTTP ...
func GetAllCardsHTTP(rsw http.ResponseWriter, req *http.Request) {

}

// GetAllCardsWS ...
func GetAllCardsWS(message *server.EventMessage) {
	message.Client.SendResponse <- []byte(message.Data.(string))
}
