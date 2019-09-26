package handlers

import (
	"go-stunning-garbanzo/server"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

	groupID := mux.Vars(r)["group_id"]
	if groupID == "" {
		newClientSes.Group = addClientGroup(groupID, newClientSes, hub)
	} else {
		group, isAdded := hub.ClientGroups[groupID]
		if isAdded {
			group.AddClientSession(newClientSes)
		} else {
			addClientGroup(groupID, newClientSes, hub)
		}
		newClientSes.Group = groupID
	}

	go newClientSes.WriteToSocket()
	newClientSes.ReadFromSocket()
}

func addClientGroup(groupID string, newClientSes *server.ClientSession, hub *server.EventHub) string {
	newClientGroup := server.NewClientGroup(groupID)
	newClientGroup.AddClientSession(newClientSes)
	hub.AddGroup(newClientGroup)
	return newClientGroup.ID
}
