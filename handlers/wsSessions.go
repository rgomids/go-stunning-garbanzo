package handlers

import (
	"go-stunning-garbanzo/server"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// ServeWs Cria as sessões
func ServeWs(hub *server.EventHub, w http.ResponseWriter, r *http.Request) {
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[ERRO] Creating websocket session: %v", err)
	}
	newClientSes := server.NewClientSession()
	newClientSes.WebsocketConnection = wsConn
	newClientSes.EventsHub = hub
	go newClientSes.WriteToSocket()
	newClientSes.ReadFromSocket()
}
