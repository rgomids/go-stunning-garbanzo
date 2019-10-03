package handlers

import (
	"log"
	"net/http"

	wsh "github.com/RafaelGomides/go-wsh"
	"github.com/gorilla/mux"
)

// ServeWs Cria as sessões
func ServeWs(hub *wsh.EventHub, w http.ResponseWriter, r *http.Request) {
	wsConn, err := wsh.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[ERRO] Creating websocket session: %v", err)
	}
	newClientSes := wsh.NewClientSession()
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

func addClientGroup(groupID string, newClientSes *wsh.ClientSession, hub *wsh.EventHub) string {
	newClientGroup := wsh.NewClientGroup(groupID)
	newClientGroup.AddClientSession(newClientSes)
	hub.AddGroup(newClientGroup)
	return newClientGroup.ID
}
