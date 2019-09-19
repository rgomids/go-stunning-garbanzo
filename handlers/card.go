package handlers

import (
	"fmt"
	"go-stunning-garbanzo/server"
	"net/http"
)

// GetAllCardsHTTP ...
func GetAllCardsHTTP(rsw http.ResponseWriter, req *http.Request) {

}

// GetAllCardsWS ...
func GetAllCardsWS(message *server.EventMessage) {
	fmt.Println("Chegou")
}
