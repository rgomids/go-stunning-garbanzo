package handlers

import (
	"fmt"
	"net/http"
)

// GetAllCards ...
func GetAllCards(rsw http.ResponseWriter, req *http.Request) {
	fmt.Println("Listando todos os cards")
}
