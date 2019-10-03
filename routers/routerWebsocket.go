package routers

import (
	"go-stunning-garbanzo/handlers"
	"log"
	"net/http"

	wsh "github.com/RafaelGomides/go-wsh"

	"github.com/gorilla/mux"
)

var (
	hub *wsh.EventHub
)

func init() {
	log.Println("[INFO] Starting Websocket Hub")
	hub = wsh.NewEventHub()
	go hub.Run()
}

func routerWebsocket(r *mux.Router) {
	eventHandlerRegistry()
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeWs(hub, w, r)
	})
	r.HandleFunc("/ws/{group_id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.ServeWs(hub, w, r)
	})
}

func eventHandlerRegistry() {
	cardRoutesWS(hub)
}
